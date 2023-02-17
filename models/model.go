package models

type Color string

const (
	Red   Color = "red"
	Blue  Color = "blue"
	Green Color = "green"
	Other Color = "other"
)

var ColorValues = []Color{
	Red, Green, Blue, Other,
}

func (c Color) Valid() bool {
	for _, v := range ColorValues {
		if c == v {
			return true
		}
	}
	return false
}

// Stringer interface
func (c Color) String() string {
	if c.Valid() {
		return string(c)
	}
	return ""
}

type Payload struct {
	Color Color  `json:"color" validate:"enum"`
	Other string `json:"other" validate:"required_if=Color other"`
}

type MulPayload struct {
	Colors []Color `json:"colors" validate:"dive,enum"`
	Other  string  `json:"other" validate:"required_if_element=Colors other"`
}
