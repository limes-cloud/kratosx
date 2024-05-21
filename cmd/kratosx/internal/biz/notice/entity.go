package notice

type Classify struct {
}

type Notice struct {
	Id         uint32      `json:"id"`
	ClassifyId uint32      `json:"classifyId"`
	Title      string      `json:"title"`
	Classifies []*Classify `json:"classifies"`
}
