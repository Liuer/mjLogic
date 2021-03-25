# 用 go 实现的简单麻将胡牌算法

## 构建
 
```
cd <project-dir>
go build -o bin
```
使用了 go 1.16 版本， 在 windows 操作系统下会在项目的 bin 目录下生成 mahjong.exe 程序, 或者直接使用已经在 bin 目录构建好的程序

## 命令行
```
./mahjong.exe -pai="T4, B3, W7, B3, W8, T4, T4, W5, W6, W7, T7, T8, T9, W6" -op=FindAllWins -jiang="3,5,8"
```
输出：
```
 pai: [B3 B3 T4 T4 T4 T7 T8 T9 W5 W6 W6 W7 W7 W8] 
 op: FindAllWins
 jiang: [W3 T3 B3 W5 T5 B5 W8 T8 B8]
 result: (0): Gangs=[], Pengs=[[T4 T4 T4]], Shunzis=[[W5 W6 W7] [W6 W7 W8] [T7 T8 T9]], Jiang=[B3 B3]

```
```
合法的麻将牌：
万 = W1 W2 W3 W4 W5 W6 W7 W8 W9
饼 = B1 B2 B3 B4 B5 B6 B7 B8 B9
条 = T1 T2 T3 T4 T5 T6 T7 T8 T9
字 = ZD ZN ZX ZB ZZ ZF ZM
```

## 作为 http 服务
```
./mahjong.exe -addr :8080
```
程序会监听 8080 端口 
```
// use curl 
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{
  "pai": "T4, B3, W7, B3, W8, T4, T4, W5, W6, W7, T7, T8, T9, W6",
  "op": "FindAllWins",
  "jiang": "3,5,8"
}' \
 'http://localhost:8080/mj'
```
输出：
```
{"result":" pai: [B3 B3 T4 T4 T4 T7 T8 T9 W5 W6 W6 W7 W7 W8] \n op: FindAllWins \n jiang: [W3 T3 B3 W5 T5 B5 W8 T8 B8]\n result: (0): Gangs=[], Pengs=[[T4 T4 T4]], Shunzis=[[W5 W6 W7] [W6 W7 W8] [T7 T8 T9]], Jiang=[B3 B3] \n\n","time":"0"}
```