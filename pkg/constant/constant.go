package constant

const (
	MsgNone     = 0
	MsgText     = 1
	MsgImage    = 2
	MsgAudio    = 3
	MsgVideo    = 4
	MsgFile     = 5
	MsgLocation = 6
	MsgQuote    = 7  // 引用消息
	MsgMerger   = 8  // 合并消息
	MsgTip      = 9  // 提示消息
	MsgRecall   = 10 // 撤回消息
	// 以下为透传消息
	MsgTyping          = 3000 // 正在输入
	MsgDeliveryReceipt = 3001 // 送达回执
	MsgReadReceipt     = 101  // 已读回执
)

const (
	ConvTypeC2C   = 1 // 单聊
	ConvTypeGroup = 2 // 群聊
	// ConvTypeSystem        = 3 // 系统
	// ConvTypeCustomService = 4 // 客服
)

const (
	MsgKeepDays  = 7   // 离线消息保留天数
	ConvKeepDays = 180 // 会话保留天数
)

const (
	Online     = 1
	PushOnline = 2
	Offline    = 3
)

const (
	PushOnlineKeepDays = 7 // 推送在线状态保持天数
)
