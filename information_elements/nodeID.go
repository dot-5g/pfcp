package information_elements

type NodeIDType int

type NodeID struct {
	Type        int
	Length      int
	NodeIDType  NodeIDType
	NodeIDValue []byte
}
