package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	var a string
	fmt.Println("Введите адрес сайта")
	fmt.Scan(&a)
	res, err := http.Get(a)
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(res.Body)

	stringBody := string(data)
	re := regexp.MustCompile(`<!--.*-->`)
	matches := re.FindAllString(stringBody, -1)
	for _, v := range matches {
		fmt.Println(v)
	}

}
