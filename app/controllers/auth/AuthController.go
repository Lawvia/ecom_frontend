package auth

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"ecom_frontend/app/dto"
	"ecom_frontend/app/modules"
)

const (
	userkey = "user"
)

//ControllerAuth is a class of library routing api
type ControllerAuth struct{}

//Routes for configure database connection
func (cp ControllerAuth) Routes(router *gin.Engine) {
	routerPanels := router.Group("/auth")
	routerPanels.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth/authentication.html", gin.H{
			"error": "",
			"logged": false,
		})
	})

	routerPanels.POST("/login", func(c *gin.Context) {
		session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("password")
		users := new(dto.UserPanelDto)

		respLogin := new(modules.BackRequest).BackendRequestGet("auth/login", username, password)
		fmt.Println(respLogin)

		if respLogin == "failed" {
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"errorLogin": "Sorry, there is a problem on our end!", "logged": false,})
			return
		}

		var DataScResp dto.AuthResponse
		ok := json.Unmarshal([]byte(string(respLogin)), &DataScResp)
		fmt.Println(ok)
		fmt.Println(DataScResp)

		if DataScResp.Error == 0{
			//save token, continue
			users.Token = DataScResp.Token
			users.Username = username

		}else{
			fmt.Println(DataScResp.Message)
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"errorLogin": DataScResp.Message, "logged": false,})
			return
		}

		gob.Register(dto.UserPanelDto{})
		session.Set(userkey, users)
		err := session.Save()

		if err != nil {
			fmt.Println(err)
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"error": "Failed to save session", "logged": false,})
			return
		}
		c.Redirect(301, "/")
		return
	})

	routerPanels.GET("/logout", func(c *gin.Context) {
		fmt.Println("logout")
		session := sessions.Default(c)
		user := session.Get(userkey)
		if user == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			return
		}
		session.Delete(userkey)
		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		c.Redirect(301, "/auth")
		return
	})

	routerPanels.POST("/register", func(c *gin.Context) {
		// session := sessions.Default(c)
		username := c.PostForm("username")
		password := c.PostForm("new_password")
		repassword := c.PostForm("confirm_password")

		if password != repassword {
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"errorRegister": "Password do not match!", "logged": false,})
			return
		}

		respReg := new(modules.BackRequest).BackendRequestGet("auth/register", username, password)
		fmt.Println(respReg)

		if respReg == "failed" {
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"errorRegister": "Sorry, there is a problem on our end!", "logged": false,})
			return
		}

		var DataScResp dto.AuthResponse
		ok := json.Unmarshal([]byte(string(respReg)), &DataScResp)
		fmt.Println(ok)
		fmt.Println(DataScResp)

		if DataScResp.Error == 0{
			fmt.Println(DataScResp.Message)
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"success": DataScResp.Message + ", login to continue :)", "logged": false,})
			return
		}else{
			fmt.Println(DataScResp.Message)
			c.HTML(http.StatusOK, "auth/authentication.html", gin.H{"errorRegister": DataScResp.Message, "logged": false,})
			return
		}
	})
}

// AuthRequired is a simple middleware to check the session
func (cp ControllerAuth) AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	fmt.Println("authreq", user)
	if user == nil {
		c.Redirect(301, "/auth")
		return
	}
	c.Next()
}

// GetUserSession is a simple function to get user session data
func (cp ControllerAuth) GetUserSession(c *gin.Context) *dto.UserPanelDto {
	session := sessions.Default(c)
	user := session.Get(userkey)
	userData, ok := user.(dto.UserPanelDto)

	if ok {
		return &userData
	}
	return nil
}

// IsAdmin is a simple function to get user session data
func (cp ControllerAuth) IsAdmin(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	userData, _ := user.(dto.UserPanelDto)
	if userData.Token != "" {
		c.Redirect(301, "/panels/welcome")
		return
	}
	c.Next()
}
