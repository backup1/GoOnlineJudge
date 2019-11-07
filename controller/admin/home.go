package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/restweb"
)

type AdminHome struct {
	class.Controller
} //@Controller

//@URL: /admin/ @method: GET
func (hc *AdminHome) Home() {
	restweb.Logger.Debug("Admin Home")

	hc.Output["IsHome"] = true
	hc.Output["Title"] = "Admin - Home"
	hc.RenderTemplate("view/admin/layout.tpl", "view/admin/home.tpl")
}
