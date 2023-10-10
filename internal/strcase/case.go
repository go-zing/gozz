package strcase

import (
	"strings"
)

// KebabCase converts a string into kebab case.
func KebabCase(s string) string {
	return delimiterCase(s, '-', toLower)
}

// SnakeCase converts a string into snake case.
func SnakeCase(s string) string {
	return delimiterCase(s, '_', toLower)
}

// delimiterCase converts a string into snake_case or kebab-case depending on the delimiter passed
// as second argument. When upperCase is true the result will be UPPER_SNAKE_CASE or UPPER-KEBAB-CASE.
func delimiterCase(s string, delimiter rune, adjustCase func(rune) rune) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s)+3)

	var prev rune
	var curr rune
	for _, next := range s {
		if isDelimiter(curr) {
			if !isDelimiter(prev) {
				buffer = append(buffer, delimiter)
			}
		} else if isUpper(curr) {
			if isLower(prev) || (isUpper(prev) && isLower(next)) {
				buffer = append(buffer, delimiter)
			}
			buffer = append(buffer, adjustCase(curr))
		} else if curr != 0 {
			buffer = append(buffer, adjustCase(curr))
		}
		prev = curr
		curr = next
	}

	if len(s) > 0 {
		if isUpper(curr) && isLower(prev) && prev != 0 {
			buffer = append(buffer, delimiter)
		}
		buffer = append(buffer, adjustCase(curr))
	}

	return string(buffer)
}

// CamelCase converts a string into camel case starting with a lower case letter.
func CamelCase(s string) string {
	return camelCase(s, false)
}

func UpperCamelCase(s string) string {
	return camelCase(s, true)
}

func camelCase(s string, upper bool) string {
	s = strings.TrimSpace(s)
	buffer := make([]rune, 0, len(s))

	stringIter(s, func(prev, curr, next rune) {
		if !isDelimiter(curr) {
			if isDelimiter(prev) || (upper && prev == 0) {
				buffer = append(buffer, toUpper(curr))
			} else if isLower(prev) {
				buffer = append(buffer, curr)
			} else if isUpper(prev) && isUpper(curr) && isLower(next) {
				// Assume a case like "R" for "XRequestId"
				buffer = append(buffer, curr)
			} else {
				buffer = append(buffer, toLower(curr))
			}
		}
	})

	return string(buffer)
}
