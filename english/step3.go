package english

import (
	"github.com/aaaton/snowball/snowballword"
)

var step3Suffixes = [][]rune{[]rune("ational"), []rune("tional"), []rune("alize"), []rune("icate"), []rune("ative"),
	[]rune("iciti"), []rune("ical"), []rune("ful"), []rune("ness")}

// Step 3 is the stemming of various longer sufficies
// found in R1.
//
func step3(w *snowballword.SnowballWord) bool {

	suffixRunes := w.FirstRuneSuffix(step3Suffixes)

	// If it is not in R1, do nothing
	if len(suffixRunes) == 0 || len(suffixRunes) > len(w.RS)-w.R1start {
		return false
	}
	suffix := string(suffixRunes)

	// Handle special cases where we're not just going to
	// replace the suffix with another suffix: there are
	// other things we need to do.
	//
	if suffix == "ative" {

		// If in R2, delete.
		//
		if len(w.RS)-w.R2start >= 5 {
			w.RemoveLastNRunes(len(suffixRunes))
			return true
		}
		return false
	}

	// Handle a suffix that was found, which is going
	// to be replaced with a different suffix.
	//
	var repl string
	switch suffix {
	case "ational":
		repl = "ate"
	case "tional":
		repl = "tion"
	case "alize":
		repl = "al"
	case "icate", "iciti", "ical":
		repl = "ic"
	case "ful", "ness":
		repl = ""
	}
	w.ReplaceSuffixRunes(suffixRunes, []rune(repl), true)
	return true

}
