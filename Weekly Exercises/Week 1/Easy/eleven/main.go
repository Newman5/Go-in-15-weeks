package main

import (
	"Easy/eleven/utils"
	"fmt"
)

func main() {
	var user utils.User
	var user1 utils.User
	utils.StructChange(&user, "John", 20)
	user1 = utils.StructCopyChange(user, "Jane", 21)
	fmt.Println(user)
	fmt.Println(user1)
	fmt.Println(user)

}
