package util

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"text/template"

	"github.com/nfnt/resize"
)

const ascii_artT = `
<html>
	<pre>
		{{ .Art }}
	</pre>
</html>
`

func LoadAsciiArtT(name string) (*template.Template, error) {
	tmpl, err := template.New(name).Parse(ascii_artT)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return tmpl, nil
}

func AsciiArt(f io.Reader) (string, error) {
	var ascii_art string

	var grayRamp = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/|()1{}[]?-_+~<>i!lI;:,\"^`\\"
	var rampLength = len(grayRamp)

	imData, imType, err := image.Decode(f)
	if err != nil {
		return ascii_art, err
	}
	log.Println(imType)
	newImage := resize.Resize(80, 0, imData, resize.Lanczos3)

	for y := newImage.Bounds().Min.Y; y < newImage.Bounds().Max.Y; y++ {
		for x := newImage.Bounds().Min.X; x < newImage.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(newImage.At(x, y)).(color.Gray)
			level := (rampLength - 1) * int(c.Y) / 255
			ascii_art = fmt.Sprint(ascii_art + string(grayRamp[level]) + string(grayRamp[level]))
		}
		ascii_art = fmt.Sprint(ascii_art + "\n")
	}
	return ascii_art, nil
}
