package session

import (
	"fmt"
	"net/http"

	"github.com/ftsog/ecom/customerrors"
)

func (rdSession *Session) CreateSession(w http.ResponseWriter, r *http.Request, user string) error {
	session, err := rdSession.RediStore.Get(r, "session_id")
	if err != nil {
		return err
	}

	if session.Values["user"] != nil {
		return customerrors.SessionExistsError
	}

	session.Options.Path = "/"
	session.Values["user"] = user

	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (rdSession *Session) DeleteSession(w http.ResponseWriter, r *http.Request) error {
	session, err := rdSession.RediStore.Get(r, "session_id")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		return err
	}

	return nil
}

func (rdSession *Session) GetUserFromSession(w http.ResponseWriter, r *http.Request) (*string, error) {
	session, err := rdSession.RediStore.Get(r, "session_id")
	if err != nil {
		return nil, err
	}

	user, ok := session.Values["user"]
	if !ok {
		return nil, nil
	}

	usr := fmt.Sprintf("%v", user)

	return &usr, nil
}

func (rdSession *Session) CheckSession(w http.ResponseWriter, r *http.Request) (*bool, error) {
	right := true
	wrong := false
	session, err := rdSession.RediStore.Get(r, "session_id")
	if err != nil {
		return nil, err
	}

	_, ok := session.Values["user"]
	if !ok {
		return &wrong, customerrors.InvalidSessionError
	}

	return &right, nil

}
