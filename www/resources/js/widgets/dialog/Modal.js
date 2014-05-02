define(
	[
		"jquery",
		"modules/util/WidgetTools",
		"jqueryui"
	],

	function($, WidgetTools) {
		"use strict";

		var
			_template = "<div id=\"{{id}}\">{{body}}</div>",
			_iconStyle = "style=\"font-size: 24px; margin-right: 8px; margin-bottom: 8px;\"",

			_buildDom = function(id, body, config) {
				var $body = _template.replace("{{id}}", id).replace("{{body}}", body);

				$("body").append($body);
				$("#" + id).dialog(config)
			},

			_buildMessage = function(alertClass, glyphiconClass, message) {
				return "<div class=\"alert " + alertClass + "\"><span class=\"glyphicon " + glyphiconClass + "\"" + _iconStyle + "></span> " + message + "</div>";
			},

			_destroyDom = function(id) {
				$("#" + id).remove();
			},

			_getConfig = function(config, title, buttons) {
				config = $.extend({
					message   : "",

					title     : title,
					width     : 350,
					height    : 200,
					modal     : true,
					resizable : false,
					autoOpen  : true,
					buttons: buttons
				}, config);

				return config;
			},

			modal = {
				error: function(config) {
					var
						id = WidgetTools.generateId("widgets-dialog-modal-error-");

					config = _getConfig(config, "Error", [
						{
							text : "OK",
							click: function() { $(this).dialog("close"); _destroyDom(id); }
						}
					]);

					_buildDom(id, _buildMessage("alert-danger", "glyphicon-remove-sign", config.message), config);
				},

				information: function(config) {
					var
						id = WidgetTools.generateId("widgets-dialog-modal-information-");

					config = _getConfig(config, "Information", [
						{
							text : "OK",
							click: function() { $(this).dialog("close"); _destroyDom(id); }
						}
					]);

					_buildDom(id, _buildMessage("alert-info", "glyphicon-info-sign", config.message), config);
				}

			};

		return modal;
	}
);