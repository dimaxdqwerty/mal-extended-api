package models

type Statistics struct {
	Status       Status `json:"status"`
	NumListUsers int    `json:"num_list_users"`
}
