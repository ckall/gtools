package ding

const dingTalkURL = "https://oapi.dingtalk.com/robot/send?"
const dtmdFormat = "[%s](dtmd://dingtalkclient/sendMessage?content=%s)"
const formatSpliter = "$$"

const (
	h1    MarkType = "h1"
	h2    MarkType = "h2"
	h3    MarkType = "h3"
	h4    MarkType = "h4"
	h5    MarkType = "h5"
	h6    MarkType = "h6"
	red   MarkType = "red"
	blue  MarkType = "blue"
	green MarkType = "green"
	gold  MarkType = "gold"
	n     MarkType = ""
)

var hMap = map[MarkType]string{
	h1:    "# %s",
	h2:    "## %s",
	h3:    "### %s",
	h4:    "#### %s",
	h5:    "##### %s",
	h6:    "###### %s",
	red:   "<font color=#ff0000>%s</font>",
	blue:  "<font color=#1E90FF>%s</font>",
	green: "<font color=#00CD66>%s</font>",
	gold:  "<font color=#FFD700>%s</font>",
}
