package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	. "dungou.cn/controller"
	. "dungou.cn/datasource"
	. "dungou.cn/def"
	. "dungou.cn/task"
	. "dungou.cn/util"
	"os"
)

var orm Orm
var mssql Mssql
var tjMssql TjMssql
func main() {

	MODE = Trim(os.Getenv("mode"))
	beego.BConfig.Listen.HTTPPort = 9700                     //端口设置
	beego.BConfig.RecoverPanic = true                        //开启异常捕获
	beego.BConfig.EnableErrorsShow = true
	beego.BConfig.CopyRequestBody = true
	beego.InsertFilter("/*", beego.BeforeRouter, BaseFilter) //路由过滤

	//自动匹配路由
	beego.AutoRouter(&ApiController{})
	beego.InsertFilter("/rpc", beego.BeforeRouter, RpcFilter)
	//beego.InsertFilter("/api/*", beego.BeforeRouter, WhiteListFilter)
	Mkdir("./logs")                                        //创建日志文件夹
	beego.SetLogger("file", `{"filename":"logs/run.log"}`) //定义日志文件
	beego.BeeLogger.SetLogFuncCallDepth(4)
	//调用以下函数处理接口数据
	orm.Init()
	mssql.Init()
	tjMssql.Init()
	//a()
	Insert()
	go func() {
		//开启协程
		InitCache() //初始化
		crontab()   //开启定时任务
	}()
	beego.Run() //启动项目
}

func crontab() {
	toolbox.AddTask("pd", toolbox.NewTask("pd", "0 */15 * * * *", func() error {
		//每10分钟运行以下函数
		Dhq <- func() {

		}
		return nil
	}))
	toolbox.StartTask() //开启定时任务
}

func a() {
	jd,err :=Upload("http://106.75.33.170:16680/api/upload","D:/lon.xlsx")
	if err != nil {
		Debug(err)
	}
	//json := *JsonDecode([]byte(jd))
	Debug(string(jd))
}