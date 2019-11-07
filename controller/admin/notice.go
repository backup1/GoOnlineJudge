package admin

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/restweb"
	"html/template"
	"net/http"
	"os"
)

type AdminNotice struct {
	class.Controller
} //@Controller

//@URL: /admin/notice/ @method: GET
func (n *AdminNotice) Index() {
	restweb.Logger.Debug("Admin notice index")
	n.Output["Msg"] = string(n.Output["Msg"].(template.HTML))
	n.Output["IsNotice"] = true
	n.RenderTemplate("view/admin/layout.tpl", "view/admin/notice.tpl")
}

//@URL: /admin/notice/ @method: POST
func (n *AdminNotice) Edit() {
	restweb.Logger.Debug("Admin notice edit")

	msg := n.Input.Get("msg")
	file, err := os.Create("view/admin/msg.txt")
	if err != nil {
		return
	}
	defer file.Close()
	file.Write([]byte(msg))
	n.Redirect("/", http.StatusFound)
}
