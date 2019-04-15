package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	fmt.Println("INIT: run go test")
}

func TestAuth_Greet(t *testing.T) {
	a := assert.New(t)
	as := NewAuthenticator()
	a.NotNil(as)

	tests := []struct {
		request string
		expect  string
		err     error
	}{
		{
			request: "",
			expect:  fmt.Sprintf("%s %s", greeting, ""),
			err:     nil,
		},
		{
			request: "name",
			expect:  fmt.Sprintf("%s %s", greeting, "name"),
			err:     nil,
		},
		{
			request: "name - long",
			expect:  fmt.Sprintf("%s %s", greeting, "name - long"),
			err:     nil,
		},
	}

	for _, v := range tests {
		res := as.Greet(v.request)
		a.Equal(v.expect, res)
	}
}
