package xslice

// SetIntersection returns the elements in `a` and `b`. (a∩b)
func SetIntersection(a []string, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; found {
			diff = append(diff, x)
		}
	}
	return diff
}

// SetUnion returns the elements in `a` or `b`. (a∪b)
func SetUnion(a []string, b []string) []string {
	mb := make(map[string]struct{}, len(b)+len(a))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	for _, x := range a {
		mb[x] = struct{}{}
	}
	var res []string
	for x := range mb {
		res = append(res, x)
	}
	return res
}
