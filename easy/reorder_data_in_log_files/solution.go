package reorder_data_log_files

import (
	"sort"
	"strings"
)

func reorderLogFiles(logs []string) []string {
	lets := []string{}
	digs := []string{}

	for _, l := range logs {
		split := strings.Split(l, " ")
		// cant just attempt to convert string to int and check err,
		// because string might be a numerical value too large to fit
		// in an integer, and thus returning an error
		firstCharOfLastWord := []byte(split[len(split)-1])[0]
		if firstCharOfLastWord >= '0' && firstCharOfLastWord <= '9' {
			digs = append(digs, l)
		} else {
			lets = append(lets, l)
		}
	}

	// letter logs must be returned in lexicographical order
	sort.Slice(lets, func(i, j int) bool {
		aSplit := strings.Split(lets[i], " ")
		bSplit := strings.Split(lets[j], " ")

		a := strings.Join(aSplit[1:], " ")
		b := strings.Join(bSplit[1:], " ")

		if lex := strings.Compare(a, b); lex == -1 {
			return true
		} else if lex == 1 {
			return false
		} else {
			// in case of a tie, we use the identifier
			idLex := strings.Compare(aSplit[0], bSplit[0])
			if idLex == -1 {
				return true
			}
			return false
		}
	})

	return append(lets, digs...)
}
