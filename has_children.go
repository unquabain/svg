package svg

type HasChildren struct {
	Children ElementArray
}

func (hc *HasChildren) Add(children ...any) {
	hc.Children = append(hc.Children, children...)
}
