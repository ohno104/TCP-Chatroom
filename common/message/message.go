package message

const (
	LoginMsgType            = "LoginMsg"
	LoginResMsgType         = "LoginResMeg"
	RegisterMsgType         = "RegisterMsg"
	RegisterResType         = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType              = "SmsMsg"
)

const (
	UserOnline = iota
	UserOffline
	UserBusy
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

type LoginResMsg struct {
	Code    int    `json:"code"`
	UserIds []int  `json:"userIds"`
	Error   string `json:"error"`
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

type NotifyUserStatusMsg struct {
	UserId int `json:"usesrId"`
	Status int `json:"status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User
}
