package main

import (
	// "errors"
	"fmt"
	seelog "github.com/cihub/seelog"
	// "io"
)

var Logger seelog.LoggerInterface

func loadAppConfig() {
	appConfig := `
<seelog minlevel="warn">
	<outputs formatid="common">
		<rollingfile type="size" filename="./logs/roll.log" maxsize="100000" maxrolls="5"/>
		<filter levels="critical">
			<file path="./logs/critical.log" formatid="critical" />
			<smtp formatid="criticalemail" senderaddress="li_yunteng@163.com" sendername="ShortUrl API" hostname="smtp.163.com" hostport="25" username="li_yunteng" password="yun1988">
				<recipient address="li_yunteng@163.com" />
			</smtp>
		</filter>
	</outputs>
	<formats>
		<format id="common" format="%Date/%Time [%LEV] %Msg%n" />
		<format id="critical" format="%File %FullPath %Func %Msg%n" />
		<format id="criticalemail" format="Critical error on our server!\n     %Time %Date %RelFile %Func %Msg \nSent by Seelog" />
	</formats>
</seelog>
`
	logger, err := seelog.LoggerFromConfigAsBytes([]byte(appConfig))
	if err != nil {
		fmt.Println(err)
		return
	}

	UseLogger(logger)
}

func init() {
	DisableLog()
	loadAppConfig()
}

func DisableLog() {
	Logger = seelog.Disabled
}

func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func main() {
	Logger.Info("this is  a test!")
	Logger.Critical("this is an error!")
	Logger.Trace("this is a Trace")
}
