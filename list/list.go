package list

func Of[k any](elements ...k) []k {
	return elements
}