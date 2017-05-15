package model

/*
Message
---------------------------------------------
| Token | SessionId | From | To | parameter |
---------------------------------------------

from
-------------
| Type | Id |
-------------

to
-------------
| Type | Id |
-------------

parameter
-------------------------------------------------
| Action | Target | Data1 | Data2 | ... | DataN |
-------------------------------------------------

data
-----------------------
| Type | Id | Content |
-----------------------
*/

const (
	// ================================================= message role type
	// 类型
	// for user 1 ~ 100
	MSG_ROLE_TYPE_USER_WEB         = 1 // web客户端
	MSG_ROLE_TYPE_USER_PC_APP      = 2 // PC应用程序客户端
	MSG_ROLE_TYPE_USER_IOS_APP     = 3 // IOS移动客户端
	MSG_ROLE_TYPE_USER_ANDROID_APP = 4 // 安卓移动客户端
	// for system
	//MSG_ROLE_TYPE_SYS_OPCODE        = 0x05
	// for system account 101 ~ 200
	MSG_ROLE_TYPE_SYS_ACCOUNT = 101
	// for system env 201 ~ 300
	MSG_ROLE_TYPE_SYS_ENV = 201
	// for system data 301 ~ 400
	MSG_ROLE_TYPE_SYS_DATA = 301
	// for system job 401 ~ 500
	MSG_ROLE_TYPE_SYS_JOB = 401
	// for system job processor 501 ~ 600
	MSG_ROLE_TYPE_SYS_JOB_PROCESSOR = 501

	// for app service
	MSG_ROLE_TYPE_SERVICE = 600

	ACCOUNT_SERVICE   = 601
	ALGORITHM_SERVICE = 602
	BUSSINESS_SERVICE = 603
	DATA_SERVICE      = 604
	LOG_SERVICE       = 605
	LOGIN_SERVICE     = 606
	STATUS_SERVICE    = 607
	SUMMARY_SERVICE   = 608

	// ================================================= message parameter action
	//增删改查

	MSG_PARAM_ACTION_CREATE   = 1
	MSG_PARAM_ACTION_GET      = 2
	MSG_PARAM_ACTION_DELETE   = 3
	MSG_PARAM_ACTION_UPDATE   = 4
	MSG_PARAM_ACTION_REGISTER = 5
	MSG_PARAM_ACTION_LOGIN    = 6
	MSG_PARAM_ACTION_STOP     = 7
	MSG_PARAM_ACTION_START    = 8

	// ================================================= message parameter target
	// 处理目标对象
	// for user 1 ~ 100
	MSG_PARAM_TYPE_USER         = 1
	MSG_PARAM_TYPE_ACTION       = 2
	MSG_PARAM_TYPE_STATUS       = 3
	MSG_PARAM_TYPE_MEMORY_USAGE = 4
	MSG_PARAM_TYPE_CPU_USAGE    = 5
	MSG_PARAM_TYPE_SUMMARY      = 6
	MSG_PARAM_TYPE_ALGORITHM    = 7
	MSG_PARAM_TYPE_INPUT_FILES  = 8
	MSG_PARAM_TYPE_USER_LIST    = 9

	// for project 201 ~ 300
	MSG_PARAM_TYPE_PROJECT        = 201
	MSG_PARAM_TYPE_PROJECT_LIST   = 202
	MSG_PARAM_TYPE_PROJECT_LOG    = 203
	MSG_PARAM_TYPE_PROJECT_STATUS = 204

	// for job 301 ~ 400
	MSG_PARAM_TYPE_JOB             = 301
	MSG_PARAM_TYPE_JOB_LIST        = 302
	MSG_PARAM_TYPE_JOB_LOG         = 303
	MSG_PARAM_TYPE_JOB_STATUS      = 304
	MSG_PARAM_TYPE_JOB_LIST_STATUS = 305
	MSG_PARAM_TYPE_JOB_WRITE       = 306
	MSG_PARAM_TYPE_JOB_READ        = 307
	MSG_PARAM_TYPE_JOB_CHECK       = 308
	MSG_PARAM_TYPE_JOB_UPDATE      = 309

	// for job processor 401 ~ 500
	MSG_PARAM_TYPE_POD         = 401
	MSG_PARAM_TYPE_POD_STATUS  = 402
	MSG_PARAM_TYPE_POD_LOG     = 403
	MSG_PARAM_TYPE_CURRENT_POD = 404

	// for env 501 ~ 600
	MSG_PARAM_TYPE_ENV = 501

	// for data 601 ~ 700
	MSG_PARAM_TYPE_DATA = 601

	// ================================================= message result code
	MSG_RESULTCODE_SUCCESS = 1
	MSG_RESULTCODE_FAILED  = -1
)

type MessageParameterData struct {
	Type    int    `json:"type"`    // 数据类型
	Content string `json:"content"` // 数据内容
}

type MessageParameter struct {
	Action int                     `json:"action"` // 增删改查
	Target int                     `json:"target"` // 处理目标对象
	Data   []*MessageParameterData `json:"data"`   // 请求数据
}

type MessageRole struct {
	Type int   `json:"type"` // 类型
	Id   int64 `json:"id"`   // 唯一ID
}

type MessageRequest struct {
	Token     string            `json:"token"`     // message有效性判断
	SessionId string            `json:"sessionId"` // 会话ID
	From      *MessageRole      `json:"from"`      // 请求发起者
	To        *MessageRole      `json:"to"`        // 请求接受者
	Parameter *MessageParameter `json:"parameter"` // 参数
}

type MessageResponse struct {
	Token      string `json:"token"`      // message有效性判断
	SessionId  string `json:"sessionId"`  // 会话ID
	ResultCode int    `json:"resultCode"` // 成功失败标识号
	Reason     string `json:"reason"`     // 成功失败原因
	Target     int    `json:"target"`     // 返回数据类型
	Result     string `json:"result"`     // 返回数据内容
}

type Response struct {
	Status     int    `json:"status"`
	RetryCount int    `json:"retryCount"`
	Reason     string `json:"reason"`
	Result     string `json:"result"`
}
