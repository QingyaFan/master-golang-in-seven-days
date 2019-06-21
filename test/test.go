package main

import (
	"errors"
	"fmt"
)

func GetAge() (age int, err error) {
	return -1, errors.New("get age failed, error is not null")
}

func main() {
	_, err := GetAge()
	if err != nil {
		fmt.Println(err)
	}
}
