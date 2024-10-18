package of

func ListOf[k any](elements ...k) []k {
	var res []k
	for _, e := range elements {
		res = append(res, e)
	}
	return res
}
