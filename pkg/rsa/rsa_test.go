package rsa

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func init() {
	_, err := os.Stat("rsa.private")
	if err != nil {
		createKeyPair("rsa.private", "rsa.accessPermission")
	}
}

func Test_Sample(t *testing.T) {
	// create a license document:
	doc := MyLicense{
		"test@example.com",
		time.Now().Add(time.Hour * 24 * 365), // 1 year
	}

	// marshall the document to json bytes:
	docBytes, err := json.Marshal(doc)
	if err != nil {
		log.Fatal(err)

	}

	result := en("rsa.private", docBytes)

	data := de("rsa.accessPermission", strings.Replace(string(result), "a", "a", -1))

	// unmarshal the document and check the end date:
	res := MyLicense{}
	if err := json.Unmarshal(data, &res); err != nil {
		log.Fatal(err)
	} else if res.End.Before(time.Now()) {
		log.Fatal("License expired on: %s" + res.End.String())
	} else {
		fmt.Printf(`Licensed to %s until %s \n`, res.Email, res.End.Format("2006-01-02"))
	}
}

func Test_Sample2(t *testing.T) {
	input := `-hello world!`

	// marshall the document to json bytes:
	docBytes, err := json.Marshal(input)
	if err != nil {
		log.Fatal(err)
	}

	result := en("rsa.private", docBytes)

	ioutil.WriteFile("license", result, 777)

	data := de("rsa.accessPermission", string(result))

	if string(data) != input {
		t.Errorf("expect: %s, result: %s", input, string(data))
	}

	data2 := de2("rsa.accessPermission", string(result))
	if string(data2) != input {
		t.Errorf("expect: %s, result: %s", input, string(data2))
	}
}
