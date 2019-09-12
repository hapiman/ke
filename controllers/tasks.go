package controllers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hapiman/ke/models"
	"github.com/hapiman/ke/utils"
	"github.com/robfig/config"
	"github.com/robfig/cron"
	"github.com/tidwall/gjson"
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

func SyncHouseTask() {
	c := cron.New()
	spec := "10 10 10,16,18 * * *"
	c.AddFunc(spec, func() {
		fmt.Println("sync task start again")
		syncXiaoQuOverview()
		syncOverview()
		SyncTransRecords()
		fmt.Println("sync task end")
	})
	c.Start()
}

func findEx(quCode, hName string) map[string]interface{} {
	dUrl := fmt.Sprintf("https://bj.ke.com/api/listtop?type=resblock&resblock_id=%s&community_id=0&district_id=&bizcircle_id=&subway_station_id=&word=%s&source=ershou_xiaoqu", quCode, hName)
	fmt.Println("dUrl =>", dUrl)

	content := utils.HTTPDo("GET", dUrl, []byte{}, map[string]string{})
	errno := gjson.Get(content, "errno").Num
	bks := map[string]interface{}{}
	if errno == 0 {
		soldNumInNinety64 := strconv.FormatInt(gjson.Get(content, "data.info.90saleCount").Int(), 10)
		bks["soldNumInNinety"] = utils.ConvertStr2Num(soldNumInNinety64)
		visitNumInThirty64 := strconv.FormatInt(gjson.Get(content, "data.info.day30See").Int(), 10)
		bks["visitNumInThirty"] = utils.ConvertStr2Num(visitNumInThirty64)
	}
	return bks
}

/*
同步交易记录
*/
func SyncTransRecords() {
	pageNo := 0
	for pageNo < 40 {
		tUrl := fmt.Sprintf("https://bj.ke.com/chengjiao/pg%d/", pageNo)
		fmt.Println("syncTransRecords pageNo: ", pageNo, ", url: ", tUrl)
		res, err := http.Get(tUrl)
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

		doc.Find(".leftContent ul.listContent li").Each(func(i int, s *goquery.Selection) {
			recoBody := &models.TabTransRecords{}
			recoBody.Link, _ = s.Find(".info .title a").Attr("href")
			regx := regexp.MustCompile(`https://bj.ke.com/chengjiao/(\d{1,}).html`)
			params := regx.FindStringSubmatch(recoBody.Link)
			if len(params) >= 2 {
				recoBody.HouseCode = params[1]
			}
			recoBody.TxId, _ = s.Find(".info .title a").Attr("data-maidian")
			h := strings.TrimSpace(s.Find(".info .title a").Text())
			recoBody.HouseName = strings.Split(h, " ")[0]
			recoBody.Date = strings.TrimSpace(s.Find(".info .address .dealDate").Text())
			recoBody.TotalPrice = utils.ConvertStr2Num(strings.TrimSpace(s.Find(".info .address .totalPrice .number").Text()))
			recoBody.AvgPrice = utils.ConvertStr2Num(strings.TrimSpace(s.Find(".info .flood .unitPrice .number").Text()))
			count := 0
			models.ConnKe().Where("date=? AND house_code=?", recoBody.Date, recoBody.HouseCode).Find(&models.TabTransRecords{}).Count(&count)
			if count > 0 {
				return
			}
			err := models.ConnKe().Create(recoBody).Error
			if err != nil {
				fmt.Println("syncTransRecords err: ", err.Error())
			}
		})
		pageNo++
	}
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
			quId, _ := s.Attr("data-id")
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
			fmt.Print("2222xxxx \n")
			models.ConnKe().Where("date=? AND qu_code=?", curYYYYMMDD, quId).Find(&models.TabXiaoQuOverview{}).Count(&count)
			if count < 1 {
				fmt.Printf("3333xxxx %s %s\n", quId, quName)
				exinfo := findEx(quId, quName)
				fmt.Print("44444xxxx \n")
				house := &models.TabXiaoQuOverview{
					CityCode:         "bj",
					DistrictCode:     "chaoyang",
					AreaCode:         "wangjing",
					QuCode:           quId,
					Date:             curYYYYMMDD,
					Name:             quName,
					OnsaleNumCurrent: onsaleNum,
					AvgPrice:         avgPrice,
					SoldNumInNinety:  exinfo["soldNumInNinety"].(int),
					SoldNumInThirty:  soldNumInThirty,
					VisitNumInThirty: exinfo["visitNumInThirty"].(int),
				}
				models.ConnKe().Create(house)
			}
		})
		fmt.Println("4444 =>")
		pageNo++
		fmt.Printf("current date: %s, pageNo: %d\n", curYYYYMMDD, pageNo)
		// time.Sleep(time.Millisecond * 300) // 暂停0.3s
	}
}

