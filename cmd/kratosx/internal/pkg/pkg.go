package pkg

type ListType interface {
	~string | ~int | ~uint32 | ~rune | ~float64 | ~int64 | ~float32
}

func InList[ListType comparable](list []ListType, val ListType) bool {
	for _, v := range list {
		if v == val {
			return true
		}
	}
	return false
}
