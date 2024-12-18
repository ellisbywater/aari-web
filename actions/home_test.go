package actions

import (
	"net/http"

	"aari_web/models"
)

func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(http.StatusFound, res.Code)
	as.Equal(res.Location(), "/auth/new")
}

func (as *ActionSuite) Test_HomeHandler_LoggedIn() {
	u := &models.User{
		Email:                "mark@example.com",
		Password:             "password",
		PasswordConfirmation: "password",
	}
	verrs, err := u.Create(as.DB)
	as.NoError(err)
	as.False(verrs.HasAny())
	as.Session.Set("current_user_id", u.ID)

	res := as.HTML("/auth").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Sign Out")

	as.Session.Clear()
	res = as.HTML("/auth").Get()
	as.Equal(http.StatusOK, res.Code)
	as.Contains(res.Body.String(), "Sign In")
}
