package main

import (
	"encoding/json"
	"errors"
	"flag"
	"time"
	"fmt"
	"net/http"
	"strings"
)

var fAddr = flag.String("addr", "", "开启一个 http 服务 eg: -addr=127.0.0.1:8080, 将开启 127.0.0.1 的 8080 端口")

var fMjPai = flag.String("pai", "", fmt.Sprintf(`指定一组有效的麻将牌, eg: -pai="T4, W3, W7, W3, W8, T4, T4, W5, W6, W7, T7, T8, T9, W6"
合法的麻将牌：
    万=%v
    饼=%v
    条=%v
    字=%v`, MJ_W, MJ_B,MJ_T,MJ_Z))

var fOp = flag.String("op", "", `对输入麻将牌的操作 eg：-op=FindAllWins
有效的操作为：
1. FindAllWins 查找所有胡牌的组合, 要求 len % 3 == 2
2. CanWin 是否可以胡牌, 要求 len % 3 == 2
3. HandTips 手牌胡牌提示, 要求 len % 3 == 1
4. PlayTips 打出某个手牌后可以听牌提示, 要求 len % 3 == 2
`)

var fJiang = flag.String("jiang", "", "将 eg: 只有 3,5,8 才能作将 -jiang=\"3,5,8\"")
var parsedJiang []string


type resultResp struct{
	Result string `json:"result,omitempty"`
	FinishTime string `json:"time,omitempty"`
	Err string `json:"err,omitempty"`
}
type mjReq struct{
	Pai string `json:"pai"`
	Jiang string `json:"jiang,omitempty"`
	Op string `json:"op"`
}

func init(){
	http.HandleFunc("/mj", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			return
		}

		w.Header().Set("Content-Type","application/json")

		var resp resultResp
		var req mjReq
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			resp.Err = err.Error()
			json.NewEncoder(w).Encode(resp)
			return
		}

		t1 := time.Now()

		mjArr, err := parseMJTile(req.Pai)
		if err != nil {
			resp.Err = err.Error()
			json.NewEncoder(w).Encode(resp)
			return
		}
		// fmt.Printf("mj pai %+v \n", mjArr)

		res, err := parseOp(mjArr, req.Op, req.Jiang)
		if err != nil {
			resp.Err = err.Error()
			json.NewEncoder(w).Encode(resp)
			return
		}
		
		t2 := time.Now()

		resp.Result = fmt.Sprintf(" pai: %+v \n op: %v \n jiang: %v\n result: %v\n", mjArr, req.Op, parsedJiang, res)
		// resp.FinishTime = fmt.Sprintf("%.4f's", ft.Seconds())
		resp.FinishTime = fmt.Sprintf("%v", t2.Sub(t1).Seconds())

		json.NewEncoder(w).Encode(resp)
	})
}

func main() {

	flag.Parse()

	if (*fAddr) != "" {
		fmt.Println("listen addr",*fAddr)
		err := http.ListenAndServe(*fAddr, nil)
		if err != nil {
			fmt.Println("listen error: ", err)
			return
		}
	}
	
	if (*fMjPai) != "" {
		mjArr, err := parseMJTile("")
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			return
		}
		// fmt.Printf("mj pai %+v \n", mjArr)

		res, err := parseOp(mjArr, "", "")
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			return
		}
		fmt.Printf(" pai: %+v \n op: %v \n jiang: %v\n result: %v\n", mjArr, *fOp, parsedJiang, res)
		return
	}

	fmt.Println("输入示例:\n  查找手牌所有能胡牌的组合 mahjong.exe -pai=\"T4, B3, W7, B3, W8, T4, T4, W5, W6, W7, T7, T8, T9, W6\" -op=FindAllWins")
}

func parseMJTile(str string) ([]string, error){
	if str == "" {
		str = *fMjPai
	}
	mjPaiArr := strings.Split(str, ",")
	for i,v := range mjPaiArr {
		v = strings.TrimSpace(v)
		if ValidMjTileStr(v) {
			mjPaiArr[i] = v
		} else {
			return nil, errors.New("输入了错误的麻将牌: " + v)
		}
	}

	return mjPaiArr, nil
}

func parseOp(mjArr []string, opStr string, jiangStr string) (string, error){
	if opStr == "" {
		opStr = *fOp
	}
	if jiangStr == "" {
		jiangStr = *fJiang
	}

	if opStr != "" {
		res := ""

		parsedJiang = []string{}
		if jiangStr != "" {
			jStrArr := strings.Split(jiangStr, ",")
			for _, v := range jStrArr {
				v = strings.TrimSpace(v)
				if len(v) == 1 {
					for _, v2 := range MJ_TYPES {
						s := string(v2) + v
						if ValidMjTileStr(s) {
							parsedJiang = append(parsedJiang, s)
						}
					}
				} else if len(v) == 2 {
					if ValidMjTileStr(v) {
						parsedJiang = append(parsedJiang, v)
					}
				}
			}
		}

		op := strings.TrimSpace(opStr)
		switch op {
			case "FindAllWins":
				allCase := FindAllWins(mjArr, parsedJiang)
				if allCase != nil && allCase.hasResult() {
					res = fmt.Sprintf("%v", allCase.StringResult())
				} else {
					res = "no cases"
				}
			case "CanWin":
				ok := CanWin(mjArr, parsedJiang)
				if ok {
					res = "ok"
				} else {
					res = "not ok"
				}
			case "HandTips":
				allTile := HandTips(mjArr, false, parsedJiang)
				if len(allTile) > 0 {
					res = fmt.Sprintf("%v", allTile)
				} else {
					res = "no tips"
				}
			case "PlayTips":
				allTile := PlayTips(mjArr, parsedJiang)
				if len(allTile) > 0 {
					res = fmt.Sprintf("%v", allTile)
				} else {
					res = "no tips"
				}
			default:
				return "", errors.New("解析操作输入错误")
		}

		return res, nil
	}
	return "", errors.New("请输入必要的操作")
}