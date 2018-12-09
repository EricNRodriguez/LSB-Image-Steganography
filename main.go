package main

import (
  "fmt"
  "os"
  _ "io/ioutil"
  "image"
  _ "image/png"
  // "math/bits"
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
  d := []byte{}
  for _, b := range []byte(s) {
    d = append(d, b>>7&1)
    d = append(d, b>>6&1)
    d = append(d, b>>5&1)
    d = append(d, b>>4&1)
    d = append(d, b>>3&1)
    d = append(d, b>>2&1)
    d = append(d, b>>1&1)
    d = append(d, b&1)
  }
  return d
}

func encodeImage(message []byte, i [][]Pixel) [][]Pixel {
  index := 0
  for _, row := range i {
    for _, pix := range row {
      pix.R = encodePixel(pix.R, message[index])
      index ++

      pix.G = encodePixel(pix.G, message[index])
      index ++

      pix.B = encodePixel(pix.B, message[index])
      index ++

    }
  }
  return i
}

//works
func encodePixel(colourValue int, LSB byte) int {
  if LSB == 0  && colourValue%2==1{
    // fmt.Println(colourValue&(colourValue-1)%2)
    return colourValue&(colourValue-1)
  } else if LSB == 1 && colourValue%2==0{
    // fmt.Println(((colourValue&(colourValue-1)) + 1)%2)
    return (colourValue&(colourValue-1)) + 1
  } else {
    // fmt.Println(colourValue%2)
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
  // fmt.Println(imagePixels)
  // fmt.Println(encodeMessage("eric"))
  imagePixels = encodeImage(encodeMessage("eric"), imagePixels)
  //it doesnt seem to be iterating throuhg the pixels row by row
  fmt.Println(
  imagePixels[0][0].R%2,
  imagePixels[0][0].G%2,
  imagePixels[0][0].B%2,
  imagePixels[0][1].R%2,
  imagePixels[0][1].G%2,
  imagePixels[0][1].B%2)


}
