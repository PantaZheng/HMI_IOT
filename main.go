package HMI_IoT

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/middleware/logger"
	"HMI_IoT/controller"
)

func newApp() (api *iris.Application){
	api = iris.New()
	api.Use(logger.New())

	api.PartyFunc("/anon", func(anon router.Party){
		anon.PartyFunc("/wechat", func(wechat router.Party) {
			wechat.Get("/",controller.Login)
			//wechat.Post("/",control)
		})
	})


	return
}
func main() {
	
}
