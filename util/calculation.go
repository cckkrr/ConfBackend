package util

// Intersection returns the intersection of two string slices
func Intersection(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	intersection := []string{}

	for _, str := range slice1 {
		set[str] = true
	}

	for _, str := range slice2 {
		if set[str] {
			intersection = append(intersection, str)
		}
	}

	return intersection
}

// Difference 返回两个集合的差集（对于sliceA）
func Difference(sliceA, sliceB []string) []string {
	set := make(map[string]bool)

	// 记录集合B中的元素
	for _, str := range sliceB {
		set[str] = true
	}

	difference := []string{}

	// 遍历集合A，如果元素不在集合B中，则添加到差集结果中
	for _, str := range sliceA {
		if !set[str] {
			difference = append(difference, str)
		}
	}

	return difference
}

// Union returns the union of two string slices
func Union(slice1, slice2 []string) []string {
	set := make(map[string]bool)
	union := []string{}

	for _, str := range slice1 {
		set[str] = true
	}

	for _, str := range slice2 {
		set[str] = true
	}

	for str := range set {
		union = append(union, str)
	}

	return union
}
