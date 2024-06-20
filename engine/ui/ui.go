package ui

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Layout int
type Alignment int
type WidgetState int

const (
	LayoutPile = iota
	LayoutVertical
	LayoutHorizontal
)

const (
	AlignmentFill = iota
	AlignmentCenter
	AlignmentNear
	AlignmentFar
)

const (
	WidgetStateIdle = iota
	WidgetStateHover
	WidgetStateActive
)

type Thickness struct {
	Left, Top, Right, Bottom float32
}

type Color struct {
	R, G, B, A uint8
}

type Style struct {
	FillColor        Color
	ContentColor     Color
	OutlineColor     Color
	HoverFillColor   Color
	ActiveFillColor  Color
	OutlineThickness float32
}

func (c Color) RL() rl.Color {
	return rl.NewColor(c.R, c.G, c.B, c.A)
}

type Border struct {
	Thickness
	Color
}

func NewThickness(l, t, r, b float32) Thickness {
	return Thickness{
		Left:   l,
		Top:    t,
		Right:  r,
		Bottom: b,
	}
}

func NewColor(r, g, b uint8) Color {
	return Color{r, g, b, 255}
}
func NewColorA(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}
func (c Color) WithAlpha(a uint8) Color {
	c.A = a
	return c
}

var (
	Beige       = NewColor(211, 176, 131)
	Black       = NewColor(0, 0, 0)
	Blue        = NewColor(0, 121, 241)
	Brown       = NewColor(127, 106, 79)
	DarkBlue    = NewColor(0, 82, 172)
	DarkBrown   = NewColor(76, 63, 47)
	DarkGray    = NewColor(80, 80, 80)
	DarkGreen   = NewColor(0, 117, 44)
	DarkPurple  = NewColor(112, 31, 126)
	Gold        = NewColor(255, 203, 0)
	Gray        = NewColor(130, 130, 130)
	Green       = NewColor(0, 228, 48)
	Lime        = NewColor(0, 158, 47)
	LightGray   = NewColor(200, 200, 200)
	Magenta     = NewColor(255, 0, 255)
	Maroon      = NewColor(190, 33, 55)
	Orange      = NewColor(255, 161, 0)
	Pink        = NewColor(255, 109, 194)
	Purple      = NewColor(200, 122, 255)
	Red         = NewColor(230, 41, 55)
	SkyBlue     = NewColor(102, 191, 255)
	Transparent = NewColor(0, 0, 0)
	Violet      = NewColor(135, 60, 190)
	White       = NewColor(255, 255, 255)
	Yellow      = NewColor(253, 249, 0)
)
