package bufferDemo

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"testing"
)

func TestPrintBuffer(t *testing.T) {
	fmt.Println("hh")

	b := []byte("good")
	bu, _ := json.Marshal(b)
	fmt.Printf("%s\n", string(bu))

	s := fmt.Sprintf("%s", string(bu))
	fmt.Println(s)
	r, _ := base64.StdEncoding.DecodeString(s)
	fmt.Println(r)
	fmt.Println(string(r))
}

// "Z29vZA=="

// As per the docs: https://golang.org/pkg/encoding/json/#Marshal

// Array and slice values encode as JSON arrays, except that []byte encodes as a base64-encoded string, and a nil slice encodes as the null JSON object.

// The value AAAAAQID is a base64 representation of your byte slice - e.g.

// b, err := base64.StdEncoding.DecodeString("AAAAAQID")
// if err != nil {
//     log.Fatal(err)
// }

// fmt.Printf("%v", b)
// // Outputs: [0 0 0 1 2 3]
