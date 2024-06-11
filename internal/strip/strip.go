// internal/strip/strip.go
package strip

import (
    "regexp"
)

// StripStyles removes all styles (CSS/SCSS/SASS) from the input.
func StripStyles(input string) string {
    re := regexp.MustCompile(`(?s)<style.*?>.*?</style>`)
    return re.ReplaceAllString(input, "")
}

// StripScripts removes all scripts (JavaScript) from the input.
func StripScripts(input string) string {
    re := regexp.MustCompile(`(?s)<script.*?>.*?</script>`)
    return re.ReplaceAllString(input, "")
}

// StripMarkup removes all HTML/XML markup from the input.
func StripMarkup(input string) string {
    re := regexp.MustCompile(`(?s)<.*?>`)
    return re.ReplaceAllString(input, "")
}

// StripMarkdown removes all Markdown syntax from the input.
func StripMarkdown(input string) string {
    re := regexp.MustCompile(`(?m)(^\s*#.*?$|^\s*[*\-].*$|^\s*\d+\..*$|^\s*>\s.*?$|!\[.*?\]\(.*?\)|\[.*?\]\(.*?\))`)
    return re.ReplaceAllString(input, "")
}

// StripCode removes all code (inline and blocks) from the input.
func StripCode(input string) string {
    re := regexp.MustCompile("(?s)```.*?```|`.*?`")
    return re.ReplaceAllString(input, "")
}

// StripAll removes all artifacts, leaving only the plain text.
func StripAll(input string) string {
    input = StripStyles(input)
    input = StripScripts(input)
    input = StripMarkup(input)
    input = StripMarkdown(input)
    input = StripCode(input)
    return input
}
