package model

type Name struct {
	Common string `json:"common"`
}

type Country struct {
	Name   Name   `json:"name"`
	Region string `json:"region"`
}
