package main

import (
	"fmt"
	"testing"
	"time"
)

func TestMj(t *testing.T) {
	jiang := []string{}
	// fmt.Println(AllMJPai())
	// shufPai := Shuffle(AllMJPai())
	// fmt.Println(shufPai)
	// SortMjPai(shufPai)
	// fmt.Println(shufPai)

	// pai := []string{ "B3", "B2", "B4" }
	// pai2 := []string{ "B3", "B7", "B4" }
	// fmt.Println("isSameType pai: ", isSameType(pai), ", isSunZi: ", isShunZi(pai))
	// fmt.Println("isSameType pai2: ", isSameType(pai2), ", isSunZi: ", isShunZi(pai2))

	// pai3 := []string{ "B3", "B3", "B3" }
	// pai4 := []string{ "B3", "B3", "B3", "B3" }
	// fmt.Println("isGang pai: ", isGang(pai3), isGang(pai4))
	// fmt.Println("isPeng pai: ", isPeng(pai3), isPeng(pai4))

	// pai5 := []string{ "B3", "B3", "B3", "B3" }
	// pai6 := []string{ "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3"}
	// fmt.Println("is7Pair pai: ", is7Pair(pai5), is7Pair(pai6))

	// pai7 := []string{ "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3", "B3"}
	// pai7 := []string{"B3", "B1", "B2", "B3", "B4", "B5", "B5", "B5", "B5", "W3", "W3", "W3", "ZD", "ZD"}
	// pai7 := []string{"T4", "W3", "W7", "W3", "W8", "T4", "T4", "W5", "W6", "W7", "T7", "T8", "T9", "W6"}
	// pai7 := []string{"B4", "B1", "B2", "B5", "B6"}
	// SortMjPai(pai7)
	// fmt.Printf("getOneShunZi: %+v \n", getOneShunZi(pai7))
	// fmt.Printf("getOneShunZi: %+v \n", excludeMjPai(pai7, []string{"B6"}))
	// fmt.Printf("getOneShunZi: %+v \n", getOneShunZi(pai7))
	// fmt.Printf("getMatchTypes: %+v \n", getMatchTypes(pai7, false, jiang))
	// fmt.Printf("getMatchTypes: %+v \n", CanWin(pai7, jiang))

	// pai8 := []string{"T4", "W3", "W7", "W3", "W8", "T4", "T4", "W5", "W6", "W7", "T7", "T8", "T9"}
	// // pai8 := []string{"B4", "B2", "B5", "B6"}
	// // pai8 := []string{"B3", "B1", "B2", "B3", "B4", "B5", "B5", "B5", "B5", "W3", "W3", "W3", "ZD"}
	// t1 := time.Now()
	// res := HandTips(pai8, false, jiang)
	// t2 := time.Now()
	// fmt.Printf("HandTips: %+v, t=%d \n", res, t2.Sub(t1).Milliseconds())

	pai9 := []string{"T4", "W3", "W7", "W3", "W8", "T4", "T4", "W5", "W6", "W7", "T7", "T8", "T9", "W6"}
	// pai8 := []string{"B4", "B2", "B5", "B6"}
	// pai8 := []string{"B3", "B1", "B2", "B3", "B4", "B5", "B5", "B5", "B5", "W3", "W3", "W3", "ZD"}
	t1 := time.Now()
	res := PlayTips(pai9, jiang)
	t2 := time.Now()
	fmt.Printf("PlayTips: %+v, t=%f \n", res, t2.Sub(t1).Seconds())
}
