package server

import (
	S "ConfBackend/services"
	"ConfBackend/view"
	"github.com/gin-gonic/gin"
	"log"
)

func StartApi() {
	s := gin.Default()
	//s.Use(cors())
	//s.Use(printRequest)

	// Human's terminal url group
	pt := s.Group("/pterm")

	//com := s.Group("/com")

	// node api group
	node := s.Group("/node")

	/*			// The car's api group
				car := s.Group("/hero")*/

	// The internal_model center's api
	cc := s.Group("/cc")

	// pad PTerm Services
	{
		pt.GET("/ping", view.PTerm)
		pt.POST("/file", view.FileReceived)
	}

	{
		node.POST("/update_location", view.UpdateLocation)
	}

	// pad Com url
	{
		cc.GET("/hero_control", view.HeroControl)
	}

	{
		test := s.Group("/test")
		test.POST("/db", view.TestDb)
	}
	// set release mode
	err := s.Run(":" + S.S.Conf.App.Port)
	if err != nil {
		log.Fatalln(err)
	}

}
