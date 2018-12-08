package main

import (
  "fmt"
  // "image"
  "os"
  // "io/ioutil"
  // "strconv"
  "image"
  _ "image/png"
)

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

  bounds := image.Bounds()
  wid, hei := bounds.Max.X, bounds.Max.Y


  var imagePixels [][][]uint32
  for r := 0 ; r < hei ; r++ {
    var row [][]uint32
    for c := 0 ; c < wid ; c++ {
      a, b, c, d := image.At(c, r).RGBA()
      e := []uint32{a,b,c,d}
      row = append(row, e)
    }
    imagePixels = append(imagePixels, row)
  }

  for _, row := range imagePixels {
    fmt.Println(row)
  }
}
  func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
      return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
  }

  // Pixel struct example
  type Pixel struct {
      R int
      G int
      B int
      A int
  }
  // fb, err := ioutil.ReadAll(file)

  // if err != nil {
  //   panic(err)
  // }
  //
  // img,d,err := image.Decode(file)
  // fmt.Println(d, err)
  // bounds := img.Bounds()
  // fmt.Println(bounds)
  // width, height := bounds.Max.X, bounds.Max.Y
  // fmt.Println(width, height)

