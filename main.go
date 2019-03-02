package main

import "fmt"

type Features struct {
    Id          string `json:"id"`
    Description string `json:"description"`
    Url         string `json:"url"`
    ImageUrl    string `json:"image_url"`
}

type Action string

const (
    CLICK_ACTION Action = "click"
    VIEW_ACTION  Action = "view"
)

type InteractionRequest struct {
    Id     string `json:"id"`
    Action Action `json:"action"`
}

type Interaction struct {
    View  uint64 `json:"view"`
    Click uint64 `json:"click"`
}

type Ad struct {
    Features    Features           `json:"features"`
    Tags        map[string]float64 `json:"tags"`
    Interaction Interaction        `json:"interaction"`
}


func main() {
	fmt.Println("Hello, world")
}
