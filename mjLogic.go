package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type MjSpecialCase int

const (
	MJ_TYPES = "WTB" // W: 万, T: 条, B: 大饼
	MJ_TYPES_NUM = "_123456789"

	SpecialCase_None MjSpecialCase = 0
	SpecialCase_SevenPair MjSpecialCase = 1 // 特殊类型 7小对
)

// 万
var MJ_W = []string{"W1", "W2", "W3", "W4", "W5", "W6", "W7", "W8", "W9"}
// 条
var MJ_T = []string{"T1", "T2", "T3", "T4", "T5", "T6", "T7", "T8", "T9"}
// 大饼(筒)
var MJ_B = []string{"B1", "B2", "B3", "B4", "B5", "B6", "B7", "B8", "B9"}
// 字, ZD: 东, ZN: 南, ZX: 西, ZB: 北, ZZ: 中, ZF: 發, ZM: 白板(门)
var MJ_Z = []string{"ZD", "ZN", "ZX", "ZB", "ZZ", "ZF", "ZM"}
// "HC": 春, "HX": 夏, "HQ": 秋, "HD": 冬, "HM": 梅, "HL": 兰, "HZ": 竹, "HJ": 菊
// 花, 每种只有一个
var MJ_H = []string{"HC", "HX", "HQ", "HD", "HM", "HL", "HZ", "HJ"}

var MJ_ALL []string

// 所有有效的麻将牌
func GetAllValidMjTile() []string {
	if MJ_ALL == nil {
		MJ_ALL = []string{}
		MJ_ALL = append(MJ_ALL, MJ_W...)
		MJ_ALL = append(MJ_ALL, MJ_T...)
		MJ_ALL = append(MJ_ALL, MJ_B...)
		MJ_ALL = append(MJ_ALL, MJ_Z...)
		MJ_ALL = append(MJ_ALL, MJ_H...)
	}
	
	return MJ_ALL[:]
}

func ValidMjTileStr(tileStr string) bool {
	allMjTile := GetAllValidMjTile()
	for _, v2 := range allMjTile {
		if tileStr == v2 {
			return true
		}
	}

	return false
}

// 生成一副完整牌
func AllMJPai() []string {
	arrType := []string{}
	arrType = append(arrType, MJ_W...)
	arrType = append(arrType, MJ_T...)
	arrType = append(arrType, MJ_B...)
	arr := make([]string, 0, len(arrType) * 4 + len(MJ_H))
	for _, v := range arrType {
		arr = append(arr, v, v, v, v)
	}
	arr = append(arr, MJ_H...) // 花, 每种只有一个

	return arr
}

// 生成一副牌,不包含花
func AllMJPaiWithoutHua() []string {
	arrType := []string{}
	arrType = append(arrType, MJ_W...)
	arrType = append(arrType, MJ_T...)
	arrType = append(arrType, MJ_B...)
	arrType = append(arrType, MJ_Z...)
	arr := make([]string, 0, len(arrType) * 4)
	for _, v := range arrType {
		arr = append(arr, v, v, v, v)
	}

	return arr
}

// 生成一副牌,不包含花和字
func AllMJPaiWithoutHuaZi() []string {
	arrType := []string{}
	arrType = append(arrType, MJ_W...)
	arrType = append(arrType, MJ_T...)
	arrType = append(arrType, MJ_B...)
	arr := make([]string, 0, len(arrType) * 4)
	for _, v := range arrType {
		arr = append(arr, v, v, v, v)
	}

	return arr
}

// 打乱一副牌
func Shuffle(allPai []string) []string {
	aLen := len(allPai)
	rand.Seed(time.Now().Unix())
	rand.Shuffle(aLen, func(i, j int){
		allPai[i], allPai[j] = allPai[j], allPai[i]
	})
	return allPai
}

// 给牌排序
func SortMjPai(mjPai []string){
	sort.Strings(mjPai) // 这里按字母顺序排列
	// fmt.Println("SortMjPai: ", mjPai)
}

