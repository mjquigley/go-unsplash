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
	"net/http"
	"strconv"
	"time"
)

//The following are implementing error interface

// IllegalArgumentError occurs when the argument to a function are
// messed up
type IllegalArgumentError struct {
	ErrString string
}

func (e IllegalArgumentError) Error() string {
	return e.ErrString
}

// JSONUnmarshallingError occurs due to a unmarshalling error
type JSONUnmarshallingError struct {
	ErrString string
}

func (e JSONUnmarshallingError) Error() string {
	return e.ErrString
}

// AuthorizationError occurs for an Unauthorized request
type AuthorizationError struct {
	ErrString string
}

func (e AuthorizationError) Error() string {
	return e.ErrString
}

// NotFoundError occurs when the resource queried returns a 404.
type NotFoundError struct {
	ErrString string
}

func (e NotFoundError) Error() string {
	return e.ErrString
}

// ServerError occurs when the resource queried returns an unknown error.
type ServerError struct {
	ErrString  string
	StatusCode int
	RetryAfter time.Duration
}

func (e ServerError) Error() string {
	return e.ErrString
}

func (e ServerError) Status() int {
	return e.StatusCode
}

func (e ServerError) Header() http.Header {
	if e.RetryAfter == 0 {
		return http.Header{}
	}
	return http.Header{"Retry-After": {strconv.Itoa(int(e.RetryAfter / time.Second))}}
}

// InvalidPhotoOptError occurs when PhotoOpt.Valid() fails.
type InvalidPhotoOptError struct {
	ErrString string
}

func (e InvalidPhotoOptError) Error() string {
	return e.ErrString
}

// InvalidListOptError occurs when ListOpt.Valid() fails.
type InvalidListOptError struct {
	ErrString string
}

func (e InvalidListOptError) Error() string {
	return e.ErrString
}

// InvalidStatsOptError occurs when StatsOpt.Valid() fails.
type InvalidStatsOptError struct {
	ErrString string
}

func (e InvalidStatsOptError) Error() string {
	return e.ErrString
}

// RateLimitError occurs when rate limit is reached for the API key.
type RateLimitError struct {
	ErrString string
}

func (e RateLimitError) Error() string {
	return e.ErrString
}
