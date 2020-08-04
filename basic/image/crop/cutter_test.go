package crop

import (
	"fmt"
	"github.com/oliamb/cutter"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func getImage() image.Image {
	return image.NewGray(image.Rect(0, 0, 1600, 1437))
}

type Source struct {
	Rows      int
	Cols      int
	imagePath string
}

func TestCrop(t *testing.T) {
	root := "/Users/kang/Downloads/动物"
	dir, err := os.Stat(root)
	if err != nil {
		panic(err)
	}

	if dir.IsDir() {
		files, err := ioutil.ReadDir(root)
		if err != nil {
			panic(err)
		}

		for _, file := range files {
			if file.IsDir() {
				continue
			}

			ext := path.Ext(file.Name())
			if !strings.Contains(".jpg.jpeg.png.gif", ext) {
				continue
			}

			fileName := strings.TrimSuffix(file.Name(), ext)
			Crop(3, 3, path.Join(root, file.Name()), path.Join(root, fileName+"_3_3"))
			Crop(3, 4, path.Join(root, file.Name()), path.Join(root, fileName+"_3_4"))
			Crop(4, 4, path.Join(root, file.Name()), path.Join(root, fileName+"_4_4"))
		}
	}
}

func Crop(rows, cols int, file string, saveDir string) {
	err := os.Mkdir(saveDir, os.ModePerm)
	if err != nil {
		return
	}

	src, err := os.Open(file)

	if err != nil {
		panic(err)
	}
	img, _, err := image.Decode(src)
	if err != nil {
		//panic(err)
		fmt.Println("error:", file, err.Error())
		return
	}

	rect := img.Bounds()
	x := rect.Size().X
	y := rect.Size().Y

	width := x / cols
	height := y / rows

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			c := cutter.Config{
				Width:  width,
				Height: height,
				Anchor: image.Point{col * width, row * height},
			}
			if row == rows-1 {
				c.Height = y - row*height
			}
			if col == cols-1 {
				c.Width = x - col*width
			}

			//fmt.Printf("%+v\n", c)
			r, err := cutter.Crop(img, c)
			if err != nil {
				panic(err)
			}
			// save r to file
			f, err := os.Create(fmt.Sprintf("%s/%d%d.png", saveDir, row+1, col+1))
			if err != nil {
				panic(err)
			}
			defer f.Close()
			png.Encode(f, r)
		}
	}
}
