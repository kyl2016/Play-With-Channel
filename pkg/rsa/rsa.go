package rsa

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/hyperboloide/lk"
)

// MyLicense struct
type MyLicense struct {
	Email string
	End   time.Time
}

func en(privateKeyFile string, data []byte) []byte {
	buffer, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := lk.PrivateKeyFromB64String(string(buffer))
	if err != nil {
		panic(err)
	}

	// generate your license with the private key and the document:
	license, err := lk.NewLicense(privateKey, data)
	if err != nil {
		log.Fatal(err)

	}

	// encode the new license to b64, this is what you give to your customer.
	str64, err := license.ToB64String()
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(str64)

	return []byte(str64)
}

func de(publicKeyFile, licenseStr string) []byte {
	publicKeyStr, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		panic(err)
	}

	// get the accessPermission key. The accessPermission key should be hardcoded in your app to check licences.
	// Do not distribute the private key!
	publicKey, err := lk.PublicKeyFromB64String(string(publicKeyStr))
	if err != nil {
		panic(err)
	}

	license, err := lk.LicenseFromB64String(licenseStr)
	if err != nil {
		panic(err)
	}

	// validate the license:
	if ok, err := license.Verify(publicKey); err != nil {
		log.Fatal(err)
	} else if !ok {
		log.Fatal("Invalid license signature")
	}

	return license.Data
}

func createKeyPair(privateKeyFile, publicKeyFile string) {
	// create a new Private key:
	privateKey, err := lk.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyStr, err := privateKey.ToB64String()
	if err != nil {
		log.Fatal(err)
	}

	ioutil.WriteFile(privateKeyFile, []byte(privateKeyStr), 777)

	str := privateKey.GetPublicKey().ToB64String()
	ioutil.WriteFile(publicKeyFile, []byte(str), 777)
}
