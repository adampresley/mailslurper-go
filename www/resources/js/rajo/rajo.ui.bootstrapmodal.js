/*
 * Class: rajo.ui.bootstrapmodal
 * Provides the ability to create arbitrary Bootstrap Modal dialogs
 * without having the supporting HTML markup. These functions will
 * generate the correct markup and handle display and teardown.
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
 *    * jquery
 *    * bootstrap
 *
 * License (BSD 2-Clause):
 *    > Copyright 2013 Adam Presley. All rights reserved.
 *    >
 *    > Redistribution and use in source and binary forms, with or without
 *    > modification, are permitted provided that the following conditions are met:
 *    >
 *    > 1. Redistributions of source code must retain the above copyright notice, this
 *    >    list of conditions and the following disclaimer.
 *    >
 *    > 2. Redistributions in binary form must reproduce the above copyright notice,
 *    >    this list of conditions and the following disclaimer in the documentation
 *    >    and/or other materials provided with the distribution.
 *    >
 *    > THIS SOFTWARE IS PROVIDED BY Adam Presley "AS IS" AND ANY EXPRESS OR IMPLIED
 *    > WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
 *    > MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO
 *    > EVENT SHALL Adam Presley OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT,
 *    > INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 *    > LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
 *    > PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
 *    > LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE
 *    > OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF
 *    > ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

define(["jquery", "bootstrap"], function($) {
	"use strict";

	var RajoUiBootstrapModal = {
		dialogInformationImage: "/resources/images/dialog-information.png",
		dialogErrorImage: "/resources/images/dialog-error.png"
	};

	/**
	 * Class: rajo.ui.bootstrapmodal.Modal
	 * Wrapper for the Twitter Bootstrap modal plugin. This object
	 * generates all the necessary markup for you so you don't have to.
	 *
	 * Author:
	 *    Adam Presley
	 *
	 * Constructor:
	 * Takes a config object.
	 *
	 *    * id - ID of the generated top-level DOM element. Defaults to bsp-modal-{datetimestamp}
	 *    * header - Title of the modal dialog
	 *    * headerStyle - Additional style information for the header DIV
	 *    * body - Contents of the body DIV for this dialog
	 *    * bodyStyle - Additional style information for the body DIV
	 *    * buttons - Object where each key is the button text. The value is an object with information about the button. See below
	 *    * keyboard - True/false to allow keyboard shortcuts
	 *    * backdrop - True/false to show a backdrop when this dialog is displayed
	 *    * show - True/false to show on instantiation
	 *    * animate - True/false to animate the display of this dialog
	 *
	 * A button config takes an object with up to two keys.
	 *    * handler - Function to handle the click of this button. Required
	 *    * type - "Primary" to make this a primary styled button.
	 *
	 * Example:
	 *    > require(["rajo.ui.bootstrapmodal"], function(RajoModal) {
	 *    >    var modal = new RajoModal.Modal({
	 *    >       header: "Example Dialog",
	 *    >       body: "<p>This is a modal dialog box.</p>",
	 *    >       buttons: {
	 *    >          "Thanks, Got It": {
	 *    >             type: "primary",
	 *    >             handler: { window.location = '/'; }
	 *    >          }
	 *    >       }
	 *    >    });
	 */
	RajoUiBootstrapModal.Modal = function(config) {
		/**
		 * Function: getId
		 * Return the ID assigned to this dialog
		 */
		this.getId = function() {
			return __config.id;
		};

		/**
		 * Function: close
		 * Closes this dialog by calling the "hide" Bootstrap method
		 */
		this.close = function() {
			__$modalDiv.modal("hide");
		};

		/**
		 * Function: show
		 * Shows this dialog by called the "show" Bootstrap method
		 */
		this.show = function() {
			__$modalDiv.modal("show");
		};

		/**
		 * Function: toggle
		 * Toggles the visibility of this dialog
		 */
		this.toggle = function() {
			__$modalDiv.modal("toggle");
		};

		var
			__init = function() {
				var
					$dialog,
					$content,
					$header,
					$closeAnchor,
					$headerDiv,
					$body,
					$footer,

					item,
					cls;

				__$modalDiv = $("<div />").attr({
					"class": "modal " + ((__config.animate) ? " fade" : ""),
					id: __config.id
				});

				/*
				 * Dialog
				 */
				$dialog = $("<div />").attr({
					"class": "modal-dialog"
				});

				/*
				 * Content div
				 */
				$content = $("<div />").attr({
					"class": "modal-content"
				});

				/*
				 * Header
				 */
				$closeAnchor = $("<button />").attr({
					"type": "button",
					"class": "close",
					"data-dismiss": "modal",
					"aria-hidden": true
				}).html("&times;");

				$header = $("<h4 />").html(__config.header);

				$headerDiv = $("<div />").attr({ "class": "modal-header" });
				if (__config.headerStyle.length > 0) $headerDiv.attr({ "style": __config.headerStyle });
				$headerDiv.append($closeAnchor);
				$headerDiv.append($header);

				/*
				 * Body
				 */
				$body = $("<div />").attr({ "class": "modal-body" }).html(__config.body);
				if (__config.bodyStyle.length > 0) $body.attr({ "style": __config.bodyStyle });

				/*
				 * Footer
				 */
				$footer = $("<div />").attr({ "class": "modal-footer" });

				for (item in __config.buttons) {
					cls = "btn" + ((__config.buttons[item].type && __config.buttons[item].type === "primary") ? " btn-primary" : "");
					var $btn = $("<button />").attr({ type: "button", href: "#", "class": cls, id: __generateId("item"), name: __generateId("item") });
					var txt = "";

					if (__config.buttons[item].icon) {
						txt += "<span class=\"glyphicon glyphicon-" + __config.buttons[item].icon + "\"></span> ";
					}

					txt += item;
					$btn.html(txt).on("click", __config.buttons[item].handler).appendTo($footer);
				}

				/*
				 * Craft the final modal
				 */
				$content.append($headerDiv);
				$content.append($body);
				$content.append($footer);
				$dialog.append($content);
				__$modalDiv.append($dialog);

				$("body").append(__$modalDiv);

				__$modalDiv.modal({ keyboard: __config.keyboard, backdrop: __config.backdrop, show: __config.show });
				__$modalDiv.on("hidden", __destroy);

				__attachEvents();
			},

			__attachEvents = function() {
				var makeHiddenHandler = function(e) {
					return function() {
						__destroy();
						__config.events[e]();
					};
				};

				for (var e in __config.events) {
					if (e === "hidden") {
						__$modalDiv.on(e, makeHiddenHandler(e));
					}
					else {
						__$modalDiv.on(e, __config.events[e]);
					}
				}
			},

			__destroy = function() {
				__$modalDiv.remove();
			},

			__generateId = function(prefix) {
				return prefix + (new Date().getTime());
			},

			__this = this,
			__config = $.extend({
				id: __generateId("bsp-modal-"),
				header: "Header",
				headerStyle: "",
				body: "",
				bodyStyle: "",
				buttons: {
					"Close": { type: "primary", handler: function(target) { __this.close(); } }
				},

				keyboard: true,
				backdrop: true,
				show: true,
				animate: true,

				events: {}
			}, config),

			__$modalDiv;

		__init();
	};

	/**
	 * Class: rajo.ui.bootstrapmodal.Modal.YesNo
	 * A specialized version of the <rajo.ui.bootstrapmodal.Modal> class which offers
	 * a pre-built dialog with Yes and No buttons.
	 *
	 * For config information see <rajo.ui.bootstrapmodal.Modal>
	 *
	 * Author:
	 *    Adam Presley
	 */
	RajoUiBootstrapModal.Modal.YesNo = function(config) {
		var
			__this = this,
			__config = $.extend({
				header: config.header || "Header",
				body: config.body || "Are you sure?",
				buttons: {
					"Yes": {
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "yes", target);
								} else {
									config.handler("yes", target);
								}
							}

							__this.modal.close();
						},
						icon: "thumbs-up"
					},
					"No": {
						type: "primary",
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "no", target);
								} else {
									config.handler("no", target);
								}
							}

							__this.modal.close();
						},
						icon: "thumbs-down"
					}
				},
				show: config.show || true
			}, config);

		this.modal = new RajoUiBootstrapModal.Modal(__config);
	};

	/**
	 * Class: rajo.ui.bootstrapmodal.Modal.OK
	 * A specialized version of the <rajo.ui.bootstrapmodal.Modal> class which offers
	 * a pre-built dialog with an OK button.
	 *
	 * For config information see <rajo.ui.bootstrapmodal.Modal>
	 *
	 * Author:
	 *    Adam Presley
	 */
	RajoUiBootstrapModal.Modal.OK = function(config) {
		var
			__this = this,
			__config = $.extend({
				header: config.header || "Confirmation",
				body: config.body || "OK",
				buttons: {
					"OK": {
						type: "primary",
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "ok", target);
								} else {
									config.handler("ok", target);
								}
							}

							__this.modal.close();
						}
					}
				},
				show: config.show || true
			}, config);

		this.modal = new RajoUiBootstrapModal.Modal(__config);
	};

	/**
	 * Class: rajo.ui.bootstrapmodal.Modal.OkCancel
	 * A specialized version of the <rajo.ui.bootstrapmodal.Modal> class which offers
	 * a pre-built dialog with OK and Cancel buttons.
	 *
	 * For config information see <rajo.ui.bootstrapmodal.Modal>
	 *
	 * Author:
	 *    Adam Presley
	 */
	RajoUiBootstrapModal.Modal.OkCancel = function(config) {
		var
			__this = this,
			__config = $.extend({
				header: config.header || "Confirmation",
				body: config.body || "OK",
				buttons: {
					"OK": {
						type: "primary",
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "ok", target);
								} else {
									config.handler("ok", target);
								}
							}

							__this.modal.close();
						}
					},
					"Cancel": {
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "cancel", target);
								} else {
									config.handler("cancel", target);
								}
							}

							__this.modal.close();
						}
					}
				},
				show: config.show || true
			}, config);

		this.modal = new RajoUiBootstrapModal.Modal(__config);
	};

	/**
	 * Class: rajo.ui.bootstrapmodal.Modal.YesNo
	 * A specialized version of the <rajo.ui.bootstrapmodal.Modal> class which offers
	 * a pre-built dialog with an OK button and an error icon.
	 *
	 * For config information see <rajo.ui.bootstrapmodal.Modal>
	 *
	 * Author:
	 *    Adam Presley
	 */
	RajoUiBootstrapModal.Modal.Error = function(config) {
		var
			__this = this,
			__config = $.extend({
				header: config.header || "Error",
				body: config.body || "Error",
				buttons: {
					"OK": {
						type: "primary",
						handler: function(target) {
							if (config.handler) {
								if (config.hasOwnProperty("scope")) {
									config.handler.call(config.scope, "ok", target);
								} else {
									config.handler("ok", target);
								}
							}

							__this.modal.close();
						}
					}
				},
				show: config.show || true
			}, config);

		__config.body = "<img src=\"" + RajoUiBootstrapModal.dialogErrorImage + "\" style=\"float: left; margin-right: 10px\" />" + __config.body;
		this.modal = new RajoUiBootstrapModal.Modal(__config);
	};

	return RajoUiBootstrapModal;
});