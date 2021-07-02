package calc

const (
	CtUnknown         = 0x00000
	CtPingHU          = 0x00001
	CtPengPeng        = 0x00002
	CtYitiaoLong      = 0x00004
	CtKaZhang         = 0x00008
	CtNormal7Dui      = 0x00010
	CtSuper7DuiSolo   = 0x00020
	CtSuper7DuiDouble = 0x00040
	CtSuper7DuiThree  = 0x00100
	Ct13Yao           = 0x00200
	CtQingYise        = 0x00400
	CtHunYise         = 0x00800
	CtZiYise          = 0x01000
	CtHuaLong         = 0x02000
)

const (
	OpEatLeft = iota + 1
	OpEatCenter
	OpEatRight
	OpPeng
	OpMingGang
	OpAnGang
	OpBuGang
)

//组合子项
type TagWeaveItem struct {
	cbWeaveKind  int //组合类型
	cbCenterCard int //中心扑克
	cbPublicCard int //公开标志
	wProvideUser int //供应用户
}
