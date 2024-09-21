package azuread_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zachowj/goth"
	"github.com/zachowj/goth/providers/azuread"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	provider := azureadProvider()

	a.Equal(provider.ClientKey, os.Getenv("AZUREAD_KEY"))
	a.Equal(provider.Secret, os.Getenv("AZUREAD_SECRET"))
	a.Equal(provider.CallbackURL, "/foo")
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	p := azureadProvider()
	a.Implements((*goth.Provider)(nil), p)
}

func Test_BeginAuth(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	provider := azureadProvider()
	session, err := provider.BeginAuth("test_state")
	s := session.(*azuread.Session)
	a.NoError(err)
	a.Contains(s.AuthURL, "login.microsoftonline.com/common/oauth2/authorize")
	a.Contains(s.AuthURL, "https%3A%2F%2Fgraph.windows.net%2F")
}

func Test_SessionFromJSON(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := azureadProvider()
	session, err := provider.UnmarshalSession(`{"AuthURL":"https://login.microsoftonline.com/common/oauth2/authorize","AccessToken":"1234567890"}`)
	a.NoError(err)

	s := session.(*azuread.Session)
	a.Equal(s.AuthURL, "https://login.microsoftonline.com/common/oauth2/authorize")
	a.Equal(s.AccessToken, "1234567890")
}

func azureadProvider() *azuread.Provider {
	return azuread.New(os.Getenv("AZUREAD_KEY"), os.Getenv("AZUREAD_SECRET"), "/foo", nil)
}
