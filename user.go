// Copyright 2013 The Fblib Authors.  All rights reserved.
// Use of this source code is governed by the Simplified BSD
// license that can be found in the LICENSE file.

package fblib

import (
	"errors"
	"net/url"
)

var (
	UserNotConnectedError = errors.New("User Not Connected")
)

// Useful if all you want is to have an Id to perform actions
// This prevents having to query a user object from FB.
func NewUser(Id string, fc *FacebookClient) *User {
	return &User{Id: Id, Client: fc}
}

func (user *User) PostStatus(message string) error {
	if user.Client == nil {
		return UserNotConnectedError
	}
	if message == "" {
		return errors.New("Missing message parameter")
	}
	u := make(url.Values)
	u.Set("message", message)
	_, err := user.Client.Call("POST", user.Id+"/feed", u)
	return err
}

func (user *User) PostLink(link, message string) error {
	if user.Client == nil {
		return UserNotConnectedError
	}
	if link == "" {
		return errors.New("Missing link parameter")
	}
	u := make(url.Values)
	if message != "" {
		u.Set("message", message)
	}
	u.Set("link", link)
	_, err := user.Client.Call("POST", user.Id+"/links", u)
	return err
}
