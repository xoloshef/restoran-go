package model

type User struct {
	Id      int
	Name    string
	Surname string
}

func GetAllUsers() (users []User, err error) {
	users = []User{
		{1, "Клиент", "Клиенттыч"},
		{2, "Официант", "Официантович"},
		{3, "Повар", "Поваренкич"},
		{4, "Админ", "Админыч"},
		{5, "Райан", "Гослинг"},
	}
	return
}
