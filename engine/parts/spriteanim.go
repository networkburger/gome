package parts

import "strings"

type SpriteAnimation struct {
	RTexPacker
	Frames []RTexPackerSprite
}

func NewSpriteAnimation(sheet RTexPacker, animationName string) SpriteAnimation {
	frames := []RTexPackerSprite{}

	for _, sprite := range sheet.Entries {
		if strings.Index(sprite.NameId, animationName) == 0 {
			frames = append(frames, sprite)
		}
	}

	return SpriteAnimation{
		RTexPacker: sheet,
		Frames:     frames,
	}
}