func isSameType(mj []string) bool{
	if len(mj) == 0 {
		log.Println("isSameType mj len is 0")
		return true
	}

	t := mj[0][0]
	for _, v := range mj {
		if t != v[0]  {
			return false
		}
	}

	return true
}

// 顺子
var tmpShunZiArr = make([]string, 3)
func isShunZi(sortedMjArr []string) bool{
	if isSameType(sortedMjArr) && len(sortedMjArr) == 3 {
		copy(tmpShunZiArr, sortedMjArr)
		SortMjPai(tmpShunZiArr)
		
		i1 := strings.Index(MJ_TYPES_NUM, string(tmpShunZiArr[0][1]))
		i2 := strings.Index(MJ_TYPES_NUM, string(tmpShunZiArr[1][1]))
		i3 := strings.Index(MJ_TYPES_NUM, string(tmpShunZiArr[2][1]))
		// fmt.Println("isSunZi tmpShunZiArr: ", mjArr, tmpShunZiArr, i1,i2,i3)
		if i1 + 1 == i2 && i2 + 1 == i3 {
			return true
		}
	}

	if len(sortedMjArr) != 3 {
		fmt.Println("isSunZi len mjArr != 3, ", sortedMjArr)
	}
	return false
}
// 从 sortedMjArr 中选出一个顺子, sortedMjArr 必需是用 SortMjPai 排好序的
func selectOneShunZi(sortedMjArr []string) []string {
	aLen := len(sortedMjArr);
	if aLen >= 3 {
		for j := 0; j < aLen - 2; j ++ {
			for j1 := j + 1; j1 < aLen; j1++ {
				for j2 := j1 + 1; j2 < aLen; j2++ {
					i1 := strings.Index(MJ_TYPES_NUM, string(sortedMjArr[j][1]))
					i2 := strings.Index(MJ_TYPES_NUM, string(sortedMjArr[j1][1]))
					i3 := strings.Index(MJ_TYPES_NUM, string(sortedMjArr[j2][1]))
					if i1 + 1 == i2 && i2 + 1 == i3 {
						mjArr := []string{sortedMjArr[j], sortedMjArr[j1], sortedMjArr[j2]}
						// fmt.Println("selectOneShunZi: ", mjArr, sortedMjArr)
						return mjArr
					}
				}
			}
		}
	}
	return nil
}

// 杠
func isGang(mjArr []string) bool {
	if len(mjArr) == 4 {
		first := mjArr[0]
		for i := 1; i < len(mjArr); i++ {
			if first != mjArr[i] {
				return false
			}
		}

		return true
	}

	fmt.Println("isGang len mjArr != 4, ", mjArr)
	return false
}
// 从 mjArr 中选出一个杠
func selectOneGang(mjArr []string) []string {
	aLen := len(mjArr)
	if aLen >= 4 {
		for i := 0; i < aLen - 3; i++ {
			for i2 := i + 1; i2 < aLen; i2++ {
				for i3 := i2+1; i3 < aLen; i3++ {
					for i4 := i3+1; i4 < aLen; i4++ {
						if mjArr[i] == mjArr[i2] && mjArr[i2] == mjArr[i3] && mjArr[i3] == mjArr[i4] {
							return []string{mjArr[i], mjArr[i2], mjArr[i3], mjArr[i4]}
						}
					}
				}
			}

		}
	}
	return nil
}

// 碰
func isPeng(mjArr []string) bool {
	if len(mjArr) == 3 {
		first := mjArr[0]
		for i := 1; i < len(mjArr); i++ {
			if first != mjArr[i] {
				return false
			}
		}

		return true
	}

	fmt.Println("isGang len mjArr != 3, ", mjArr)
	return false
}
// 从 mjArr 中选出一个碰
func selectOnePeng(mjArr []string) []string {
	aLen := len(mjArr)
	if aLen >= 3 {
		for i := 0; i < aLen - 2; i++ {
			for i2 := i + 1; i2 < aLen; i2++ {
				for i3 := i2+1; i3 < aLen; i3++ {
					if mjArr[i] == mjArr[i2] && mjArr[i2] == mjArr[i3] {
						return []string{mjArr[i], mjArr[i2], mjArr[i3]}
					}
				}
			}

		}
	}
	return nil
}

