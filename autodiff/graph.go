package autodiff

type Edge struct {
	Grad float64
	Uid  int64
}

type Graph struct {
	nodes map[int64][]Edge
}

var graphInstance *Graph
var uidCounter int64

func NewGraph() *Graph {
	return &Graph{nodes: make(map[int64][]Edge)}
}

func GetGraph() *Graph {
	if graphInstance == nil {
		graphInstance = NewGraph()
	}
	return graphInstance
}

func NextUID() int64 {
	uidCounter++
	return uidCounter
}

func (g *Graph) Connect(uid int64, grad float64, nodeUid int64) {
	g.nodes[uid] = append(g.nodes[uid], Edge{Grad: grad, Uid: nodeUid})
}

func (g *Graph) IsConnected(uid int64) bool {
	_, ok := g.nodes[uid]
	return ok
}

func (g *Graph) Get(uid int64) []Edge {
	return g.nodes[uid]
}

func (g *Graph) NewRecording() {
	g.nodes = make(map[int64][]Edge)
}
