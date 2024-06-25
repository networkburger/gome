package parts

import "strings"

type SpriteAnimation struct {
	Spritesheet
	Frames []SpritesheetSprite
}

func NewSpriteAnimation(sheet Spritesheet, animationName string) SpriteAnimation {
	frames := []SpritesheetSprite{}

	for _, sprite := range sheet.Entries {
		if strings.Index(sprite.NameId, animationName) == 0 {
			frames = append(frames, sprite)
		}
	}

	return SpriteAnimation{
		Spritesheet: sheet,
		Frames:      frames,
	}
}
