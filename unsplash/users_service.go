// Copyright (c) 2017 Hardik Bagdi <hbagdi1@binghamton.edu>
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unsplash

import (
	"context"
	"encoding/json"
	"fmt"
)

// ProfileImageOpt denotes properties of any Image
type ProfileImageOpt struct {
	Height int `json:"h,omitempty" url:"h"`
	Width  int `json:"w,omitempty" url:"w"`
}

// UsersService interacts with /users endpoint
type UsersService service

// User returns a User with username and optional profile image size ImageOpt
func (us *UsersService) User(ctx context.Context, username string, imageOpt *ProfileImageOpt) (*User, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, imageOpt, nil)
	if err != nil {
		return nil, err
	}
	resp, err := us.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	var user User
	err = json.Unmarshal(*resp.body, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

type urlWrapper struct {
	URL *URL `json:"url"`
}

// Portfolio returns a User with username and optional profile image size ImageOpt
func (us *UsersService) Portfolio(ctx context.Context, username string) (*URL, error) {
	if "" == username {
		return nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	endpoint := fmt.Sprintf("%v/%v/portfolio", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, nil, nil)
	if err != nil {
		return nil, err
	}
	resp, err := us.client.do(ctx, req)
	if err != nil {
		return nil, err
	}
	var portfolio urlWrapper
	err = json.Unmarshal(*resp.body, &portfolio)
	if err != nil {
		return nil, err
	}
	return portfolio.URL, nil
}

// Photos return an array of photos uploaded by the user.
func (us *UsersService) Photos(ctx context.Context, username string, opt *ListOpt) (*[]Photo, *Response, error) {
	if "" == username {
		return nil, nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	s := (service)(*us)
	endpoint := fmt.Sprintf("%v/%v/%v", getEndpoint(users), username, getEndpoint(photos))
	return s.getPhotos(ctx, opt, endpoint)
}

// LikedPhotos return an array of liked photos
func (us *UsersService) LikedPhotos(ctx context.Context, username string, opt *ListOpt) (*[]Photo, *Response, error) {
	if "" == username {
		return nil, nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	s := (service)(*us)
	endpoint := fmt.Sprintf("%v/%v/%v", getEndpoint(users), username, "likes")
	return s.getPhotos(ctx, opt, endpoint)
}

// Collections return an array of user's collections.
func (us *UsersService) Collections(ctx context.Context, username string, opt *ListOpt) (*[]Collection, *Response, error) {
	if "" == username {
		return nil, nil, &IllegalArgumentError{ErrString: "Username cannot be null"}
	}
	s := (service)(*us)
	endpoint := fmt.Sprintf("%v/%v/%v", getEndpoint(users), username, getEndpoint(collections))
	return s.getCollections(ctx, opt, endpoint)
}

// Statistics return a stats about a photo with id.
func (us *UsersService) Statistics(ctx context.Context, username string, opt *StatsOpt) (*UserStatistics, *Response, error) {
	if "" == username {
		return nil, nil, &IllegalArgumentError{ErrString: "Photo ID cannot be null"}
	}
	if opt == nil {
		opt = defaultStatsOpt
	}
	if !opt.Valid() {
		return nil, nil, &InvalidStatsOptError{ErrString: "opt provided is not valid."}
	}
	endpoint := fmt.Sprintf("%v/%v/statistics", getEndpoint(users), username)
	req, err := newRequest(GET, endpoint, opt, nil)
	if err != nil {
		return nil, nil, err
	}
	resp, err := us.client.do(ctx, req)
	if err != nil {
		return nil, nil, err
	}
	var stats UserStatistics
	err = json.Unmarshal(*resp.body, &stats)
	if err != nil {
		return nil, nil, err
	}
	return &stats, resp, nil
}
