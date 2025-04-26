package parsing

import "strings"

func SplitSpaces(line *string) []string {
	return strings.Fields(*line)
}
