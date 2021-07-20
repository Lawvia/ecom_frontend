package controllers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	// "time"

	"ecom_frontend/app/controllers/auth"
	"ecom_frontend/app/controllers/panels"
	"ecom_frontend/app/modules"
	"ecom_frontend/app/dto"

	"github.com/foolin/goview"
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
)

func Init() {
	new(modules.Configuration).Init()
	config := new(modules.Configuration).GetConfig()

	r := gin.Default()
	store, _ := redis.NewStore(10, "tcp", config.GetString("REDIS_HOST")+":"+config.GetString("REDIS_PORT"), config.GetString("REDIS_PASS"), []byte("secret"))
	// store.Options(sessions.Options{MaxAge: 900})
	r.Use(sessions.Sessions("app_session", store))

	r.HTMLRender = ginview.New(goview.Config{
		Root:      "app/views",
		Extension: ".html",
		Master:    "layouts/master",
	})
	r.Use(static.Serve("/", static.LocalFile("./public", false)))

	r.GET("/", func(c *gin.Context) {
		userData := new(auth.ControllerAuth).GetUserSession(c)
		respData := new(modules.BackRequest).BackendRequestGet("resource/get_items", "", "")
		fmt.Println("user session",userData)
		isUser := false
		
		if userData != nil {
			isUser = true
		}

		err := ""
		var dataItems []dto.ItemsDetail
		if respData == "failed" {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"error": err,
				"user":   userData,
				"data": dataItems,
				"logged": isUser,
			})
			return
		}
		
		var DataScResp dto.ItemsResponse
		ok := json.Unmarshal([]byte(string(respData)), &DataScResp)
		fmt.Println(ok)
		// fmt.Println(DataScResp)
		if (DataScResp.Error != 0){
			err = DataScResp.Status
		}else{
			dataItems = DataScResp.Data
		}
		
		c.HTML(http.StatusOK, "index.html", gin.H{
			"error": err,
			"user":   userData,
			"data": dataItems,
			"logged": isUser,
		})
	})

	new(auth.ControllerAuth).Routes(r)
	new(panels.ControllerPanels).Routes(r)

	//prevent error, add 404 page on unknown request
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "pages/page_404.html", nil)
	})

	log.Fatal(r.Run(":" + config.GetString("app_port")))
}

