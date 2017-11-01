package swedish

import "github.com/aaaton/snowball/snowballword"

// Step 1a is the deleting of the s-ending
//
func step1a(w *snowballword.SnowballWord) bool {

	// Possible sufficies for this step, longest first.
	suffixes := []string{"s"}

	// Using FirstSuffixIn since there are overlapping suffixes, where some might not be in the R1,
	// while another might. For example: "ärade"
	suffix, suffixRunes := w.FirstSuffixIn(w.R1start, len(w.RS), suffixes...)

	// If it is not in R1, do nothing
	if suffix == "" || len(suffixRunes) > len(w.RS)-w.R1start {
		return false
	}

	// Delete if preceded by a valid s-ending. Valid s-endings inlude the
	// following charaters: bcdfghjklmnoprtvy.
	//
	rsLen := len(w.RS)
	if rsLen >= 2 {
		switch w.RS[rsLen-2] {
		case 'b', 'c', 'd', 'f', 'g', 'h', 'j', 'k',
			'l', 'm', 'n', 'o', 'p', 'r', 't', 'v', 'y':
			w.RemoveLastNRunes(len(suffixRunes))
			return true
		}
	}
	return false
}