// 特殊牌型判断, 7小对
func is7Pair(sortedMjArr []string) bool {
	if len(sortedMjArr) == 14 {
		for i := 0; i < len(sortedMjArr); i += 2 {
			if sortedMjArr[i] != sortedMjArr[i+1] {
				return false
			}
		}
		return true;
	}

	return false
}

// 一种匹配类型
type matchType struct {
	Pengs [][]string
	Shunzis [][]string
	Gangs [][]string
	Jiang []string
}

func (mt *matchType) IsSame(other *matchType) bool {
	if other == nil {
		return false
	}

	if len(mt.Gangs) == len(other.Gangs) {
		if len(mt.Gangs) > 0 {
			for i, v := range mt.Gangs {
				v2 := other.Gangs[i]
				if v[0] != v2[0] {
					return false
				}
			}
		}
	} else {
		return false
	}

	if len(mt.Shunzis) == len(other.Shunzis) {
		if len(mt.Shunzis) > 0 {
			for i, v := range mt.Shunzis {
				v2 := other.Shunzis[i]
				if v[0] != v2[0] {
					return false
				}
			}
		}
	} else {
		return false
	}

	if len(mt.Pengs) == len(other.Pengs) {
		if len(mt.Pengs) > 0 {
			for i, v := range mt.Pengs {
				v2 := other.Pengs[i]
				if v[0] != v2[0] {
					return false
				}
			}
		}
	} else {
		return false
	}

	if len(mt.Jiang) == len(other.Jiang) {
		if len(mt.Jiang) > 0 &&  mt.Jiang[0] != other.Jiang[0] {
			return false
		}
	} else {
		return false
	}

	return true
}

func (mt *matchType) IsEmpty() bool {
	if len(mt.Gangs) == 0 && len(mt.Jiang) == 0 && len(mt.Shunzis) == 0 && len(mt.Pengs) == 0 {
		return true
	}

	return false
}

func (mt *matchType) HasJiang() bool {
	return len(mt.Jiang) != 0
}

// 每个种类的匹配类型集合
type matchTypeArr struct {
	arr []*matchType
	dontMatchShunzi bool
	mjArr []string
}

// 过滤相同组合
func (mta *matchTypeArr) filterSameArr() {
	aLen := len(mta.arr)
	if aLen >= 2 {
		for i := 0; i < aLen - 1; i++ {
			for i2 := i + 1; i2 < aLen; i2++ {
				v := mta.arr[i]
				v2 := mta.arr[i2]
				if v.IsSame(v2) {
					mta.arr[i] = nil
					break
				}
			}
		}
		newArr := make([]*matchType, 0)
		for _, v := range mta.arr {
			if v != nil {
				newArr = append(newArr, v)
			}
		}
		mta.arr = newArr
	}
}

func (mta *matchTypeArr) String() string {
	str := "";
	if len(mta.arr) > 0 {
		for i, v := range mta.arr {
			str = str + fmt.Sprintf("%v: %+v \n", i, v)
		}
		
	} else {
		str = "arr is 0 \n"
	}

	return str
}

func (mta *matchTypeArr) notMatch() bool {
	return len(mta.mjArr) != 0 && len(mta.arr) == 0
}

// 排除麻将牌
func excludeMjPai(src []string, exclude []string) []string {
	res := []string{}

	tmpExDump := make([]string, len(exclude))
	copy(tmpExDump, exclude)

	for _, v := range src {
		isExclude := false
		for i, ev := range tmpExDump {
			if ev != "" && ev == v {
				isExclude = true
				tmpExDump[i] = ""
				break
			}
		}
		if !isExclude {
			res = append(res, v)
		}
	}

	return res
}

