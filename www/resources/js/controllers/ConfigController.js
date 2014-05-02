require(
	[
		"jquery", "modules/util/Logger", "modules/util/Blocker", "rajo.pubsub", "rajo.ui.bootstrapmodal", "Ractive",
		"app/ConfigService"
	],
	function($, logger, Blocker, PubSub, BootstrapModal, Ractive, ConfigService) {
		"use strict";

		Blocker.block("Loading config...");

		var
			ractive = new Ractive({
				el: "content",
				template: "#template",
				data: {
					config: {
						www: "www/",
						wwwPort: 8080,
						smtpAddress: "0.0.0.0",
						smtpPort: 8000
					}
				}
			});

		ractive.on({
			save: function(e) {
				var data = e.context;
				ConfigService.save(data.www, data.wwwPort, data.smtpAddress, data.smtpPort)
					.done(function() {
						BootstrapModal.Modal.OK({
							body: "<p>Your settings have been saved.</p> <p class=\"alert alert-warning\"><strong>Please " +
								"note that you must restart MailSlurper for these changes to take effect!</strong></p>"
						});
					})
					.fail(function() {
						BootstrapModal.Modal.Error({
							body: "<p>There was an error trying to save your settings!</p>"
						});
					});
			}
		});

		/*
		 * Go get our configuration settings
		 */
		ConfigService.get()
			.done(function(data) {
				ractive.set("config", data, function() {
					Blocker.unblock(function() { $("#www").focus(); });
				});
			})
			.fail(function() {
				BootstrapModal.Modal.Error({
					body: "<p>There was an error trying to retrieve your settings!</p>"
				});
			});
	}
);