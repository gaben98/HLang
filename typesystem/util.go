package typesystem

func forall(nodes []*HType, a func(*HType) string) []string {
	out := make([]string, 0)
	for _, c := range nodes {
		out = append(out, a(c))
	}
	return out
}

func hptr(val *HType) *IHType {
	var ih IHType = *val
	return &ih
}
