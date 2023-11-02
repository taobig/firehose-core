// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package firecore

import (
	"fmt"
	"strings"
)

const (
	// PrintOutputModeText is a PrintOutputMode of type Text.
	PrintOutputModeText PrintOutputMode = iota
	// PrintOutputModeJSON is a PrintOutputMode of type JSON.
	PrintOutputModeJSON
	// PrintOutputModeJSONL is a PrintOutputMode of type JSONL.
	PrintOutputModeJSONL
)

var ErrInvalidPrintOutputMode = fmt.Errorf("not a valid PrintOutputMode, try [%s]", strings.Join(_PrintOutputModeNames, ", "))

const _PrintOutputModeName = "TextJSONJSONL"

var _PrintOutputModeNames = []string{
	_PrintOutputModeName[0:4],
	_PrintOutputModeName[4:8],
	_PrintOutputModeName[8:13],
}

// PrintOutputModeNames returns a list of possible string values of PrintOutputMode.
func PrintOutputModeNames() []string {
	tmp := make([]string, len(_PrintOutputModeNames))
	copy(tmp, _PrintOutputModeNames)
	return tmp
}

var _PrintOutputModeMap = map[PrintOutputMode]string{
	PrintOutputModeText:  _PrintOutputModeName[0:4],
	PrintOutputModeJSON:  _PrintOutputModeName[4:8],
	PrintOutputModeJSONL: _PrintOutputModeName[8:13],
}

// String implements the Stringer interface.
func (x PrintOutputMode) String() string {
	if str, ok := _PrintOutputModeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("PrintOutputMode(%d)", x)
}

// IsValid provides a quick way to determine if the typed value is
// part of the allowed enumerated values
func (x PrintOutputMode) IsValid() bool {
	_, ok := _PrintOutputModeMap[x]
	return ok
}

var _PrintOutputModeValue = map[string]PrintOutputMode{
	_PrintOutputModeName[0:4]:                   PrintOutputModeText,
	strings.ToLower(_PrintOutputModeName[0:4]):  PrintOutputModeText,
	_PrintOutputModeName[4:8]:                   PrintOutputModeJSON,
	strings.ToLower(_PrintOutputModeName[4:8]):  PrintOutputModeJSON,
	_PrintOutputModeName[8:13]:                  PrintOutputModeJSONL,
	strings.ToLower(_PrintOutputModeName[8:13]): PrintOutputModeJSONL,
}

// ParsePrintOutputMode attempts to convert a string to a PrintOutputMode.
func ParsePrintOutputMode(name string) (PrintOutputMode, error) {
	if x, ok := _PrintOutputModeValue[name]; ok {
		return x, nil
	}
	// Case insensitive parse, do a separate lookup to prevent unnecessary cost of lowercasing a string if we don't need to.
	if x, ok := _PrintOutputModeValue[strings.ToLower(name)]; ok {
		return x, nil
	}
	return PrintOutputMode(0), fmt.Errorf("%s is %w", name, ErrInvalidPrintOutputMode)
}

// MarshalText implements the text marshaller method.
func (x PrintOutputMode) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *PrintOutputMode) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParsePrintOutputMode(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
