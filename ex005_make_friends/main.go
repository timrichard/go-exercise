package main

import (
	"fmt"
)

var flist = make(map[string]map[string]bool)


func listFriends(name string) {
	for friendName := range flist[name] {
		fmt.Println(friendName)
	}
}

func linkFriends(nameOne, nameTwo string) {
	personFriendList, defined := flist[nameOne]

	if (!defined) {
		personFriendList = make(map[string]bool)
		flist[nameOne] = personFriendList
	}

	personFriendList[nameTwo] = true
}

func makeFriends(nameOne, nameTwo string) {
	linkFriends(nameOne, nameTwo)
	linkFriends(nameTwo, nameOne)
}

func main() {
	makeFriends("userEight", "userTwo")
	makeFriends("userEight", "userThree")
	listFriends("userEight")
	listFriends("userThree")
}