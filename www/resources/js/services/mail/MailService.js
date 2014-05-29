// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery",
		"modules/util/FuncTools",
		"moment"
	],
	function($, FuncTools, moment) {
		"use strict";

		var
			service = {
				formatMailDate: function(dateString) {
					return moment(dateString).format("MMMM Do YYYY, h:mm:ss a");
				},

				isMultipartMail: function(contentType, boundary) {
					return (contentType.indexOf("multipart") > -1 && boundary.length > 0);
				},

				parseHTMLPart: function(body) {
					var
						split = body.split("\n"),
						foundHeaderSplit = false,
						trimmedBody = "";

						FuncTools.each(split, function(item) {
							if (!foundHeaderSplit) {
								if ($.trim(item) === "") {
									foundHeaderSplit = true;
								}
							} else {
								trimmedBody += item;
							}
						});

						return trimmedBody;
				},

				parseMailBody: function(contentType, boundary, body) {
					if (!service.isMultipartMail(contentType, boundary)) {
						return body;
					}

					boundary = "--" + boundary;

					var
						split = body.split(boundary);

					/*
					 * The first item will be blank because the mail body starts
					 * with the boundary marker. The last item will be '--'
					 * because the very last boundary marker has two dashes at the
					 * end. We want the next to last body because the RFC-1341
					 * specifies that the fanciest version of the mail will be last.
					 *
					 * TODO: This logic is flawed. The LAST-LAST items will be attachments
					 */
					if (split.length > 2 && ((split.length - 2) > 0)) {
						return service.parseHTMLPart($.trim(split[split.length - 2]));
					} else {
						return split[0];
					}
				},

				parseMailItem: function(mailItem) {
					mailItem.body = service.parseMailBody(mailItem.contentType, mailItem.boundary, mailItem.body);
					return mailItem;
				}
			};

		return service;
	}
);