package app

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/chenkan123456/go-real-pkg-001/pkg/e"
	"github.com/gin-gonic/gin"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {

	err := c.Bind(form)

	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(form)

	if err != nil {
		return http.StatusInternalServerError, e.ERROR
	}
	if !check {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS

}
