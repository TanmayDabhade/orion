package ranking

import (
	"sort"
	"strings"

	"orion/internal/history"
	"orion/internal/shortcuts"
)

func RankedKeys(keys []string, usage map[string]history.Usage) []string {
	ranked := make([]string, 0, len(keys))
	ranked = append(ranked, keys...)

	sort.SliceStable(ranked, func(i, j int) bool {
		left := usage[shortcuts.Normalize(ranked[i])]
		right := usage[shortcuts.Normalize(ranked[j])]
		if left.Count != right.Count {
			return left.Count > right.Count
		}
		if left.LastUnix != right.LastUnix {
			return left.LastUnix > right.LastUnix
		}
		return strings.ToLower(ranked[i]) < strings.ToLower(ranked[j])
	})

	return ranked
}
