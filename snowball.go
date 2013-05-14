package snowball

import (
	"fmt"
	"github.com/kljensen/snowball/english"
	"github.com/kljensen/snowball/spanish"
)

const (
	VERSION string = "v0.1.0"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {
	switch language {
	case "english":
		stemmed = english.Stem(word, stemStopWords)
	case "spanish":
		stemmed = spanish.Stem(word, stemStopWords)
	default:
		err = fmt.Errorf("Unknown language: %s", language)
	}
	return

}
