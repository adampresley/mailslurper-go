require(["./resources/js/config"], function() {
	require(
		[
			"jquery", "rajo.pubsub", "rajo.ui.bootstrapmodal", "rajo.singlepage", "Ractive",
			"app/MailCollection",
			"jquery.blockUI", "jquery.jlayout"
		],
		function($, PubSub, BootstrapModal, SinglePage, Ractive, MailCollection) {
			"use strict";

			$.blockUI({ message: "<h3>Loading mails...</h3>" });

			var
				ractive = new Ractive({
					el: "content",
					template: "#template",
					data: {
						mails: [],
						mailView: "",
						subject: "",
						dateSent: "",
						fromAddress: "",

						compressTo: function(toAddresses) {
							return toAddresses.join("; ");
						}
					}
				}),

				websocketConnection = undefined,
				container = $(".layout"),

				/**
				 * Adds a new mail item to the mails array, which is bound to the interface
				 * and will display the mail item in a table.
				 */
				addMailItemToTable = function(mailItem) {
					var data = ractive.get("mails");

					data.unshift(mailItem);
					ractive.set("mails", data);
				},

				/**
				 * Wrapper for console.log
				 */
				log = function(msg) {
					if (window.hasOwnProperty("console")) {
						console.log(msg);
					}
				},

				/**
				 * Resizes the jLayout border layout container. This is called
				 * by the window resize event.
				 */
				relayout = function() {
					container.layout({resize: false});
				},

				/**
				 * Sets up a websocket connection to the web server. Hooks up the
				 * close, message, and error events. The *onmessage* event adds
				 * the passed in mail item to our table.
				 */
				setupWebsocket = function() {
					if (window.hasOwnProperty("WebSocket")) {
						websocketConnection = new WebSocket("ws://" + location.host + "/ws");

						websocketConnection.onclose = function(e) { log("Websocket closed"); websocketConnection = null; }
						websocketConnection.onmessage = function(e) { addMailItemToTable($.parseJSON(e.data)); }
						websocketConnection.onerror = function(e) { log("An error occurred on the websocket. Closing."); websocketConnection.close(); websocketConnection = null; }
					}
				};

			ractive.on({
				viewMailItem: function(e) {
					ractive.set("subject", e.context.subject);
					ractive.set("dateSent", e.context.dateSent);
					ractive.set("fromAddress", e.context.fromAddress);
					ractive.set("mailView", e.context.body);

					$(".mailrow").removeClass("highlight-row");
					$(e.node).addClass("highlight-row");
				}
			});

			/*
			 * Go get our mail items from the webserver.
			 */
			MailCollection.get().done(function(data) {
				ractive.set("mails", data);
				PubSub.publish("unblock");
			});

//			$("#content").css("height", "100%");

			relayout();
			setupWebsocket();

			$(window).resize(relayout);
			$.unblockUI();
		}
	);
});
