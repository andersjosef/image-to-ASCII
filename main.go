package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Import for JPEG decoding
	_ "image/png"  // Import for PNG decoding
	"os"
	"strconv"
)

// Characters from dim to bright
const signs = " .-_,^+r/?sJ7(Fi1tloxj2ShdbUHm0MQ%&@"

var step int = 2 // Step for looping over pixels

func main() {
  fileSRC, fileDST, err := getArguments()
  if err != nil {
    fmt.Println(err)
    return
  }

  // Open the file
  file, err := os.Open("./" + fileSRC)
  if err != nil {
    fmt.Println("Error: ", err)
    return
  }
  defer file.Close()

  // Decode the image
  img, format, err := image.Decode(file)
  if err != nil {
    fmt.Println("Error decoding image: ", err)
    return
  }

  bounds := img.Bounds()
  width := bounds.Dx()
  height := bounds.Dy()

  // print format of pic
  fmt.Println("Image format", format)
  fmt.Printf("Image dimensions: %dx%d\n", width, height)
  fmt.Println("Creating...")

  // Main creation loop
  ascii := createASCII(img, height, width)

  // Writing to file
  err = os.WriteFile("./" + fileDST, ascii, 0644)
  if err != nil {
    fmt.Println("Error writing:", err)
    return
  }
  fmt.Println("Done!")
  fmt.Printf("number of pixels in total: %d\n", width*height)

}

// Function for looping over each pixel (or per step)
// making a byte slice of the characters
func createASCII(img image.Image, height, width int) []byte {
  // ASCII variable
  ascii := []byte{}

  // loop over each pixel
  for y:=0; y<height;y+=step {
    for x:=0; x<width;x+=step {
      
      color := img.At(x, y)
      r, g, b, _ := color.RGBA()
      sum := int(((r+g+b)/3 * 255) / 65535) // Change to 0-255 span
      ascii = append(ascii, getASCIICharacter(sum))
    }
      // new line
      ascii = append(ascii, "\n"...)
  }

  return ascii
}

// Choosing ASCII character best suited for the brigtness
func getASCIICharacter(avg int) byte {
  // integer for the closest sign in the const 
  integer := (avg *(len(signs)-1)) / 255
  return signs[integer]
}


// Handeling the arguments for sourcefile destination file and stepsize
func getArguments() (fileSRC, fileDST string, err error){
  if (len(os.Args) < 3 || len(os.Args) > 4) {
    return "", "", fmt.Errorf("Usage: ./program /source/path /destination/path stepnumber (stepnumber optional)")
  }
  if len(os.Args) == 4 {
    step, err = strconv.Atoi(os.Args[3])
    if err != nil {
    return "", "", fmt.Errorf("Did not get int stepnumer")
    }
  }
  fileSRC = os.Args[1]
  fileDST = os.Args[2]

  return fileSRC, fileDST, nil
}
