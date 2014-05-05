// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

/**
 * File: Logger.js
 * Provides a utility module for logging to the console. This wrapper
 * will make sure the console exists in the Window object prior to
 * making the call.
 */
define(
	[],
	function() {
		"use strict";

		return function() {
			if (window.hasOwnProperty("console")) {
				Function.prototype.apply.call(console.log, console, arguments)
			}
		}
	}
);