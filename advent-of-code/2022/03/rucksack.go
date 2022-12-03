package main

func getPriority(char byte) int {
	if char >= 97 {
		return int(char - 96)
	}
	return int(char - 38)
}

func getGroupPriority(itemDatas []string) int {
	intersection := getIntersection(itemDatas)
	return getPriority(intersection)
}

func getDuplicateItemPriority(itemData string) int {
	comps := getCompartments(itemData, 2)
	return getGroupPriority(comps)
}

func getCompartments(itemData string, compartmentCount int) []string {
	partLength := len(itemData) / compartmentCount
	parts := make([]string, compartmentCount)

	var idx int
	for i := 0; i < compartmentCount; i += 1 {
		parts[i] = itemData[idx : idx+partLength]
		idx += partLength
	}
	return parts
}

func getIntersection(groups []string) byte {
	// The key of charMap is the char; the value is
	// the last group idx it was seen in.
	charMap := make(map[byte]int)
	groupCount := len(groups)

	// Iterate over each item in each group
	for groupIdx, groupItems := range groups {
		for i := 0; i < len(groupItems); i += 1 {
			item := groupItems[i]

			// Check if item is already reported in the char map
			lastCharGroupIdx, ok := charMap[item]

			if !ok {
				// If item is not reported and this is the first group, count it.
				if groupIdx == 0 {
					charMap[item] = 0
				}
				// If item is not reported and this is not the first group, skip it.
				// If it isn't already in the first group, it can't be in the intersection anyway.
				continue
			}

			// If we get here, the char is in the char map

			// If this char is not in the previous group, it can't be part of the intersection
			if lastCharGroupIdx != groupIdx-1 {
				continue
			}

			// If this is the last group, this is the intersecting char.
			if groupIdx == groupCount-1 {
				return item
			}

			// If this is not the last group, it has the potential to be
			// the intersecting char. Set its last seen group idx to this one.
			charMap[item] = groupIdx
		}
	}
	return 0
}
