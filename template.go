package main

import (
	"embed"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

//go:embed templates/*.tmpl
var templateFS embed.FS
var nixTemplates = template.New("nix").Funcs(sprig.FuncMap()).Funcs(funcMap)

func execTemplate(t *template.Template) func(string, any) (string, error) {
	_ = "STUB: not implemented"
	return nil
}

// indentNonEmpty indents the given text by the provided number of spaces while
// skipping empty lines.
func indentNonEmpty(spaces int, text string) string { _ = "STUB: not implemented"; return "" }

func derefInt(v *int) int { _ = "STUB: not implemented"; return 0 }

func toNixValue(v any) any { _ = "STUB: not implemented"; return *new(any) }

func toNixList(s []string) string { _ = "STUB: not implemented"; return "" }

// We purposefully do not use %q to avoid Go's built-in string escaping.

func escapeNixString(s string) string {
	_ = "STUB: not implemented"
	// https://nix.dev/manual/nix/latest/language/syntax#string-literal
	return ""
}

func escapeIndentedNixString(s string) string {
	_ = "STUB: not implemented"
	// https://nix.dev/manual/nix/latest/language/syntax#string-literal
	return ""
}

var funcMap template.FuncMap = template.FuncMap{
	"derefInt":                derefInt,
	"toNixValue":              toNixValue,
	"toNixList":               toNixList,
	"escapeNixString":         escapeNixString,
	"escapeIndentedNixString": escapeIndentedNixString,
}
