package engine

type Scene struct {
	G        *Engine
	RootNode *Node
	Paused   bool
	DT       float32
	T        float64

	// WallClock times update even while "paused"
	WallClockT  float64
	WallClockDT float32
	Camera      *Camera
}
