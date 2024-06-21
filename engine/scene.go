package engine

type Physics interface {
	Register(*Node)
	Unregister(*Node)
	Solve(*Scene)
}

type Scene struct {
	*Engine
	*Node
	Physics
	Camera

	Paused      bool
	DT          float32
	T           float64
	WallClockT  float64 // WallClock times update even while "paused"
	WallClockDT float32
}