/*
同步城市概览数据
*/
func syncOverview() {
	overUrl := "https://bj.ke.com/fangjia/"
	res, err := http.Get(overUrl)
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
	dura, _ := time.ParseDuration("-24h")
	curYYYYMMDD := time.Now().Add(dura).Format("2006-01-02")
	eleBox := doc.Find(".g-main .m-tongji .box-l")
	eleTop := eleBox.Find(".box-l-t .qushi .qushi-2")
	avgPriceLastMonth := utils.ConvertStr2Num(strings.TrimSpace(eleTop.Find(".num").Text()))
	onsaleNumCurrent := 0
	soldNumInNinety := 0
	newHouseNum := 0
	newPeopleNum := 0
	visitNum := 0
	eleTop.Find(".txt").Each(func(i int, s *goquery.Selection) {
		elem := strings.TrimSpace(s.Text())
		if strings.Contains(elem, "在售") {
			// 在售房源97892套
			regx := regexp.MustCompile(`在售房源(\d{1,})套`)
			params := regx.FindStringSubmatch(elem)
			if len(params) >= 2 {
				onsaleNumCurrent = utils.ConvertStr2Num(params[1])
			}
		}
		if strings.Contains(elem, "成交") {
			// 最近90天内成交房源18468套
			regx := regexp.MustCompile(`最近90天内成交房源(\d{1,})套`)
			params := regx.FindStringSubmatch(elem)
			if len(params) >= 2 {
				soldNumInNinety = utils.ConvertStr2Num(params[1])
			}
		}
	})
	eleBox.Find(".item").Each(func(i int, s *goquery.Selection) {
		valueTxt := strings.TrimSpace(s.Find(".num").Text())
		valueNum := utils.ConvertStr2Num(valueTxt)
		if strings.Contains(strings.TrimSpace(s.Find(".text").Text()), "新增房") {
			newHouseNum = valueNum
		}
		if strings.Contains(strings.TrimSpace(s.Find(".text").Text()), "新增客") {
			newPeopleNum = valueNum
		}
		if strings.Contains(strings.TrimSpace(s.Find(".text").Text()), "带看量") {
			visitNum = valueNum
		}
	})

	fmt.Println("newHouseNum: ", newHouseNum, "; newPeopleNum: ", newPeopleNum, ";visitNum: ", visitNum)
	count := 0
	models.ConnKe().Where("date=?", curYYYYMMDD).Find(&models.TabOverview{}).Count(&count)
	if count > 0 {
		return
	}
	ov := &models.TabOverview{
		CityCode:          "bj",
		Date:              curYYYYMMDD,
		NewHouseNum:       newHouseNum,
		NewPeopleNum:      newPeopleNum,
		VisitNum:          visitNum,
		SoldNumInNinety:   soldNumInNinety,
		OnsaleNumCurrent:  onsaleNumCurrent,
		AvgPriceLastMonth: avgPriceLastMonth,
	}
	err = models.ConnKe().Create(ov).Error
	if err != nil {
		fmt.Println("err message: ", err.Error())
	}
}
