package templates

const (
	// GitIgnore .gitignore template.
	GitIgnore = `# Binaries for programs and plugins
*.exe
*.bin
*.log
*.exe~
*.dll
*.so
*.dylib

*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out
`

	// GitAttributes .gitattributes template.
	GitAttributes = `* text=auto eol=lf`
)
