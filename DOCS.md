### 每日统计数据
`curl -H 'Authorization: MjAxODAxMTFfaW9zOmFkMDBhNDdjMTYxMGUwZTJmZDkwOGMwYTBiZWE4MmQ0MDc0YzIzMzU=' 'https://app.api.ke.com/config/home/content?city_id=510100&request_ts=1539823818&type=iPhone'
`
分析`market`层级下数据
```go
"market": {
  "title": "贝壳指数",
  "list": [{
    "name": "全市均价",
    "count": "14694",
    "unit": "元\/平"
  }, {
    "name": "昨日成交",
    "count": "56",
    "unit": "套"
  }],
  "m_url": "https:\/\/m.ke.com\/cd\/fangjia",
  "more_desc": "查看更多"
},
```

### 新上房源
`curl -H 'Authorization: MjAxODAxMTFfaW9zOmFkNWNjZjY1MzUzZDNmYWQzMjEyNjE3MmMwMGZmN2VhYzJlNTUyZGI=' 'https://app.api.ke.com/house/ershoufang/searchv4?cityId=510100&condition=tt2&fullFilters=1&hasRecommend=0&limitCount=20&limitOffset=0&request_ts=1539825002'`

分析 data.list数组中某项数据houseCode存在, 获取`infoList`中的`name`为`挂牌`的获取挂牌时间, 按天归类数据
```json
{
      "houseCode": "106101711899",
      "title": "启明花园大套二可改套三，户型方正，满两年，对中庭",
      "desc": "2室2厅|113.03㎡|东南|启明花园",
      "priceStr": "95",
      "priceUnit": "万",
      "unitPriceStr": "8,405元\/平",
      "coverPic": "http:\/\/ke-image.ljcdn.com\/510100-inspection\/test-d49c92b0-22e7-43b3-91e0-a93cccbc100e.png.280x210.jpg?from=ke.com",
      "colorTags": [{
        "desc": "VR房源",
        "color": "849AAE",
        "boldFont": 0
      }, {
        "desc": "新上房源",
        "color": "849AAE",
        "boldFont": 0
      }],
      "isVr": true,
      "isVideo": false,
      "frameStr": "2室2厅",
      "cardType": "house",
      "isFocus": false,
      "basicList": [{
        "name": "售价",
        "value": "95万"
      }, {
        "name": "房型",
        "value": "2室2厅"
      }, {
        "name": "建筑面积",
        "value": "113.03㎡"
      }],
      "infoList": [{
        "name": "单价：",
        "value": "8405元\/平"
      }, {
        "name": "挂牌：",
        "value": "2018.10.11"
      }, {
        "name": "朝向：",
        "value": "东南"
      }, {
        "name": "楼层：",
        "value": "高楼层\/6层"
      }, {
        "name": "楼型：",
        "value": "板塔结合"
      }, {
        "name": "电梯：",
        "value": "无"
      }, {
        "name": "装修：",
        "value": "其他"
      }, {
        "name": "年代：",
        "value": "2003年"
      }, {
        "name": "用途：",
        "value": "普通住宅"
      }, {
        "name": "权属：",
        "value": "商品房"
      }],
      "communityId": "3011053471852",
      "communityName": "启明花园",
      "baiduLa": 30.773253,
      "baiduLo": 103.970861,
      "blueprintHallNum": 2,
      "blueprintBedroomNum": 2,
      "area": 113.03,
      "price": 950000,
      "unitPrice": 8405,
      "strategyInfo": "{\"fb_query_id\":\"105230904660111360\",\"fb_expo_id\":\"105230904660115456\",\"fb_item_location\":\"0\",\"fb_service_id\":\"1011710017\",\"fb_ab_test_flag\":\"[\\\"ab-test-exp-249-group-2\\\",\\\"ab-test-exp-279-group-3\\\"]\",\"fb_item_id\":\"106101711899\"}",
      "houseSource": "104",
      "serviceCommitment": "1",
      "isTop": 0
    }
```

### 每日成交数据
`curl -H 'Authorization: MjAxODAxMTFfaW9zOmY0NDA1MDRjZTI5MDk0NzY1ZTgzZjBlNjc4NjllNWJlYzMwOTM0OGY=' 'https://app.api.ke.com/house/chengjiao/searchv2?channel=sold&city_id=510100&limit_count=20&limit_offset=0&request_ts=1539825667'`

`limit_count`为每页数量,
`limit_offset`为那一条数据开始查询
`request_ts`为每次访问时间

直接分析`data.list`数组中数据即可
```json
{
  "house_code": "106101510602",
  "title": "仁厚街38号 3室1厅",
  "desc": "东南\/高楼层\/7层\/65.88㎡",
  "resblock_id": "1611061558346",
  "sign_date": "2018.10.03",
  "price_str": "86",
  "price_unit": "万",
  "unit_price_str": "13055元\/平",
  "require_login": 0
}
```






