package routers


import (
	"aaDe/aaDeZhou/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/testSng", &controllers.Webservice{})

}
