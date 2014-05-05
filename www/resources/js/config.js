// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

require.config({
	baseUrl: "resources/js",
	paths: {
		"text": "requirejs/text",

		"jqueryui": "jquery/jquery-ui-1.10.4.min",
		"bootstrap": "jquery/bootstrap",
		"jquery.blockUI": "jquery/jquery.blockUI",
		"Ractive": "jquery/Ractive",
		"daterangepicker": "jquery/daterangepicker",
		"moment": "jquery/moment",
		"layout": "jquery/jquery.layout"
	},
	shim: {
		"jqueryui": { deps: ["jquery"] },
		"jquery.blockUI": { deps: ["jquery"] },
		"bootstrap": { deps: ["jquery"] },
		"daterangepicker": { deps: ["moment", "bootstrap"] },
		"layout": { deps: ["jqueryui"] }
	}
});
