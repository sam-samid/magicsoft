package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type array struct {
	fol1 []string
	fol2 []string
}

type location struct {
	folder1 string
	folder2 string
}

func main() {
	var arr array
	var loc location

	loc.folder1 = "./fol1"
	loc.folder2 = "./fol2"

	files1, err := ioutil.ReadDir(loc.folder1)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files1 {
		arr.fol1 = append(arr.fol1, f.Name())
	}

	files2, err := ioutil.ReadDir(loc.folder2)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files2 {
		arr.fol2 = append(arr.fol2, f.Name())
	}

	for _, value := range arr.fol2 {
		if contains(arr.fol1, value) == false {
			deleteFile(loc.folder2 + "/" + value)
			fmt.Println(loc.folder2 + "/" + value + " DELETED")
		}
	}

	for _, value := range arr.fol1 {
		if contains(arr.fol2, value) == false {
			copyFile(loc.folder1+"/"+value, loc.folder2+"/"+value)
			fmt.Println(loc.folder2 + "/" + value + " NEW")
		} else {
			compareFile(loc.folder1+"/"+value, loc.folder2+"/"+value)
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

func deleteFile(path string) {
	var err = os.Remove(path)
	if err != nil {
		return
	}
}

func copyFile(path1, path2 string) {
	from, err := os.Open(path1)
	if err != nil {
		log.Fatal(err)
	}
	defer from.Close()

	to, err := os.OpenFile(path2, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer to.Close()

	_, err = io.Copy(to, from)
	if err != nil {
		log.Fatal(err)
	}
}

func compareFile(path1, path2 string) {
	f1, err1 := ioutil.ReadFile(path1)

	if err1 != nil {
		log.Fatal(err1)
	}

	f2, err2 := ioutil.ReadFile(path2)

	if err2 != nil {
		log.Fatal(err2)
	}

	if bytes.Equal(f1, f2) == false {
		deleteFile(path2)
		copyFile(path1, path2)
		fmt.Println(path2 + " MODIFIED")
	}
}
