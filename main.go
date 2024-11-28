package main

import (
	"bookeng/contentful"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func roomsHandler(cf *contentful.Contentful) func(c *gin.Context) {
	return func(c *gin.Context) {
		hotelID := c.Param("hotelID")
		b, err := cf.FetchRoomsByHotelID(hotelID)
		if err != nil {
			c.Error(err)
			return
		}

		rooms, err := cf.ConvertRawToRoom(b)
		if err != nil {
			c.Error(err)
			return
		}

		c.JSON(200, gin.H{
			"rooms": rooms,
		})
	}
}

func main() {
	cf := contentful.NewContentful()

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("./web/dist", true)))
	r.GET("/hotels/:hotelID/rooms", roomsHandler(cf))
	r.NoRoute(func(c *gin.Context) {
		c.File("./web/dist/index.html")
	})

	r.Run()
}
