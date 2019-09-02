# sago

西米露,一种好吃且富有营养的物质,需要手工加工


cli并行框架
提供一套轻量可控的脚本开发脚手架,降低并发程序开发难度
#### 软件架构
- 支持mysql,redis,orm采用xorm
- 支持测试/生产环境配置文件区分
- 统一日志输出,日志切分
#### 软件目录
- console  存放具体项目业务代码
- internal  项目核心代码
    - app 项目初始化代码
    - db 数据库组件
    - logger 日志组件
    - tool  工具组件(文件,网络,变量转换等)
    - pipe  并发组件(支持协程复用模式和单次释放模式,支持并发数和超时时间设置)
- cmd  脚手架工具等(生成orm-model文件)
 

### 使用说明
- go mod 安装
    - 需要安装 mod 包,安装成功后配置环境变量  
        - export GO111MODULE=on 
        - export GOPROXY=https://goproxy.cn
- IDE选择:建议使用goland版本 >= 2019.2,安装好后会自动检测并引入依赖,go mod是golang新出特性,旧版本不支持



### 新项目初始化
- 1 利用脚手架工具生成一个新项目,如下
```go
cd cmd
./sago_cmd -pname 项目名 
```
&emsp;&emsp;&emsp;table参数如不指定，则不生成model下的文件
- 2 将生成的目录拷贝到console下面,结束

###  项目部署
- 交叉编译:mac,windows,linux之间是不兼容的,在mac下生成linux上可执行文件命令
```go
GOOS=linux GOARCH=amd64 go build -o worker_pcpro  main.go
```
- 生产环境:
```go
//生产环境-prod.toml配置文件
./sago  -env=prod

//开发环境-对应etc下dev.toml配置文件
./sago  -env=dev
```


### 其他examples
####  常规使用（参考export项目）
```go
package main

import (
	"fmt"
	"gitee.com/xiawucha365/sago"
	"strconv"
)

func main() {

defer logger.flush()
//使用mysql 
var priceSeekoriItems []models.THotelPriceSeekori
app.GMysql.Select("id,ctime").Where("ctime > ? ", lastTime, ).Order("id asc ").Find(&priceSeekoriItems)

//使用redis
app.GRedis.Set("test_key","message")

reply,err := app.GRedis.Get("test_key")

if(err != nill){
    panic(err)
}else{
    fmt.Println(reply)
}
//日志输出/记录到文件
logger.Warn("hello,world")
}
```

####  并发相关 （参考spider_worker）
```go
package main

import (
	"fmt"
	"gitee.com/xiawucha365/sago"
	"time"
)

//并发池实例
var wpool *pipe.WPool

//任务实例
type worker struct {
	name string
}

//要执行的任务列表
var name_slices = []string{"001", "002", "003", "004", "005", "006", "007", "008", "009"}

//单任务的业务逻辑代码
func (m *worker) Task() error {

	fmt.Println("job:" + m.name + " start")
	time.Sleep(time.Second * 3)
	fmt.Println("job:" + m.name + " end")
	return nil
}

//例子演示
func main() {

	logger.Warn("并发开始")
	wool = pipe.NewWPool(4, len(name_slices))
	wool.Run()

	for _, name := range name_slices {
		np := worker{
			name: name,
		}
		wool.Commit(&np)
	}

	wool.Shutdown()

}


```

![gopthercn](https://gitlab.mfwdev.com/WebDev/hotel/uploads/23836c8a5c7695fc860d976cde8f7f79/gopthercn.png?~/w/150)