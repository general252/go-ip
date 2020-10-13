# go-ip
查询ip的地址


### 1. 百度 open data
http://opendata.baidu.com/api.php?query=114.114.114.114&co=&resource_id=6006&t=1602555423&ie=utf8&oe=utf8&format=json

### 1. 百度
https://sp0.baidu.com/8aQDcjqpAAV3otqbppnN2DJv/api.php?query=114.114.114.114&co=&resource_id=5809&ie=utf8&oe=utf8&t=1602555423&format=json

### 2. pcOnline太平洋电脑网API(更多有用信息http://whois.pconline.com.cn)
http://whois.pconline.com.cn/ip.jsp?ip=114.114.114.114
http://whois.pconline.com.cn/ipJson.jsp?ip=114.114.114.114

### 3 网易IP查询API接口
http://ip.ws.126.net/ipquery?ip=114.114.114.114

### 4 淘宝
http://ip.taobao.com/outGetIpInfo?ip=114.114.114.114&accessKey=alibaba-inc


```
go get github.com/general252/go-ip
```

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/general252/go-ip"
)

func main() {
	if result, err := go_ip.GetIpAddress("47.121.184.45"); err != nil {
		fmt.Println(err)
	} else {
		data, _ := json.MarshalIndent(result, "", "  ")
		fmt.Println(string(data))
	}
}

```

输出
```
{
  "ip": "47.121.184.45",
  "address": "浙江省杭州市 阿里云",
  "type": "Baidu OpenData",
  "timeout": 62
}
```