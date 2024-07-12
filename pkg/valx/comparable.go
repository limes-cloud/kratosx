package valx

type In[T ListType] interface {
	Has(T) bool
}

type _comparable[T ListType] struct {
	m map[T]struct{}
}

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

func New[T ListType](list []T) In[T] {
	m := make(map[T]struct{})
	for _, item := range list {
		m[item] = struct{}{}
	}
	return &_comparable[T]{
		m: m,
	}
}

func (c *_comparable[T]) Has(t T) bool {
	_, ok := c.m[t]
	return ok
}
