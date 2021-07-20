package panels

import (
	// "log"
	"net/http"
	"fmt"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"ecom_frontend/app/modules"
	"ecom_frontend/app/dto"
	"ecom_frontend/app/controllers/auth"
)

type ControllerPanels struct{}

func (cp ControllerPanels) Routes(router *gin.Engine) {
	routerPanels := router.Group("/panels")
	controllerAuth := new(auth.ControllerAuth)
	routerPanels.Use(controllerAuth.AuthRequired)

	routerPanels.GET("/review", func(c *gin.Context) {
		userData := new(auth.ControllerAuth).GetUserSession(c)
		isUser := false	
		if userData != nil {
			isUser = true
			fmt.Println("di panel",userData.Username)
		}

		err := ""
		var dataCart []dto.UpdateDetails
		respData := new(modules.BackRequest).BackendRequestGet("transaction/get_cart", "", userData.Token)
		var DataScResp dto.UpdateResponse
		ok := json.Unmarshal([]byte(string(respData)), &DataScResp)
		fmt.Println(ok)
		fmt.Println("datasc",DataScResp)
		if (DataScResp.Error != 0){
			err = DataScResp.Message
		}else{
			dataCart = DataScResp.Data
		}
		
		c.HTML(http.StatusOK, "panels/review.html", gin.H{
			"error": err,
			"user":   userData,
			"logged": isUser,
			"cart": dataCart,
			"qty": len(dataCart),
			"total": DataScResp.GrandTotal,
		})
	})

	routerPanels.GET("/list", func(c *gin.Context) {
		userData := new(auth.ControllerAuth).GetUserSession(c)
		isUser := false	
		if userData != nil {
			isUser = true
			fmt.Println("di panel",userData.Username)
		}

		err := ""
		var dataCart []dto.ListDetails
		respData := new(modules.BackRequest).BackendRequestGet("transaction/get_transaction", "", userData.Token)
		var DataScResp dto.ListTransac
		ok := json.Unmarshal([]byte(string(respData)), &DataScResp)
		fmt.Println(ok)
		fmt.Println("datasc",DataScResp)
		if (DataScResp.Error != 0){
			err = DataScResp.Message
		}else{
			dataCart = DataScResp.Data
		}
		
		c.HTML(http.StatusOK, "panels/list.html", gin.H{
			"error": err,
			"user":   userData,
			"logged": isUser,
			"cart": dataCart,
			"qty": len(dataCart),
		})
	})

	routerPanels.POST("/addcart", func(c *gin.Context) {
		res := new(dto.AuthResponse)
		userData := new(auth.ControllerAuth).GetUserSession(c)

		var data dto.UpdateCart
		c.Bind(&data)
		fmt.Println(data)

		//update cart and proceed to review page
		respData := new(modules.BackRequest).BackendRequestPost("transaction/update_cart", userData.Token, data)
		if respData == "failed" {
			res.Status = "FAILED"
			res.Message = "Failed to update cart!"
			return
		}

		var DataScResp dto.ItemsResponse
		ok := json.Unmarshal([]byte(string(respData)), &DataScResp)
		fmt.Println(ok)

		if (DataScResp.Error == 0){
			res.Status = "SUCCESS"
		}else{
			res.Status = "FAILED"
		}
		c.JSON(http.StatusOK, res)
	})

	routerPanels.POST("/checkout", func(c *gin.Context) {
		res := new(dto.AuthResponse)
		userData := new(auth.ControllerAuth).GetUserSession(c)

		var bindable struct {
			Address string `json:"address"`
		}
		c.Bind(&bindable)

		var data dto.UpdateCart
		data.Address = bindable.Address

		fmt.Println(data)
		fmt.Println(bindable)

		//update cart and proceed to review page
		respData := new(modules.BackRequest).BackendRequestPost("transaction/checkout", userData.Token, data)
		if respData == "failed" {
			res.Status = "FAILED"
			res.Message = "Failed to checkout!"
			return
		}

		var DataScResp dto.AuthResponse
		ok := json.Unmarshal([]byte(string(respData)), &DataScResp)
		fmt.Println(ok)

		if (DataScResp.Error == 0){
			res.Status = "SUCCESS"
		}else{
			res.Status = "FAILED"
		}
		c.JSON(http.StatusOK, res)	
	})
}
