package demo

import (
	"fmt"

	bingxgo "github.com/WolffunService/bingx-go"
)

func Call() {
	client := bingxgo.NewClient("API_KEY", "SECRET_KEY")
	spot := bingxgo.NewSpotClient(client)
	fmt.Println(spot.GetBalance())
}
