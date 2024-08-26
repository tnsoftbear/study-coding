package logs

import (
	"strings"
	"unicode/utf8"
)

func markers() map[rune]string {
	return map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}
}

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, char := range log {
		for rune, appl := range markers() {
			if rune == char {
				return appl
			}
		}
	}
	return "default"
}

// Replace replaces all occurances of old with new, returning the modified log
// to the caller.
func Replace(log string, old, new rune) string {
	var sb strings.Builder
	for _, r := range log {
		if r == old {
			sb.WriteString(string(new))
		} else {
			sb.WriteString(string(r))
		}
	}
	return sb.String()
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}

// func Application(log string) string {
//     occurrance := strings.IndexAny(log, "❗🔍☀")
//     if occurrance < 0 {
//         return "default"
//     }
// 	return map[rune]string{
//         '❗': "recommendation",
//         '🔍': "search",
//         '☀': "weather",
//     }[[]rune(log)[occurrance]]
// }

// func Replace(log string, old, new rune) string {
// 	return strings.ReplaceAll(log, fmt.Sprintf("%c", old), fmt.Sprintf("%c", new))
// }
