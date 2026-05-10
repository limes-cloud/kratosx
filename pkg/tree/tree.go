package tree

type Tree[T any] interface {
	ID() uint32
	Parent() uint32
	AppendChildren(T)
	ChildrenNode() []T
}

func BuildArrayTree[T Tree[T]](array []T) []T {
	index := make(map[uint32]T, len(array))
	for _, item := range array {
		index[item.ID()] = item
	}

	rootSet := FindRootSet(array)
	var rootArray []T
	for _, item := range array {
		if parent, ok := index[item.Parent()]; ok {
			parent.AppendChildren(item)
		}
		if rootSet[item.ID()] {
			rootArray = append(rootArray, item)
		}
	}
	return rootArray
}

func BuildTree[T Tree[T]](array []T) T {
	index := make(map[uint32]T, len(array))
	for _, item := range array {
		index[item.ID()] = item
	}

	var rootNode T
	for _, item := range array {
		if parent, ok := index[item.Parent()]; ok {
			parent.AppendChildren(item)
		} else if item.Parent() == 0 {
			rootNode = item
		}
	}
	return rootNode
}

func BuildTreeByID[T Tree[T]](array []T, id uint32) T {
	index := make(map[uint32]T, len(array))
	for _, item := range array {
		index[item.ID()] = item
	}

	var rootNode T
	for _, item := range array {
		if parent, ok := index[item.Parent()]; ok {
			parent.AppendChildren(item)
		}
		if item.ID() == id {
			rootNode = item
		}
	}
	return rootNode
}

func getTreeID[T Tree[T]](t T, ids *[]uint32) {
	*ids = append(*ids, t.ID())

	for _, item := range t.ChildrenNode() {
		getTreeID(item, ids)
	}
}

func GetTreeID[T Tree[T]](tree T) []uint32 {
	var ids []uint32
	getTreeID(tree, &ids)
	return ids
}

func FindRoots[T Tree[T]](array []T) []uint32 {
	idSet := make(map[uint32]bool)
	for _, item := range array {
		idSet[item.ID()] = true
	}

	rootSet := map[uint32]struct{}{}
	for _, item := range array {
		if !idSet[item.Parent()] {
			rootSet[item.ID()] = struct{}{}
		}
	}

	var roots []uint32
	for id := range rootSet {
		roots = append(roots, id)
	}

	return roots
}

func FindRootSet[T Tree[T]](array []T) map[uint32]bool {
	idSet := make(map[uint32]bool)
	for _, item := range array {
		idSet[item.ID()] = true
	}

	rootSet := map[uint32]bool{}
	for _, item := range array {
		if !idSet[item.Parent()] {
			rootSet[item.ID()] = true
		}
	}
	return rootSet
}
