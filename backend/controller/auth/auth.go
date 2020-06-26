package auth

import (
	"StackBackend/db"
	"StackBackend/model"
	"fmt"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

const BaseURL = "/todoapi/auth"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}
func (c *Controller) Initialize(e echo.Echo) (err error) {
	e.POST(BaseURL+"/userLogin", Login)
	return nil
}

func Login(c echo.Context) error {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	var user model.User
	fmt.Println(c)
	err := c.Bind(&user)
	fmt.Println(user)
	username := user.Name
	password := user.Password
	var result model.User
	err = collection.Find(bson.M{"name": username}).One(&result)
	fmt.Println(result)
	if err != nil {
		return c.String(http.StatusUnauthorized, "用户名不存在!")
	}
	if password != result.Password {
		return c.String(http.StatusUnauthorized, "密码不正确!")
	} else {
		return c.String(http.StatusOK, "Welcome!")
	}
}


