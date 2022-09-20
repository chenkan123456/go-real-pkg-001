package util

import (
	"math"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

var pageSize = 30

type PageModel struct {
	Row        int         `json:"row"`
	Page       int         `json:"page"`
	Totalcount int         `json:"totalcount"`
	Totalpage  int         `json:"totalpage"`
	Items      interface{} `json:"items"`
}

func GetPage(pageNumber int, c *gin.Context) (int, int, int) {
	result := 0
	row := pageNumber
	if row == 0 {
		row = pageSize
	}
	page, _ := com.StrTo(c.Param("page")).Int()
	if page > 0 {
		result = (page - 1) * row
	}

	return page, row, result
}

func GetPageModel(pageSize, pageNumber, countTotal int, data interface{}) PageModel {
	var (
		row       = pageNumber
		kind      = reflect.TypeOf(data).Kind()
		len       = reflect.ValueOf(data).Len()
		totalpage = int(math.Ceil(float64(countTotal) / float64(row)))
	)
	if kind == reflect.Slice && len == 0 {
		totalpage = 0
		countTotal = 0
	}

	if row == 0 {
		row = pageSize
	}

	var model = PageModel{
		Row:        row,
		Page:       pageSize,
		Totalcount: countTotal,
		Totalpage:  totalpage,
		Items:      data,
	}

	return model

}
