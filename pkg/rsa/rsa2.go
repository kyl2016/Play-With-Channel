package rsa

import (
	"fmt"
	"io/ioutil"

	"github.com/farmerx/gorsa"
)

func de2(publicKeyFile, license string) string {
	buffer, err := ioutil.ReadFile("rsa.accessPermission")
	if err != nil {
		panic(err)
	}

	err = gorsa.RSA.SetPublicKey(string(buffer))
	if err != nil {
		panic(err)
	}

	r, err := gorsa.RSA.PubKeyDECRYPT([]byte(license))
	if err != nil {
		panic(err)
	}

	fmt.Println(string(r))

	return string(r)
}
