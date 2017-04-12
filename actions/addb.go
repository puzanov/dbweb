package actions

import (
	"fmt"

	"github.com/Unknwon/i18n"
	"github.com/go-xorm/dbweb/middlewares"
	"github.com/go-xorm/dbweb/models"
	"github.com/tango-contrib/binding"
	"github.com/tango-contrib/flash"
	"github.com/tango-contrib/renders"
	"github.com/tango-contrib/xsrf"
)

type Addb struct {
	RenderBase

	binding.Binder
	xsrf.Checker
	middlewares.Auther
	flash.Flash
}

func (c *Addb) Get() error {
	engines, err := models.FindEngines()
	if err != nil {
		return err
	}

	return c.Render("add.html", renders.T{
		"dbs":          SupportDBs,
		"flash":        c.Flash.Data(),
		"engines":      engines,
		"XsrfFormHtml": c.XsrfFormHtml(),
		"IsLogin":      c.IsLogin(),
		"isAdd":        true,
	})
}

func (c *Addb) Post() {
	var engine models.Engine
	engine.Name = c.Form("name")
	engine.Driver = c.Form("driver")
	host := c.Form("host")
	port := c.Form("port")
	dbname := c.Form("dbname")
	username := c.Form("username")
	passwd := c.Form("passwd")

	engine.DataSource = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		username, passwd, host, port, dbname)

	if engine.Driver == "sqlite3" {
		engine.DataSource = dbname
	}

	/*if err := c.MapForm(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_param"))
		c.Redirect("/addb")
		return
	}*/

	if err := models.AddEngine(&engine); err != nil {
		c.Flash.Set("ErrAdd", i18n.Tr(c.CurLang(), "err_add_failed"))
		c.Redirect("/addb")
		return
	}

	c.Redirect("/")
}
