package engine

type GameState struct {
	G         *Engine
	Paused    bool
	Terminate bool
	DT        float32
	T         float64

	// WallClock times update even while "paused"
	WallClockT        float64
	WallClockDT       float32
	WindowPixelHeight int
	WindowPixelWidth  int
	Camera            *Camera
}
