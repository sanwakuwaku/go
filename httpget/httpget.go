package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fp, err := os.Open("./links.txt")
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer fp.Close()

	reader := bufio.NewReader(fp)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("%v", err)
		}

		// if !isPrefix {
		// 	log.Fatalf("error readline")
		// }

		url := string(line)
		data, err := httpGet(url)
		if err != nil {
			log.Printf("%v", err)
			continue
		}
		log.Printf("http get success:%s\n", url)

		fragments := strings.Split(url, "/")
		if len(fragments) == 0 {
			log.Fatalf("can't split url:%s", url)
		}

		fileName := fragments[len(fragments)-1]
		if err := saveFile(fileName, data); err != nil {
			log.Fatalf("save error :%v", err)
		}

		log.Printf("save success\n")
	}
}

func httpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func saveFile(fileName string, data []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}
