package tmplhelp

import "strings"

// Indent takes a string, and indents it the given number of spaces.
func Indent(pad int, in string) string {
	spaces := strings.Repeat(" ", pad)
	return spaces + strings.Replace(in, "\n", "\n"+spaces, -1)
}
