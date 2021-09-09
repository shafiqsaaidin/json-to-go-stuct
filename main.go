package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Inbound struct {
	Listen   string `json:"listen"`
	Port     int    `json:"port"`
	Protocol string `json:"protocol"`
}

type Outbound struct {
	Protocol string `json:"protocol"`
	Settings Setting
}

type Setting struct {
	Vnext []Vnext
}

type Vnext struct {
	Address string `json:"address"`
	Port    int    `json:"port"`
	Users   []User
}

type User struct {
	Encryption string `json:"encryption"`
	Flow       string `json:"flow"`
	Id         string `json:"id"`
	Level      int    `json:"level"`
	Security   string `json:"security"`
}

type Config struct {
	Inbounds  []Inbound
	Outbounds []Outbound
}

func main() {
	jsonFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		fmt.Println(err)
	}

	var conf Config

	err1 := json.Unmarshal(jsonFile, &conf)
	if err1 != nil {
		fmt.Println(err1)
	}

	fmt.Println(conf)
}
