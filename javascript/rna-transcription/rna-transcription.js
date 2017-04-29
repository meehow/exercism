var DnaTranscriber = function() {
	this.map = {
		G: 'C',
		C: 'G',
		T: 'A',
		A: 'U',
	};
};

DnaTranscriber.prototype.toRna = function(dna) {
	var rna = [];
	for (var nuc of dna) {
		rna.push(this.map[nuc]);
	}
	return rna.join('');
};

module.exports = DnaTranscriber;
