package main

import (
	"fmt"
	"image"
	"log"
	"os"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/host"

	"github.com/fogleman/gg"
)

func main() {
	args := os.Args[1:]
	if len(args) != 5 {
		fmt.Print("usage: oled path-to-font line1 line2 line3 line4")
		return
	}
	fontPath := args[0]
	line1 := args[1]
	line2 := args[2]
	line3 := args[3]
	line4 := args[4]

	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
		return
	}

	// Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer b.Close()

	// initialize ssd1306
	opts := ssd1306.Opts{
		W:          128,
		H:          32,
		Sequential: true,
	}
	dev, err := ssd1306.NewI2C(b, &opts)
	if err != nil {
		log.Fatalf("failed to initialize ssd1306: %v", err)
		return
	}

	// Draw message to device
	top := 8.0
	ctx := gg.NewContext(dev.Bounds().Dx(), dev.Bounds().Dy())
	if err := ctx.LoadFontFace(fontPath, 8); err != nil {
		log.Fatalf("error loading font: %v", err)
		return
	}
	ctx.SetRGB(1, 1, 1)
	ctx.DrawString(line1, 0, top)
	ctx.DrawString(line2, 0, top+8)
	ctx.DrawString(line3, 0, top+16)
	ctx.DrawString(line4, 0, top+24)
	dev.Draw(dev.Bounds(), ctx.Image(), image.Point{})
}
