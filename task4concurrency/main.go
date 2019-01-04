package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	go get_data_museum()
}

func get_json(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return body
}

func get_data_museum() {
	res := map[string][]string{}
	var a []string
	json.Unmarshal(get_json("http://xxxxxx"), &res)
	for k, v := range res {
		files1, err := ioutil.ReadDir("./")
		if err != nil {
			log.Fatal(err)
		}

		for _, f := range files1 {
			a = append(a, f.Name())
		}

		if contains(a, res.Kota+".csv") == true {
			writer := csv.NewWriter(res.Kota + ".csv")
			defer writer.Flush()

			for _, value := range res {
				err := writer.Write(value)
				if err != nil {
					log.Println("Cannot write to file", err)
				}
			}
		} else {
			file, err := os.Create(res.Kota + ".csv")
			if err != nil {
				log.Println("Cannot create file", err)
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

			for _, value := range res {
				err := writer.Write(value)
				if err != nil {
					log.Println("Cannot write to file", err)
				}
			}
		}

	}
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
