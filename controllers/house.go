package controllers

import (
	"fmt"
	"time"

	"github.com/hapiman/ke/utils"
	"github.com/mikemintang/go-curl"
	"github.com/tidwall/gjson"
)

func FetchDailyStats() (string, string) {
	avgPrice, yestTxNum := "", ""
	headers := map[string]string{
		"Authorization": utils.Authorization,
		"Content-Type":  "application/json",
	}
	url := "https://app.api.ke.com/config/home/content?city_id=510100&request_ts=1539823818&type=iPhone"
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).SetHeaders(headers).Get()
	if err != nil {
		fmt.Println("FetchDailyStats error =>", err)
		return avgPrice, yestTxNum
	}
	if resp.IsOk() {
		listRaw := gjson.Get(resp.Body, "data.market.list")
		yestCount := listRaw.Get("1.count").Value()
		avgCount := listRaw.Get("0.count").Value()
		yestTxNum = yestCount.(string)
		avgPrice = avgCount.(string)
	} else {
		fmt.Println("resp is not ok.", resp.Raw)
	}
	return avgPrice, yestTxNum
}

func FetchDailyNew(searchSeds int64) []map[string]string {
	if searchSeds == 0 {
		searchSeds = utils.GetCurrentSeds()
	}
	var houseList []map[string]string
	offSet := 0
	for {
		url := fmt.Sprintf("https://app.api.ke.com/house/ershoufang/searchv4?cityId=510100&condition=tt2&hasRecommend=0&limitCount=20&limitOffset=%d&order=co32&request_ts=1541411669", offSet)
		fmt.Println("url => %s", url)
		headers := map[string]string{
			"Authorization":        utils.Authorization2,
			"Lianjia-Access-Token": "2.0012633e536b1a0ff303ce17625175c2b4",
			"Lianjia-Device-Id":    "89B621AE-A099-46C9-A172-4BC69F74445F",
			"Lianjia-Im-Version":   "1",
			"Content-Type":         "application/json",
		}
		req := curl.NewRequest()
		resp, err := req.SetUrl(url).SetHeaders(headers).Get()

		if err != nil {
			fmt.Println("FetchDailyNew error =>", err)
		}
		if resp.IsOk() {
			errNo := gjson.Get(resp.Body, "errno").Int()
			fmt.Printf("errNo's Value: %d \n", errNo)
			if errNo > 0 {
				fmt.Printf("errNo: %d \n", errNo)
				break
			}
			listRaw := gjson.Get(resp.Body, "data.list")
			listRaw.ForEach(func(key, value gjson.Result) bool {
				houseCode := value.Get("houseCode").Value()
				houseTitle := value.Get("title").Value()
				if houseCode != nil {
					houseCodeStr := houseCode.(string)
					houseTitleStr := houseTitle.(string)
					onlineDateStr := value.Get("infoList.1.value").Value().(string)
					fmt.Printf("houseCode=%s, houseTitle=%s, onlineDateStr=%s\n", houseCode, houseTitle, onlineDateStr)
					houseList = append(houseList, map[string]string{
						"houseCode":      houseCodeStr,
						"houseTitle":     houseTitleStr,
						"houseOnLineStr": onlineDateStr,
					})
				}
				return true
			})
			// 判断时间
			if len(houseList) > 0 {
				lastOne := houseList[len(houseList)-1]
				realSeds, err := utils.TimeToTimestamp("2006.01.02", lastOne["houseOnLineStr"])
				if err != nil {
					fmt.Println("TimeToTimestamp error =>", err.Error())
					break
				}
				fmt.Sprintf("searchSeds => %d, realSeds => %d", searchSeds, realSeds)
				if searchSeds > realSeds {
					time.Sleep(time.Second * 5)
					offSet += 20
					fmt.Printf("offSet => %d \n", offSet)
					continue
				}
				break
			}
		} else {
			fmt.Println("resp is not ok.", resp.Raw)
			break
		}
	}
	return houseList
}

func FetchDailyTxList(day string) []map[string]string {
	const url string = "https://app.api.ke.com/house/chengjiao/searchv2?channel=sold&city_id=510100&limit_count=20&limit_offset=0&request_ts=1541052066"
	headers := map[string]string{
		"Authorization":        utils.Authorization3,
		"Lianjia-Access-Token": "2.0012633e536b1a0ff303ce17625175c2b4",
		"Lianjia-Device-Id":    "89B621AE-A099-46C9-A172-4BC69F74445F",
		"Lianjia-Im-Version":   "1",
		"Content-Type":         "application/json",
	}
	fmt.Println("utils.Authorization3 =>", utils.Authorization3)
	req := curl.NewRequest()
	resp, err := req.SetUrl(url).SetHeaders(headers).Get()

	if err != nil {
		fmt.Println("FetchDailyNew error =>", err)
	}
	var houseList []map[string]string
	if resp.IsOk() {
		listRaw := gjson.Get(resp.Body, "data.list")
		listRaw.ForEach(func(key, value gjson.Result) bool {
			houseCode := value.Get("house_code").Value()
			houseTitle := value.Get("title").Value()
			housePrice := value.Get("price_str").Value()
			if houseCode != nil {
				houseCodeStr := houseCode.(string)
				houseTitleStr := houseTitle.(string)
				housePriceStr := housePrice.(string)
				fmt.Printf("houseCode=%s, houseTitle=%s, housePriceStr=%s\n", houseCode, houseTitle, housePriceStr)
				houseList = append(houseList, map[string]string{
					"houseCode":  houseCodeStr,
					"houseTitle": houseTitleStr,
					"housePrice": housePriceStr,
				})
			}
			return true
		})
		fmt.Println(resp.Body)
	} else {
		fmt.Println("resp is not ok.", resp.Raw)
	}
	return houseList
}
