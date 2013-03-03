package web

import (
	"text/template"
)

type Document struct {
	Close              bool	//关闭文档
	GenerateHtml       bool	//生成Html
	Static             string
	Theme              string
	Attr               map[string]string
	Css                map[string]string
	Js                 map[string]string
	Img                map[string]string
	GlobalCssFile      string
	GlobalJsFile       string
	GlobalIndexCssFile string
	GlobalIndexJsFile  string
	IndexCssFile       string
	IndexJsFile        string
	Hide               bool //是否渲染模板
	Func               template.FuncMap
	Title              string
	Body               interface{}
}
