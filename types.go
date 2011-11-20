// twitterlib - A simple, fully oauth-authenticated Twitter library

// Copyright (c) 2011, Roberto Teixeira <r@robteix.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package facebooklib

// User object. See http://developers.facebook.com/docs/reference/api/user/
type User struct {
	Id string
	Name string
	First_name string
	Middle_name string
	Last_name string
	Gender string
	Locale string
	Link string
	Username string
	Third_party_id string
	Client *FacebookClient
}
