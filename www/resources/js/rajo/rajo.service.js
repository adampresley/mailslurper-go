/**
 * Class: rajo.service
 * This object gives a JavaScript programmer the ability to define service
 * objects which provides a layer between the JavaScript application
 * and a RESTful service layer. An object created using this class
 * has helper methods which aide in performing common HTTP actions
 * such as GET, POST, DELETE, and PUT. A simple service object that
 * provides a basic GET method might look something like this.
 *
 *    > var ExampleService = RajoService.create({
 *    >    endpoint: "/example",
 *    >
 *    >    getRecord: function(id) {
 *    >        return this.$get([ id ]);
 *    >    }
 *    > });
 *
 * The above code will create a new service object (class) named
 * *ExampleService*. This object will have a function named
 * *getRecord()* which takes an ID, makes a GET request to the
 * endpoint found at the URL "/example", and returns a jQuery AJAX
 * promise object. The resulting GET URL would look like this if
 * the id parameter had a value of 100.
 *
 *    > /example/100
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
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
define(["rajo.util", "jquery"], function(RajoUtil, $) {
	"use strict";

	var RajoService = {
		/**
		 * Function: $paramsToSlashes
		 * This function will turn a JavaScript object or array into a slash-delimited
		 * URL string. This is useful for RESTful style interfaces where clean/pretty
		 * URLs are used.
		 *
		 * If params is an array, simply separate each item by slash. If
		 * it is an object separate each key/value pair with slahes, as well
		 * as slashes between the key and value.
		 *
		 * For example,
		 *    > { name: "bob", age: 30 }
		 *
		 * Would turn to
		 *    > /name/bob/age/30
		 *
		 * Parameters:
		 *    params - An object/array of parameters
		 *
		 * Returns:
		 *    A string of the parameters turned into a slash-delimited URL string
		 */
		$paramsToSlashes: function(params) {
			/*
			 */
			var
				result = [];

			if (!params) return "";


			if (params.hasOwnProperty("length")) {
				return "/" + params.join("/");
			} else if (typeof params === "object") {
				RajoUtil.eachKvp(params, function(obj) { result.push(obj.key); result.push(obj.value); });
				return "/" + result.join("/");
			} else {
				return "";
			}
		},

		/**
		 * Function: $delete
		 * Sends a DELETE request to the server endpoint. If this object
		 * is configured to use slashes the parameters sent in are added
		 * to the URL delimited by slashes. For example, the following
		 * object would be converted to a URL as show in the example
		 * below. Otherwise they are send in the request body.
		 *
		 *    > { "id": 1, "name": "Joe" }
		 *    > // URL would be /example/id/1/name/joe
		 *
		 * Parameters:
		 *    params - Object/array of parameters
		 *
		 * Returns:
		 *    A jQuery promise object
		 */
		$delete: function(params) {
			var
				packet = {},
				newUrl = this.endpoint;

			if (this.useSlashWithGetParams) {
				newUrl += RajoService.$paramsToSlashes(params);
			} else if (params) {
				packet = params;
			}

			return $.ajax({
				url: newUrl,
				dataType: "json",
				data: packet,
				type: "DELETE"
			});
		},

		/**
		 * Function: $get
		 * Sends a GET request to the server endpoint. If this object
		 * is configured to use slashes the parameters sent in are added
		 * to the URL delimited by slashes. For example, the following
		 * object would be converted to a URL as shown in the example
		 * below. Otherwise they are appended after a standard question
		 * mark separated by ampersand.
		 *
		 *    > { "id": 1, "name": "Joe" }
		 *    > // URL would be /example/id/1/name/joe
		 *
		 * Parameters:
		 *    params - Object/array of parameters
		 *
		 * Returns:
		 *    A jQuery promise object
		 */
		$get: function(params) {
			var
				packet = {},
				newUrl = this.endpoint;

			if (this.useSlashWithGetParams) {
				newUrl += RajoService.$paramsToSlashes(params);
			} else if (params) {
				packet = params;
			}

			return $.ajax({
				url: newUrl,
				dataType: "json",
				data: packet,
				type: "GET"
			});
		},

		/**
		 * Function: $post
		 * Sends a POST request to the server endpoint. Parameters
		 * are send in the request body.
		 *
		 * Parameters:
		 *    params - Object/array of parameters
		 *
		 * Returns:
		 *    A jQuery promise object
		 */
		$post: function(params) {
			var packet = params || {};

			return $.ajax({
				url: this.endpoint,
				dataType: "json",
				data: packet,
				type: "POST"
			});
		},

		/**
		 * Function: $put
		 * Sends a PUT request to the server endpoint. If this object
		 * is configured to use slashes the parameters sent in are added
		 * to the URL delimited by slashes. For example, the following
		 * object would be converted to a URL as show in the example
		 * below. Otherwise they are send in the request body.
		 *
		 *    > { "id": 1, "name": "Joe" }
		 *    > // URL would be /example/id/1/name/joe
		 *
		 * Parameters:
		 *    params - Object/array of parameters
		 *
		 * Returns:
		 *    A jQuery promise object
		 */
		$put: function(params, bodyParams) {
			var
				packet = bodyParams || {},
				newUrl = this.endpoint;

			if (this.useSlashWithGetParams) {
				newUrl += RajoService.$paramsToSlashes(params);
			} else if (params) {
				packet = params;
			}

			return $.ajax({
				url: newUrl,
				dataType: "json",
				data: packet,
				type: "PUT"
			});
		},

		/**
		 * Function: create
		 * Constructor function which creates a new service object instance. This function
		 * takes a config object in which you must provide an endpoint URL, as well as
		 * any function definitions that your service will provide.
		 *
		 * Example:
		 *
		 *    > var ExampleService = RajoService.create({
		 *    >    endpoint: "/example",
		 *    >
		 *    >    getRecord: function(id) {
		 *    >        return this.$get([ id ]);
		 *    >    }
		 *    > });
		 *
		 * Parameters:
		 *    config - Configuration object
		 *
		 * Returns:
		 *    A new service object
		 */
		create: function(config) {
			var obj = {};

			if (!config.hasOwnProperty("endpoint")) throw "You must specify an endpoint when creating services!";

			obj.$delete = RajoService.$delete;
			obj.$get = RajoService.$get;
			obj.$post = RajoService.$post;
			obj.$put = RajoService.$put;

			obj.useSlashWithGetParams = true;

			RajoUtil.eachKvp(config, function(kvp) { obj[kvp.key] = kvp.value; });
			return obj;
		}
	};

	return RajoService;
});