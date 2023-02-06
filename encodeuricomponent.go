package encodeuricomponent

import (
	"net/url"
	"strings"
)

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
