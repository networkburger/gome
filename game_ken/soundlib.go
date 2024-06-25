package game_ken

import (
	"jamesraine/grl/engine/parts"
	"jamesraine/grl/engine/sound"
	"math/rand"
)

type SoundClass int32

const (
	SoundJump SoundClass = iota
	SoundCoin
	SoundHurt
	SoundDeath
	SoundHit
	SoundCrash
	SoundPowerup
)

type Soundlib struct {
	sounds map[SoundClass][]sound.Sound
}

func NewSoundlib(assets *parts.Assets) Soundlib {
	lib := make(map[SoundClass][]sound.Sound)
	lib[SoundJump] = []sound.Sound{
		assets.Sound("jump.wav"),
		assets.Sound("jump1.wav"),
	}
	lib[SoundCoin] = []sound.Sound{
		assets.Sound("coin.wav"),
	}
	lib[SoundHurt] = []sound.Sound{
		assets.Sound("hurt.wav"),
	}
	lib[SoundDeath] = []sound.Sound{
		assets.Sound("death.wav"),
	}
	lib[SoundPowerup] = []sound.Sound{
		assets.Sound("powerup.wav"),
	}
	lib[SoundHit] = []sound.Sound{
		assets.Sound("tap.wav"),
	}
	lib[SoundCrash] = []sound.Sound{
		assets.Sound("bang1.wav"),
	}
	return Soundlib{
		sounds: lib,
	}
}

func (s *Soundlib) PlaySound(sc SoundClass) {
	s.PlaySoundAtVolume(sc, 1)
}

func (s *Soundlib) PlaySoundAtVolume(sc SoundClass, vol float32) {
	n := len(s.sounds[sc])
	if n == 0 {
		return
	} else if n == 1 {
		sound.PlaySoundAtVolume(s.sounds[sc][0], vol)
	} else {
		index := rand.Intn(n)
		sound.PlaySoundAtVolume(s.sounds[sc][index], vol)
	}
}
