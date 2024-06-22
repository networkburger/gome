package sound

import rl "github.com/gen2brain/raylib-go/raylib"

type Sound rl.Sound

func LoadSound(fpath string) Sound {
	return Sound(rl.LoadSound(fpath))
}

func UnloadSound(snd Sound) {
	rl.UnloadSound(rl.Sound(snd))
}

func PlaySound(snd Sound) {
	rl.PlaySound(rl.Sound(snd))
}
