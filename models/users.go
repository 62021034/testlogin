package models

type Users struct {
	UsersId       string `json:"usersID"`
	LoginEmail    string `json:"loginEmail"`
	LoginPassword string `json:"loginPassword"`
	UsersName     string `json:"userName"`
	UserSurname   string `json:"userSurname"`
}

type Userses struct{
	Userses []Users
}