package model

var (
	APPID                          = "okeyhall-service"
	TimeLayout                     = "2006-01-02 15:04:05"
	OkeyHall_UserInfo_Hall_Key     = "HOkeyHall:%d:UserInfo:Hall"     // HOkeyHall:[GameID]:UserInfo:Hall key:[userid] value:[time]
	OkeyHall_UserInfo_Game_Key     = "HOkeyHall:%d:UserInfo:Game"     // HOkeyHall:[GameID]:UserInfo:Game key:[userid] value:[gameinfo]
	OkeyHall_UserInfo_GameWait_Key = "HOkeyHall:%d:UserInfo:GameWait" // HOkeyHall:[GameID]:UserInfo:GameWait key:[userid] value:[time]
	OkeyHall_TableSN_Key           = "HOkeyHall:%d:TableSN"           // HOkeyHall:[GameID]:TableSN key:[TableSN] value:[TableInfo] 正在使用的房间号
	OkeyHall_RobotBalance_Key      = "HOkeyHall:%d:RobotBalance"      // HOkeyHall:[GameID]:RobotBalance key:[GroupId] value:[int64] 机器人当前输赢池
)

type Kratos struct {
	Hello string
}
