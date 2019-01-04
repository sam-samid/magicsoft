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

		if contains(a, v[index_nama_kota]+".csv") == true {
			var column []string
			f, err := os.Open(v[index_nama_kota].Kota + ".csv")
			if err != nil {
				log.Fatal(err)
			}
			read := csv.NewReader(f)
			lines, err := read.ReadAll()
			if err != nil {
				log.Fatal(err)
			}
			if err = f.Close(); err != nil {
				log.Fatal(err)
			}

			// add column
			l := len(lines)
			if len(column) < l {
				l = len(column)
			}
			for i := 0; i < l; i++ {
				lines[i] = append(lines[i], column[i])
			}

			writer := csv.NewWriter(v[index_nama_kota] + ".csv")
			defer writer.Flush()

			w := csv.NewWriter(f)
			if err = w.WriteAll(lines); err != nil {
				f.Close()
				log.Fatal(err)
			}

		} else {
			file, err := os.Create(v[index_nama_kota] + ".csv")
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

// konsep penyimpana ini :
// 1. mengambil data dari api
// 2. melakukan pengecekan ada atau tidak file dari kota yang dilooping
// 3. kalau belum ada membuat file csv.nya dan ditulis
// 4. kalau sudah ada melakukang pengeditan, dengan penambahan column
