# OLED control for Adafruit PiOLED

A command line tool to control Adafruit [PiOLED](https://www.adafruit.com/product/3527).

## Installation

1. If you have not do it already, configure I2C on your Raspberry Pi (https://learn.adafruit.com/adafruits-raspberry-pi-lesson-4-gpio-setup/configuring-i2c)
2. Download the program at [Release]()
3. Transfer it into your Raspberry Pi and copy it to `/usr/bin` or any path under `$PATH`Adafruit PiOLED - 128x32 Mini OLED for Raspberry Pi
4. Install a font that display good with 8pt size. I recommend [Press Start](https://www.dafont.com/press-start.font).

## Usage

```
sudo oled ./start.ttf "Adafruit" "PiOLED" "128x32" "For RaspberryPi"
```

![Preview](./img/preview.JPG)