// 找出一种类型的所有可能组合
func findMatchType(sortedMjArr []string, mt matchType, mtArr *matchTypeArr) {
	if len(sortedMjArr) == 0  {
		if !mt.IsEmpty() {
			mtArr.arr = append(mtArr.arr, &mt)
		}
		return
	}
	// 杠
	gang := selectOneGang(sortedMjArr)
	if len(gang) == 4 {
		findMatchType(excludeMjPai(sortedMjArr, gang), matchType{
			Pengs: mt.Pengs,
			Shunzis: mt.Shunzis,
			Gangs: append(mt.Gangs, gang),
			Jiang: mt.Jiang,
		}, mtArr)
	}
	// 碰
	peng := selectOnePeng(sortedMjArr)
	if len(peng) == 3 {
		findMatchType(excludeMjPai(sortedMjArr, peng), matchType{
			Pengs: append(mt.Pengs, peng),
			Shunzis: mt.Shunzis,
			Gangs: mt.Gangs,
			Jiang: mt.Jiang,
		}, mtArr)
	}
	// 顺子
	if !mtArr.dontMatchShunzi {
		shunzi := selectOneShunZi(sortedMjArr)
		if len(shunzi) == 3 {
			findMatchType(excludeMjPai(sortedMjArr, shunzi), matchType{
				Pengs: mt.Pengs,
				Shunzis: append(mt.Shunzis, shunzi),
				Gangs: mt.Gangs,
				Jiang: mt.Jiang,
			}, mtArr)
		}
	}
	// 将 只能存在一个
	if len(sortedMjArr) >= 2 && sortedMjArr[0] == sortedMjArr[1] && len(mt.Jiang) == 0 {
		findMatchType(excludeMjPai(sortedMjArr, []string{sortedMjArr[0], sortedMjArr[1]}), matchType{
			Pengs: mt.Pengs,
			Shunzis: mt.Shunzis,
			Gangs: mt.Gangs,
			Jiang: []string{sortedMjArr[0], sortedMjArr[1]},
		}, mtArr)
	}
}

func findMatchTypes(sortedMjArr []string, dontMatchShunzi bool) *matchTypeArr {
	mtArr := &matchTypeArr{
		arr: make([]*matchType, 0),
		dontMatchShunzi: dontMatchShunzi,
		mjArr: sortedMjArr,
	}
	findMatchType(sortedMjArr, matchType{}, mtArr)
	mtArr.filterSameArr()
	return mtArr
}

type allMatchTypes struct {
	arr [][]*matchType // 所有类型集合, 每条 []*matchType 就是一种类型的所有组合
	results [][]*matchType // 查找结果, 每条 []*matchType 就是一种胡牌组合
	onlyNeedOneResult bool // true 为只找出一个结果
	specialCase MjSpecialCase // 是否特殊牌型
	includeJiang []string // 非空，则表示必需要在这里面的值做*将*才可以胡牌
}

// 如果是特殊牌型，返回 true
func (amt *allMatchTypes) checkSpecialCase(mjArr []string) bool {
	sortedArr := make([]string, len(mjArr))
	copy(sortedArr, mjArr)
	SortMjPai(sortedArr)

	amt.specialCase = SpecialCase_None
	// 七小对
	if is7Pair(sortedArr) {
		amt.specialCase = SpecialCase_SevenPair
		return true
	}

	return false
}

func (amt *allMatchTypes) IsSpecialCase() bool {
	return amt.specialCase != SpecialCase_None
}

func (amt *allMatchTypes) String() string {
	str := ""

	if len(amt.results) > 0 {
		str = "allMatchTypes result: "
		for i, v := range amt.results {
			str += fmt.Sprintf("[%v] = ", i)
			for _, mt := range v {
				str += fmt.Sprintf("%+v, ", mt)
			}
			str += "\n"
		}
	}

	return str
}

