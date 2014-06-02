// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery",
		"modules/util/FuncTools",
		"modules/util/Http",
		"moment"
	],
	function($, FuncTools, Http, moment) {
		"use strict";

		var
			service = {
				formatMailDate: function(dateString) {
					return moment(dateString).format("MMMM Do YYYY, h:mm:ss a");
				},

				getMailItem: function(id) {
					return Http.get("/mail?id=" + id);
				},

				parseMailItem: function(mailItem) {
					mailItem.attachmentIcon = (mailItem.attachmentCount > 0) ? "<span class=\"glyphicon glyphicon-paperclip\"></span>" : "&nbsp;";
					return mailItem;
				}
			};

		return service;
	}
);