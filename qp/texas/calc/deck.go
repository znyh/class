package calc

var (
	oneDeck = []int32{
		01, 02, 03, 04, 05, 06, 07, 8, 9, 10, 11, 12, 13, //方块
		14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, //梅花
		27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, //红桃
		40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, //黑桃
	}

	descDeck = map[int32]string{
		1: "方块A", 2: "方块2", 3: "方块3", 4: "方块4", 5: "方块5", 6: "方块6", 7: "方块7", 8: "方块8", 9: "方块9", 10: "方块10", 11: "方块J", 12: "方块Q", 13: "方块K",
		14: "梅花A", 15: "梅花2", 16: "梅花3", 17: "梅花4", 18: "梅花5", 19: "梅花6", 20: "梅花7", 21: "梅花8", 22: "梅花9", 23: "梅花10", 24: "梅花J", 25: "梅花Q", 26: "梅花K",
		27: "红桃A", 28: "红桃2", 29: "红桃3", 30: "红桃4", 31: "红桃5", 32: "红桃6", 33: "红桃7", 34: "红桃8", 35: "红桃9", 36: "红桃10", 37: "红桃J", 38: "红桃Q", 39: "红桃K",
		40: "黑桃A", 41: "黑桃2", 42: "黑桃3", 43: "黑桃4", 44: "黑桃5", 45: "黑桃6", 46: "黑桃7", 47: "黑桃8", 48: "黑桃9", 49: "黑桃10", 50: "黑桃J", 51: "黑桃Q", 52: "黑桃K",
	}
)

func OneDeck() []int32 {
	dst := make([]int32, len(oneDeck))
	copy(dst, oneDeck)
	return dst
}

func ShuffleDeck() []int32 {
	return Shuffle(OneDeck())
}
