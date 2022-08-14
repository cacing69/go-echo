// package api

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
package handler

import (
  "fmt"
  "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Hello from Go!</h1>")
}
