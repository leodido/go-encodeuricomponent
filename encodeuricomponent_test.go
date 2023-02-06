package encodeuricomponent

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEncodeURIComponent(t *testing.T) {
	t.Helper()

	type test struct {
		input  string
		output string
	}

	cases := []test{
		// UNESCAPE
		{
			"alphanumeric123",
			"alphanumeric123",
		},
		{
			"ALPHANUMERIC098",
			"ALPHANUMERIC098",
		},
		{
			"asteri*",
			"asteri*",
		},
		{
			"dash-",
			"dash-",
		},
		{
			"dot.",
			"dot.",
		},
		{
			"_underscore",
			"_underscore",
		},
		{
			"~user",
			"~user",
		},
		{
			"single'quote",
			"single'quote",
		},
		{
			"open(",
			"open(",
		},
		{
			"close)",
			"close)",
		},
		{
			"exclamation!",
			"exclamation!",
		},
		// ESCAPE
		{
			"space ",
			"space%20",
		},
		{
			"a+b c",
			"a%2Bb%20c",
		},
		{
			"\"A\" B ± \"",
			"%22A%22%20B%20%C2%B1%20%22",
		},
		{
			"s/l/a/s/h/e/s",
			"s%2Fl%2Fa%2Fs%2Fh%2Fe%2Fs",
		},
		{
			"with:colon",
			"with%3Acolon",
		},
		{
			"money$",
			"money%24",
		},
		{
			"money€",
			"money%E2%82%AC",
		},
		{
			"escape@",
			"escape%40",
		},
		{
			"+/",
			"%2B%2F",
		},
		{
			"힇",
			"%ED%9E%87",
		},
		{
			"ࠑ",
			"%E0%A0%91",
		},
		{
			"힙", // rune(0xD799)
			"%ED%9E%99",
		},
		{
			"🤣见見",
			"%F0%9F%A4%A3%E8%A7%81%E8%A6%8B",
		},
		{
			"Hello World!@#$%^&*()_+-={}[]\\|;':\",.<>/?`~",
			"Hello%20World!%40%23%24%25%5E%26*()_%2B-%3D%7B%7D%5B%5D%5C%7C%3B'%3A%22%2C.%3C%3E%2F%3F%60~",
		},
		{
			"Hello World áéíóú ü",
			"Hello%20World%20%C3%A1%C3%A9%C3%AD%C3%B3%C3%BA%20%C3%BC",
		},
		{
			"日a本b語ç日ð本Ê語þ日¥本¼語i日©",
			"%E6%97%A5a%E6%9C%ACb%E8%AA%9E%C3%A7%E6%97%A5%C3%B0%E6%9C%AC%C3%8A%E8%AA%9E%C3%BE%E6%97%A5%C2%A5%E6%9C%AC%C2%BC%E8%AA%9Ei%E6%97%A5%C2%A9",
		},
		{
			"💩",
			"%F0%9F%92%A9",
		},
		// FIXME > high-low surrogate pair is okay
		// {
		//  string([]rune{0xD800, 0xDFFF}),
		//  "%F0%90%8F%BF",
		// },
		// ERRORS
		// FIXME > lone high/low surrogates return error
		// {
		//  string([]rune{0xD800}),
		//  "",
		//  fmt.Errorf("URI malformed"),
		// },
		// {
		//  string([]rune{DFF}),
		//  "",
		//  fmt.Errorf("URI malformed"),
		// },
	}

	for _, tc := range cases {
		tc := tc

		t.Run(tc.output, func(t *testing.T) {
			output := EncodeURIComponent(tc.input)

			assert.Equal(t, tc.output, output)
		})
	}
}
