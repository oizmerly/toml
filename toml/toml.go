package toml

import (
	"regexp"
	"io/ioutil"
	"strings"
	"errors"
)

type (
	Title string // stanzas title
	Key string // data key
	Value string // data value

	Stanza map[Key]Value
	Data map[Title]Stanza
)

func (data Data) SetStanza(title Title, stanza Stanza) {
	data[title] = stanza
}

func (data Data) GetStanza(title Title) (Stanza, bool) {
	stanza, exists := data[title]
	return stanza, exists
}

func (data Data) SetValue(title Title, key Key, value Value) {
	var stanza Stanza
	var exists bool
	if stanza, exists = data[title]; !exists {
		stanza = Stanza{}
		data[title] = stanza
	}
	stanza[key] = value
}

func (data Data) GetValue(title Title, key Key) (Value, bool) {
	var stanza Stanza
	var exists bool
	if stanza, exists = data.GetStanza(title); !exists {
		return "", false
	}
	return stanza[key], true
}

func (data Data) Read(filename string) error {
	var ( // supported expressions
		stanzaTitleExpr = regexp.MustCompile(`^\s*\[(.*)\]\s*$`)
		keyValueExpr = regexp.MustCompile(`^\s*(\w*)\s*=\s*(.*)\s*$`)
		commentExpr = regexp.MustCompile(`^\s*#.*$`)
	)
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	// read data from file
	var stanza Stanza
	for _, line := range strings.Split(string(bytes), "\n") {
		if match := stanzaTitleExpr.FindStringSubmatch(line); len(match) > 0 {
			stanza = Stanza{}
			data.SetStanza(Title(match[1]), stanza)
		} else if match := keyValueExpr.FindStringSubmatch(line); len(match) > 0 {
			if stanza == nil {
				return errors.New("key-value pair without stanza: " + line)
			}
			stanza[Key(match[1])] = Value(match[2])
		} else if commentExpr.MatchString(line) {
			// a comment - do nothing
		} else {
			errors.New("unknown expression: " + line)
		}
	}

	return nil
}

func (data Data) Write(filename string) error {
	var result strings.Builder

	for stanza, content := range data {
		result.WriteString("[" + string(stanza) +"]\n")
		for key, value := range content {
			result.WriteString(string(key) + "=" + string(value) + "\n")
		}
	}

	return ioutil.WriteFile(filename, []byte(result.String()), 0644) // rw-r--r--
}