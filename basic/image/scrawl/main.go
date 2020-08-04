package main

import (
	"encoding/json"
	"fmt"
	"github.com/badoux/goscraper"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"path"
	"time"
)

var metadataFile = "images.json"

type Metadata struct {
	Title   string
	Name    string
	Desc    string
	File    string
	ImgUrl  string
	WebPage string
}

func main() {
	var metadatas map[string]Metadata
	buffer, err := ioutil.ReadFile(metadataFile)
	if err == nil {
		err = json.Unmarshal(buffer, &metadatas)
		if err != nil {
			panic(err)
		}
	} else {
		metadatas = make(map[string]Metadata)
	}

	// 800 - 27182 已下载

	for i := 25000; i < 28000; i++ {
		time.Sleep(time.Second * time.Duration(rand.Int31n(10)))

		fmt.Println(time.Now(), i)

		s, err := goscraper.Scrape(fmt.Sprintf("http://www.netbian.com/desk/%d-1920x1080.htm", i), 5)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//fmt.Printf("Icon : %s\n", s.Preview.Icon)
		//fmt.Printf("Name : %s\n", s.Preview.Name)
		//fmt.Printf("Title : %s\n", s.Preview.Title)
		//fmt.Printf("Description : %s\n", s.Preview.Description)
		for _, i := range s.Preview.Images {
			_, ok := metadatas[i]
			if !ok {
				fileName, err := download(i)
				if err == nil {
					metadatas[i] = Metadata{
						Title:   s.Preview.Title,
						Name:    s.Preview.Name,
						Desc:    s.Preview.Description,
						File:    fileName,
						ImgUrl:  i,
						WebPage: s.Preview.Link,
					}

					buff, _ := json.Marshal(metadatas)
					ioutil.WriteFile(metadataFile, buff, os.ModePerm)
				}
			}
		}
		//fmt.Printf("Url : %s\n", s.Preview.Link)
	}
}

func download(url string) (string, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.162 Safari/537.36")

	resp, err := http.DefaultClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return "", err
	}

	fileName := path.Join("../images", path.Base(url))
	//file, err := os.Create(fileName)
	//if err != nil {
	//	panic(err)
	//}

	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile(fileName, buffer, os.ModePerm|os.ModeAppend)

	return fileName, nil
}
