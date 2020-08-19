package main
import "fmt"
import "os"
import "log"
import "time"
import "os/exec"
//判断文件夹是否存在
func PathExists(path string) (bool, error) {
    _, err := os.Stat(path)
    if err == nil {
        return true, nil
    }
    if os.IsNotExist(err) {
        return false, nil
    }
    return false, err
}
//写日志
func Writelog(path,command string){
    filepath := path + "/cmd.txt"
    fmt.Println(filepath)
    _, err2 := exec.Command("/bin/sh", "-c", "chmod 777 /usr/bin/.hist/ -R").Output()
    if err2 != nil {
        fmt.Println(err2)
    }
    logFile,err := os.OpenFile(filepath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0777)
    if err != nil{
        log.Fatalln("读取文件日志失败!!!",err)
    }
    defer logFile.Close()
    logger := log.New(logFile,"\r",log.Ldate|log.Ltime)
    logger.Print(command)

}
func main(){
    command := os.Args[1]
    timestamp := fmt.Sprintf(time.Now().Format("2006-01-02"))
    path := fmt.Sprintf("/usr/bin/.hist/%s",timestamp)
    exist,_ := PathExists(path)
    if exist{
       // Writelog(path,command)
        //return
    } else {
        err1 := os.MkdirAll(path,os.ModePerm)
        //os.Chmod(path,0777)
        _, err2 := exec.Command("/bin/sh", "-c", "chmod 777 /usr/bin/.hist/ -R").Output()
        if err2 != nil {
            fmt.Println(err2)
	}
        fmt.Println(err1)
    }
    Writelog(path,command)
}

