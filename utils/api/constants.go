package api

var MapMap = map[string]string{
	"Olympus":       "奥林匹斯",
	"World's Edge":  "世界尽头",
	"Storm Point":   "风暴点",
	"Party crasher": "派对破坏者",
	"Phase runner":  "相位穿梭器",
	"Overflow":      "熔岩流",
	"Habitat":       "4号栖息地",
	"Drop Off":      "原料场",
	"King's Canyon": "诸王峡谷",
	"Kings Canyon":  "诸王峡谷",
	"Encore":        "再来一次",
	"Broken Moon":   "殘月",
	"Unknown":       "不可用",
}

func MaptoZh(mapName string) string {
	if MapMap[mapName] != "" {
		return MapMap[mapName]
	} else {
		return mapName
	}
}
