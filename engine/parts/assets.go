package parts

import (
	"jamesraine/grl/engine/render"
	"jamesraine/grl/engine/sound"
	"log/slog"
	"os"
)

type Assets struct {
	Folder       string
	Textures     map[string]render.Texture2D
	SpriteSheets map[string]Spritesheet
	Sounds       map[string]sound.Sound
	Images       map[string]render.PixelBuffer
	Fonts        map[string]FontRenderer
}

func NewAssets(folder string) Assets {
	if folder[len(folder)-1] != '/' {
		folder = folder + "/"
	}
	return Assets{
		Folder:       folder,
		Textures:     make(map[string]render.Texture2D),
		SpriteSheets: make(map[string]Spritesheet),
		Sounds:       make(map[string]sound.Sound),
		Images:       make(map[string]render.PixelBuffer),
		Fonts:        make(map[string]FontRenderer),
	}
}

func (a *Assets) Close() {
	for _, t := range a.Textures {
		render.UnloadTexture(t)
	}
	for _, s := range a.Sounds {
		sound.UnloadSound(s)
	}
	a.Textures = make(map[string]render.Texture2D)
	a.Sounds = make(map[string]sound.Sound)
	a.Images = make(map[string]render.PixelBuffer)
	a.SpriteSheets = make(map[string]Spritesheet)
	a.Fonts = make(map[string]FontRenderer)
}

func (a *Assets) Path(fname string) string {
	return a.Folder + fname
}

func (a *Assets) Texture(fname string) render.Texture2D {
	t, ok := a.Textures[fname]
	if ok {
		return t
	}
	fpath := a.Path(fname)
	slog.Info("Assets.Texture", "fname", fname, "fpath", fpath)
	tex := render.LoadTexture(fpath)
	if tex.ID == 0 {
		slog.Warn("Assets.Texture FAIL", "fname", fname)
	} else {
		slog.Info("Assets.Texture DONE", "fname", fname, "ID", tex.ID, "W", tex.Width, "H", tex.Height)
	}
	a.Textures[fname] = tex
	return tex
}

func (a *Assets) Pixels(fname string) render.PixelBuffer {
	t, ok := a.Images[fname]
	if ok {
		return t
	}
	fpath := a.Path(fname)
	slog.Info("Assets.Image", "fname", fname, "fpath", fpath)
	buf := render.LoadPixels(fpath)
	slog.Info("Assets.Image DONE", "fname", fname, "W", buf.W, "H", buf.H)
	a.Images[fname] = buf
	return buf
}

func (a *Assets) SpriteSheet(fname string) (Spritesheet, error) {
	existing, ok := a.SpriteSheets[fname]
	if ok {
		return existing, nil
	}
	sheetData, _ := a.FileBytes(fname)
	s, err := SpritesheetRead(sheetData)
	if err != nil {
		return s, err
	}
	a.SpriteSheets[fname] = s
	return s, nil
}

func (a *Assets) Sound(fname string) sound.Sound {
	s, ok := a.Sounds[fname]
	if ok {
		return s
	}
	fpath := a.Path(fname)
	slog.Info("Assets.Sound", "fname", fname, "fpath", fpath)
	snd := sound.LoadSound(fpath)
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

func (a *Assets) Font(fname string) (FontRenderer, error) {
	s, ok := a.Fonts[fname]
	if ok {
		return s, nil
	}

	spriteSheet, err := a.SpriteSheet(fname)
	if err != nil {
		return FontRenderer{}, err
	}
	fontSpec, err := NewFont(spriteSheet)
	if err != nil {
		return FontRenderer{}, err
	}

	font := FontRenderer{
		Font:    fontSpec,
		Texture: a.Texture(fontSpec.ImagePath),
	}
	return font, nil
}
