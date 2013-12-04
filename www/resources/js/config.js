require.config({
	baseUrl: "resources/js",
	paths: {
		"app": "app",
		"persona": "//login.persona.org/include.js",
		"bootstrap": "jquery/bootstrap",
		"rajo.dom": "rajo/rajo.dom",
		"rajo.identity.persona": "rajo/rajo.identity.persona",
		"rajo.pubsub": "rajo/rajo.pubsub",
		"rajo.service": "rajo/rajo.service",
		"rajo.singlepage": "rajo/rajo.singlepage",
		"rajo.ui.bootstrapmodal": "rajo/rajo.ui.bootstrapmodal",
		"rajo.util": "rajo/rajo.util",
		"jquery.blockUI": "jquery/jquery.blockUI",
		"Ractive": "jquery/Ractive",
		"daterangepicker": "jquery/daterangepicker",
		"moment": "jquery/moment",
		"jlayout.border": "jquery/jlayout.border",
		"jlayout.flexgrid": "jquery/jlayout.flexgrid",
		"jlayout.flow": "jquery/jlayout.flow",
		"jlayout.grid": "jquery/jlayout.grid",
		"jquery.jlayout": "jquery/jquery.jlayout",
		"jquery.sizes": "jquery/jquery.sizes"
	},
	shim: {
		"jquery.blockUI": { deps: ["jquery"] },
		"bootstrap": { deps: ["jquery"] },
		"daterangepicker": { deps: ["moment", "bootstrap"] },
		"jquery.sizes": { deps: ["jquery"] },
		"jquery.jlayout": { deps: ["jquery", "jlayout.border"] },
		"jlayout.border": { deps: ["jquery.sizes"] },
		"jlayout.flexgrid": { deps: ["jquery.sizes"] },
		"jlayout.flow": { deps: ["jquery.sizes"] },
		"jlayout.grid": { deps: ["jquery.sizes"] }
	}
});
