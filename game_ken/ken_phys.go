package game_ken

import (
	en "jamesraine/grl/engine"
	"log/slog"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type PhysicsObstacle struct {
	*PhysicsSolver
	*en.TilemapComponent
}

func (s *PhysicsObstacle) Tick(gs *en.GameState, n *en.Node) {}
func (s *PhysicsObstacle) Event(e en.NodeEvent, n *en.Node) {
	if e == en.NodeEventLoad {
		s.TilemapComponent = *en.FindComponent[*en.TilemapComponent](n.Components)
		s.PhysicsSolver.world = s
		s.PhysicsSolver.worldNode = n
	} else if e == en.NodeEventUnload {
		s.PhysicsSolver.world = nil
		s.PhysicsSolver.worldNode = nil
	}
}

type PhysicsBody struct {
	*PhysicsSolver
	Radius   float32
	OnGround bool
}

func (s *PhysicsBody) Tick(gs *en.GameState, n *en.Node) {}
func (s *PhysicsBody) Event(e en.NodeEvent, n *en.Node) {
	if e == en.NodeEventLoad {
		s.PhysicsSolver.AddBody(n, s)
	} else if e == en.NodeEventUnload {
		s.PhysicsSolver.RemoveBody(n)
	}
}

type phys_body struct {
	*en.Node
	*PhysicsBody
	*en.BallisticComponent
}

type PhysicsSolver struct {
	worldNode *en.Node
	world     *PhysicsObstacle
	bodies    []phys_body
}

func NewPhysicsSolver() PhysicsSolver {
	return PhysicsSolver{
		bodies: make([]phys_body, 0),
	}
}

func (s *PhysicsSolver) AddBody(n *en.Node, pb *PhysicsBody) {
	ballistics := *en.FindComponent[*en.BallisticComponent](n.Components)
	if ballistics == nil {
		slog.Warn("PhysicsSolver.AddBody: Node has no BallisticComponent")
		return
	}
	b := phys_body{
		Node:               n,
		BallisticComponent: ballistics,
		PhysicsBody:        pb,
	}
	s.bodies = append(s.bodies, b)
}

func (s *PhysicsSolver) RemoveBody(n *en.Node) {
	s.bodies = en.SliceRemoveAll(s.bodies, func(_ int, b phys_body) bool {
		return b.Node == n
	})
}

func (s *PhysicsSolver) Solve(gs *en.GameState) {
	tmap := s.world.TilemapComponent.GetTilemap()
	hits := make([]rl.Rectangle, 4)
	nhits := 0

	for _, b := range s.bodies {
		bpos := b.Node.Position
		radius := b.PhysicsBody.Radius * b.Node.Scale

		for layeri, layer := range tmap.Layers {
			if layeri > 0 {
				// TODO: decide whether this is a collidable layer
				continue
			}
			for chunki, chunk := range layer.Chunks {
				chunkArea := tmap.ChunkPosition(layeri, chunki, s.worldNode.Scale)
				hitsChunk := circleOverlapsRect(bpos, radius, chunkArea)
				if hitsChunk {
					for tilei := range chunk.Data {
						if chunk.Data[tilei] == 0 {
							continue
						}
						// if layeri == 0 && chunki == 0 && tilei == 210 {
						// 	fmt.Println("??")
						// }
						tileArea := tmap.TilePosition(layeri, chunki, tilei, s.worldNode.Scale)
						if circleOverlapsRect(bpos, radius, tileArea) {
							hits[nhits] = tileArea
							nhits++
						}

						if nhits == 4 {
							break
						}
					}
				}
			}
		}

		for i := 0; i < nhits; i++ {
			hit := hits[i]
			dt := bpos.Y - hit.Y
			db := hit.Y + hit.Height - bpos.Y
			dl := bpos.X - hit.X
			dr := hit.X + hit.Width - bpos.X
			if dt < db && dt < dl && dt < dr { // TOP
				b.Node.Position = rl.NewVector2(b.Node.Position.X, hit.Y-radius)
				b.BallisticComponent.Velocity = rl.NewVector2(b.BallisticComponent.Velocity.X, 0)
				b.OnGround = true
			} else if db < dt && db < dl && db < dr { // BOTTOM
				b.Node.Position = rl.NewVector2(b.Node.Position.X, hit.Y+hit.Height+radius)
				b.BallisticComponent.Velocity = rl.NewVector2(b.BallisticComponent.Velocity.X, 0)
			} else if dl < dt && dl < db && dl < dr { // LEFT
				b.Node.Position = rl.NewVector2(hit.X-radius, b.Node.Position.Y)
				b.BallisticComponent.Velocity = rl.NewVector2(0, b.BallisticComponent.Velocity.Y)
			} else if dr < dt && dr < db && dr < dl { // RIGHT
				b.Node.Position = rl.NewVector2(hit.X+hit.Width+radius, b.Node.Position.Y)
				b.BallisticComponent.Velocity = rl.NewVector2(0, b.BallisticComponent.Velocity.Y)
			}

		}
	}
}

func circleOverlapsRect(circlePos rl.Vector2, radius float32, rectangle rl.Rectangle) bool {
	// Find the closest point to the circle within the rectangle
	closestX := circlePos.X
	closestY := circlePos.Y

	if circlePos.X < rectangle.X {
		closestX = rectangle.X
	} else if circlePos.X > rectangle.X+rectangle.Width {
		closestX = rectangle.X + rectangle.Width
	}

	if circlePos.Y < rectangle.Y {
		closestY = rectangle.Y
	} else if circlePos.Y > rectangle.Y+rectangle.Height {
		closestY = rectangle.Y + rectangle.Height
	}

	// Determine collision
	dist := rl.Vector2DistanceSqr(circlePos, rl.NewVector2(closestX, closestY))
	r2 := radius * radius
	return dist < r2
}
