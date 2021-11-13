package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"io/ioutil"
	"net/http"
	"path"
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
			videoList = append(videoList, path.Join(filePath, video.Name()))
		}
	}
	return c.JSON(http.StatusOK, echo.Map{
		"videos": videoList,
	})
}

func GetVideo(c echo.Context) error {
	videoName := c.Param("name")
	videos, err := ioutil.ReadDir(filePath)
	if err != nil {
		return err
	}
	for _, v := range videos {
		if v.Name() == videoName {
			return c.File(path.Join(filePath, v.Name()))
		}
	}
	return c.NoContent(http.StatusNotFound)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger(), middleware.Recover(), middleware.CORS(), middleware.RemoveTrailingSlash())
	e.GET("/records", GetVideos)
	e.GET("/records/:name", GetVideo)
	e.Logger.Fatal(e.Start(":1323"))
}
