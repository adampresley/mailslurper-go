/**
 * Class: rajo.singlepage
 * This class provide a simple way to build single-page, dynamic applications.
 * It is suitable for small-medium applications where you wish to load only
 * fragements of documents via URL hash fragments. This is inspired by
 * the Pagify jQuery plugin by Chris Polis (https://github.com/cmpolis/Pagify).
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
 *    * <rajo.pubsub>
 *    * <rajo.util>
 *    * jquery
 *
 * License (BSD 2-Clause):
 *    > Copyright 2013 Adam Presley. All rights reserved.
 *    >
 *    > Redistribution and use in source and binary forms, with or without
 *    > modification, are permitted provided that the following conditions are met:
 *    >
 *    > 1. Redistributions of source code must retain the above copyright notice, this
 *    >    list of conditions and the following disclaimer.
 *    >
 *    > 2. Redistributions in binary form must reproduce the above copyright notice,
 *    >    this list of conditions and the following disclaimer in the documentation
 *    >    and/or other materials provided with the distribution.
 *    >
 *    > THIS SOFTWARE IS PROVIDED BY Adam Presley "AS IS" AND ANY EXPRESS OR IMPLIED
 *    > WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 *    > MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO
 *    > EVENT SHALL Adam Presley OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 *    > INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *    > LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 *    > PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 *    > LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
 *    > OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 *    > ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */
