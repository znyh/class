//
//	  是否满足 M*AAA+N*ABC+DD
//    cards[]代表index->num的键值对,num之和为sum，sum满足sum%3 == 2
//    guiIndex: 代表癞子的索引,CheckHu(cards,MaxCard)计算不带癞子的判胡

package hulib

var GlobalTableMgr *TableMgr

func init() {
	GlobalTableMgr = &TableMgr{}
	GlobalTableMgr.Init()
	GlobalTableMgr.Gen()
}

func CheckHu(cards []int, guiIndex int) bool {

	sum := 0
	for _, c := range cards {
		sum = sum + c
	}

	if sum%3 != 2 {
		return false
	}

	if guiIndex >= MaxCard || guiIndex < 0 {
		guiIndex = MaxCard
	}

	tmpCards := [MaxCard]int{}
	for i := 0; i < len(cards) && i < MaxCard; i++ {
		if cards[i] > 0 {
			tmpCards[i] = cards[i]
		}
	}
	return GetHuInfo(GlobalTableMgr, tmpCards[:], MaxCard, MaxCard, guiIndex)
}
