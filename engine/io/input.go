package io

import (
	"jamesraine/grl/engine/v"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type InputActionMapping struct {
	ActionID
	Triggers []InputVector
}

type InputVector struct {
	MouseAxis
	GamePadAxis
	GamePadAxisScale      float32
	Const                 float32
	GamePadButtonPressed  GamePadButton
	GamePadButtonReleased GamePadButton
	GamePadButtonDown     GamePadButton
	MouseButtonPressed    MouseButton
	MouseButtonReleased   MouseButton
	MouseButtonDown       MouseButton
	KeyPressed            Key
	KeyDown               Key
	KeyReleased           Key
}

func ProcessInputs(mapping []InputActionMapping, process func(ActionID, float32)) {
	gamepad := int32(-1)
	if rl.IsGamepadAvailable(0) {
		gamepad = 0
	}

	var mouseDelta = rl.GetMouseDelta()
	for _, m := range mapping {
		for _, iv := range m.Triggers {
			if iv.GamePadAxis != GamepadAxisNone && gamepad != -1 {
				delta := rl.GetGamepadAxisMovement(gamepad, int32(iv.GamePadAxis)-1) // compensate for our "none" item at 0
				if v.Absf(delta) > 0.1 {
					process(m.ActionID, delta*iv.GamePadAxisScale)
				}
			} else if iv.GamePadButtonDown != GamepadButtonUnknown {
				if rl.IsGamepadButtonDown(gamepad, int32(iv.GamePadButtonDown)) {
					process(m.ActionID, iv.Const)
				}
			} else if iv.GamePadButtonPressed != GamepadButtonUnknown {
				if rl.IsGamepadButtonPressed(gamepad, int32(iv.GamePadButtonPressed)) {
					process(m.ActionID, iv.Const)
				}
			} else if iv.GamePadButtonReleased != GamepadButtonUnknown {
				if rl.IsGamepadButtonDown(gamepad, int32(iv.GamePadButtonReleased)) {
					process(m.ActionID, iv.Const)
				}
			} else if iv.KeyDown != rl.KeyNull && rl.IsKeyDown(int32(iv.KeyDown)) {
				process(m.ActionID, iv.Const)
			} else if iv.KeyPressed != rl.KeyNull && rl.IsKeyDown(int32(iv.KeyPressed)) {
				process(m.ActionID, iv.Const)
			} else if iv.KeyReleased != rl.KeyNull && rl.IsKeyDown(int32(iv.KeyReleased)) {
				process(m.ActionID, iv.Const)
			} else if iv.MouseAxis == MouseAxisX && v.Absf(mouseDelta.X) > 0.1 {
				process(m.ActionID, mouseDelta.X)
			} else if iv.MouseAxis == MouseAxisY && v.Absf(mouseDelta.Y) > 0.1 {
				process(m.ActionID, mouseDelta.Y)
			} else if iv.MouseButtonDown != MouseButtonNone && rl.IsMouseButtonDown(int32(iv.MouseButtonDown)-1) {
				process(m.ActionID, iv.Const)
			}
		}
	}

	gamepad = -1
}
