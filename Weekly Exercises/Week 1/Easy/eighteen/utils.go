package main

type User struct {
	Name string
	Age  int
}

func (u User) SetCopyname(name string) User {
	u.Name = name
	return u
}

func (u *User) SetName(name string) {
	u.Name = name
}
