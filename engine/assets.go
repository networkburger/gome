package engine

import (
	"log/slog"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Assets struct {
	Folder       string
	Textures     map[string]rl.Texture2D
	SpriteSheets map[string]*Spritesheet
	Sounds       map[string]rl.Sound
}

func NewAssets(folder string) Assets {
	if folder[len(folder)-1] != '/' {
		folder = folder + "/"
	}
	return Assets{
		Folder:       folder,
		Textures:     make(map[string]rl.Texture2D),
		SpriteSheets: make(map[string]*Spritesheet),
		Sounds:       make(map[string]rl.Sound),
	}
}

func (a *Assets) Path(fname string) string {
	return a.Folder + fname
}

func (a *Assets) Texture(fname string) rl.Texture2D {
	t, ok := a.Textures[fname]
	if ok {
		return t
	}
	fpath := a.Path(fname)
	slog.Info("Assets.Texture", "fname", fname, "fpath", fpath)
	tex := rl.LoadTexture(fpath)
	if tex.ID == 0 {
		slog.Warn("Assets.Texture FAIL", "fname", fname)
	} else {
		slog.Info("Assets.Texture DONE", "fname", fname, "ID", tex.ID, "W", tex.Width, "H", tex.Height)
	}
	a.Textures[fname] = tex
	return tex
}

func (a *Assets) SpriteSheet(fname string) *Spritesheet {
	existing, ok := a.SpriteSheets[fname]
	if ok {
		return existing
	}
	sheetData, _ := a.FileBytes(fname)
	s := SpritesheetRead(sheetData)
	a.SpriteSheets[fname] = &s
	return &s
}

func (a *Assets) Sound(fname string) rl.Sound {
	s, ok := a.Sounds[fname]
	if ok {
		return s
	}
	fpath := a.Path(fname)
	slog.Info("Assets.Sound", "fname", fname, "fpath", fpath)
	snd := rl.LoadSound(fpath)
	a.Sounds[fname] = snd
	return snd
}

func (a *Assets) FileBytes(fname string) ([]byte, error) {
	fpath := a.Path(fname)
	slog.Info("Assets.FileBytes", "fname", fname, "fpath", fpath)
	dat, err := os.ReadFile(fpath)
	if err != nil {
		slog.Warn("Assets.FileBytes", "error", err)
	} else {
		slog.Info("Assets.FileBytes", "fname", fname, "bytes", len(dat))
	}

	return dat, err
}
