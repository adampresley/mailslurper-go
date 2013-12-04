/**
 * Class: rajo.dom
 * Class offering small, quick DOM utilities.
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
 *
 * Dependencies:
 *    * <rajo.util>
 *    * jquery
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
define(["rajo.util", "jquery"], function($util, $) {
	"use strict";

	var $dom = {
		/**
		 * Function: getEl
		 * Retrieves a single DOM element by ID.
		 *
		 * Parameters:
		 *    id - ID of a DOM element
		 *
		 * Returns:
		 *    A matching DOM element or undefined
		 */
		getEl: function(id) {
			return document.getElementById(id);
		},

		/**
		 * Function: getEls
		 * This function takes an array of DOM element IDs (strings) and
		 * returns an array of DOM element nodes.
		 *
		 * Parameters:
		 *    elementArray - Array of DOM element IDs
		 *
		 * Retuns:
		 *    Array of DOM element nodes
		 */
		getEls: function(elementArray) {
			return $util.map(elementArray, $dom.getEl);
		},

		/**
		 * Function: makePacket
		 * Convienence function that takes an array of DOM element IDs and
		 * returns an object whos key is the element's ID and the value
		 * is the element's retrieved value. This is most useful for things
		 * like AJAX POSTs.
		 *
		 * Paramters:
		 *    elements - Array of DOM element IDs
		 *
		 * Returns:
		 *    JavaScript object
		 *
		 * Example:
		 *    > // Give two text boxes with values, this would give you a
		 *    > // structure similar to this.
		 *    > var packet = RojoDom.makePacket(["txt1", "txt2"]);
		 *    >
		 *    > // {
		 *    > //    txt1: "value1",
		 *    > //    txt2: "value2"
		 *    > // }
		 */
		makePacket: function(elements) {
			return $util.mapArrayToObject($dom.getEls(elements), function(el) {
				return {
					key: el.id,
					value: $(el).val()
				};
			});
		}
	};

	return $dom;
});