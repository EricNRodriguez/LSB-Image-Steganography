package main

import (
  "fmt"
  "os"
  _ "io/ioutil"
  _ "strconv"
  "image"
  _ "image/png"
)

type Pixel struct {
    R int
    G int
    B int
    A int
}

func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
  return Pixel{ scale(r, 65535, 0, 255, 0),
                scale(g, 65535, 0, 255, 0),
                scale(b, 65535, 0, 255, 0),
                scale(a, 65535, 0, 255, 0),
              }
}

func scale(num uint32, max uint32, min uint32, newMax uint32, newMin uint32) int {
  oldRange := max - min
  newRange := newMax - newMin
  return int((((num - min) * newRange) / oldRange) + newMin)

}

func imageToRGBA(i image.Image) [][]Pixel {
  var imagePixels [][]Pixel
  for r := 0 ; r < i.Bounds().Max.Y ; r++ {
    var row []Pixel
    for c := 0 ; c < i.Bounds().Max.X ; c++ {
      a, b, c, d := i.At(c, r).RGBA()
      row = append(row, rgbaToPixel(a,b,c,d))
    }
    imagePixels = append(imagePixels, row)
  }
  return imagePixels

}


func main() {
  file, err := os.Open("picture.png")
  defer file.Close()

  if err != nil {
    panic(err)
  }

  image, _, err := image.Decode(file)

  if err != nil {
    panic(err)
  }

  imagePixels := imageToRGBA(image)

}
