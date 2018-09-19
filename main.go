package main

import (
	"image"
	"log"

	"periph.io/x/periph/conn/i2c/i2creg"
	"periph.io/x/periph/devices/ssd1306"
	"periph.io/x/periph/host"

	"github.com/fogleman/gg"
)

func main() {
	// Make sure periph is initialized.
	if _, err := host.Init(); err != nil {
		log.Fatal(err)
	}

	// Use i2creg I²C bus registry to find the first available I²C bus.
	b, err := i2creg.Open("")
	if err != nil {
		log.Fatal(err)
	}
	defer b.Close()

	opts := ssd1306.Opts{
		W: 128,
		H: 32,
		Sequential: true,
	}

	dev, err := ssd1306.NewI2C(b, &opts)
	if err != nil {
		log.Fatalf("failed to initialize ssd1306: %v", err)
	}

	top := 8.0
	ctx := gg.NewContext(dev.Bounds().Dx(), dev.Bounds().Dy())
	if err := ctx.LoadFontFace("./start.ttf", 8); err != nil {
		log.Fatalf("error loading font: %v", err)
	}
	ctx.SetRGB(0, 0, 0)
	ctx.Clear()
	ctx.SetRGB(1, 1, 1)
	ctx.DrawString("15:04:01", 0, top)
	ctx.DrawString("15:04:02", 0, top+8)
	ctx.DrawString("15:04:03", 0, top+16)
	ctx.DrawString("15:04:04", 0, top+24)
	dev.Draw(dev.Bounds(), ctx.Image(), image.Point{})

}