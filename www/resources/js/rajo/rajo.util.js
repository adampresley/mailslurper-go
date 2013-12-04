/**
 * Class: rajo.util
 * This class provides handy utility methods for looping, number
 * ranges, and more.
 *
 * This class is a part of the RAJO, or Random Assortment of JavaScript Objects
 * library.
 *
 * Author:
 *    Adam Presley
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
define([], function() {
	"use strict";

	var $util = {
		/**
		 * Function: each
		 * Provides a functional-style looping function which can iterate over objects (maps)
		 * and arrays alike. For an array the provided function is called with the current item
		 * as the single provided argument. For an object each key is iterated over and the value
		 * is the argument passed to the function.
		 *
		 * Parameters:
		 *    items - Array or object to iterate over
		 *    fn - Function to be called for each item
		 */
		each: function(items, fn) {
			var i = 0;

			if (items.hasOwnProperty("length")) {
				for (i = 0; i < items.length; i++) {
					fn(items[i]);
				}
			} else if (typeof items === "object") {
				for (i in items) {
					if (items.hasOwnProperty(i)) {
						fn(items[i]);
					}
				}
			}
		},

		/**
		 * Function: eachKvp
		 * Similar to <each> but more specific to iterating over an object. The function
		 * passed in is called passing in an object with two keys, *key* and *value*, populated
		 * from the input object.
		 *
		 * Parameters:
		 *    obj - Object to iterated over
		 *    fn - Function to be called for each item
		 */
		eachKvp: function(obj, fn) {
			var i = 0;

			for (i in obj) {
				if (obj.hasOwnProperty(i)) {
					fn({ key: i, value: obj[i] });
				}
			}
		},

		/**
		 * Function: map
		 * Applies a function to each item in the input array or object. This function may return
		 * the input transformed in some way, and the final result is an array of each transformed
		 * item.
		 *
		 * Parameters:
		 *    items - Array/object of items
		 *    fn - Function to apply to each item
		 *
		 * Returns:
		 *    An array of transformed items
		 */
		map: function(items, fn) {
			var result = [];

			$util.each(function(item) {
				result.push(fn(item));
			}, items);

			return result;
		},

		/**
		 * Function: mapArrayToObject
		 * Iterates over an array calling a key/value pair function for each item.
		 * This function should return an object with two keys, *key* and *value*.
		 * The result is an object of those keys matched up with their values.
		 *
		 * Parameters:
		 *    items - Array of items
		 *    kvpFn - Key/value pair function returning an object with key/value pair
		 *
		 * Returns:
		 *    Object containing *key* and *value*
		 *
		 * Example:
		 *    > var domItemsAndValues = RojoUtil.mapArrayToObject(
		 *    >    ["domItem1", "domeItem2"],
		 *    >    function(item) { return { key: item, value: document.getElementById(item).value }; }
		 *    > );
		 */
		mapArrayToObject: function(items, kvpFn) {
			var
				result = {},
				kvp = [];

			$util.each(function(item) {
				kvp = kvpFn(item);
				result[kvp.key] = kvp.value;
			}, items);

			return result;
		},

		/**
		 * Function: range
		 * This function takes 1 to 3 arguments that returns an array of
		 * sequential numbers.
		 *
		 * Parameters:
		 *    start/end - First number in the sequence, or total number if this is the only argument
		 *    end - Last number in the sequence
		 *    step - Increment by
		 *
		 * Returns:
		 *    An array of numbers in the sequence described by input arguments
		 *
		 * Example:
		 *    > var range1 = RojoUtil.range(10); // [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
		 *    > var range2 = RojoUtil.range(1, 5); // [1, 2, 3, 4]
		 *    > var range3 = RojoUtil.range(0, 10, 2); // [0, 2, 4, 6, 8]
		 */
		range: function() {
			var
				result = [],
				start = 0,
				end = 0,
				step = 1,
				i = 0;

			if (arguments.length > 0) {
				if (arguments.length === 1) {
					end = arguments[0];
				} else if (arguments.length === 2) {
					start = arguments[0];
					end = arguments[1];
				} else if (arguments.length > 2) {
					start = arguments[0];
					end = arguments[1];
					step = arguments[2];
				}

				for (i = 0; i < end; i += step) {
					result.push(i);
				}
			}

			return result;
		},

		/**
		 * Function: reduce
		 * Reduces a set of items from a starting point to a single result by
		 * applying the function *fn()* which takes the current starting point
		 * and the next item in each iteration.
		 *
		 * Parameters:
		 *    start - Starting item
		 *    items - Set of items
		 *    fn - Function to apply to each pair of items, returning the combined, reduced value
		 *
		 * Returns:
		 *    A final, reduced value
		 */
		reduce: function(start, items, fn) {
			var base = start;

			$util.eachKvp(items, function(item) { base = fn(base, item); });
			return base;
		}
	};

	return $util;
});