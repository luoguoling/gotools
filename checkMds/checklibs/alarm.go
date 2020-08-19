package checklibs
import (
	"checkMds/config"
	"fmt"
	"net/http"
	"strings"
)
//钉钉报警
func SendDingMsg1(msg string) {
	//webHook := `https://oapi.dingtalk.com/robot/send?access_token=187918ed0afc579ee5e458f0bf23c84a1bafdd1782b683ad42873b4d41bba0d7`
	webHook := config.GetConfig().WebHook
	fmt.Println("报警地址打印")
	fmt.Println(webHook)

	msg =  config.GetConfig().ProjectName +   "  " + msg
	content := `{"msgtype": "text",
		"text": {"content": "`+ msg + `"}
	}`
	//创建一个请求
	for _,webhook1 := range  webHook {
		fmt.Println(webhook1)
		req, err := http.NewRequest("POST", webhook1, strings.NewReader(content))
		if err != nil {
			fmt.Println(err)
			fmt.Println("钉钉报警请求异常")
		}

		client := &http.Client{}
		//设置请求头
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		//发送请求
		resp, err := client.Do(req)
		//关闭请求
		defer resp.Body.Close()

		if err != nil {
			// handle error
			fmt.Println(err)
			fmt.Println("顶顶报发送异常!!!")
		}
		//logger.MyLogger.Error("aaerw")
		//logrus.WithFields(logrus.Fields{"aa":"aa","username":"rolin"}).Info("aaaa")
		//logrus.Error("aaaa")
	}

}


