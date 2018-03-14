package protein

var codonMap = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func stop(codon string) bool {
	switch codon {
	case "UAA", "UAG", "UGA":
		return true
	default:
		return false
	}
}

// FromCodon translates from codon to protein name
func FromCodon(input string) string {
	return codonMap[input]
}

// FromRNA tralslates from RNA to protein sequence
func FromRNA(input string) []string {
	output := make([]string, 0, len(input)/3)
	for i := 0; i < len(input); i += 3 {
		codon := input[i : i+3]
		if stop(codon) {
			break
		}
		output = append(output, FromCodon(codon))
	}
	return output
}
