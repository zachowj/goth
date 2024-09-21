package onedrive_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachowj/goth"
	"github.com/zachowj/goth/providers/onedrive"
)

func Test_Implements_Session(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &onedrive.Session{}

	a.Implements((*goth.Session)(nil), s)
}

func Test_GetAuthURL(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &onedrive.Session{}

	_, err := s.GetAuthURL()
	a.Error(err)

	s.AuthURL = "/foo"

	url, _ := s.GetAuthURL()
	a.Equal(url, "/foo")
}

func Test_ToJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &onedrive.Session{}

	data := s.Marshal()
	a.Equal(data, `{"AuthURL":"","AccessToken":"","RefreshToken":"","ExpiresAt":"0001-01-01T00:00:00Z"}`)
}

func Test_String(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	s := &onedrive.Session{}

	a.Equal(s.String(), s.Marshal())
}
