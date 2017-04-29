var Year = function(year) {
	this.year = year;
};

Year.prototype.isLeap = function() {
	if (this.year % 4) {
		return false;
	}
	if (this.year % 400 === 0) {
		return true;
	}
	if (this.year % 100 === 0) {
		return false;
	}
	return true;
};

module.exports = Year;
