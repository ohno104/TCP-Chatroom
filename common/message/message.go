package message

const (
	LoginMsgType    = "LoginMes"
	LoginResMsgType = "LoginResMes"
	RegisterMsgType = "RegisterMsg"
)

var (
	SUCCESS      = 200
	UNREGISTERED = 500
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMsg struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginRes struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
