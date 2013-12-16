define([ "jquery", "rajo.pubsub", "jquery.blockUI" ], function($, PubSub) {
	"use strict";

	PubSub.subscribe("mailslurper.block", function(msg) {
		$.blockUI({ message: "<h3>" + msg + "</h3>" });
	});

	PubSub.subscribe("mailslurper.unblock", function(fn) { 
		$.unblockUI(); 
		if (fn !== undefined) {
			fn();
		}
	});

	return {
		log: function(msg) {
			if (window.hasOwnProperty("console")) {
				console.log(msg);
			}
		}
	}
});
