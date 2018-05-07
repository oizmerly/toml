package toml

import (
	"strings"
	"regexp"
	"errors"
)

type (
	Stanza map[string]string
	Data map[string]Stanza
)

var ( // possible expressions
	stanzaExpr = regexp.MustCompile(`^\s*\[(.*)\]\s*$`)
	keyValueExpr = regexp.MustCompile(`^\s*(\w*)\s*=\s*(.*)\s*$`)
	commentExpr = regexp.MustCompile(`^\s*#.*$`)
)

func Decode(input string) (Data, error)  {
	result := make(Data)
	var stanza Stanza

	for _, line := range strings.Split(input, "\n") {
		if match := stanzaExpr.FindStringSubmatch(line); len(match) > 0 {
			stanza = make(map[string]string)
			result[match[1]] = stanza
		} else if match := keyValueExpr.FindStringSubmatch(line); len(match) > 0 {
			if stanza == nil {
				return nil, errors.New("key-value pair without stanza: " + line)
			}
			stanza[match[1]] = match[2]
		} else if commentExpr.MatchString(line) {
			// a comment - do nothing
		} else {
			errors.New("unknown expression: " + line)
		}
	}

	return result, nil
}

func Encode(data Data) string  {
	var result strings.Builder

	for stanza, content := range data {
		result.WriteString("[" + stanza +"]\n")
		for key, value := range content {
			result.WriteString(key + "=" + value + "\n")
		}
	}

	return result.String()
}
