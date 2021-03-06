package main

import (
	"github.com/henrylee2cn/teleport"
	"log"
)

// 有标识符UID的demo
// func main() {
// 	tp := teleport.New().SetUID("S")
// 	tp.SetAPI(teleport.API{
// 		"报到": func(receive *teleport.NetData) *teleport.NetData {
// 			log.Printf("报到：%v", receive.Body)
// 			resp := &teleport.NetData{
// 				Operation: "报到回复",
// 				Body:      "我是服务器，我已经收到你的来信",
// 			}
// 			return resp
// 		},
// 	})

// 	tp.Server(":20125")
// 	select {}
// }

// 无标识符UID的demo
func main() {
	tp := teleport.New()
	tp.SetAPI(teleport.API{
		"报到": func(receive *teleport.NetData) *teleport.NetData {
			log.Printf("报到：%v", receive.Body)
			return teleport.ReturnData("我是服务器，我已经收到你的来信")
		},
	})

	tp.Server(":20125")
	select {}
}

// func main() {
// 	tp := teleport.New()
// 	tp.SetAPI(teleport.API{
// 		"报到": func(receive *teleport.NetData) *teleport.NetData {
// 			log.Printf("报到：%v", receive.Body)
// 			return teleport.ReturnData("我是服务器，我已经收到你的来信", "回复")
// 		},
// 	})

// 	tp.Server(":20125")
// 	select {}
// }
