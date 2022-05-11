package lexers

import (
	. "github.com/alecthomas/chroma/v2" // nolint
)

var Salt = Register(DelegatingLexer(Get("YAML"), MustNewLexer(
	&Config{
		Name:      "Salt",
		Aliases:   []string{"sls", "salt"},
		Filenames: []string{"*.sls"},
		MimeTypes: []string{"text/x-yaml+jinja", "text/x-sls"},
		DotAll:    true,
	},
	func() Rules {
		return Get("Django/Jinja").(*RegexLexer).MustRules()
	},
)))
