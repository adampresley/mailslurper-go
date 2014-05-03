require(
	[
		/* Injected dependencies */
		"jquery", "modules/util/Logger", "modules/util/Blocker",
		"Ractive", "modules/util/FuncTools",

		"services/mail/MailCollectionService",
		"services/mail/MailService",

		/* Templates */
		"text!/resources/templates/mail-list.html",
		"text!/resources/templates/mail-view.html",

		/* Other non-injected dependencies */
		"layout"
	],
	function($, logger, Blocker, Ractive, FuncTools, MailCollectionService, MailService, MailListPartial, MailViewPartial) {
		"use strict";

		Blocker.block("Loading mails...");

		$("#homeNav").addClass("active");
		$("#configNav").removeClass("active");

		$("body").layout({
			north__resizable: false,
			north__closable: false,
			south__resizable: false,
			south__closable: false,
			east__size: "40%"
		});

		var
			mails = [],

			/*
			 * Ractive instance to handle the list of mail items
			 */
			mailListRactive = new Ractive({
				el: "mailList",
				template: "{{>mailList}}",
				partials: {
					mailList: MailListPartial
				},
				data: {
					mails: mails,
					sortColumn: "dateSent",
					sortDirection: "desc",

					compressTo: function(toAddresses) {
						return toAddresses.join("; ");
					},

					/*
					 * Called when clicking on a header column to sort.
					 * This method will sort the array of data based on a passed
					 * in column and current sort order.
					 *
					 * There is an event attached to this Ractive instance
					 * that will swap the current sort direction.
					 */
					sort: function(array, column) {
						var
							dir = this.get("sortDirection"),
							result1 = (dir === "desc") ? 1 : -1,
							result2 = (dir === "desc") ? -1 : 1;

						array = array.slice();
						logger("Sorting mails on %s:%s", column, dir);

						return array.sort(function(a, b) {
							return a[column] < b[column] ? result1 : result2;
						});
					},

					/*
					 * Returns the correct CSS classes for a column
					 * based on if it is the current sort column and
					 * what the direction is.
					 */
					getSortIcon: function(column) {
						var
							sc = this.get("sortColumn"),
							sd = this.get("sortDirection");

						if (sc !== column) {
							return "";
						} else {
							if (sd === "desc") {
								return "glyphicon glyphicon-arrow-down";
							} else {
								return "glyphicon glyphicon-arrow-up";
							}
						}
					}
				}
			}),

			/*
			 * Ractive to handle viewing a single mail item's details
			 */
			mailViewRactive = new Ractive({
				el: "mailView",
				template: "{{>mailView}}",
				partials: {
					mailView: MailViewPartial
				},
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
				mails.unshift(MailService.parseMailItem(mailItem));
			},

			/**
			 * Sets up a websocket connection to the web server. Hooks up the
			 * close, message, and error events. The *onmessage* event adds
			 * the passed in mail item to our table.
			 */
			setupWebsocket = function() {
				if (window.hasOwnProperty("WebSocket")) {
					websocketConnection = new WebSocket("ws://" + location.host + "/ws");

					websocketConnection.onclose = function(e) { logger("Websocket closed"); websocketConnection = null; }
					websocketConnection.onmessage = function(e) { addMailItemToTable($.parseJSON(e.data)); }
					websocketConnection.onerror = function(e) { logger("An error occurred on the websocket. Closing."); websocketConnection.close(); websocketConnection = null; }
				}
			};

		mailListRactive.on({
			viewMailItem: function(e) {
				mailViewRactive.set("subject", e.context.subject);
				mailViewRactive.set("dateSent", MailService.formatMailDate(e.context.dateSent));
				mailViewRactive.set("fromAddress", e.context.fromAddress);
				mailViewRactive.set("mailView", e.context.body);

				$(".mailrow").removeClass("highlight-row");
				$(e.node).addClass("highlight-row");
			},

			sort: function(e, column) {
				if (this.get("sortColumn") === column) {
					this.set("sortDirection", (this.get("sortDirection") === "desc") ? "asc" : "desc");
				} else {
					this.set("sortDirection", "desc");
				}

				this.set("sortColumn", column);
			}
		});

		/*
		 * Go get our mail items from the webserver.
		 */
		MailCollectionService.get().done(function(data) {
			mails = data;
			mails = FuncTools.map(mails, MailService.parseMailItem);

			mailListRactive.set("mails", mails);
			Blocker.unblock();
		});

		setupWebsocket();
	}
);