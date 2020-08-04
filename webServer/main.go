package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

var templ = template.Must(template.New("qr").Parse(templateStr))

func main() {
	flag.Parse()

	http.Handle("/", http.HandlerFunc(ReadData))
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func ReadData(w http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadFile("./data.pb")
	if err != nil {
		w.Write([]byte("read data.pb error"))
	} else {
		w.Write(buf)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))
}

const templateStr = `
<html>
	<head>
		<title>QR Link Generator</title>
	</head>
	<body>
		{{if .}}
		<img src="http://chart.apis.google.com/chart?chs=300*300&cht=qr&choe=UTF-8&chl={{.}}" />
		<br>
		{{.}}
		<br>
		<br>
		{{end}}
		<form action="/" name=f method="GET">
			<input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
			<input type=submit value="Show QR" name=qr>
		</form>
	</body>
</html>
`
