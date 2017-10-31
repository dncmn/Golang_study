package controllers

import (
	"github.com/astaxie/beego"
	"encoding/json"
	"log"
	"fmt"
	"time"
	"aaDe/aaDeZhou/models"
)

type Webservice struct {
	beego.Controller
}

func (this *Webservice) Get() {
	//handRequest(this)
	data := this.GetString("data") //获取一条数据
	var res models.Result
	parseErr := json.Unmarshal([]byte(data), &res)
	if parseErr != nil {
		log.Println("解析数据错误", parseErr)
		this.Ctx.WriteString("解析数据失败.....")
		return
	}
	msg := fmt.Sprint("   ", res.Name, "   ", res.Password)
	log.Println(time.Now().Format("20016-01-02 15:04:05"), "  ", msg)
	this.Ctx.WriteString("ok...............ok")
	this.Data["ok"] = "ok.................ok"

}

func handRequest(this *Webservice) {

}