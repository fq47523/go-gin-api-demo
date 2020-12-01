package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VeaLogin(c *gin.Context) {


	// .....验证请求的{"username":"admin","password":"111111"}

	LoginRespons := map[string]interface{}{}
	token,_ := generateToken()
	fmt.Printf("%T:%v \n",token,token)
	LoginResponsData := map[string]string{"token":token}
	LoginRespons["code"] = 20000
	LoginRespons["data"] = LoginResponsData


	c.JSON(http.StatusOK, LoginRespons)
}

func VeaLoginInfo(c *gin.Context){
	LoginResponsInfo := map[string]interface{}{}
	LoginResponsInfoRoles := []string{"admin"}
	LoginResponsInfoData := map[string]interface{}{"roles":LoginResponsInfoRoles,"introduction":"I am a super administrator","avatar":"https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif","name":"Super Admin"}
	LoginResponsInfo["code"] = 20000
	LoginResponsInfo["data"] = LoginResponsInfoData

	c.JSON(http.StatusOK, LoginResponsInfo)
}

func VeaLogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code":20000,"data":"success"})
}