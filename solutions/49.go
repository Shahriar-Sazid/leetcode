package solutions

import (
	"slices"
)

func groupAnagrams(strs []string) [][]string {
	groups := map[string][]string{}

	for _, str := range strs {
		strRunes := []rune(str)
		slices.Sort(strRunes)
		groupID := string(strRunes)

		if _, ok := groups[groupID]; !ok {
			groups[groupID] = []string{}
		}

		groups[groupID] = append(groups[groupID], str)
	}

	groupList := [][]string{}
	for _, group := range groups {
		groupList = append(groupList, group)
	}

	return groupList
}