define(["rajo.pubsub", "jquery", "rajo.util"], function(RajoPubSub, $, RajoUtil) {
	"use strict";

	return (function() {
		var
			self = this,

			/*
			 * Event names
			 */
			_BEFORE_LOAD = "rajo.singlepage.beforeload",
			_LOAD = "rajo.singlepage.load",
			_AFTER_LOAD = "rajo.singlepage.afterload",


			/*
			 * Private methods and variables
			 */

			_assemblePublishData = function(hashParserFn, requestedView) {
				var result = hashParserFn(requestedView);
				result.config = _config;
				return result;
			},

			_getFullPath = function(path) {
				return _config.baseViewPath + path + "." + _config.viewExtension;
			},

			_getHash = function() {
				var subject = "#";
				if (_config.useHashBang) subject += "!";

				return window.location.hash.replace(subject, "");
			},

			_hashParser = function(requestedView) {
				if (!requestedView) {
					requestedView = _getHash() || _config.defaultView;
				}

				return _parsePage(requestedView);
			},

			/*
			 * Breaks the hash into parts. Page, URL params after a ?, or even slash-delimited.
			 * If there is a ? get the into a key/value object. If there are just slashes
			 * after the page then place them into an array.
			 */
			_parsePage = function(path) {
				var
					page = path,
					params = null,
					split = [];

				if (path.indexOf("?") > -1) {
					params = {};
					split = path.split("?");
					page = split[0];

					split = split[1].split("&");

					RajoUtil.each(split, function(item) {
						var temp = item.split("=");
						params[temp[0]] = (temp.length > 1) ? temp[1] : null;
					});
				} else if (path.indexOf("/") > -1) {
					params = [];
					split = path.split("/");
					page = split[0];

					if (split.length > 1) {
						RajoUtil.each(split, function(item) { params.push(item); });
					}
				}

				return {
					page: page,
					params: params
				};
			},

			_loadViewFromFile = function(data, callback) {
				$.get(_getFullPath(data.page), function(content) {
					/*
					 * Prime cache if caching is enabled.
					 */
					if (_config.cacheViews) {
						_config.views[_getFullPath(data.page)] = content;
					}

					callback(data, content);
				});
			},

			_loadViewFromCache = function(data, callback) {
				callback(data, _config.views[_getFullPath(data.page)]);
			},

			_render = function(data, content) {
				var $el = $(_config.el);

				$el[_config.animationOut](_config.animationOutSpeed, function() {
					$el.html(content)[_config.animationIn](_config.animationInSpeed, function() {
						RajoPubSub.publish(_AFTER_LOAD, data);
					});
				});
			},


			/*
			 * Default configuration
			 */
			_config = {

				/* Settings */
				el: null,
				baseViewPath: "",
				viewExtension: "html",
				cacheViews: false,
				useHashBang: false,
				defaultView: null,
				views: [],

				/* Animation */
				animationIn: "show",
				animationInSpeed: 0,
				animationOut: "hide",
				animationOutSpeed: 0,

				/* Events */
				beforeLoad: function(data) {
					RajoPubSub.publish(_LOAD, data);
				},
				load: function(data) {
					if (_config.cacheViews && _getFullPath(data.page) in _config.views) {
						_loadViewFromCache(data, _render);
					} else {
						_loadViewFromFile(data, _render);
					}
				},
				afterLoad: function(data) {}
			};

		/*
		 * Public interface
		 */
		return {
			/**
			 * Enum: PAGE_EVENTS
			 *
			 * BEFORE_LOAD - Called prior to any page loading occurs
			 * LOAD - Called to handle page loading
			 * AFTER_LOAD - Called after page loading has finished
			 */
			BEFORE_LOAD: _BEFORE_LOAD,
			LOAD: _LOAD,
			AFTER_LOAD: _AFTER_LOAD,

			/**
			 * Function: getHash
			 * Returns the current hash portion of the URI. It will strip any preceding pound sign.
			 */
			getHash: function() {
				return _getHash();
			},

			/**
			 * Function: getPublishData
			 * This function returns an object with information needed for the various
			 * events in the SinglePage class. It will provide the original configuration
			 * information as well as requested page data. This includes
			 *
			 *    * page (parsed from the hash)
			 *    * parameters
			 *
			 * Parameters:
			 *    view - The requested view/page
			 *
			 * Returns:
			 *    An object with config and requested view/page information
			 */
			getPublishData: function(view) {
				return _assemblePublishData(_hashParser, view);
			},

			/**
			 * Function: setup
			 * Sets up a listener for handling URL hash changes. It also prepares events
			 * fired before, during, and after hash change and page load. When a hash
			 * change occurs a BEFORE_LOAD event is fired, which in turn fires the
			 * LOAD event, which finally ends by calling the AFTER_LOAD event.
			 *
			 * The configuration object passed in to this function has the following
			 * possible values.
			 *
			 *    * el - Element on the page for holding view content (required)
			 *    * baseViewPath - Prefix path for where view pages are stored
			 *    * viewExtension - View file extension (defaults to "html")
			 *    * cacheViews - true/false to cache views after first load. Reduces number of AJAX calls
			 *    * views - Array of view page names
			 *
			 * Example:
			 *    > $singlepage.setup({
			 *    >    el: "#content",
			 *    >    baseViewPath: "/views",
			 *    >    viewExtension: "html",
			 *    >    views: [
			 *    >       "home",
			 *    >       "about-me",
			 *    >       "contact-us"
			 *    >    ]
			 *    > });
			 *
			 * Parameters:
			 *    config - Configuration object
			 */
			setup: function(config) {
				var iface = this;
				_config = $.extend(_config, config);

				/*
				 * Clean our base page path. Ensure it has a trailing slash.
				 */
				if (_config.baseViewPath.indexOf("/", _config.baseViewPath.length - 2) === -1) _config.baseViewPath += "/";

				/*
				 * Setup view event subscribers
				 */
				RajoPubSub.subscribe(this.BEFORE_LOAD, _config.beforeLoad);
				RajoPubSub.subscribe(this.LOAD, _config.load);
				RajoPubSub.subscribe(this.AFTER_LOAD, _config.afterLoad);

				/*
				 * Listen for hash changes, then load the initial page.
				 */
				$(window).bind("hashchange", function() {
					RajoPubSub.publish(iface.BEFORE_LOAD, iface.getPublishData());
				});

				if (window.location.hash)
					RajoPubSub.publish(this.BEFORE_LOAD, this.getPublishData());
				else if(_config.defaultView)
					RajoPubSub.publish(this.BEFORE_LOAD, this.getPublishData(_config.defaultView));
			}
		};
	}());
});
