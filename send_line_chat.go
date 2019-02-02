package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/linebot"
	"time"
)

func main() {
	//Userid := "U229e9a449179d8c8a9efc3e2d00aa15a"
	//Roomid := "Rdd3a9e3bc5e73ed5ec283cd84fc7fd8d"
	Roomid := "C493aef71d982c6e4d9b6996c548bc131"
	var group string
	for {
		t := time.Now()
		bot, _ := linebot.New("7b1317123f743c47984eefd74a66180f", "84RSL9DoYB7CYToRF/7n38pidk4YqaylvtyULojYQYLSRpyOf7dHIYPGK/GwtrCzXJo0FoaDznyJ+jYiygmWc0azPyO+i1QBlqj2kNobdAoVGhGWQzFmI1pZQ8/F9IxEwLw5Heoofae684hdMDobzAdB04t89/1O/w1cDnyilFU=")
		//var messages []linebot.Message
		//leftBtn := linebot.NewMessageAction("left", "left clicked")
		//rightBtn := linebot.NewMessageAction("right", "right clicked")
		switch t.Hour(){
		case 10:
			group = "桃園兵器保修"
		case 15:
			group = "桃園食勤, 成功嶺一營"
		case 16:
			group = "桃園彈藥, 兵器保修"
		case 20:
			group = "所有弟兄們"
		default:
			group = ""
		}
		if (group != ""){
		text := fmt.Sprintf("現在時間 %d-%02d-%02d %02d:%02d\n %s \n 等等記得回報一下訊息\n 祝大家新年快樂", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), group)
		message := linebot.NewTextMessage(text)
		bot.PushMessage(Roomid, message).Do()
		}
		time.Sleep(time.Duration(30 * time.Minute))
	}

}
