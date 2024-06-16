package game_ken

import (
	en "jamesraine/grl/engine"
	cm "jamesraine/grl/engine/component"
	"jamesraine/grl/engine/contact"
	pt "jamesraine/grl/engine/parts"
	ph "jamesraine/grl/engine/physics"
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Health     int
	snd        rl.Sound
	sprite     *cm.SpritesheetComponent
	ballistics *cm.BallisticComponent
	body       *cm.PhysicsBodyComponent
}

func NewPlayerNode(assets *pt.Assets, solver *ph.PhysicsSolver) *en.Node {
	sheet := assets.SpriteSheet("knight.spritesheet")

	player := Player{
		Health: 100,
		snd:    assets.Sound("jump.wav"),
		sprite: &cm.SpritesheetComponent{
			Spritesheet: sheet,
			Texture:     assets.Texture(sheet.ImageRef),
		},
		ballistics: &cm.BallisticComponent{
			VelocityDamping: v.V2(0.1, 0.1),
			AngularDamping:  0.8,
			Gravity:         v.V2(0, 300),
		},
	}

	player.sprite.SetSprite("idle")
	playerNode := en.NewNode("Player")
	en.G.AddComponent(playerNode, player.sprite)
	en.G.AddComponent(playerNode, &player)
	en.G.AddComponent(playerNode, player.ballistics)

	player.body = &cm.PhysicsBodyComponent{
		PhysicsManager: solver,
		Radius:         8,
		SurfaceProperties: contact.SurfaceProperties{
			Friction:    0,
			Restitution: 0.025,
		},
	}
	en.G.AddComponent(playerNode, player.body)

	// en.G.AddComponent(playerNode, &cm.CircleComponent{
	// 	Radius: 8,
	// 	Color:  rl.Green,
	// })

	playerNode.Scale = 2
	return playerNode
}

func (s Player) String() string {
	return "Player"
}

func (s *Player) Event(e en.NodeEvent, n *en.Node) {}

func (p *Player) Tick(gs *en.GameState, n *en.Node) {
	en.ProcessInputs(InputOverworld, func(action en.ActionID, power float32) {
		thrust := float32(500)
		if !p.body.IsOnGround(gs.T) {
			thrust = 80
		}
		switch action {
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
				rl.PlaySound(p.snd)
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
}
