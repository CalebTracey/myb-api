package models

type User struct {
	FirstName string `json:"FirstName,omitempty"`
	LastName  string `json:"LastName,omitempty"`
	Username  string `json:"Username,omitempty"`
	Password  string `json:"Password,omitempty"`
	Email     string `json:"Email,omitempty"`
	Addr1     string `json:"Addr1,omitempty"`
	Addr2     string `json:"Addr2,omitempty"`
}
