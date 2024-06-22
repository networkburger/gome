package v

type Color struct {
	R, G, B, A uint8
}

func NewColor(r, g, b uint8) Color {
	return Color{r, g, b, 255}
}
func NewColorA(r, g, b, a uint8) Color {
	return Color{r, g, b, a}
}
func NewColorInt(r, g, b int) Color {
	return NewColor(uint8(r), uint8(g), uint8(b))
}

var (
	Beige       = NewColorInt(211, 176, 131)
	Black       = NewColorInt(0, 0, 0)
	Blue        = NewColorInt(0, 121, 241)
	Brown       = NewColorInt(127, 106, 79)
	DarkBlue    = NewColorInt(0, 82, 172)
	DarkBrown   = NewColorInt(76, 63, 47)
	DarkGray    = NewColorInt(80, 80, 80)
	DarkGreen   = NewColorInt(0, 117, 44)
	DarkPurple  = NewColorInt(112, 31, 126)
	Gold        = NewColorInt(255, 203, 0)
	Gray        = NewColorInt(130, 130, 130)
	Green       = NewColorInt(0, 228, 48)
	Lime        = NewColorInt(0, 158, 47)
	LightGray   = NewColorInt(200, 200, 200)
	Magenta     = NewColorInt(255, 0, 255)
	Maroon      = NewColorInt(190, 33, 55)
	Orange      = NewColorInt(255, 161, 0)
	Pink        = NewColorInt(255, 109, 194)
	Purple      = NewColorInt(200, 122, 255)
	Red         = NewColorInt(230, 41, 55)
	SkyBlue     = NewColorInt(102, 191, 255)
	Transparent = NewColorInt(0, 0, 0)
	Violet      = NewColorInt(135, 60, 190)
	White       = NewColorInt(255, 255, 255)
	Yellow      = NewColorInt(253, 249, 0)
)
