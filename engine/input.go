package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type ActionID int16
type Key int32
type MouseButton int32

type MouseAxis int32
type GamePadAxis int32
type GamePadButton int32

const (
	MouseAxisNone = iota
	MouseAxisX
	MouseAxisY
)

const (
	MouseButtonNone = iota
	MouseButtonLeft
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonSide
	MouseButtonExtra
	MouseButtonForward
	MouseButtonBack
)

const (
	GamepadAxisNone = iota
	GamepadAxisLeftX
	GamepadAxisLeftY
	GamepadAxisRightX
	GamepadAxisRightY
	GamepadAxisLeftTrigger
	GamepadAxisRightTrigger
)

type InputActionMapping struct {
	ActionID
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

	var mouseDeltaX = float32(0)
	var mouseDeltaY = float32(0)

	for i := 0; i < len(mapping); i++ {
		if rl.IsKeyDown(int32(mapping[i].KeyDown)) {
			process(mapping[i].ActionID, mapping[i].Const)
		}
		if rl.IsKeyPressed(int32(mapping[i].KeyPressed)) {
			process(mapping[i].ActionID, mapping[i].Const)
		}
		if rl.IsKeyReleased(int32(mapping[i].KeyReleased)) {
			process(mapping[i].ActionID, mapping[i].Const)
		}

		if mapping[i].MouseButtonDown != MouseButtonNone && rl.IsMouseButtonDown(int32(mapping[i].MouseButtonDown)-1) {
			process(mapping[i].ActionID, mapping[i].Const)
		}
		if mapping[i].MouseButtonDown != MouseButtonNone && rl.IsMouseButtonPressed(int32(mapping[i].MouseButtonPressed)-1) {
			process(mapping[i].ActionID, mapping[i].Const)
		}
		if mapping[i].MouseButtonDown != MouseButtonNone && rl.IsMouseButtonReleased(int32(mapping[i].MouseButtonReleased)-1) {
			process(mapping[i].ActionID, mapping[i].Const)
		}

		if mapping[i].MouseAxis == MouseAxisX {
			process(mapping[i].ActionID, mouseDeltaX)
		}
		if mapping[i].MouseAxis == MouseAxisY {
			process(mapping[i].ActionID, mouseDeltaY)
		}

		if gamepad != -1 {
			if rl.IsGamepadButtonDown(gamepad, int32(mapping[i].GamePadButtonDown)) {
				process(mapping[i].ActionID, mapping[i].Const)
			}
			if rl.IsGamepadButtonPressed(gamepad, int32(mapping[i].GamePadButtonPressed)) {
				process(mapping[i].ActionID, mapping[i].Const)
			}
			if rl.IsGamepadButtonReleased(gamepad, int32(mapping[i].GamePadButtonReleased)) {
				process(mapping[i].ActionID, mapping[i].Const)
			}

			if mapping[i].GamePadAxis == GamepadAxisLeftX {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisLeftX)*mapping[i].GamePadAxisScale)
			}
			if mapping[i].GamePadAxis == GamepadAxisLeftY {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisLeftY)*mapping[i].GamePadAxisScale)
			}
			if mapping[i].GamePadAxis == GamepadAxisRightX {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisRightX)*mapping[i].GamePadAxisScale)
			}
			if mapping[i].GamePadAxis == GamepadAxisRightY {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisRightY)*mapping[i].GamePadAxisScale)
			}
			if mapping[i].GamePadAxis == GamepadAxisLeftTrigger {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisLeftTrigger)*mapping[i].GamePadAxisScale)
			}
			if mapping[i].GamePadAxis == GamepadAxisRightTrigger {
				process(mapping[i].ActionID, rl.GetGamepadAxisMovement(gamepad, rl.GamepadAxisRightTrigger)*mapping[i].GamePadAxisScale)
			}
		}
	}

	gamepad = -1
}
