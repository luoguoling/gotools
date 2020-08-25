package main

import (
	"errors"
	"fmt"
	"github.com/robfig/cron"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"
	"time"
)

//获取本机ip
func getLocalIp() (net.IP, error){
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}
//格式化ip格式
func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}
//获取公有ip
func getPubIp() string{
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Stderr.WriteString("\n")
		os.Exit(1)
	}
	defer resp.Body.Close()
	//fmt.Println(resp.Body,os.Stdout)
	//io.Copy(os.Stdout, resp.Body)

	//os.Exit(0)
	data,err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(data))
	return string(data)


}
//发送钉钉消息
func SendDingMsg(msg string) {
	webHook := `https://oapi.dingtalk.com/robot/send?access_token=d6c65b2c15aaae24da3c506a8fe329f1565b09c2e3186ef93f10a078eb66f872`
	//webHook := config.GetConfig().WebHook

	//msg =  "工会项目ip不符合规则呀!!!" + "  " + msg
	content := `{"msgtype": "text",
		"text": {"content": "`+ msg + `"}
	}`
	//创建一个请求
	req, err := http.NewRequest("POST", webHook, strings.NewReader(content))
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
//判断是否存在某个元素
func IsExistInArray(value string,array []string) bool{
	for _,v := range array{
		if v == value {
			return true
		}
	}
	return false
}
func main(){
	unix := time.Now().Unix()
	timeLayout := "2006-01-02 15:04:05"
	timeStr := time.Unix(unix, 0).Format(timeLayout)
	fmt.Println(timeStr)
	ip, err := getLocalIp()
	if err != nil {
		fmt.Println(err)
	}

	//fmt.Println(ip.String())
	localip := ip.String()
	fmt.Println("开始执行检查任务!!!")
	c := cron.New()
	spec := "0 58 15 * * ?"
	c.AddFunc(spec,func(){
		//i++
		msg := "工会项目: " + timeStr + "   " + localip + "这个ip的定时任务已经开始启动了!!!"
		SendDingMsg(msg)
		//fmt.Println("cron is runing")
	})
	c.Start()
	for {
		//sip := "115.159.89.39"
		//bip := "58.87.73.248"
		ip := getPubIp()
		result := "工会项目IP不符合规则:" + ip + "  此人私有ip是:" + localip
		//fmt.Println(result)
		//cip := "222.190.107.198"
		dip := []string{"58.87.73.248","115.159.89.39","222.190.107.198"}
		if !IsExistInArray(ip,dip){
			SendDingMsg(result)
		}
		time.Sleep(60*time.Second)
		//计划任务开始


		//select{}
	}
	//定时任务调用
	//crontab := cron.New()
	//task := func() {
	//	//fmt.Println("hello world")
	//	msg := "工会项目: " + timeStr + ":" + localip + "这个ip的定时任务已经开始启动了!!!"
	//	fmt.Println(msg)
	//	SendDingMsg(msg)
	//}
	//// 添加定时任务, * * * * * 是 crontab,表示每分钟执行一次
	//crontab.AddFunc("0 11 15 * * *", task)
	//// 启动定时器
	//crontab.Start()
	//// 定时任务是另起协程执行的,这里使用 select 简答阻塞.实际开发中需要
	//// 根据实际情况进行控制
	//select {}
	//i := 0


}
