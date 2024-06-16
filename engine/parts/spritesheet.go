package parts

import (
	"jamesraine/grl/engine/v"
	"log/slog"
	"strconv"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type SpritesheetFrame struct {
	SpriteName string
	Source     rl.Rectangle
	Origin     v.Vec2
}

type Spritesheet struct {
	ImageRef string
	Frames   []SpritesheetFrame
}

func SpritesheetRead(filedata []byte) Spritesheet {
	fileDataString := string(filedata)
	sanitised := strings.ReplaceAll(fileDataString, "\r\n", "\n")
	lines := strings.Split(sanitised, "\n")
	var ss Spritesheet
	for _, line := range lines {
		switch {
		case len(line) < 1:
			// pass
		case line[0] == 'i':
			ss = readImageRefLine(ss, line)
		case line[0] == 'f':
			ss = readFrameLine(ss, line)
		}
	}
	return ss
}

func (s *Spritesheet) NumberOfFrames(spritename string) int {
	n := 0
	for i := range s.Frames {
		if strings.Compare(spritename, s.Frames[i].SpriteName) == 0 {
			n++
		}
	}
	return n
}

func (s *Spritesheet) GetFrame(spritename string, index int) SpritesheetFrame {
	n := -1
	for i := range s.Frames {
		if strings.Compare(spritename, s.Frames[i].SpriteName) == 0 {
			n++
		}
		if n == index {
			return s.Frames[i]
		}
	}
	return SpritesheetFrame{}
}

//
//
//

func readImageRefLine(ss Spritesheet, line string) Spritesheet {
	args := strings.Split(line, " ")
	if len(args) != 2 {
		slog.Warn("SpriteSheet.readImageRefLine: expected 2 args", "line", line)
		return ss
	}

	ss.ImageRef = args[1]
	return ss
}

func readFrameLine(ss Spritesheet, line string) Spritesheet {
	args := strings.Split(line, " ")
	if len(args) != 8 {
		slog.Warn("SpriteSheet.readFrameLine: expected 6 args", "line", line)
		return ss
	}

	x, err := strconv.Atoi(args[2])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: x arg not int")
		return ss
	}
	y, err := strconv.Atoi(args[3])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: y arg not int")
		return ss
	}
	w, err := strconv.Atoi(args[4])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: w arg not int")
		return ss
	}
	h, err := strconv.Atoi(args[5])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: h arg not int")
		return ss
	}
	ox, err := strconv.Atoi(args[6])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: originx arg not int")
		return ss
	}
	oy, err := strconv.Atoi(args[7])
	if err != nil {
		slog.Warn("SpriteSheet.readFrameLine: originy arg not int")
		return ss
	}

	frame := SpritesheetFrame{
		SpriteName: args[1],
		Source:     rl.NewRectangle(float32(x), float32(y), float32(w), float32(h)),
		Origin:     v.V2(float32(ox), float32(oy)),
	}

	ss.Frames = append(ss.Frames, frame)

	return ss
}
