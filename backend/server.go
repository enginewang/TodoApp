package StackBackend

import (
	"StackBackend/controller/user"
	"StackBackend/controller/auth"
	"echo/middleware"
	"github.com/labstack/echo"
	"net/http"
)

type Server struct {
	Addr string
	e    *echo.Echo
}

func NewServer(addr string) *Server {
	return &Server{
		Addr: addr,
		e:    echo.New(),
	}
}

func (s *Server) Init() (err error) {
	s.e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Stack!")
	})
	//g := s.e.Group("")
	e := s.e
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowHeaders:     []string{"*"},
	}))

	err = user.NewController().Initialize(*e)
	err = auth.NewController().Initialize(*e)
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) StartServer ()  {
	s.e.Logger.Fatal(s.e.Start(":1323"))
}
