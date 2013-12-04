/**
 * Class: rajo.identity.persona
 * This class provides a simple object for quickly defining the setup for using
 * the Mozilla Persona identity management library. It offers functions for
 * signing in and logging out.
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
 *    * //login.persona.org/include.js
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
define(["rajo.pubsub", "//login.persona.org/include.js"], function(PubSub) {
	"use strict";

	return {
		/**
		 * Function: setup
		 * Initializes the Mozilla Persona service interactions. You provide it a current
		 * email address, if any.
		 *
		 * This method also sets up events to be published when a log in and log out
		 * occurs. These events are published using the <rajo.pubsub> object.
		 *
		 * Events:
		 *    identity.persona.login - When a login attempt occurs. The Persona assertion object is passed to the subscriber
		 *    identity.persona.logout - When a logout attempt occurs.
		 *
		 * Parameters:
		 *    email - Email address. Blank if none provided yet
		 */
		setup: function(email) {
			navigator.id.watch({
				loggedInUser: email,
				onlogin: function(assertion) {
					PubSub.publish("identity.persona.login", assertion);
				},
				onlogout: function() {
					PubSub.publish("identity.persona.logout");
				}
			});
		},

		/**
		 * Function: signIn
		 * Attempts to verify the current email address as logged in with Persona.
		 */
		signIn: function() {
			navigator.id.request();
		},

		/**
		 * Function: signOut
		 * Attempts to sign the current email address out of Persona.
		 */
		signOut: function() {
			navigator.id.logout();
		}
	};
});