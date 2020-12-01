package sd

import (
	"apidemo/conf"
	d "apidemo/dao"
	u "apidemo/handler/utils"
	m "apidemo/models"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)







// HealthCheck shows `OK` as the ping-pong result.
func HealthCheck(c *gin.Context) {
	message := "OK"
	c.String(http.StatusOK, "\n"+message)
}


func StudentList(c *gin.Context) {
	q := c.Query("page")
	l := c.Query("limit")
	s := c.Query("sort")
	page,_ := strconv.Atoi(q)
	limit,_ := strconv.Atoi(l)
	fmt.Println("qqqqqqq:",s,page,limit)
	StudentList,StudentListCount := d.QueryStudentAll(page,limit,s)
	ResponseResult := u.ResponseFmt(conf.ResponseSuccess,"",StudentList,StudentListCount)
	c.JSON(http.StatusOK, ResponseResult)
}
/*
func QueryStudentSingle(c *gin.Context) {
	par := c.Query("id")
	stuid,_ := strconv.Atoi(par)
	stuobj := d.QueryStudentSingle(stuid)

	c.JSON(http.StatusOK, &u.Response{Code:200,Message:"OK",Data:stuobj})
}
*/

func CreateStudent(c *gin.Context) {
	
	if err := d.CreateStudent(c);err !=nil{
		c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseFailure,Message:err})
		return
	}

	c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseSuccess,Message:"OK"})
}

func UpdateStudent(c *gin.Context) {
	stuu := m.Student{}
	buf := make([]byte, 1024)
	n,_ := c.Request.Body.Read(buf)
	json.Unmarshal(buf[0:n], &stuu)
	// 很关键,把读过的字节流重新放到body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(buf[0:n]))
	
	if err := d.UpdateStudent(c,stuu.ID);err !=nil{
		c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseFailure,Message:err})
		return
	}
	
	c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseSuccess,Message:"OK"})
}

func DeleteStudent(c *gin.Context){
	DeleteStudent := m.Student{}
	c.ShouldBindJSON(&DeleteStudent)
	if err := d.DeleteStudent(&DeleteStudent); err != nil{
		c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseFailure,Message:err})
	}
	c.JSON(http.StatusOK,&u.Response{Code:conf.ResponseSuccess,Message:"OK"})
}

func ClassList(c *gin.Context){
	classlist := d.QueryClassList()
	ResponseResult := u.ResponseFmt(conf.ResponseSuccess,"",classlist,len(classlist))
	c.JSON(http.StatusOK,ResponseResult)
}


func TeacherList(c *gin.Context){
	teacherlist := d.QueryTeacherList()
	ResponseResult := u.ResponseFmt(conf.ResponseSuccess,"",teacherlist,len(teacherlist))
	c.JSON(http.StatusOK,ResponseResult)
}