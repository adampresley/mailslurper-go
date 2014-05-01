define(
	[
		"modules/util/Http"
	],
	function(Http) {
		"use strict";

		var
			endpoint = "/mails";

		return {
			get: function() {
				return Http.get(endpoint);
			}
		}
	}
);