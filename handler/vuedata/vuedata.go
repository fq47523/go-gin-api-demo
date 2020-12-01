package vuedata

import (
	d "apidemo/dao"
	m "apidemo/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type response struct {
	Code int `json:"code"`
	Date interface{} `json:"data"`
}

type DeleteData struct {
 Sqlid int	`json:"sqlid"`
 Tableindex int `json:"tableindex"`
}

type QueryParam struct {
	Page int `json:"page"`
	Limit int `json:"limit"`
	Sort string `json:"sort"`
}



// @Title 新增VueData
// @Author 289557905@qq.com
// @Description 新增VueData
// @Tags release template
// @Param vue body	models.VueData true "JSON数据"
// @Success 20000 {object} response
// @Router	/api/v1/release/template/add [post]
func CreateVueData(c *gin.Context) {
	create := m.VueData{}
	buf := make([]byte, 1024)
	n,_ := c.Request.Body.Read(buf)
	json.Unmarshal(buf[0:n], &create)

	d.CreateVueData(create)
	c.JSON(http.StatusOK,response{Code:20000})
}


func QueryVueDataAll(c *gin.Context) {
	q := c.Query("page")
	l := c.Query("limit")
	s := c.Query("sort")
	page,_ := strconv.Atoi(q)
	limit,_ := strconv.Atoi(l)


	query_list,count := d.QueryVueDataAll(page,limit,s)

	data_fmt := make(map[string]interface{})
	data_fmt["total"] = count
	data_fmt["items"] = query_list
	c.JSON(http.StatusOK, response{Code:20000,Date:data_fmt})
}



func UpdateVueData(c *gin.Context) {
	update := m.VueData{}
	c.ShouldBindJSON(&update)

	d.UpdateVueData(update.ID,update)
	c.JSON(http.StatusOK, response{Code:20000})
}

func DeleteVueData(c *gin.Context){
	delete := DeleteData{}
	c.ShouldBindJSON(&delete)
	err := d.DeleteVueData(delete.Sqlid)
	if err != nil {
		c.JSON(http.StatusOK, response{Code:40000})
		return
	}
	c.JSON(http.StatusOK, response{Code:20000})

}