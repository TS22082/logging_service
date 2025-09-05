package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func FormatDocTitle(subject string) string {

	if subject == "apiKeys" {
		return "Api Keys"
	}

	strCase := cases.Title(language.AmericanEnglish)
	titleCase := strCase.String(subject)

	return titleCase
}
