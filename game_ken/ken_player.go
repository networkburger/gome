package game_ken

import (
	"jamesraine/grl/engine"
	"jamesraine/grl/engine/component"
	"jamesraine/grl/engine/io"
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/sound"
	"jamesraine/grl/engine/v"
	"jamesraine/grl/game_shared"
)

type Player struct {
	Health     int
	snd        sound.Sound
	sprite     *component.SpritesheetComponent
	ballistics *physics.BallisticComponent
	body       *physics.PhysicsBodyComponent
}

func NewPlayerNode(e *engine.Engine, assets *parts.Assets) *engine.Node {
	sheet := assets.SpriteSheet("knight.spritesheet")

	player := Player{
		Health: 100,
		snd:    assets.Sound("jump.wav"),
		sprite: &component.SpritesheetComponent{
			Spritesheet: sheet,
			Texture:     assets.Texture(sheet.ImageRef),
		},
		ballistics: &physics.BallisticComponent{
			VelocityDamping: v.V2(0.1, 0.1),
			AngularDamping:  0.8,
			Gravity:         v.V2(0, 300),
		},
	}

	player.sprite.SetSprite("idle")
	playerNode := e.NewNode("Player")
	playerNode.AddComponent(player.sprite)
	playerNode.AddComponent(&player)
	playerNode.AddComponent(player.ballistics)

	player.body = &physics.PhysicsBodyComponent{
		Radius: 8,
		SurfaceProperties: physics.SurfaceProperties{
			Friction:    0,
			Restitution: 0,
		},
	}
	playerNode.AddComponent(player.body)

	// engine.G.AddComponent(playerNode, &component.CircleComponent{
	// 	Radius: 8,
	// 	Color:  rl.Green,
	// })

	playerNode.Scale = 2
	return playerNode
}

func (s Player) String() string {
	return "Player"
}

func (p *Player) Event(event engine.NodeEvent, gs *engine.Scene, n *engine.Node) {
	if event == engine.NodeEventTick {
		io.ProcessInputs(InputOverworld, func(action io.ActionID, power float32) {
			thrust := float32(500)
			if !p.body.IsOnGround(gs.T) {
				thrust = 80
			}
			switch action {
			case Pause:
				game_shared.ShowPauseMenu(gs)
			case Move:
				p.ballistics.Impulse = p.ballistics.Impulse.Add(v.V2(power*thrust, 0))
			case Jump:
				if p.body.IsOnGroundIsh(gs.T, 0.5) {
					// cancel out any downward velocity so we get a full speed jump
					// useful in double jump, and "grace period" jumps where you've fell off
					// a ledge and hav some downward velocity that would dampen the jump
					// impulse
					if p.ballistics.Velocity.Y < 0 {
						p.ballistics.Velocity = v.V2(p.ballistics.Velocity.X, 0)
					}
					p.ballistics.Impulse = p.ballistics.Impulse.Add(v.V2(0, -9000))
					// reset back to zero to avoid triggering multiple jumps
					p.body.OnGround = 0
					sound.PlaySound(p.snd)
				}
			}
		})

		if p.body.IsOnGround(gs.T) {
			p.ballistics.VelocityDamping = v.V2(3, 3)
		} else {
			p.ballistics.VelocityDamping = v.V2(0.2, 0)
		}

		moving := p.ballistics.Velocity.LenLen() > 0.1
		if !p.body.IsOnGround(gs.T) {
			p.sprite.SetSprite("roll")
		} else if moving {
			p.sprite.SetSprite("run")
			p.sprite.FlipX = p.ballistics.Velocity.X < 0
		} else {
			p.sprite.SetSprite("idle")
		}

		gs.Camera.Position.X = n.Position.X - float32(gs.Engine.WindowPixelWidth)/2
		gs.Camera.Position.Y = n.Position.Y - float32(gs.Engine.WindowPixelHeight)/2
		if gs.Camera.Position.X < gs.Camera.Bounds.X {
			gs.Camera.Position.X = gs.Camera.Bounds.X
		}
		if gs.Camera.Position.X+gs.Camera.Position.W > gs.Camera.Bounds.X+gs.Camera.Bounds.W {
			gs.Camera.Position.X = (gs.Camera.Bounds.X + gs.Camera.Bounds.W) - gs.Camera.Position.W
		}

		if gs.Camera.Position.Y+gs.Camera.Position.H > gs.Camera.Bounds.Y+gs.Camera.Bounds.H {
			gs.Camera.Position.Y = (gs.Camera.Bounds.Y + gs.Camera.Bounds.H) - gs.Camera.Position.H
		}
		if gs.Camera.Position.Y < gs.Camera.Bounds.Y {
			gs.Camera.Position.Y = gs.Camera.Bounds.Y
		}
	}
}
