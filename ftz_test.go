package ftz

import (
	"testing"
)

func TestFTZ(t *testing.T) {
	simplifiedString := "我爱学习，我爱吃饭。"
	traditionalString := SimplifiedToTraditional(simplifiedString)

	if TraditionalToSimplified(traditionalString) != simplifiedString {
		t.Fatal()
	}

	if TraditionalToSimplified("吃飯") != "吃饭" {
		t.Fatal()
	}
	if SimplifiedToTraditional("干燥") != "乾燥" {
		t.Fatal()
	}

	if TraditionalToSimplified("乾燥") != "干燥" {
		t.Fatal()
	}
	if SimplifiedToTraditional("吃饭") != "吃飯" {
		t.Fatal()
	}

	if !ContainsTraditional("吃飯") {
		t.Fatal()
	}
	if ContainsTraditional("吃饭") {
		t.Fatal()
	}
}

func TestCharTableLength(t *testing.T) {
	if len(s2tSimplifiedChars) != len(s2tTraditionalChars) {
		t.Fatalf("length mismatch: simplified=%d traditional=%d", len(s2tSimplifiedChars), len(s2tTraditionalChars))
	}
	if len(t2sTraditionalChars) != len(t2sSimplifiedChars) {
		t.Fatalf("t2s length mismatch: traditional=%d simplified=%d", len(t2sTraditionalChars), len(t2sSimplifiedChars))
	}
}

func TestT2STableInvariants(t *testing.T) {
	simplifiedSources := make(map[rune]struct{}, len(s2tSimplifiedChars))
	for _, c := range s2tSimplifiedChars {
		simplifiedSources[c] = struct{}{}
	}

	seen := make(map[rune]struct{}, len(t2sTraditionalChars))
	for _, c := range t2sTraditionalChars {
		if _, exists := seen[c]; exists {
			t.Fatalf("duplicate character %q in t2sTraditionalChars", c)
		}
		seen[c] = struct{}{}

		if _, exists := simplifiedSources[c]; exists {
			t.Fatalf("t2s source %q is also an s2t simplified source", c)
		}
	}
}

func TestKnownDuplicateTraditionalChoices(t *testing.T) {
	tests := map[rune]rune{
		'壟': '垄',
		'墊': '垫',
		'濫': '滥',
		'強': '强',
		'滾': '滚',
		'線': '线',
		'謔': '谑',
		'鍾': '钟',
		'飆': '飙',
		'餘': '余',
		'鯰': '鲶',
		'鹼': '碱',
		'堿': '碱',
	}

	for traditional, simplified := range tests {
		if got := t2sMap[traditional]; got != simplified {
			t.Fatalf("t2sMap[%q] = %q, want %q", traditional, got, simplified)
		}
	}
}

func TestTraditionalToSimplifiedLiteraryMappings(t *testing.T) {
	tests := map[string]string{
		"鉅癧痾糹鷽鰺鏇":      "钜疬疴纟鸴鲹镟",
		"鉅鬁痾糹鷽鯵鏇":      "钜疬疴纟鸴鲹镟",
		"瀋陽諮詢逕自鹼性韝":    "沈阳咨询径自碱性鞲",
		"唿嗬堿屭崳嶴簷鬁鯵鼇":   "唿嗬碱屭崳嶴檐疬鲹鳌",
		"乾燥的頭髮與飯":      "干燥的头发与饭",
		"天乾地支，看著他唱著歌。": "天干地支，看着他唱着歌。",
	}

	for traditional, simplified := range tests {
		if got := TraditionalToSimplified(traditional); got != simplified {
			t.Fatalf("TraditionalToSimplified(%q) = %q, want %q", traditional, got, simplified)
		}
	}
}

func TestTraditionalToSimplifiedCommonSupplementChars(t *testing.T) {
	tests := map[string]string{
		"麵包與麵條":  "面包与面条",
		"鐘聲與時鐘":  "钟声与时钟",
		"製造複製":   "制造复制",
		"一隻貓":    "一只猫",
		"臺灣颱風檯面": "台湾台风台面",
		"匯款與飢餓":  "汇款与饥饿",
		"鬍鬚":     "胡须",
		"抽籤儘管收穫": "抽签尽管收获",
		"僕人與蘿蔔":  "仆人与萝卜",
		"採用佈置輕鬆": "采用布置轻松",
		"兇手颳風老闆": "凶手刮风老板",
	}

	for traditional, simplified := range tests {
		if got := TraditionalToSimplified(traditional); got != simplified {
			t.Fatalf("TraditionalToSimplified(%q) = %q, want %q", traditional, got, simplified)
		}
	}
}

func TestTraditionalToSimplifiedCommonVariants(t *testing.T) {
	tests := map[string]string{
		"閒聊啓動衆人祕密": "闲聊启动众人秘密",
		"羣眾嘗試嚐鮮剛纔": "群众尝试尝鲜刚才",
		"席捲擡頭喫飯香菸": "席卷抬头吃饭香烟",
		"山峯屋簷鰲頭鼇頭": "山峰屋檐鳌头鳌头",
		"麽剋裊勛衞嫺":   "么克袅勋卫娴",
		"贊綫牀竈巖":    "赞线床灶岩",
	}

	for traditional, simplified := range tests {
		if got := TraditionalToSimplified(traditional); got != simplified {
			t.Fatalf("TraditionalToSimplified(%q) = %q, want %q", traditional, got, simplified)
		}
	}
}

