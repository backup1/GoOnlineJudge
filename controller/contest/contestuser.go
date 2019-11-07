package contest

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/restweb"
)

//@URL: /contests/(\d+)/password @method: GET
func (c *Contest) PasswordPage(Cid string) {
	c.InitContest(Cid)
	restweb.Logger.Debug("herehr")
	if c.ContestDetail.Encrypt != config.EncryptPW {
		c.Error("No such page", 400)
	}
	c.RenderTemplate("view/layout.tpl", "view/contest/passwd.tpl")
}

//@URL: /contests/(\d+)/password @method: POST
func (c *Contest) Password(Cid string) {
	c.InitContest(Cid)

	if c.ContestDetail.Encrypt != config.EncryptPW {
		c.Error("No such page", 400)
	}

	passwd := c.R.FormValue("password")
	restweb.Logger.Debug(c.ContestDetail.Argument.(string), passwd)
	if passwd == c.ContestDetail.Argument.(string) {
		c.SetSession(Cid+"pass", passwd)
		c.W.WriteHeader(200)
	} else {
		c.Error("incorrect password", 400)
	}
}
