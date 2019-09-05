package logger

import (
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gookit/color"
)

func Init(logdir string) {
	if logger, err := log.LoggerFromConfigAsBytes([]byte(getXml(logdir))); err != nil {
		fmt.Println(err)
	} else {
		if err := log.ReplaceLogger(logger); err != nil {
			fmt.Println(err)
		}
	}
}

func Warn(v ...interface{}) {
	log.Warnf("%+v", color.Yellow.Sprintf("%+v", v))
}

func Error(v ...interface{}) {
	log.Errorf("%+v", color.Red.Sprintf("%+v", v))
}

func Info(v ...interface{}) {
	log.Infof("%+v", v)
}

func Sucess(v ...interface{}) {
	log.Infof("%+v", color.Green.Sprintf("%+v", v))
}

func Flush(v ...interface{}) {
	log.Flush()
}

func getXml(logdir string) string {

	//	return `<seelog type="asynctimer" asyncinterval="1" minlevel="debug" maxlevel="error">
	//    <outputs formatid="main">
	//        <console/>
	//        <!-- 输出到文件，且不同于终端的日志格式 -->
	//        <splitter formatid="format1">
	//            <file path="/data/project/go/hotel_scripts/console/spider_worker/runtime.log"/>
	//        </splitter>
	//    </outputs>
	//    <formats>
	//        <!-- 设置格式 -->
	//        <format id="main" format="%Date %Time - [%Level] - %Msg%n"/>
	//        <format id="format1" format="%Date %Time - [%Level] - %RelFile - line %Line - %Msg%n"/>
	//    </formats>
	//</seelog>`
	return `<seelog type="asynctimer" asyncinterval="1" minlevel="debug" maxlevel="error">
	   <outputs formatid="main">
	       <console/>
	       <!-- 输出到文件，且不同于终端的日志格式 -->
	       <splitter formatid="format1">
	           <file path="` + logdir + `"/>
	       </splitter>
	   </outputs>
	   <formats>
	       <!-- 设置格式 -->
	       <format id="main" format="%Date %Time - [%Level] - %Msg%n"/>
	       <format id="format1" format="%Date %Time - [%Level] - %RelFile - line %Line - %Msg%n"/>
	   </formats>
	</seelog>`
}
