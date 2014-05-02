define([], function() {
	"use strict";

	return {
		generateId: function(prefix) {
			return prefix + (new Date().getTime());
		}
	};
});