// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define(
	[
		"jquery",
		"jquery.blockUI"
	],
	function($) {
		"use strict";

		return {
			/**
			 * Function: block
			 * Blocks the page or an element with a specified message.
			 */
			block: function(message, el) {
				var config = { message: "<h3>" + message + "</h3>" };

				if (el === undefined) {
					$.blockUI(config);
				} else {
					$(el).blockUI(config);
				}
			},

			/**
			 * Function: unblock
			 * Unblocks the page or an element
			 */
			unblock: function(el, fn) {
				if (el === undefined) {
					$.unblockUI({ onUnblock: fn });
				} else {
					if (typeof el === "function") {
						$.unblockUI({ onUnblock: el });
					} else {
						$(el).unblockUI({ onUnblock: fn });
					}
				}
			}
		};
	}
);