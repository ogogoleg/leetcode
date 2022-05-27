package problems

import "strings"

func longestCommonPrefix(strs []string) string {

	firstString := strs[0] // according to 
	remainingStrings := strs[1:]
	var sb strings.Builder

	for i, char := range firstString{  //m

		for _, stringToMatch := range remainingStrings{ //n

			if len(stringToMatch) < i+1 || rune(stringToMatch[i]) != char {
				return sb.String()
			}
		}
		
		sb.WriteRune(char)
	}

	return firstString
}
