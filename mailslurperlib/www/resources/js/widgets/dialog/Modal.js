// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery",
		"modules/util/WidgetTools",
		"jqueryui"
	],

	function($, WidgetTools) {
		"use strict";

		var
			_template = "<div id=\"{{id}}\">{{body}}</div>",
			_iconStyle = "style=\"font-size: 24px; margin-right: 8px; margin-bottom: 8px;\"",

			_buildDom = function(id, body, config) {
				var $body = _template.replace("{{id}}", id).replace("{{body}}", body);

				$("body").append($body);
				$("#" + id).dialog(config)
			},

			_buildMessage = function(alertClass, glyphiconClass, message) {
				var
					result = "<div ";

				if (alertClass.length > 0) {
					result += "class=\"alert " + alertClass + "\"";
				}

				result += ">";

				if (glyphiconClass.length > 0) {
					result += "<span class=\"glyphicon " + glyphiconClass + "\"" + _iconStyle + "></span> ";
				}

				result += message + "</div>";
				return result;
			},

			_destroyDom = function(id) {
				$("#" + id).remove();
			},

			_getConfig = function(config, title, buttons) {
				config = $.extend({
					message   : "",

					title     : title,
					width     : 350,
					height    : 200,
					modal     : true,
					resizable : false,
					autoOpen  : true,
					buttons: buttons
				}, config);

				return config;
			},

			modal = {
				error: function(config) {
					var
						id = WidgetTools.generateId("widgets-dialog-modal-error-");

					config = _getConfig(config, "Error", [
						{
							text : "OK",
							click: function() { $(this).dialog("close"); _destroyDom(id); }
						}
					]);

					_buildDom(id, _buildMessage("alert-danger", "glyphicon-remove-sign", config.message), config);
				},

				information: function(config) {
					var
						id = WidgetTools.generateId("widgets-dialog-modal-information-");

					config = _getConfig(config, "Information", [
						{
							text : "OK",
							click: function() { $(this).dialog("close"); _destroyDom(id); }
						}
					]);

					_buildDom(id, _buildMessage("alert-info", "glyphicon-info-sign", config.message), config);
				},

				yesNo: function(config) {
					var
						id = WidgetTools.generateId("widgets-dialog-modal-yesno-");

					config = _getConfig(config, "Question", [
						{
							text : "Yes",
							click: function() {
								if (config.yes !== undefined) {
									config.yes();
								}

								$(this).dialog("close");
								_destroyDom(id);
							}
						},
						{
							text : "No",
							click: function() {
								if (config.no !== undefined) {
									config.no();
								}

								$(this).dialog("close");
								_destroyDom(id);
							}
						}
					]);

					_buildDom(id, _buildMessage("", "glyphicon-question-sign", config.message), config);
				}

			};

		return modal;
	}
);