// fblib - A simple, fully oauth-authenticated Twitter library

// Copyright 2011 The Fblib Authors.  All rights reserved.
// Use of this source code is governed by the Simplified BSD
// license that can be found in the LICENSE file.

package fblib

// User object. See http://developers.facebook.com/docs/reference/api/user/
type User struct {
	Id             string
	Name           string
	First_name     string
	Middle_name    string
	Last_name      string
	Gender         string
	Locale         string
	Languages      []ObjectRef
	Link           string
	Username       string
	Third_party_id string
	Timezone       int64
	Update_time    string
	Verified       bool
	Bio            string
	Birthday       string // MM/DD/YYYY
	//Education ODO
	Email               string
	Hometown            string
	Interested_in       []string
	Location            ObjectRef
	Political           string
	Favorite_athletes   []ObjectRef
	Quotes              string
	Relationship_status string
	Religion            string
	Significant_other   ObjectRef
	// Video_upload_limits
	Website string
	//Work

	Client *FacebookClient
}

// This type is used by various facebook objects to
// represent a user ID. The reason for having both
// Id and Name is to avoid ambiguity in cases where
// a user name is a valid Id
type ObjectRef struct {
	Id   string
	Name string
}

// Link represents a link to be posted
type Link struct {
	Text  string // Additional text added to the post
	Image string // URL for the Image to be used
	Url   string // URL to be shared
}

// Photo represents a photo resource to be shared
type Photo struct {
	Source   []byte // the photo data to be uploaded
	FileName string // name of the photo file to be uploaded (e.g. "photo.png")
	Message  string // some descriptive text
}
