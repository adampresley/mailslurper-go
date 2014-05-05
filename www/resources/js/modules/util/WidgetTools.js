// Copyright 2013-2014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

define([], function() {
	"use strict";

	return {
		generateId: function(prefix) {
			return prefix + (new Date().getTime());
		}
	};
});