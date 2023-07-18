package models

type Compage struct {
	Id int64 `json:"id,omitempty"`

	Password string `json:"password,omitempty"`

	User_name string `json:"userName,omitempty"`
}
