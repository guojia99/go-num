package trees

type Node struct {
	Parent   *Node       `json:"p"`
	Children []*Node     `json:"c"`
	Data     []*NodeData `json:"d"`
}

type NodeData struct {
	Key   interface{} `json:"k"`
	Value interface{} `json:"v"`
}
