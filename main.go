package main

import (
	"encoding/json"
	"fmt"
	"io"

	"net/http"

	"github.com/emanuen/countries_info/model"
)

func main() {

	resp, err := http.Get("https://restcountries.com/v3.1/all")

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	bytes, err2 := io.ReadAll(resp.Body)

	if err2 != nil {
		panic(err2)
	}

	var countries []model.Country

	err3 := json.Unmarshal(bytes, &countries)

	if err3 != nil {
		panic(err3)
	}

	fmt.Println(countries)

}
