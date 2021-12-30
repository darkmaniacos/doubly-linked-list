package list

type IdNode_t struct {
	Parent *IdNode_t
	Child *IdNode_t
	Id int32
}

type IdList_t struct {
	First *IdNode_t;
}

func IdNode(parent *IdNode_t, child *IdNode_t, Id int32) *IdNode_t {
	return &IdNode_t { Parent: parent, Child: child, Id: Id };
}

func IdList() *IdList_t {
	return &IdList_t { First: nil };
}

func GetId_s(node *IdNode_t) int32 {
	if node == nil {
		return -1;
	}

	return node.Id;
}

func (this *IdList_t) AddToList(Id int32) {
	if this.First == nil {
		this.First = IdNode(nil, nil, Id);
		return;
	}

	var current *IdNode_t = this.First;
	for !(current.Child == nil) {
		current = current.Child;
	}

	current.Child = IdNode(current, nil, Id);
}

func (this *IdList_t) RemoveById(Id int32) {
	if this.First != nil && this.First.Id == Id {
		this.First = this.First.Child;
		if (this.First == nil) {
			return;
		}

		this.First.Parent = nil;
		return;
	}

	var current *IdNode_t = this.First.Child;
	for !(current == nil) {
		if current.Id == Id {
			current.Parent.Child = current.Child;

			if current.Child == nil {
				return;
			}

			current.Child.Parent = current.Parent;
			return;
		}
		current = current.Child;
	}
}