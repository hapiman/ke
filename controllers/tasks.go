package controllers

import (
	"fmt"

	"github.com/hapiman/ke/utils"
	"github.com/robfig/config"
	"github.com/robfig/cron"
)

// 每日08:00到22:00点，每两个小时执行一次
// func AutoSync() {
// 	ticker := time.NewTicker(time.Hour * 2)
// 	go func() {
// 		for _ = range ticker.C {
// 			h := time.Now().Hour()
// 			if h >= 8 && h < 22 {
// 				FetchDailyNew(0)
// 			}
// 		}
// 	}()
// }

// AutoSync 每日08:00到22:00点，每两个小时执行一次
func AutoSync() {
	go func() {
		fmt.Println("AutoSync Started.")
		c := cron.New()
		spec := "10 0 8-22 * * *"
		// spec := "10 31 16 * * *"
		count := 0
		cfgPath := utils.CacuCurrentConfigFile()
		cfg, _ := config.ReadDefault(cfgPath)
		from, _ := cfg.String("email", "from")
		to, _ := cfg.String("email", "to")
		fmt.Printf("from:%s, to:%s", from, to)
		c.AddFunc(spec, func() {
			fmt.Printf("AutoSync started %d times \n", count)
			count++
			list := FetchDailyNew(0)
			e := &utils.EmailEntity{
				Subject:  "同步数据",
				From:     from,
				To:       []string{to},
				Nickname: "hapiman",
			}
			if len(list) > 0 {
				e.Content = "数据正常"
			} else {
				e.Content = "数据异常"
			}
			utils.SendEmail(e)
		})
		c.Start()
	}()
}
