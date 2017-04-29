//
// This is only a SKELETON file for the "Bob" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Bob = function() {};

Bob.prototype.hey = function(input) {
	if (input == input.toUpperCase()) {
		return 'Whoa, chill out!';
	}
	if (input.slice(-1)[0] == '?') {
		return 'Sure.';
	}
	return 'Whatever.';
};

module.exports = Bob;
