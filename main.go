package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
	"strings"
)

const filePath = "/home/lighthouse/video/live"

func GetVideos(c echo.Context) error {
	// list directory files
	videos, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}
	videoList := make([]string, 0)
	for _, video := range videos {
		if !strings.HasSuffix(video.Name(), ".tmp") {
			videoList = append(videoList,  video.Name())
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"videos": videoList,
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover(), middleware.CORS(), middleware.RemoveTrailingSlash())
	e.GET("/records", GetVideos)
	e.Static("/record", filePath)
	e.Logger.Fatal(e.Start(":1323"))
}