func (amt *allMatchTypes) StringResult() string {
	str := ""

	if len(amt.results) > 0 {
		for i, v := range amt.results {
			Pengs := [][]string{}
			Shunzis := [][]string{}
			Gangs := [][]string{}
			Jiang := []string{}

			for _, mt := range v {
				if len(mt.Gangs) > 0 {
					Gangs = append(Gangs, mt.Gangs...)
				}
				if len(mt.Pengs) > 0 {
					Pengs = append(Pengs, mt.Pengs...)
				}
				if len(mt.Shunzis) > 0 {
					Shunzis = append(Shunzis, mt.Shunzis...)
				}
				if len(mt.Jiang) > 0 {
					Jiang = append(Jiang, mt.Jiang...)
				}
			}
			str += fmt.Sprintf("(%v): Gangs=%v, Pengs=%v, Shunzis=%v, Jiang=%v \n", i, Gangs, Pengs, Shunzis, Jiang)
		}
	}

	return str
}

func (amt *allMatchTypes) add(mt []*matchType) {
	if amt.arr == nil {
		amt.arr = make([][]*matchType, 0)
	}
	amt.arr = append(amt.arr, mt)
}

// 找出所有组合
func (amt *allMatchTypes) findAllResult() {
	amt.onlyNeedOneResult = false
	amt.results = make([][]*matchType, 0)
	amt.findResult(nil, 0)
}

// 找出一个
func (amt *allMatchTypes) findOneResult() {
	amt.onlyNeedOneResult = true
	amt.results = make([][]*matchType, 0)
	amt.findResult(nil, 0)
}

func (amt *allMatchTypes) isJiangRight(jiang string) bool {
	for _, v := range amt.includeJiang {
		if v == jiang {
			return true
		}
	}

	return false
}

func (amt *allMatchTypes) SetIncludeJiang(jiangs []string) {
	if amt.includeJiang == nil {
		amt.includeJiang = make([]string, 0)
	}
	amt.includeJiang = append(amt.includeJiang, jiangs...)
}

func (amt *allMatchTypes) findResult(arr []*matchType, arrIdx int) {
	if amt.onlyNeedOneResult && len(amt.results) == 1 {
		// 递归退出
		return
	}
	if arrIdx >= len(amt.arr) {
		if arr != nil {
			jiangCount := 0
			jiang := ""
			for _, v := range arr {
				if v.HasJiang() {
					jiangCount += 1
					jiang = v.Jiang[0]
				}
			}
			if jiangCount == 1 {
				// 找到一个符合胡牌的组合
				// fmt.Println("findResult find one: ", arrIdx, arr)
				if len(amt.includeJiang) > 0 {
					if amt.isJiangRight(jiang) {
						amt.results = append(amt.results, arr)
					}
				} else {
					amt.results = append(amt.results, arr)
				}
			}
		}
		// 递归退出
		return
	}

	for _, v := range amt.arr[arrIdx] {
		nextArr := make([]*matchType, 0)
		if arr != nil {
			nextArr = append(nextArr, arr...)
		}
		nextArr = append(nextArr, v)
		amt.findResult(nextArr, arrIdx+1)
	}
}

func (amt *allMatchTypes) hasResult() bool{
	return len(amt.results) > 0
}

