package tree

type Tree interface {
	ID() uint32
	Parent() uint32
	AppendChildren(any)
	ChildrenNode() []Tree
}

func BuildArrayTree(array []Tree) []Tree {
	maxLen := len(array)
	var rootArray []Tree
	rootSet := FindRootSet(array)
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].Parent() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if rootSet[array[i].ID()] {
			rootArray = append(rootArray, array[i])
		}
	}
	return rootArray
}

func BuildTree(array []Tree) Tree {
	maxLen := len(array)
	var rootNode Tree = nil
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].Parent() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if count == 0 && array[i].Parent() == 0 {
			rootNode = array[i]
		}
	}
	return rootNode
}

func BuildTreeByID(array []Tree, id uint32) Tree {
	maxLen := len(array)
	var rootNode Tree = nil
	for i := 0; i < maxLen; i++ {
		count := 0
		for j := 0; j < maxLen; j++ {
			if array[j].ID() == array[i].Parent() {
				count++
				array[j].AppendChildren(array[i])
			}
		}
		if array[i].ID() == id {
			rootNode = array[i]
		}
	}
	return rootNode
}

func getTreeID(t Tree, ids *[]uint32) {
	if t == nil {
		return
	}

	*ids = append(*ids, t.ID())

	for _, item := range t.ChildrenNode() {
		getTreeID(item, ids)
	}
}

func GetTreeID(tree Tree) []uint32 {
	var ids []uint32
	getTreeID(tree, &ids)
	return ids
}

func FindRoots(array []Tree) []uint32 {
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

func FindRootSet(array []Tree) map[uint32]bool {
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
