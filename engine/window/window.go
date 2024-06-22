package window

import rl "github.com/gen2brain/raylib-go/raylib"

func InitWindow(width, height int32, title string) {
	rl.InitWindow(width, height, title)
	rl.SetTargetFPS(15)
	rl.InitAudioDevice()
	rl.SetExitKey(rl.KeyNull)
}

func CloseWindow() {
	rl.CloseAudioDevice()
	rl.CloseWindow()
}

func SetTargetFPS(fps int32) {
	rl.SetTargetFPS(fps)
}