func TestContainsTraditionalPhraseRules(t *testing.T) {
	tests := []string{
		"乾燥",
		"風乾",
		"為甚麼",
		"天乾地支",
		"看著他",
		"唱著歌",
	}

	for _, text := range tests {
		if !ContainsTraditional(text) {
			t.Fatalf("ContainsTraditional(%q) = false, want true", text)
		}
	}
}

func TestTraditionalToSimplifiedTTSGuards(t *testing.T) {
	tests := map[string]string{
		"蓧":       "蓧",
		"看著作":     "看著作",
		"唱著名歌曲":   "唱著名歌曲",
		"看著他唱著歌。": "看着他唱着歌。",
		"乾杯、乾淨、乾旱、乾脆、乾貨、乾糧":     "干杯、干净、干旱、干脆、干货、干粮",
		"風乾、曬乾、晾乾、烘乾、吹乾、擦乾、口乾":  "风干、晒干、晾干、烘干、吹干、擦干、口干",
		"為甚麼、爲甚麼、為什麼、爲什麼、甚麼、什麼": "为什么、为什么、为什么、为什么、什么、什么",
		"乾隆與乾坤": "乾隆与乾坤",
	}

	for traditional, simplified := range tests {
		if got := TraditionalToSimplified(traditional); got != simplified {
			t.Fatalf("TraditionalToSimplified(%q) = %q, want %q", traditional, got, simplified)
		}
	}
}

func TestSimplifiedToTraditionalGanWords(t *testing.T) {
	tests := map[string]string{
		"天干地支": "天干地支",
		"天干":   "天干",
		"干支":   "干支",
		"干涉":   "干涉",
		"干预":   "干預",
		"若干":   "若干",
		"榨干":   "榨乾",
		"风干":   "風乾",
		"晒干":   "曬乾",
		"晾干":   "晾乾",
		"烘干":   "烘乾",
		"吹干":   "吹乾",
		"擦干":   "擦乾",
		"树干":   "樹幹",
		"干部":   "幹部",
	}

	for simplified, traditional := range tests {
		if got := SimplifiedToTraditional(simplified); got != traditional {
			t.Fatalf("SimplifiedToTraditional(%q) = %q, want %q", simplified, got, traditional)
		}
	}
}

func TestTraditionalToSimplifiedDoesNotBreakSimplifiedText(t *testing.T) {
	tests := []string{
		"劈开木头",
		"糊涂一点",
		"脊梁挺直",
		"镟床加工",
		"钜细靡遗",
		"矽谷",
		"爿字旁",
	}

	for _, text := range tests {
		if got := TraditionalToSimplified(text); got != text {
			t.Fatalf("TraditionalToSimplified(%q) = %q, want unchanged", text, got)
		}
		if ContainsTraditional(text) {
			t.Fatalf("ContainsTraditional(%q) = true, want false", text)
		}
	}
}

func TestSimplifiedToTraditionalDoesNotCorruptNormalText(t *testing.T) {
	tests := map[string]string{
		"呼吸": "呼吸",
		"呵呵": "呵呵",
		"哄骗": "哄騙",
		"坂本": "坂本",
		"畲族": "畲族",
		"拖沓": "拖沓",
		"藁城": "藁城",
		"憷场": "憷場",
		"愍然": "愍然",
		"癯瘦": "癯瘦",
		"咔嚓": "咔嚓",
		"嚯":  "嚯",
		"埝":  "埝",
		"莜麦": "莜麥",
	}

	for simplified, traditional := range tests {
		if got := SimplifiedToTraditional(simplified); got != traditional {
			t.Fatalf("SimplifiedToTraditional(%q) = %q, want %q", simplified, got, traditional)
		}
	}
}

func TestSimplifiedToTraditionalZhe(t *testing.T) {
	tests := map[string]string{
		"走着": "走著",
		"说着": "說著",
		"笑着": "笑著",
		"拿着": "拿著",
		"睡着": "睡著",
		"著名": "著名",
		"著作": "著作",
	}

	for simplified, traditional := range tests {
		if got := SimplifiedToTraditional(simplified); got != traditional {
			t.Fatalf("SimplifiedToTraditional(%q) = %q, want %q", simplified, got, traditional)
		}
	}
}

func TestSimplifiedToTraditionalFixes(t *testing.T) {
	tests := map[string]string{
		"钜疬疴纟鸴鲹镟碹": "鉅癧痾糹鷽鰺鏇碹",
	}

	for simplified, traditional := range tests {
		if got := SimplifiedToTraditional(simplified); got != traditional {
			t.Fatalf("SimplifiedToTraditional(%q) = %q, want %q", simplified, got, traditional)
		}
	}
}
