package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hapiman/ke/models"
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
	// FetchDailyNew(0)
	go func() {
		currentStr := utils.TimestampToTime(utils.GetCurrentSeds(), utils.TimeYyyymmddhhmmss)
		fmt.Printf("AutoSync Started At %s.", currentStr)
		c := cron.New()
		spec := "10 0 8-14 * * *"
		// spec := "10 31 16 * * *"
		count := 0
		cfgPath := utils.CacuCurrentConfigFile()
		cfg, _ := config.ReadDefault(cfgPath)
		from, _ := cfg.String("email", "from")
		to, _ := cfg.String("email", "to")
		fmt.Printf("from:%s, to:%s", from, to)
		c.AddFunc(spec, func() {
			currentStr := utils.TimestampToTime(utils.GetCurrentSeds(), utils.TimeYyyymmddhhmmss)
			fmt.Printf("AutoSync started %d times \n at %s", count, currentStr)
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

func SyncXiaoQuTask() {
	c := cron.New()
	spec := "10 10 10,16,18 * * *"
	c.AddFunc(spec, func() {
		fmt.Println("sync xiaoqu task start again")
		syncXiaoQuOverview()
		fmt.Println("sync xiaoqu task end")
	})
	c.Start()
}

/*
同步小区概览数据
*/
func syncXiaoQuOverview() {
	pageNo := 0
	loop := true
	for pageNo < 200 && loop {
		quUrl := fmt.Sprintf("https://bj.ke.com/xiaoqu/chaoyang/pg%d/", pageNo)
		res, err := http.Get(quUrl)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()
		if res.StatusCode != 200 {
			log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		}

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			fmt.Println("err message: ", err.Error())
		}

		le := doc.Find(".leftContent ul.listContent li").Length()
		if le == 0 {
			loop = false
		}
		dura, _ := time.ParseDuration("-24h")
		curYYYYMMDD := time.Now().Add(dura).Format("2006-01-02")
		doc.Find(".leftContent ul.listContent li").Each(func(i int, s *goquery.Selection) {
			quId, _ := (s.Attr("data-id"))
			quName := strings.TrimSpace(s.Find(".info .title").Text())
			soldNumInThirty := 0
			s.Find(".info .houseInfo a").Each(func(ii int, ss *goquery.Selection) {
				elem := strings.TrimSpace(ss.Text())
				if strings.Contains(elem, "成交") && !strings.Contains(elem, "暂无成交") {
					regx := regexp.MustCompile(`30天成交(\d{1,})套`)
					params := regx.FindStringSubmatch(elem)
					if len(params) >= 2 {
						soldNumInThirty = utils.ConvertStr2Num(params[1])
					}
				}
			})
			avgPrice := utils.ConvertStr2Num(strings.TrimSpace(s.Find(".xiaoquListItemRight .xiaoquListItemPrice .totalPrice span").Text()))
			onsaleNum := utils.ConvertStr2Num(strings.TrimSpace(s.Find(".xiaoquListItemRight .xiaoquListItemSellCount .totalSellCount span").Text()))
			count := 0
			models.ConnKe().Where("date=? AND qu_code=?", curYYYYMMDD, quId).Find(&models.TabXiaoQuOverview{}).Count(&count)
			if count < 1 {
				house := &models.TabXiaoQuOverview{
					CityCode:         "bj",
					DistrictCode:     "chaoyang",
					AreaCode:         "wangjing",
					QuCode:           quId,
					Date:             curYYYYMMDD,
					Name:             quName,
					OnsaleNumCurrent: onsaleNum,
					AvgPrice:         avgPrice,
					SoldNumInNinety:  soldNumInThirty,
					VisitNumInThirty: 0,
				}
				models.ConnKe().Create(house)
			}
		})
		pageNo++
		fmt.Printf("current date: %s, pageNo: %d\n", curYYYYMMDD, pageNo)
		time.Sleep(time.Millisecond * 300) // 暂停0.3s
	}
}
