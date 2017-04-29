var Hamming = function() {
	this.error = new Error('DNA strands must be of equal length.');
};

Hamming.prototype.compute = function(str1, str2) {
	if (str1.length != str2.length) {
		throw this.error;
	}
	var diff = 0;
	for (var i = str1.length - 1; i >= 0; i--) {
		if (str1[i] != str2[i]) {
			diff ++;
		}
	}
	return diff;
};

module.exports = Hamming;
