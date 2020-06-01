package english

import (
	"github.com/aaaton/snowball/snowballword"
)

var step0suffixes = [][]rune{[]rune("'s'"), []rune("'s"), []rune("'")}

// Step 0 is to strip off apostrophes and "s".
//
func step0(w *snowballword.SnowballWord) bool {
	suffixRunes := w.FirstRuneSuffix(step0suffixes)
	if len(suffixRunes) == 0 {
		return false
	}
	w.RemoveLastNRunes(len(suffixRunes))
	return true
}
