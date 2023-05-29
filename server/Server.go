package server

import (
	"ConfBackend/chat"
	com "ConfBackend/common"
	S "ConfBackend/services"
	"ConfBackend/view"
	"github.com/gin-gonic/gin"
	"log"
)

func StartApi() {
	s := gin.Default()
	//s.Use(cors())
	//s.Use(printRequest)

	// pad PTerm Services
	{
		// Human's terminal url group
		pt := s.Group("/pterm")
		pt.GET("/ping", view.PTerm)
		pt.POST("/file", view.FileReceived)
	}

	{
		// node api group
		node := s.Group("/node")
		node.POST("/update_location", view.UpdateLocation)
	}

	{
		// The car's api group
		car := s.Group("/hero")
		car.GET("/ping", func(context *gin.Context) {
			com.Ok(context)
		})
	}

	// pad Com url
	{
		// The internal_model center's api
		cc := s.Group("/cc")
		cc.GET("/hero_control", view.HeroControl)
	}

	{
		// instant messaging im common api
		im := s.Group("/im")
		im.Use(MustHasUserUUID())
		im.POST("/sendmsg", view.SendMsg)
		im.GET("/ws", chat.WsConnectionManager.WebSocketHandler)
		im.GET("/all_online", view.AllOnline)
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
