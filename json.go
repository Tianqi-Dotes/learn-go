package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Name string `json:"name"`
	Year string `json:"year"`
	Price int	`json:"price"`
	Actors []string `json:"actors"`
}

func mainccc() {
	movie:=Movie{"pianolist","1999",20,[]string{"albert","shubenhua"}}
	str,err := json.Marshal(movie)
	if err!=nil{
		fmt.Println("error!")
	}
	fmt.Printf("jsonstr: %s\n",str)

}