package main

import (
	"fmt"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println("hi")

	giovani := &Person{
		Name: "Giovani",
		Age:  25,
		SocialFollowers: &SocialFollowers{
			Twitter: 200,
			Youtube: 100,
		},
	}

	data, err := proto.Marshal(giovani)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	newGio := &Person{}
	err = proto.Unmarshal(data, newGio)
	if err != nil {
		log.Fatal("unmarshal", err)
	}

	fmt.Println(newGio.GetAge())
	fmt.Println(newGio.GetName())
	fmt.Println(newGio.SocialFollowers.Twitter)
	fmt.Println(newGio.SocialFollowers.Youtube)

}
