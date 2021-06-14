package utility

import "regexp"

func ExtractKeyFromFields(s []byte, fieldSeparator *regexp.Regexp, index int) string {
	if fieldSeparator == nil {
		return string(s)
	}

	separatorMatches := fieldSeparator.FindAllIndex(s, index+1)
	if separatorMatches == nil {
		// consider the whole string as the first field
		if index == 0 {
			return string(s)
		}
	} else if len(separatorMatches) >= index {
		// field index is available, so find the field and use it as key
		keyStart := 0
		if index > 0 {
			keyStart = separatorMatches[index-1][1]
		}

		keyEnd := len(s)
		if index < len(separatorMatches) {
			keyEnd = separatorMatches[index][0]
		}

		keyBytes := s[keyStart:keyEnd]
		return string(keyBytes)
	}

	// field index is not available in this string, so key is empty
	return ""
}
