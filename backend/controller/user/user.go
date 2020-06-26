package user

import (
	"StackBackend/db"
	"StackBackend/model"
	"fmt"
	"github.com/labstack/echo"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"strconv"
)

const BaseURL = "/todoapi"

type Controller struct{}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Initialize(e echo.Echo) (err error) {
	e.GET(BaseURL+"/all", getAllScore)
	e.GET(BaseURL+"/top", getTopScore)
	e.GET(BaseURL+"/:username/addToScore/:score", addScore)
	e.GET(BaseURL+"/allUser", getAllUser)
	e.PUT(BaseURL+"/username/:username", updateUser)
	e.POST(BaseURL+"/updateTodo/:username", updateTodo)
	e.GET(BaseURL+"/username/:username", getUserByUserName)
	e.POST(BaseURL+"/register", newUser)
	return nil
}

func getAllScore(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.SimpleUser()
	defer closeConn()
	var results []model.SimpleUser
	err = collection.Find(nil).All(&results)
	if err != nil {
		return err
	}
	fmt.Println(results)
	return c.JSON(http.StatusOK, results)
}

func getAllUser(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	var results []model.User
	err = collection.Find(nil).All(&results)
	if err != nil {
		return err
	}
	fmt.Println(results)
	return c.JSON(http.StatusOK, results)
}

func getTopScore(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.SimpleUser()
	defer closeConn()
	var results []model.SimpleUser
	err = collection.Find(nil).Sort("-score").All(&results)
	if err != nil {
		return err
	}
	fmt.Println(results)
	if len(results) > 8{
		results = results[:8]
	}
	return c.JSON(http.StatusOK, results)
}


func addScore(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.SimpleUser()
	defer closeConn()
	name := c.Param("username")
	score := c.Param("score")
	var user model.SimpleUser
	user.Name = name
	user.Score, err = strconv.Atoi(score)
	if err != nil{
		fmt.Println(err)
	}
	err = collection.Insert(user)
	if err != nil{
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, user)
}

func updateUser(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	var updatedUser model.User
	err = c.Bind(&updatedUser)
	if err != nil {
		return err
	}
	username := c.Param("username")
	err = collection.Update(bson.M{"name": username}, updatedUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func updateTodo(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	var updatedUser model.User
	var newTodo model.NewTodo
	err = c.Bind(&newTodo)
	if err != nil {
		return err
	}
	username := c.Param("username")
	err = collection.Find(bson.M{"name": username}).One(&updatedUser)
	if err != nil {
		return err
	}
	updatedUser.Todo = newTodo.Todo
	err = collection.Update(bson.M{"name": username}, updatedUser)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, updatedUser)
}

func getUserByUserName(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	username := c.Param("username")
	var result model.User
	err = collection.Find(bson.M{"name": username}).One(&result)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, result)
}

func newUser(c echo.Context) (err error) {
	collection, closeConn := db.GlobalDatabase.User()
	defer closeConn()
	var newUser model.User
	err = c.Bind(&newUser)
	fmt.Println(newUser)
	if err != nil {
		return err
	}
	var findUser model.User
	err = collection.Find(bson.M{"name": newUser.Name}).One(&findUser)
	if err != nil{
		newUser.Todo = ""
		newUser.Id = bson.NewObjectId()
		newUser.Avatar = "https://i.loli.net/2020/03/15/XsJjRomr1dy8u4D.png"
		err = collection.Insert(&newUser)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, newUser)
	} else {
		println("用户名重复")
		return c.JSON(http.StatusNoContent, "用户名重复")
	}

}