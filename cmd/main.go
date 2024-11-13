package main

import (
	"fmt"
	"net/http"
	"openuao"
	"sync"

	"github.com/labstack/echo/v4"
)

func main() {
	var i = 0

	var wg sync.WaitGroup
	wg.Add(1)
	uao := openuao.NewOrchestrator(
		// openuao.WithJsonConfig(""),
	)

	go uao.Run(&i)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, fmt.Sprintf("Number of lines: %d", i))
	})
	
	e.Logger.Fatal(e.Start("127.0.0.1:1323"))
}