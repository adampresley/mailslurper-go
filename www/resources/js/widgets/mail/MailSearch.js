define(
	[
		"jquery",
		"modules/util/WidgetTools",
		"jqueryui"
	],
	function($, WidgetTools) {
		"use strict";

		var
			_dialogs = {},

			_search = function(dialogEl) {
				var
					dialog = _dialogs[dialogEl],
					searchTermEl = $("#" + dialog.searchTermId);

				dialog.options.search(searchTermEl.val());
				searchTermEl.focus().select();
			};

		$.widget("mailslurper.mailSearch", $.ui.dialog, {
			options: {
				title    : "Search Mails",
				width    : 450,
				height   : 275,
				autoOpen : false,
				modal    : false,
				resizable: false,

				search   : function(term) { alert("You clicked search and searched for " + term); },
				clear    : function() { alert("You clicked clear"); },

				buttons  : [
					{
						text : "Search",
						click: function() { _search(this.id); }
					},
					{
						text : "Clear",
						click: function() {
							_dialogs[this.id].options.clear();
						}
					},
					{
						text : "Close",
						click: function() {
							_dialogs[this.id].dialog.close();
						}
					}
				]
			},
			_create: function() {
				/*
				 * Keep track of opened dialogs
				 */
				_dialogs[this.element.context.id] = {
					dialog: this,
					options: this.options
				};

				/*
				 * Build the DOM for this dialog
				 */
				var
					self = this,
					searchTermId = WidgetTools.generateId("mail-search-term-"),
					html = "<div class=\"alert alert-info\">Enter your search term below and press the Search button.</div>" +
						"<input type=\"text\" id=\"" + searchTermId + "\" class=\"form-control input-lg\" />";

				_dialogs[this.element.context.id].searchTermId = searchTermId;
				this.element.html(html);

				$(this.element).on("keyup", "#" + searchTermId, function(e) {
					if (e.keyCode === 13) {
						_search(self.element.context.id);
					}
				});

				this._super();
			}
		});
	}
);