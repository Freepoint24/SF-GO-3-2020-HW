package intersect

const (
	valueExist uint8 = 1
	valueAdded uint8 = 2
)

func StringSlices(a, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return []string{}
	}

	matches := make(map[string]uint8)
	for _, v := range a {
		matches[v] = valueExist
	}

	intersection := make([]string, 0)
	for _, v := range b {
		if matches[v] == valueExist {
			matches[v] = valueAdded
			intersection = append(intersection, v)
		}
	}

	return intersection
}
