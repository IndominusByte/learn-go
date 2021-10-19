package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (userSlice UserSlice) Len() int           { return len(userSlice) }
func (userSlice UserSlice) Swap(i, j int)      { userSlice[i], userSlice[j] = userSlice[j], userSlice[i] }
func (userSlice UserSlice) Less(i, j int) bool { return userSlice[i].Age < userSlice[j].Age }
func (userSlice UserSlice) Debug() {
	fmt.Println(userSlice)
}

func main() {
	users := UserSlice{
		{"oman", 20},
		{"pradipta", 21},
		{"dewantara", 14},
		{"indah", 19},
	}

	users.Debug()
	fmt.Println("=============")
	fmt.Println(users)
	sort.Sort(users)
	fmt.Println(users)
}
