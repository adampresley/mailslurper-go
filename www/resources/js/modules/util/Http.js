// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery"
	],
	function($) {
		"use strict";

		return {
			delete: function(url, data) {
				return $.ajax({ url: url, data: data, dataType: "json", type: "DELETE" });
			},

			get: function(url) {
				return $.ajax({ url: url, dataType: "json", type: "GET" });
			},

			post: function(url, data) {
				return $.ajax({ url: url, data: data, dataType: "json", type: "POST" });
			},

			put: function(url, data) {
				return $.ajax({ url: url, data: data, dataType: "json", type: "PUT" });
			}
		}
	}
);