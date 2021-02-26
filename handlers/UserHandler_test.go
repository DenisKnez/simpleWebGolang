package handlers_test

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DenisKnez/simpleWebGolang/data"
	"github.com/DenisKnez/simpleWebGolang/handlers"
)

type UserServiceMock struct{}

var isError bool = false

func (srv UserServiceMock) Users() (users []data.User, err error) {
	if isError == false {
		users = []data.User{
			{Name: "someName"},
		}
		return
	} else {
		err = errors.New("test error")
		return
	}

}
func (srv UserServiceMock) GetUserByID(id string) (user data.User, err error) {
	if isError == false {
		user = data.User{}
		return
	} else {
		err = errors.New("test error")
		return
	}
}
func (srv UserServiceMock) PagedUsers(pageSize int, pageNumber int) (users []data.User, err error) {
	users = []data.User{}
	return
}
func (srv UserServiceMock) CreateUser(user *data.User) error {
	return nil
}
func (srv UserServiceMock) DeleteUser(id string) error {
	return nil
}
func (srv UserServiceMock) UpdateUser(user *data.User) error {
	return nil
}

func TestGetUsers(t *testing.T) {

	tt := []struct {
		name           string
		status         int
		errMsg         string
		isServiceError bool
		isError        bool
	}{
		{name: "correct", status: http.StatusFound, isServiceError: false, isError: false},
		{name: "service returns error", status: http.StatusInternalServerError, errMsg: "Could not return users", isServiceError: true, isError: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/users", nil)

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			serviceMock := UserServiceMock{}
			isError = tc.isServiceError

			handler := handlers.NewUserHandler(serviceMock)
			getUsersFunction := http.HandlerFunc(handler.GetUsers)

			getUsersFunction.ServeHTTP(rr, req)

			response := rr.Result()
			defer response.Body.Close()

			responseBody, err := ioutil.ReadAll(response.Body)
			trimmedResponseBody := strings.TrimSpace(string(responseBody))

			if err != nil {
				t.Error("could not read response", err)
			}

			if isError == true && tc.errMsg != trimmedResponseBody {
				t.Errorf("handler returned wrong err message got: %s, expected: %s", responseBody, tc.errMsg)
			}

			if response.StatusCode != tc.status {
				t.Errorf("handler returned  wrong status code got: %v want: %v", response.StatusCode, tc.status)
			}

		})
	}

}

func TestGetUserByID(t *testing.T) {

	tt := []struct {
		name           string
		status         int
		errMsg         string
		isServiceError bool
		isError        bool
		id             string
	}{
		{name: "id not provided", status: http.StatusBadRequest, errMsg: "user id was not provided", isServiceError: false, isError: false},
		{name: "correct", status: http.StatusFound, id: "1", isServiceError: false, isError: false},
		{name: "service returns error", status: http.StatusNotFound, id:"1", errMsg: "The user with the id does not exist", isServiceError: true, isError: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "/users/user?id="+tc.id, nil)

			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()

			serviceMock := UserServiceMock{}
			isError = tc.isServiceError

			handler := handlers.NewUserHandler(serviceMock)
			getUsersFunction := http.HandlerFunc(handler.GetUser)

			getUsersFunction.ServeHTTP(rr, req)

			response := rr.Result()
			defer response.Body.Close()

			responseBody, err := ioutil.ReadAll(response.Body)
			trimmedResponseBody := strings.TrimSpace(string(responseBody))

			if err != nil {
				t.Error("could not read response", err)
			}
			if response.StatusCode != tc.status {
				t.Errorf("handler returned  wrong status code got: %v want: %v", response.StatusCode, tc.status)
			}

			if isError == true && tc.errMsg != trimmedResponseBody {
				t.Errorf("handler returned wrong err message got: %s, expected: %s", responseBody, tc.errMsg)
			}

		})
	}

}
