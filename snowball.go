package snowball

import (
	"fmt"


	"github.com/aaaton/snowball/swedish"
	"github.com/aaaton/snowball/english"
	"github.com/aaaton/snowball/french"
	"github.com/aaaton/snowball/russian"
	"github.com/aaaton/snowball/spanish"
)

const (
	VERSION string = "v0.5.0"
)

// Stem a word in the specified language.
//
func Stem(word, language string, stemStopWords bool) (stemmed string, err error) {

	var f func(string, bool) string
	switch language {
	case "english":
		f = english.Stem
	case "spanish":
		f = spanish.Stem
	case "french":
		f = french.Stem
	case "russian":
		f = russian.Stem
	case "swedish":
		f = swedish.Stem
	default:
		err = fmt.Errorf("Unknown language: %s", language)
		return
	}
	stemmed = f(word, stemStopWords)
	return

}
