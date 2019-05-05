package main

import (
  "os"
  "image"
  "image/png"
  "image/color"
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

func encodeMessage(s string) []byte {
  newMessage := []byte{}
  for _, b := range []byte(s) {
    newMessage = append(newMessage, b >> 7 & 1)
    newMessage = append(newMessage, b >> 6 & 1)
    newMessage = append(newMessage, b >> 5 & 1)
    newMessage = append(newMessage, b >> 4 & 1)
    newMessage = append(newMessage, b >> 3 & 1)
    newMessage = append(newMessage, b >> 2 & 1)
    newMessage = append(newMessage, b >> 1 & 1)
    newMessage = append(newMessage, b & 1)
  }
  return newMessage
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

func createImage(p [][]Pixel) {
  //creates file in current directory
  outputFile, err := os.Create("encodedPicture.png")
  if err != nil {
    panic(err)
  }
  //returns a new rgba image with the given dimensions
  outputImage := image.NewRGBA(image.Rectangle{image.Point{0,0},image.Point{len(p[0]), len(p)}})


  for y := 0 ; y < len(p) ; y++ {
    for x := 0 ; x < len(p[y]) ; x++ {
      outputImage.Set(x, y, color.RGBA{uint8(p[y][x].R),
                                       uint8(p[y][x].G),
                                       uint8(p[y][x].B),
                                       uint8(p[y][x].A)})
    }
  }

  png.Encode(outputFile, outputImage)
}

func main() {
  file, err := os.Open("picture.png")
  defer file.Close()

  if err != nil {
    panic(err)
  }

  ima, _, err := image.Decode(file)

  if err != nil {
    panic(err)
  }

  imagePixelsEncoded := encodeImage(encodeMessage("eric"), imageToRGBA(ima))
  createImage(imagePixelsEncoded)

}
