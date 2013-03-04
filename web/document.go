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
	CssPath            string //当前view对应的css目录
	JsPath             string //当前view对应的Js目录
	ImgPath            string //当前view对应的Img目录
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
