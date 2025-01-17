// Copyright 2017 Drone.IO Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scm

import (
	"context"
	"time"
)

type (
	// User represents a user account.
	User struct {
		Login   string
		Name    string
		Email   string
		Avatar  string
		Created time.Time
		Updated time.Time
	}

	// Email represents a user email.
	Email struct {
		Value    string
		Primary  bool
		Verified bool
	}

	// UserService provides access to user account resources.
	UserService interface {
		// Find returns the authenticated user.
		Find(context.Context) (*User, *Response, error)

		// FindEmail returns the authenticated user email.
		FindEmail(context.Context) (string, *Response, error)

		// FindLogin returns the user account by username.
		FindLogin(context.Context, string) (*User, *Response, error)

		// ListEmail returns the user email list.
		ListEmail(context.Context, ListOptions) ([]*Email, *Response, error)
	}
)
