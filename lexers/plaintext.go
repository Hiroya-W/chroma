package lexers

import (
	. "github.com/Hiroya-W/chroma/v2" // nolint
)

var Plaintext = Register(MustNewLexer(
	&Config{
		Name:      "plaintext",
		Aliases:   []string{"text", "plain", "no-highlight"},
		Filenames: []string{"*.txt"},
		MimeTypes: []string{"text/plain"},
		Priority:  -1,
	},
	PlaintextRules,
))
