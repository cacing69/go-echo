package handler

// import (
// 	"net/http"
// 	"github.com/labstack/echo/v4"
// )

// func Api() {
// 	e := echo.New()
// 	e.GET("/api/", func(c echo.Context) error {
// 		return c.String(http.StatusOK, "Hello, World!")
// 	})
// 	e.Logger.Fatal(e.Start(":1323"))
// }
// package handler

import (
	"net/http"
	"fmt"
	// "github.com/labstack/echo/v4"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  // e := echo.New()
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
