package main
//import "github.com/nelsonken/cos-go-sdk-v5/cos"
import "fmt"
import (
    "context"
    "net/http"
    "net/url"
    "os"
    _"strings"
    "github.com/tencentyun/cos-go-sdk-v5"
)
func main(){
    command := os.Args[1]
    u, _ := url.Parse("https://buck-1301342943.cos.ap-nanjing")
    b := &cos.BaseURL{BucketURL: u}
    fmt.Println(b)
    c := cos.NewClient(b, &http.Client{
        Transport: &cos.AuthorizationTransport{
            SecretID:  "AKIDIN2RbW6bkcwrWxo4k8G7mBJCqDJpo8tP",
            SecretKey: "d2e04sRjlX3IGxOM3UeXR3nGRMzLPOq2",
        },
    })
    //创建存储桶
    //fmt.Println("创建存储桶")
     //_, err := c.Bucket.Put(context.Background(), nil)
    //if err != nil {
     //   panic(err)
    //}
    //获取存储桶列表
    fmt.Println("获取存储桶列表")
    s, _, err := c.Service.Get(context.Background())
    if err != nil {
        panic(err)
    }

    for _, b := range s.Buckets {
        fmt.Printf("%#v\n", b)
    }
    fmt.Println("开始上传......")
    name := command
    //fmt.Println(name1)
    //name := "/root/Python-2.7.14.tgz"
    //f := strings.NewReader("root")
    //fmt.Println(context.Background())
    //_, err = c.Object.Put(context.Background(), name, f, nil)
    //if err != nil {
    //    panic(err)
   //}
    // 2.通过本地文件上传对象
    _, err = c.Object.PutFromFile(context.Background(), name, name, nil)
    if err != nil {
        panic(err)
    }







}
