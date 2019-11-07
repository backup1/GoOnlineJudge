package controller

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/restweb"
)

type OSCController struct {
	class.Controller
} //@Controller

//@URL: /osc @method: GET
func (oc *OSCController) Index() {
	restweb.Logger.Debug("OSC Page")

	oc.Output["Title"] = "ZJGSU OSC"
	oc.Output["IsOSC"] = true
	oc.RenderTemplate("view/layout.tpl", "view/osc.tpl")
}
