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
	PlaySoundAtVolume(snd, 1)
}

func PlaySoundAtVolume(snd Sound, vol float32) {
	rls := rl.Sound(snd)
	rl.SetSoundVolume(rls, vol)
	rl.PlaySound(rls)
}
