package checklibs

import (
	"checkMds/config"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)
//定义需要检查的字段
type Mds struct {
	PlatformId sql.NullInt64 `db:"platformId";null`
	PlatformAlias sql.NullString	`db:"platformAlias"`
	ServerId	sql.NullInt64	`db:"serverId"`
	ServerName	sql.NullString	`db:"serverName"`
	ServerIp	sql.NullString	`db:"serverIp"`
	GateMinaServerPort	sql.NullInt64	`db:"gateMinaServerPort";null`
	OpenTime	sql.NullInt64	`db:"openTime"`
}
var Db *sqlx.DB
func init() {
	dburl := config.GetConfig().DbUrl
	dbuser := config.GetConfig().DbUser
	dbport := config.GetConfig().DbPort
	dbpwd := config.GetConfig().DbPwd
	dbname := config.GetConfig().DbName
	fmt.Println(dburl,dbuser,dbport,dbpwd,dbname)
	//dbSourceName := fmt.Printf("%d:%d@tcp(%d:%d)/%d", (dbuser,dbpwd,dburl,dbport,dbname)
	database, err := sqlx.Open("mysql", "xx:xx@tcp(xx:3306)/xx")

	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	Db = database
	//defer Db.Close()  // 注意这行代码要写在上面err判断的下面
}
func B2S(bs []int8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
		}
	return string(ba)
}
func CheckMds(platformalias string)  {
	var mds []Mds
	err := Db.Select(&mds,"select platformId,platformAlias,serverId,serverName,serverIp,gateMinaServerPort,openTime from mds_server where platformAlias=?",platformalias)
	if err != nil{
		fmt.Println("exec failed",err)
		return
	}
	fmt.Println("select success",mds)
	NullKeyList := make([]string,0)
	//var NullKeyList []string
	for _,j := range mds{

		if !j.GateMinaServerPort.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "GateMinaServerPort" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}
		if !j.PlatformAlias.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "PlatformAlias" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}
		if !j.ServerId.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "ServerId" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}
		if !j.OpenTime.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "OpenTime" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}
		if !j.ServerIp.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "ServerIp" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}
		if !j.ServerName.Valid{
			str := j.PlatformAlias.String + ":"+ strconv.FormatInt(j.ServerId.Int64,10) + " " + "ServerName" + "这个字段为空,请及时修改!!!"
			NullKeyList = append(NullKeyList, str)
		}



	}
	fmt.Println(NullKeyList)
	NullKeyListstr,err := json.Marshal(NullKeyList)
	resultstr := string([]byte(NullKeyListstr))
	if err != nil{
		fmt.Println("解析异常")
	}
	currentime := time.Now().Format("2006-01-02 15:04:05")
	result :=  config.GetConfig().ProjectName +"项目:"+ string(resultstr) + "请注意修改!!!" +  string(currentime)
	fmt.Println(result)
	if len(NullKeyList) != 0{
		//SendDingMsg1(result)
		for _,value := range NullKeyList{
			message := config.GetConfig().ProjectName + value + string(currentime)
			SendDingMsg1(message)
		}
	}


}
func CheckAll()  {
	platformAlias := config.GetConfig().PlatformAlias
	for _, plat := range platformAlias {
		CheckMds(plat)
	}

}
