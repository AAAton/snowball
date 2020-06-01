package swedish

import (
	"github.com/aaaton/snowball/snowballword"
)

var step2Suffixes = []string{"dd", "gd", "nn", "dt", "gt", "kt", "tt", "mm"}

// Step 2: Search for one of the following suffixes in R1,
// and if found delete the last letter.
//
func step2(w *snowballword.SnowballWord) bool {

	suffix, suffixRunes := w.FirstSuffix(step2Suffixes)

	// If it is not in R1, do nothing
	if suffix == "" || len(suffixRunes) > len(w.RS)-w.R1start {
		return false
	}
	w.RemoveLastNRunes(1)
	return true
}
