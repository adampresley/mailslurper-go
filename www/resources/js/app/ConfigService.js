// Copyright 2013 Adam Presley. All rights reserved
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
 *    * <rajo.service>
 */
define(["rajo.service"], function(Service) {
	return Service.create({
		endpoint: "/config",

		/**
		 * Function: get
		 * Retrieves configuration settings
		 *
		 * Returns:
		 *     A jQuery AJAX promise object
		 */
		get: function() {
			return this.$get();
		},

		/**
		 * Function: save
		 * Saves configuration settings to the server
		 *
		 * Returns:
		 *    A jQuery AJAX promise object
		 */
		save: function(www, wwwPort, smtpAddress, smtpPort) {
			return this.$put(null, {
				www: www,
				wwwPort: wwwPort,
				smtpAddress: smtpAddress,
				smtpPort: smtpPort
			});
		}
	});
});