// 找出所有可以胡牌的组合
func FindAllWins(mjArr []string, includeJiang []string) *allMatchTypes {
	if len(mjArr) % 3 == 2 {
		amt := new(allMatchTypes)
		amt.SetIncludeJiang(includeJiang)
		// 特殊牌型
		if amt.checkSpecialCase(mjArr) {
			return amt
		}

		SortMjPai(mjArr)

		// 先把各个类型的组合算出来
		// 万
		tmpArr := make([]string, 0, len(mjArr))
		for _, v := range mjArr {
			if string(v[0]) == "W" {
				tmpArr = append(tmpArr, v)
			}
		}
		wMatchTypes := findMatchTypes(tmpArr, false)
		if wMatchTypes.notMatch() {
			return nil
		}
		// 条
		tmpArr = tmpArr[0:0]
		for _, v := range mjArr {
			if string(v[0]) == "T" {
				tmpArr = append(tmpArr, v)
			}
		}
		tMatchTypes := findMatchTypes(tmpArr, false)
		if tMatchTypes.notMatch() {
			return nil
		}
		// 大饼
		tmpArr = tmpArr[0:0]
		for _, v := range mjArr {
			if string(v[0]) == "B" {
				tmpArr = append(tmpArr, v)
			}
		}
		bMatchTypes := findMatchTypes(tmpArr, false)
		if bMatchTypes.notMatch() {
			return nil
		}
		// 字
		tmpArr = tmpArr[0:0]
		for _, v := range mjArr {
			if string(v[0]) == "Z" {
				tmpArr = append(tmpArr, v)
			}
		}
		zMatchTypes := findMatchTypes(tmpArr, true)
		if zMatchTypes.notMatch() {
			return nil
		}

		// 从所有类型的组合中选出符合胡牌的组合
		if len(wMatchTypes.arr) > 0 {
			// fmt.Println(wMatchTypes)
			amt.add(wMatchTypes.arr)
		}
		if len(tMatchTypes.arr) > 0 {
			// fmt.Println(tMatchTypes)
			amt.add(tMatchTypes.arr)
		}
		if len(bMatchTypes.arr) > 0 {
			// fmt.Println(bMatchTypes)
			amt.add(bMatchTypes.arr)
		}
		if len(zMatchTypes.arr) > 0 {
			// fmt.Println(zMatchTypes)
			amt.add(zMatchTypes.arr)
		}
		amt.findAllResult()
		if amt.hasResult() {
			// fmt.Println(amt)
			return amt
		}
	}

	return nil
}

// 检查是否可以胡牌
func CanWin(mjArr []string, includeJiang []string) bool {
	if len(mjArr) % 3 == 2 {
		amt := new(allMatchTypes)
		amt.SetIncludeJiang(includeJiang)
		// 特殊牌型
		// 七小对
		if amt.checkSpecialCase(mjArr) {
			return true
		}

		SortMjPai(mjArr)

		// 先把各个类型的组合算出来
		// 万
		tmpArr := make([]string, 0, len(mjArr)) 
		for _, v := range mjArr {
			if string(v[0]) == "W" {
				tmpArr = append(tmpArr, v)
			}
		}
		// fmt.Println("CanWin tmpArr: ", tmpArr)
		wMatchTypes := findMatchTypes(tmpArr, false)
		if wMatchTypes.notMatch() {
			return false
		}
		// 条
		tmpArr = make([]string, 0, len(mjArr)) 
		for _, v := range mjArr {
			if string(v[0]) == "T" {
				tmpArr = append(tmpArr, v)
			}
		}
		// fmt.Println("CanWin tmpArr: ", tmpArr)
		tMatchTypes := findMatchTypes(tmpArr, false)
		if tMatchTypes.notMatch() {
			return false
		}
		// 大饼
		tmpArr = make([]string, 0, len(mjArr))
		for _, v := range mjArr {
			if string(v[0]) == "B" {
				tmpArr = append(tmpArr, v)
			}
		}
		// fmt.Println("CanWin tmpArr: ", tmpArr)
		bMatchTypes := findMatchTypes(tmpArr, false)
		if bMatchTypes.notMatch() {
			return false
		}
		// 字
		tmpArr = make([]string, 0, len(mjArr)) 
		for _, v := range mjArr {
			if string(v[0]) == "Z" {
				tmpArr = append(tmpArr, v)
			}
		}
		// fmt.Println("CanWin tmpArr: ", tmpArr)
		zMatchTypes := findMatchTypes(tmpArr, true)
		if zMatchTypes.notMatch() {
			return false
		}

		// 从所有类型的组合中选出符合胡牌的组合
		if len(wMatchTypes.arr) > 0 {
			// fmt.Println(wMatchTypes)
			amt.add(wMatchTypes.arr)
		}
		if len(tMatchTypes.arr) > 0 {
			// fmt.Println(tMatchTypes)
			amt.add(tMatchTypes.arr)
		}
		if len(bMatchTypes.arr) > 0 {
			// fmt.Println(bMatchTypes)
			amt.add(bMatchTypes.arr)
		}
		if len(zMatchTypes.arr) > 0 {
			// fmt.Println(zMatchTypes)
			amt.add(zMatchTypes.arr)
		}
		amt.findOneResult() // 找出一个即可
		if amt.hasResult() {
			// fmt.Println(amt)
			return true
		}
	}

	return false
}

