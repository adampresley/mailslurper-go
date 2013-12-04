// Copyright 2013 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

/**
 * Class: MailCollection
 * Defines a service for working with collections of mail items.
 *
 * Author:
 *     Adam Presley
 *
 * Endpoint:
 *    /mails
 *
 * Dependencies:
 *    * <rajo.service>
 */
define(["rajo.service"], function(Service) {
	return Service.create({
		endpoint: "/mails",

		/**
		 * Function: get
		 * Retrieves a page of mail items.
		 *
		 * Returns:
		 *     A jQuery AJAX promise object
		 */
		get: function() {
			return this.$get();
		}
	});
});