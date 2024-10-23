package of

func Map[K any](collection []K, transformer func(K) K) []K {
	var result []K

	for _, element := range collection {
		result = append(result, transformer(element))
	}

	return result
}

func MapWithIndex[K any](collection []K, transformer func(int, K) K) []K {
	var result []K

	for i, element := range collection {
		result = append(result, transformer(i, element))
	}

	return result
}

func Filter[K any](collection []K, predicate func(K) bool) []K {
	var result []K

	for _, element := range collection {
		if predicate(element) {
			result = append(result, element)
		}
	}

	return result
}

func FilterWithIndex[K any](collection []K, predicate func(int, K) bool) []K {
	var result []K
	for i, element := range collection {
		if predicate(i, element) {
			result = append(result, element)
		}
	}

	return result
}

func ForeEach[K any](collection []K, action func(K)) {
	for _, element := range collection {
		action(element)
	}
}

func ForEachWithIndex[K any](collection []K, action func(int, K)) {
	for i, element := range collection {
		action(i, element)
	}
}

func Fold[K any, R any](collection []K, initial R, reducer func(R, K) R) R {
	var result R = initial

	for _, element := range collection {
		result = reducer(result, element)
	}

	return result
}

func Any[K comparable](collection []K, value K) bool {
	for _, element := range collection {
		if element == value {
			return true
		}
	}
	return false
}

func All[K comparable](collection []K, value K) bool {
	for _, element := range collection {
		if element != value {
			return false
		}
	}
	return true
}

func None[K comparable](collection []K, value K) bool {
	for _, element := range collection {
		if element == value {
			return false
		}
	}
	return true
}

func Count[K comparable](collection []K, value K) int {
	var result int

	for _, element := range collection {
		if element == value {
			result++
		}
	}

	return result
}

func Head[K any](collection []K) K {
	return collection[0]
}

func Tail[K any](collection []K) K {
	return collection[len(collection)-1]
}

func Zip[K any, R any](collection1 []K, collection2 []R, transformer func(K, R) K) []K {
	var result []K

	for i := 0; i < len(collection1); i++ {
		result = append(result, transformer(collection1[i], collection2[i]))
	}

	return result
}
