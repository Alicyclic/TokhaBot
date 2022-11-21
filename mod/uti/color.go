package uti

import (
	"fmt"
	"image"
	"net/http"
	"strconv"
	"strings"

	"github.com/cenkalti/dominantcolor"
)

type Color struct {
	R, G, B int
}

func (c Color) RGBToInt() int {
	return (c.R << 16) + (c.G << 8) + c.B
}

func (c Color) RGBToHex() string {
	return fmt.Sprintf("%02x%02x%02x", c.R, c.G, c.B)
}

func GetColorFromImageURL(url string) Color {
	rq, _ := http.Get(fmt.Sprint(url))
	bodyDecode, _, _ := image.Decode(rq.Body)
	palette := dominantcolor.Hex(dominantcolor.Find(bodyDecode))
	palette = strings.Replace(palette, "#", "", -1)
	value, _ := strconv.ParseInt(palette, 16, 64)
	defer rq.Body.Close()
	return Color{R: int(value >> 16), G: int((value >> 8) & 0xFF), B: int(value & 0xFF)}
}
