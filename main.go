package main

import (
	"apidemo/config"
	"apidemo/dao"
	m "apidemo/models"
	"apidemo/router"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"time"


)


var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)
// @title Golang Esign API
// @version 1.0
// @description  Golang api of demo
// @termsOfService http://github.com

// @contact.name API Support
// @contact.url http://www.cnblogs.com
// @contact.email ×××@qq.com

// @host 127.0.0.1:8081

func main() {

	pflag.Parse()

	// init config
	// cfg变量值从命令行 flag 传入，可以传值，比如./XXXX -c config.yaml，也可以为空，如果为空会默认读取conf/config.yaml
	if err := config.Init(*cfg); err != nil {
		panic(err)
	}

	/*
	for {
		fmt.Println(viper.GetString("runmode"))     # 热更新监控
		time.Sleep(4*time.Second)
	}
	*/
	// Set gin mode.
	gin.SetMode(viper.GetString("runmode"))

	// conn database
	err := dao.InitSqlite()
	if err != nil {
		fmt.Printf("Init Sqlite failed, err:%v\n", err)
		return
	}
	defer dao.Close() // 程序退出关闭数据库连接
	dao.DB.AutoMigrate(
		&m.IDCard{},
		&m.Teacher{},
		&m.Class{},
		&m.Student{},
		&m.VueData{},
		&m.Platforms{},
	)

	/*
	//  mock data
	data,err := ioutil.ReadFile("mock/VueDataMock")  //这里返回的data是一个字节切片
	if err!=nil{
		fmt.Println("File reading error", err)
	}

	var a []m.VueData
	eerr := json.Unmarshal(data, &a)

	if eerr != nil {
		fmt.Printf("unmarshal err=%v\n", eerr)
	}

	for i:=0;i<len(a);i++{
		dao.DB.Create(&a[i])
	}
	*/



	// Create the Gin engine.
	g := gin.New()

	middlewares := []gin.HandlerFunc{}
	
	// Routes.
	router.Load(
		// Cores.
		g,

		// Middlwares.
		middlewares...,
	)
	// Ping the server to make sure the router is working.
	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", "8080")
	log.Infof(http.ListenAndServe(viper.GetString("addr"), g).Error())
}

// pingServer pings the http server to make sure the router is working.
func pingServer() error {
	for i := 0; i < 2; i++ {
		// Ping the server by sending a GET request to `/health`.
		resp, err := http.Get("http://127.0.0.1:8080" + "/api/v1/student/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Infof("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
