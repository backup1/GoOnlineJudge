package class

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/restweb"
	"strconv"
	"time"
)

var specialArr = []string{"Standard", "Special"}
var judgeArr = []string{"Pengding", "Running & Judging", "Compile Error", "Accepted", "Runtime Error",
	"Wrong Answer", "Time Limit Exceeded", "Memory Limit Exceeded", "Output Limit Exceeded", "Presentation Error", "System Error"}
var languageArr = []string{"None", "C", "C++", "Java"}
var encryptArr = []string{"None", "Public", "Private", "Password"}
var privilegeArr = []string{"None", "Primary User", "Teacher", "Admin"}

// ShowStatus 根杮status确定状思是坦坯达的
func ShowStatus(status int) bool {
	return status == config.StatusAvailable
}

// ShowSim 是坦显示相似度
func ShowSim(sim int) bool {
	return sim != 0
}

// ShowRatio 显示solve/submit的比率
func ShowRatio(solve int, submit int) (ratio string) {
	if submit == 0 {
		ratio = "0.00%"
	} else {
		ratio = strconv.FormatFloat(float64(solve)/float64(submit)*100, 'f', 2, 64) + "%"
	}
	return
}

// ShowSpecial显示Judge程庝是标准或是特判
func ShowSpecial(num int) (special string) {
	special = specialArr[num]
	return
}

// ShowJudge显示判题结果
func ShowJudge(num int) (judge string) {
	judge = judgeArr[num]
	return
}

// ShowLanguage 显示代砝语言类型
func ShowLanguage(num int) (language string) {
	language = languageArr[num]
	return
}

// ShowEncrypt显示比赛的加密方弝，公开，秝有或者密砝
func ShowEncrypt(num int) (encrypt string) {
	encrypt = encryptArr[num]
	return
}

// LargePU 判断privilege是坦大于普通用户
func LargePU(privilege int) bool {
	return privilege > config.PrivilegePU
}

// ShowPrivilege 显示用户权陝
func ShowPrivilege(privilege int) string {
	return privilegeArr[privilege]
}

// 判断两个ID是坦相等
func SameID(ID1, ID2 string) bool {
	return ID1 == ID2
}

func HasPriv(priv, needpriv int) bool {
	return priv&needpriv > 0
}

func ShowErrFlag(flag uint8) bool {
	return flag == config.FLagER
}

func ShowACFlag(flag uint8) bool {
	return flag == config.FLagAC
}

// ShowTime 将unixtime转杢为北京时间
func ShowTime(unixtime int64) string {

	loc, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(unixtime, 0).In(loc).Format("2006-01-02 15:04:05")
}

// initFuncMap 初始化FuncMap
func initFuncMap() {
	restweb.AddFuncMap("ShowErrFlag", ShowErrFlag)
	restweb.AddFuncMap("ShowACFlag", ShowACFlag)
	restweb.AddFuncMap("ShowRatio", ShowRatio)
	restweb.AddFuncMap("ShowSpecial", ShowSpecial)
	restweb.AddFuncMap("ShowJudge", ShowJudge)
	restweb.AddFuncMap("ShowLanguage", ShowLanguage)
	restweb.AddFuncMap("ShowEncrypt", ShowEncrypt)
	restweb.AddFuncMap("ShowPrivilege", ShowPrivilege)
	restweb.AddFuncMap("LargePU", LargePU)
	restweb.AddFuncMap("ShowStatus", ShowStatus)
	restweb.AddFuncMap("ShowSim", ShowSim)
	restweb.AddFuncMap("HasPriv", HasPriv)
	restweb.AddFuncMap("ShowTime", ShowTime)
}
