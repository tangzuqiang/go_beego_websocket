package controllers

import (
	"github.com/astaxie/beego"
	data2 "webapi/util/data"
)

// BaseController operations for Base
type BaseController struct {
	beego.Controller
}

func (c *BaseController) success(args ...interface{}) {
	var format data2.Format
	for k, v := range args {
		switch k {
		case 0:
			format.Message = v
		case 1:
			format.Data = v
		}
	}
	if format.Message == nil {
		format.Message = "ok"
	}
	format.Status = 200
	c.Data["json"] = format
	c.ServeJSON()
}

func (c *BaseController) error(args ...interface{}) {
	var format data2.Format
	for k, v := range args {
		switch k {
		case 0:
			format.Message = v
		case 1:
			format.Data = v
		case 3:
			format.Status = v
		}
	}
	if format.Message == nil {
		format.Message = "error"
	}
	if format.Data == nil {
		format.Data = `{}`
	}
	if format.Status == nil {
		format.Status = 400
	}
	c.Data["json"] = format
	c.ServeJSON()
}
