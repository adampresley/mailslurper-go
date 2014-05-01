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
			unblock: function(el) {
				if (el === undefined) {
					$.unblockUI();
				} else {
					$(el).unblockUI();
				}
			}
		};
	}
);