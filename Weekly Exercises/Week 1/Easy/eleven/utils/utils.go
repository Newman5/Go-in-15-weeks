package utils

type User struct {
	Name string
	Age  int
}

func StructChange(user *User, name string, age int) {
	user.Name = name
	user.Age = age
}

func StructCopyChange(user User, name string, age int) User {
	user.Name = name
	user.Age = age
	return user
}
