package web

import (
	"github.com/iyf/gotool/log"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
)

type templateCache struct {
	ModTime int64
	Content string
}

type site struct {
	*base
	supportSession       bool
	supportCookieSession bool
	supportI18n          bool
	templateFunc         template.FuncMap
	templateCache        map[string]templateCache
	globalTemplate       *template.Template
	Root                 string
	Version              string
}

func (s *site) Init(w http.ResponseWriter, r *http.Request) *site {
	s.base.Init(w, r)

	return s
}

func (s *site) AddTemplateFunc(name string, i interface{}) {
	_, ok := s.templateFunc[name]
	if !ok {
		s.templateFunc[name] = i
	} else {
		log.Warn("<site.AddTemplateFunc> ", "func:"+name+" be added,do not repeat to add")
	}
}

func (s *site) DelTemplateFunc(name string) {
	if _, ok := s.templateFunc[name]; ok {
		delete(s.templateFunc, name)
	}
}

func (s *site) SetTemplateCacheObject(tmplKey, content string, modTime int64) {
	s.templateCache[tmplKey] = templateCache{
		ModTime: modTime,
		Content: content,
	}
}

func (s *site) SetTemplateCache(tmplKey, tmplPath string) {
	if tmplFi, err := os.Stat(tmplPath); err == nil {
		if b, err := ioutil.ReadFile(tmplPath); err == nil {
			s.SetTemplateCacheObject(tmplKey, string(b), tmplFi.ModTime().Unix())
		}
	}

}

func (s *site) GetTemplateCache(tmplKey string) templateCache {
	if tmpl, ok := s.templateCache[tmplKey]; ok {
		return tmpl
	}

	return templateCache{}
}

func (s *site) DelTemplateCache(tmplKey string) {
	delete(s.templateCache, tmplKey)
}
