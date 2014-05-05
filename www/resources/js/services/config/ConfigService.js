// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

/**
 * Class: ConfigService
 * Defines a service for working with config data
 *
 * Author:
 *     Adam Presley
 *
 * Endpoint:
 *    /config
 *
 * Dependencies:
 *    * <Http>
 */
define(
	[
		"modules/util/Http"
	],
	function(Http) {
		"use strict";

		var endpoint = "/config";

		return {
			get: function() {
				return Http.get(endpoint);
			},

			save: function(www, wwwPort, smtpAddress, smtpPort) {
				return Http.put(endpoint, {
					www: www,
					wwwPort: wwwPort,
					smtpAddress: smtpAddress,
					smtpPort: smtpPort
				});
			}
		};
	}
);