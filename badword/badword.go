package badword

import (
	"io/ioutil"
	"path"
	"regexp"
	"runtime"
	"strings"
)

type BadWords struct {
	*regexp.Regexp
}

func NewBadWords() (*BadWords, error) {
	badWordContent, err := getBadWordFormatter()
	if err != nil {
		return nil, err
	}

	var regexpBadWords *regexp.Regexp

	//regex like (?P<name>(Roberto|Paulo|luis))
	regexpBadWords, err = regexp.Compile("(?P<name>(" + badWordContent + "))")
	if err != nil {
		return nil, err
	}

	return &BadWords{regexpBadWords}, nil
}

func getBadWordFormatter() (string, error) {
	content, err := getFileContent()
	if err != nil {
		return "", err
	}

	content = strings.TrimSuffix(content, "\n")

	return strings.ReplaceAll(content, "\n", "|"), nil
}

func getFileContent() (string, error) {
	_, base, _, _ := runtime.Caller(1)

	pathname := path.Join(path.Dir(base), "../badword/badword.txt")

	content, err := ioutil.ReadFile(pathname)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func (badWords *BadWords) Check(data string) bool {
	return badWords.MatchString(data)
}
