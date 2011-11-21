package facebooklib

import (
	"os"
	"url"
	)

var (
	UserNotConnectedError = os.NewError("User Not Connected")
	)

// Useful if all you want is to have an Id to perform actions
// This prevents having to query a user object from FB.
func NewUser(Id string, fc *FacebookClient) *User {
	return &User{Id: Id, Client: fc}
}

func (user *User) PostStatus(message string) os.Error {
	if user.Client == nil {
		return UserNotConnectedError
	}
	if message == "" {
		return os.NewError("Missing message parameter")
	}
	u := make(url.Values)
	u.Set("message", message)
	_, err := user.Client.Call("POST", user.Id + "/feed", u)
	return err
}

func (user *User) PostLink(link, message string) os.Error {
	if user.Client == nil {
		return UserNotConnectedError
	}
	if link == "" {
		return os.NewError("Missing link parameter")
	}
	u := make(url.Values)
	if message != "" {
		u.Set("message", message)
	}
	u.Set("link", link)
	_, err := user.Client.Call("POST", user.Id + "/links", u)
	return err
}