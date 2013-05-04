/*
	This package implements various string-like methods
	over slices of runes.  It is similar to the
	exp/utf8string package.
*/
package runeslice

// Word represents a word that is going to be stemmed.
// 
type Word struct {

	// A slice of runes
	rs []rune

	// The index in rs where the R1 region begins
	r1start int

	// The index in rs where the R2 region begins
	r2start int
}

// Create a new Word struct
func New(in string) (word *Word) {
	word = &Word{rs: []rune(in)}
	word.r1start = len(word.rs)
	word.r2start = len(word.rs)
	return
}

// Resets r1start and r2start to ensure they 
// are within bounds of the current rune slice.
func (w *Word) resetR1R2() {
	rsLen := len(w.rs)
	if w.r1start > rsLen {
		w.r1start = rsLen
	}
	if w.r2start > rsLen {
		w.r2start = rsLen
	}
}

// Return a slice of w.rs, allowing the start
// and stop to be out of bounds.
//
func (w *Word) slice(start, stop int) []rune {
	startMin := 0
	if start < startMin {
		start = startMin
	}
	max := len(w.rs) - 1
	if start > max {
		start = max
	}
	if stop > max {
		stop = max
	}
	return w.rs[start:stop]
}

// Return the R1 region as a slice of runes
func (w *Word) R1() []rune {
	return w.rs[w.r1start:]
}

// Return the R1 region as a string
func (w *Word) R1String() string {
	return string(w.R1())
}

// Return the R2 region as a slice of runes
func (w *Word) R2() []rune {
	return w.rs[w.r2start:]
}

// Return the R2 region as a string
func (w *Word) R2String() string {
	return string(w.R2())
}

// Return the Word as a string
func (w *Word) String() string {
	return string(w.rs)
}

// Return the first prefix found or the empty string.
func (w *Word) FirstPrefix(prefixes ...string) string {
	found := false
	rsLen := len(w.rs)

	for _, prefix := range prefixes {
		found = true
		for i, r := range prefix {
			if i > rsLen-1 || (w.rs)[i] != r {
				found = false
				break
			}
		}
		if found {
			return prefix
		}
	}
	return ""
}

// Return the first suffix found or the empty string.
func (w *Word) FirstSuffix(sufficies ...string) (suffix string) {
	rsLen := len(w.rs)
	for _, suffix := range sufficies {
		numMatching := 0
		suffixRunes := []rune(suffix)
		suffixLen := len(suffixRunes)
		for i := 0; i < rsLen && i < suffixLen; i++ {
			if w.rs[rsLen-i-1] != suffixRunes[suffixLen-i-1] {
				break
			} else {
				numMatching += 1
			}
		}
		if numMatching == suffixLen {
			return suffix
		}
	}
	return ""
}