/**
 * Class: rajo.pubsub
 * PubSub, is a class which provies the functions necessary for implementing
 * the Publish/Subscribe Pattern (http://en.wikipedia.org/wiki/Publish%E2%80%93subscribe_pattern).
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
 *    * <rajo.util>
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
define(["rajo.util"], function($util) {
	"use strict";

	var pubsub = {
		subscribers: {},

		/**
		 * Function: subscribe
		 * This method registers a handler for a specified event.
		 *
		 * Author:
		 *    Adam Presley
		 *
		 * Parameters:
		 *    eventName - String name of the event to subscribe to
		 *    handler - Function called to handle published events of this type
		 *    scope - The scope in which to call the handler function
		 */
		subscribe: function(eventName, handler, scope) {
			var def = {
				eventName: eventName,
				handler: handler,
				scope: (scope || undefined)
			};

			if (eventName in pubsub.subscribers) {
				pubsub.subscribers[eventName].push(def);
			}
			else {
				pubsub.subscribers[eventName] = [ def ];
			}
		},

		/**
		 * Function: publish
		 * Tells the object to publish an event by name, sending a series of parameters
		 * along with the message for any subscribers to pick up.
		 *
		 * Author:
		 *    Adam Presley
		 *
		 * Parameters:
		 *    eventName - String name of the event to publish
		 *    params - An object of optional parameters to send with the message
		 */
		publish: function(eventName, params) {
			var i = 0;

			params = (params || {});

			if (pubsub.subscribers.hasOwnProperty(eventName)) {
				$util.each(pubsub.subscribers[eventName], function(subscriber) {
					if (subscriber.hasOwnProperty("scope") && subscriber.scope !== undefined) {
						subscriber.handler.call(subscriber.scope, params);
					} else {
						subscriber.handler(params);
					}
				});
			}
		}
	};

	return pubsub;
});