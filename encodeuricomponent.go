package encodeuricomponent

import (
	"net/url"
	"strings"
)

// EncodeURIComponent encodes a URI by replacing each instance of
// certain characters by one, two, three, or four escape sequences
// representing the UTF-8 encoding of the character.
//
// It will use 4 escape sequences only for characters composed of 2 surrogate characters.
func EncodeURIComponent(input string) string {
	return strings.ReplaceAll(
		strings.ReplaceAll(
			strings.ReplaceAll(
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ReplaceAll(
							url.QueryEscape(strings.ReplaceAll(input, " ", "%20")),
							"%2520",
							"%20",
						),
						"%27",
						"'",
					),
					"%28",
					"(",
				),
				"%29",
				")",
			),
			"%21",
			"!",
		),
		"%2A",
		"*",
	)
}
