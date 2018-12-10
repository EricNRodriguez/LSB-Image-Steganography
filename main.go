package main

import (
  "fmt"
  "os"
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

func encodeMessage(s string, a int, b int) []byte {
  d := []byte{}
  for _, b := range []byte(s) {
    d = append(d, b >> 7 & 1)
    d = append(d, b >> 6 & 1)
    d = append(d, b >> 5 & 1)
    d = append(d, b >> 4 & 1)
    d = append(d, b >> 3 & 1)
    d = append(d, b >> 2 & 1)
    d = append(d, b >> 1 & 1)
    d = append(d, b & 1)
  }
  return d
}


func encodeImage(message []byte, i [][]Pixel) [][]Pixel {
  index := 0
  for a := 0 ; a < len(i) ; a++ {
    for b := 0 ; b < len(i[a]) ; b++ {
      if index >= len(message) { break }
      i[a][b].R = encodePixel(i[a][b].R, message[index])
      index++
      if index >= len(message) { break }
      i[a][b].G = encodePixel(i[a][b].G, message[index])
      index++
      if index >= len(message) { break }
      i[a][b].B = encodePixel(i[a][b].B, message[index])
      index++
    }
  }

  return i
}

func encodePixel(colourValue int, LSB byte) int {
  if LSB == 0  && colourValue % 2 == 1 {
    return colourValue & (colourValue - 1)
  } else if LSB == 1 && colourValue % 2 == 0 {
    return (colourValue & (colourValue - 1)) + 1
  } else {
    return colourValue
  }
}

// func decodeImage() {
//
// }

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

  e := encodeMessage("eric", len(imagePixels), len(imagePixels[0]))
  encodeImage(e, imagePixels)
  // imagePixelsEncoded := encodeImage(e, imagePixels)
  // fmt.Println(imagePixelsEncoded)

}