// 手牌胡牌提示，输出可以胡牌的牌, needOne 表示只要一个结果
func HandTips(handPai []string, needOne bool, includeJiang []string) []string {
	if len(handPai) % 3 == 1 {
		winPaiArr := make([]string, 0)
		// 万
		hasType := false
		for _, v := range handPai {
			if string(v[0]) == "W" {
				hasType = true
				break
			}
		}
		if hasType {
			for _, v := range MJ_W {
				tmpArr := make([]string, len(handPai) + 1)
				copy(tmpArr, handPai)
				tmpArr[len(handPai)] = v
				if CanWin(tmpArr, includeJiang) {
					winPaiArr = append(winPaiArr, v)
					if needOne {
						return winPaiArr
					}
				}
			}
		}
		// 条
		hasType = false
		for _, v := range handPai {
			if string(v[0]) == "T" {
				hasType = true
				break
			}
		}
		if hasType {
			for _, v := range MJ_T {
				tmpArr := make([]string, len(handPai) + 1)
				copy(tmpArr, handPai)
				tmpArr[len(handPai)] = v
				if CanWin(tmpArr, includeJiang) {
					winPaiArr = append(winPaiArr, v)
					if needOne {
						return winPaiArr
					}
				}
			}
		}
		// 大饼(筒)
		hasType = false
		for _, v := range handPai {
			if string(v[0]) == "B" {
				hasType = true
				break
			}
		}
		if hasType {
			for _, v := range MJ_B {
				tmpArr := make([]string, len(handPai) + 1)
				copy(tmpArr, handPai)
				tmpArr[len(handPai)] = v
				if CanWin(tmpArr, includeJiang) {
					winPaiArr = append(winPaiArr, v)
					if needOne {
						return winPaiArr
					}
				}
			}
		}
		// 字
		hasType = false
		for _, v := range handPai {
			if string(v[0]) == "Z" {
				hasType = true
				break
			}
		}
		if hasType {
			for _, v := range MJ_Z {
				tmpArr := make([]string, len(handPai) + 1)
				copy(tmpArr, handPai)
				tmpArr[len(handPai)] = v
				if CanWin(tmpArr, includeJiang) {
					winPaiArr = append(winPaiArr, v)
					if needOne {
						return winPaiArr
					}
				}
			}
		}

		return winPaiArr
	}

	return nil
}

// 把重复的牌过滤掉，返回的结果里每个牌都不相同
func filterSame(mjArr []string) []string {
	filted := make([]string, 0)

	for _, v := range mjArr {
		isSame := false
		for _, v2 := range filted {
			if v == v2 {
				isSame = true
				break
			}
		}
		if !isSame {
			filted = append(filted, v)
		}
	}

	return filted
}

// 打出某个手牌后可以听牌提示，输出打出牌后可听牌的牌
func PlayTips(handPai []string, includeJiang []string) []string {
	if len(handPai) % 3 == 2 {
		playPaiArr := make([]string, 0)

		filtedHand := filterSame(handPai)
		for _, v := range filtedHand {
			tmpArr := make([]string, 0, len(handPai) - 1)
			findIdx := -1
			for i2, v2 := range handPai {
				if v == v2 {
					findIdx = i2
					tmpArr = append(tmpArr, handPai[0:i2]...)
					tmpArr = append(tmpArr, handPai[i2+1:]...)
					// fmt.Println("PlayTips tmpArr: ", tmpArr)
					break
				}
			}

			if findIdx != -1 && len(HandTips(tmpArr, true, includeJiang)) > 0 {
				playPaiArr = append(playPaiArr, v)
			}
		}
		
		return playPaiArr
	}

	return nil
}