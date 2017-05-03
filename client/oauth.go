// Code generated by goagen v1.1.0-dirty, command line:
// $ goagen
// --design=github.com/tikasan/goa-oauth2-practice/design
// --out=$(GOPATH)/src/github.com/tikasan/goa-oauth2-practice
// --version=v1.1.0
//
// API "auth": oauth Resource Client
//
// The content of this file is auto-generated, DO NOT MODIFY

package client

import (
	"fmt"
	"golang.org/x/net/context"
	"net/http"
	"net/url"
)

// CallbackOauthPath computes a request path to the callback action of oauth.
func CallbackOauthPath() string {

	return fmt.Sprintf("/login/callback")
}

// CallbackOauth makes a request to the callback action endpoint of the oauth resource
func (c *Client) CallbackOauth(ctx context.Context, path string, code string) (*http.Response, error) {
	req, err := c.NewCallbackOauthRequest(ctx, path, code)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewCallbackOauthRequest create the request corresponding to the callback action endpoint of the oauth resource.
func (c *Client) NewCallbackOauthRequest(ctx context.Context, path string, code string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	values.Set("code", code)
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// LoginOauthPath computes a request path to the login action of oauth.
func LoginOauthPath() string {

	return fmt.Sprintf("/login")
}

// LoginOauth makes a request to the login action endpoint of the oauth resource
func (c *Client) LoginOauth(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewLoginOauthRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewLoginOauthRequest create the request corresponding to the login action endpoint of the oauth resource.
func (c *Client) NewLoginOauthRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "https"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}