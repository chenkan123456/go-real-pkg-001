package app

import (
	"fmt"

	"gitlab.centaline-sc.com/real-db-project/go-pkg/pkg/e"
	"gitlab.centaline-sc.com/real-db-project/go-pkg/pkg/util"

	"github.com/gin-gonic/gin"
)

type ResponseBase struct {
	State bool   `json:"state"`
	Msg   string `json:"message"`
	Code  int    `json:"code"`
}

type ResponseModel struct {
	State bool        `json:"state"`
	Msg   string      `json:"message"`
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(state bool, httpCode, errCode int) {

	g.C.JSON(httpCode, &ResponseBase{
		State: state,
		Code:  errCode,
		Msg:   e.GetMsg(errCode),
	})

}

func (g *Gin) ResponseError(state bool, httpCode, errCode int, err error) {
	// var msg string

	// value, ok := err.(error)
	// if ok {
	// 	msg = fmt.Sprintf("%v:%v", e.GetMsg(errCode), value)
	// }

	g.C.JSON(httpCode, &ResponseBase{
		State: state,
		Code:  errCode,
		Msg:   fmt.Sprintf("%v:%v", e.GetMsg(errCode), err),
	})

}

func (g *Gin) ResponseModel(state bool, httpCode, errCode int, data interface{}) {

	val, ok := data.(util.PageModel)
	if ok && val.Totalcount == 0 {
		data = make([]string, 0)
	}

	g.C.JSON(httpCode, &ResponseModel{
		State: state,
		Code:  errCode,
		Msg:   e.GetMsg(errCode),
		Data:  data,
	})

}
