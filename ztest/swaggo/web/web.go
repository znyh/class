package web

type RspCode int

const (
	Succ RspCode = iota
	Failed
	ArgInvalid
	NotFound
)

type Login struct {
	User     string `form:"user" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type userInfo struct {
	UserID uint64 `json:"uid" form:"uid"` // 用户ID
	Name   string `json:"name" form:"name"`
	Age    int    `json:"age"  form:"age"`
	Gold   int64  `json:"gold" form:"gold"`
}

type userChange struct {
	UserID uint64 `json:"uid" form:"uid"` // 用户ID
	Typ    uint32 `json:"type" form:"type"`
	Change int64  `json:"change" form:"change"`
	Event  int    `json:"event"  form:"event"`
}
