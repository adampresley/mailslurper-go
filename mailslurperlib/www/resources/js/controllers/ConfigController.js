// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

require(
	[
		"jquery", "modules/util/Logger", "modules/util/Blocker", "Ractive", "widgets/dialog/Modal",
		"services/config/ConfigService",

		/* Templates */
		"text!/resources/templates/config-template.html",

		/* Other non-injected dependencies */
		"layout",
	],
	function($, logger, Blocker, Ractive, Modal, ConfigService, ConfigTemplate) {
		"use strict";

		Blocker.block("Loading config...");

		var
			ractive = new Ractive({
				el: "content",
				partials: { configTemplate: ConfigTemplate },
				template: "{{>configTemplate}}",
				data: {
					config: {
						www: "www/",
						wwwPort: 8080,
						smtpAddress: "0.0.0.0",
						smtpPort: 8000,
						dbEngine: "sqlite",
						dbHost: "",
						dbPort: "",
						dbDatabase: "",
						dbUserName: "",
						dbPassword: ""
					}
				},

				complete: function() {
					$("body").layout({
						north__resizable: false,
						north__closable: false,
						south__resizable: false,
						south__closable: false
					});

					$("#configNav").addClass("active");
				}
			});

		ractive.on({
			save: function(e) {
				var data = e.context;
				data.dbEngine = $("#dbEngine").val();

				ConfigService.save(data.www, data.wwwPort, data.smtpAddress, data.smtpPort, data.dbEngine, data.dbHost, data.dbPort, data.dbDatabase, data.dbUserName, data.dbPassword)
					.done(function() {
						Modal.information({
							message: "Your settings have been saved. <strong>Please note that you must restart MailSlurper for these changes to take effect!</strong>",
							height: 265
						});
					})
					.fail(function() {
						Modal.error({
							message: "There was an error trying to save your settings!"
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
					$("#dbEngine").val(data.dbEngine);
					Blocker.unblock(function() { $("#www").focus(); });
				});
			})
			.fail(function() {
				Modal.error({
					message: "There was an error trying to retrieve your settings!"
				});
			});
	}
);