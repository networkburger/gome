package game_ken

import (
	en "jamesraine/grl/engine"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Health     int
	snd        rl.Sound
	sprite     *en.SpritesheetComponent
	ballistics *en.BallisticComponent
	body       *PhysicsBody
}

func NewPlayerNode(assets *en.Assets, solver *PhysicsSolver) en.Node {
	sheet := assets.SpriteSheet("knight.spritesheet")

	player := Player{
		Health: 100,
		snd:    assets.Sound("jump.wav"),
		sprite: &en.SpritesheetComponent{
			Spritesheet: sheet,
			Texture:     assets.Texture(sheet.ImageRef),
		},
		ballistics: &en.BallisticComponent{
			VelocityDamping: 0.1,
			AngularDamping:  0.8,
			Gravity:         rl.NewVector2(0, 160),
		},
	}

	player.sprite.SetSprite("idle")
	playerNode := en.NewNode("Player")
	en.G.AddComponent(&playerNode, player.sprite)
	en.G.AddComponent(&playerNode, &player)
	en.G.AddComponent(&playerNode, player.ballistics)

	player.body = &PhysicsBody{
		PhysicsSolver: solver,
		Radius:        8,
	}
	en.G.AddComponent(&playerNode, player.body)
	return playerNode
}

func (s Player) String() string {
	return "Player"
}

func (s *Player) Event(e en.NodeEvent, n *en.Node) {}

func (p *Player) Tick(gs *en.GameState, n *en.Node) {
	en.ProcessInputs(InputOverworld, func(action en.ActionID, power float32) {
		thrust := float32(300)
		if !p.body.OnGround {
			thrust = 20
		}
		switch action {
		case Move:
			p.ballistics.Impulse = rl.Vector2Add(p.ballistics.Impulse, rl.NewVector2(power*thrust, 0))
		case Jump:
			if p.body.OnGround {
				p.ballistics.Impulse = rl.Vector2Add(p.ballistics.Impulse, rl.NewVector2(0, -6000))
				p.body.OnGround = false
				rl.PlaySound(p.snd)
			}
		}
	})

	if p.body.OnGround {
		p.ballistics.VelocityDamping = 3
	} else {
		p.ballistics.VelocityDamping = 0.2
	}

	moving := rl.Vector2LenSqr(p.ballistics.Velocity) > 0.1
	if moving {
		p.sprite.SetSprite("run")
		p.sprite.FlipX = p.ballistics.Velocity.X < 0
	} else {
		p.sprite.SetSprite("idle")
	}
}
