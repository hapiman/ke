### 时区转化

- 区别当前时区, 返回`Location`
主要使用`Time.LoadLocation`方法，`Local`表示当前服务器所在时区，`UTC`或者`""`表示0时区，
`America/Los_Angeles`表示洛杉矶时间，`Asia/Chongqing`表示重庆时间，也就是北京时间
```go
formate:="2006-01-02 15:04:05 Mon"
local1, err1 := time.Now().LoadLocation("UTC") //输入参数"UTC"，等同于""
local2, err2 := time.Now().LoadLocation("Local")
local3, err3 := time.Now().LoadLocation("America/Los_Angeles")
```

- 时间初始化
``` go
//通过字符串，默认UTC时区初始化Time
func Parse(layout, value string) (Time, error)
//通过字符串，指定时区来初始化Time; 即以什么时区来处理该字符串
func ParseInLocation(layout, value string, loc *Location) (Time, error)
//通过unix 标准时间初始化Time
func Unix(sec int64, nsec int64) Time
```

- 时间转化
```go
local, _ := time.LoadLocation("America/Los_Angeles")
timeFormat := "2006-01-02 15:04:05"
time1 := time.Unix(1480390585, 0) //通过unix标准时间的秒，纳秒设置时间
time2, _ := time.ParseInLocation(timeFormat, "2016-11-28 19:36:25", local) //将指定的时间当作洛杉矶时间处理
fmt.Println(time1.In(local).Format(timeFormat))
fmt.Println(time2.In(local).Format(timeFormat)) // 当前的In(local)没有作用， 因为time2中带了local作为时区
chinaLocal, _ := time.LoadLocation("Local") //运行时，该服务器必须设置为中国时区，否则最好是采用"Asia/Chongqing"之类具体的参数。
fmt.Println(time2.In(chinaLocal).Format(timeFormat))
//output:
//2016-11-28 19:36:25
//2016-11-28 19:36:25
//2016-11-29 11:36:25
```
- 其他
将已知字符串当作0时区时间中，转化为中国时间
``` go
newTStr := "2016-12-04 15:39:06"
timeFormat := "2006-01-02 15:04:05"
localUtc, _ := time.LoadLocation("")
t, _ := time.ParseInLocation(timeFormat, newTStr, localUtc)
chinaLocal, _ := time.LoadLocation("Asia/Chongqing")
t.In(chinaLocal).Format(timeFormat)
```
将已知的时间戳转化为某个时区的时间字符串
```go
timeFormat := "2006-01-02 15:04:05"
t := time.Unix(1480390585, 0)
local, _ := time.LoadLocation("Asia/Chongqing")
t.In(local).Format(timeFormat)
```
