package main

import (
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strings"
)

func main() {
	http.HandleFunc("/upload", uploadHandle)
	http.HandleFunc("/uploaded/", showPicHandle)
	http.HandleFunc("/sample", createSampleHandle)
	err := http.ListenAndServe(":10000", nil)
	fmt.Println(err)
}

func uploadHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
	} else {
		uploadFile, handle, err := req.FormFile("image")
		errHandle(err, w)
		defer uploadFile.Close()

		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" {
			errHandle(errors.New("Only spport .jpg image"), w)
			return
		}

		os.Mkdir("./uploaded/", 0777)
		saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		errHandle(err, w)
		defer saveFile.Close()

		io.Copy(saveFile, uploadFile)

		w.Write([]byte("upload image: <a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a>"))
	}
}

func showPicHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errHandle(err, w)
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	errHandle(err, w)

	w.Write(buf)
}

func errHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	}
}

const html = `
<html>
	<head></head>
	<body>
		<form method="post" enctype="multipart/form-data">
			<input type="file" name="image" />
			<input type="submit" />
		</form>
	</body>
</html>
`

func createSampleHandle(w http.ResponseWriter, req *http.Request) {
	pixels := make([]byte, 100*100)
	img := image.NewGray(image.Rect(0, 0, 100, 100))
	img.Pix = pixels
}
