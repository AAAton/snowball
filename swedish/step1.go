package swedish

import (
	"github.com/aaaton/snowball/snowballword"
)

// Possible sufficies for this step, longest first.
var step1suffixes = []string{
	"heterna", "hetens", "anden", "heten", "heter", "arnas",
	"ernas", "ornas", "andes", "arens", "andet", "arna", "erna",
	"orna", "ande", "arne", "aste", "aren", "ades", "erns", "ade",
	"are", "ern", "ens", "het", "ast", "ad", "en", "ar", "er",
	"or", "as", "es", "at", "et", "an", "a", "e",
}

// Step 1 is the stemming of various endings found in
// R1 including "heterna", "ornas", and "andet".
//
func step1b(w *snowballword.SnowballWord) bool {

	// Using FirstSuffixIn since there are overlapping suffixes, where some might not be in the R1,
	// while another might. For example: "ärade"
	suffix, suffixRunes := w.FirstSuffixIn(w.R1start, len(w.RS), step1suffixes)

	// If it is not in R1, do nothing
	if suffix == "" || len(suffixRunes) > len(w.RS)-w.R1start {
		return false
	}

	// Remove the suffix
	w.RemoveLastNRunes(len(suffixRunes))
	return true
}
