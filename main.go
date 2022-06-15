package main

import (
	"fmt"

	client "./client"
)

func main() {

	client := client.NewClient("Alex G", 42, client.Avatar{"https://loremflickr.com/640/480/abstract", 330})
	fmt.Println(client.HasAvatar())
	fmt.Printf("%#v\n", client)
	client.UpdateAvatar()
	fmt.Printf("%#v\n", client)

	// i := new(int64)
	// _ = i

	// client := Client{}
	// fmt.Printf()
	// fmt.Println(client)
	// updateAvatar(&client)
	// fmt.Println(client)
	// updateClient(&client)
	// fmt.Println(client)

}

// func updateAvatar(client *Client) {
// 	client.IMG.URL = "https://loremflickr.com/640/480/abstract"
// }

// func updateClient(client *Client) {
// 	client.Name = "Alexey G"
// }
