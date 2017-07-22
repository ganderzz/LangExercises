package main

import (
	"fmt"
	"os"
	"bytes"
	"image"
	"image/png"
	_ "image/jpeg"
)

func HammingDistance(b1 , b2 []byte) int {
	distance := 0
	largest := b1
	smallest := b2

	// Find the larger of the two byte arrays
	if len(b2) > len(b1) {
		largest = b2
		smallest = b1
	}

	// Get the difference of large to small if they
	// have different lengths
	if len(largest) != len(smallest) {
		distance = len(largest) - len(smallest)
	}

	// Count the differences between the byte arrays
	for i := 0; i < len(smallest); i++ {
		if smallest[i] != largest[i] {
			distance++
		}
	}

	return distance
}

func ReadImage(url string) ([]byte, string) {
	imgFile, err := os.Open(url);
	if err != nil {
		return nil, "Error reading file."
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, "Error decoding file."
	}

	buffer := &bytes.Buffer{}
	png.Encode(buffer, img)

	return buffer.Bytes(), ""
}

func main() {
	pwd, _ := os.Getwd()
	img1, err := ReadImage(pwd + "/out.png")
	img2, _ := ReadImage(pwd + "/img.jpg")

	fmt.Printf("%s", err)

	dist := HammingDistance(img1, img2)

	fmt.Printf("Got a hamming distance of: %d \n", dist)

	if dist > 10 {
		fmt.Printf("The images are nothing alike!")
	} else if dist <= 10 {
		fmt.Printf("These images are pretty similar.")
	} else {
		fmt.Printf("They're the same image!")
	}
}