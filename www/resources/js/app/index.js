require(["./resources/js/config"], function() {
	require(
		[
			"jquery", "app/MailSlurper", "rajo.pubsub", "rajo.ui.bootstrapmodal", "Ractive",
			"app/MailCollection",
			"layout"
		],
		function($, MailSlurper, PubSub, BootstrapModal, Ractive, MailCollection) {
			"use strict";

			PubSub.publish("mailslurper.block", "Loading mails...");
			$("body").layout({
				applyDemoStyles: false,
				north__resizable: false,
				north__closable: false,
				south__resizable: false,
				south__closable: false,
				east__size: "40%"
			});

			var
				mails = [],

				mailListRactive = new Ractive({
					el: "mailList",
					template: "#mailListTemplate",
					data: {
						mails: mails,
						sortColumn: "dateSent",

						compressTo: function(toAddresses) {
							return toAddresses.join("; ");
						},

						sort: function(array, column) {
							console.log("in sort %o", arguments);
							array = array.slice();

							return array.sort(function(a, b) {
								return a[column] < b[column] ? -1 : 1;
							});
						}
					}
				}),

				mailViewRactive = new Ractive({
					el: "mailView",
					template: "#mailViewTemplate",
					data: {
						mailView: "",
						subject: "",
						dateSent: "",
						fromAddress: ""
					}
				}),

				websocketConnection = undefined,
				container = $(".layout"),

				/**
				 * Adds a new mail item to the mails array, which is bound to the interface
				 * and will display the mail item in a table.
				 */
				addMailItemToTable = function(mailItem) {
					mails.unshift(mailItem);
				},

				/**
				 * Sets up a websocket connection to the web server. Hooks up the
				 * close, message, and error events. The *onmessage* event adds
				 * the passed in mail item to our table.
				 */
				setupWebsocket = function() {
					if (window.hasOwnProperty("WebSocket")) {
						websocketConnection = new WebSocket("ws://" + location.host + "/ws");

						websocketConnection.onclose = function(e) { MailSlurper.log("Websocket closed"); websocketConnection = null; }
						websocketConnection.onmessage = function(e) { addMailItemToTable($.parseJSON(e.data)); }
						websocketConnection.onerror = function(e) { MailSlurper.log("An error occurred on the websocket. Closing."); websocketConnection.close(); websocketConnection = null; }
					}
				};

			mailListRactive.on({
				viewMailItem: function(e) {
					mailViewRactive.set("subject", e.context.subject);
					mailViewRactive.set("dateSent", e.context.dateSent);
					mailViewRactive.set("fromAddress", e.context.fromAddress);
					mailViewRactive.set("mailView", e.context.body);

					$(".mailrow").removeClass("highlight-row");
					$(e.node).addClass("highlight-row");
				},

				sort: function(e, column) {
					this.set("sortColumn", column);
				}
			});

			/*
			 * Go get our mail items from the webserver.
			 */
			MailCollection.get().done(function(data) {
				mails = data;
				mailListRactive.set("mails", mails);

				PubSub.publish("mailslurper.unblock");
			});

			setupWebsocket();
		}
	);
});
