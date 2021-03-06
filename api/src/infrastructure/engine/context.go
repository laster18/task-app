package engine

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Context struct {
	w      http.ResponseWriter
	r      *http.Request
	ps     httprouter.Params
	enable bool
}

type Response struct {
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func (c *Context) JSON(code int, msg string, data interface{}) {
	fmt.Println("response data: ", data)
	res := &Response{
		Msg:  msg,
		Data: data,
	}

	b, err := json.Marshal(res)
	if err != nil {
		c.error()
		return
	}
	c.w.Header().Set("Content-Type", "application/json")
	c.w.WriteHeader(code)
	c.w.Write(b)
}

func (c *Context) error() {
	c.w.WriteHeader(http.StatusInternalServerError)
	c.w.Write([]byte("Internal server error."))
}

func (c *Context) Param(key string) string {
	return c.ps.ByName(key)
}

func (c *Context) Set(key string, value interface{}) {
	c.r.Context()
	Context := context.WithValue(c.r.Context(), key, value)
	c.r = c.r.WithContext(Context)
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	v := c.r.Context().Value(key)
	return v, v != nil
}

func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}

	panic("Key \"" + key + "\" does not exist")
}

func (c *Context) Bind(value interface{}) error {
	// TODO: validtorを使ってvalidationする
	return json.NewDecoder(c.r.Body).Decode(value)
}

func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get("q")
}

func (c *Context) Header(key string) string {
	return c.r.Header.Get(key)
}

func (c *Context) Abort() {
	c.enable = false
}

func (c *Context) SetCookie(name, value string) {
	cookie := &http.Cookie{
		Name:  name,
		Value: value,
		Path:  "/", // これがないと/apiに対してセットするためfrontサーバにcookieが送信されない
		// Domain: "localhost:4000",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
		// Secure: true, // localではfalseにしておく
	}

	http.SetCookie(c.w, cookie)
}

func (c *Context) GetCookie(name string) (string, error) {
	cookie, err := c.r.Cookie(name)

	if err != nil {
		if err == http.ErrNoCookie {
			return "", nil
		}

		return "", err
	}

	return cookie.Value, nil
}
