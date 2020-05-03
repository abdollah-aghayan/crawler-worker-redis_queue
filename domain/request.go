package domain

import "crawler/utils/errorh"

//Request struct
type Request struct {
	Username string   `json:"username"`
	Urls     []string `urls`
}

// ValidateRequest check weather requested data for login is ok or not
func (u *Request) ValidateRequest() *errorh.Errorh {
	// check username
	if len(u.Username) == 0 {
		return errorh.BadRequestError("Username can not be empty")
	}

	if len(u.Urls) == 0 {
		return errorh.BadRequestError("urls can not be empty.please pass list of valid urls")

	}

	return nil
}
