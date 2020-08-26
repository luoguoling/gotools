package main

import (
	"fmt"
	"os"
	"runtime/trace"

)
import "sync"
func mockSendToServer(url string){
	fmt.Printf("SERVER URL:%s \n",url)
}
func main()  {
	f,err := os.Create("trace.out")
	if err != nil{
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil{
		panic(err)
	}
	defer trace.Stop()
	urls := []string{"0.0.0.0:1000","0.0.0.0:2000","0.0.0.0:3000","0.0.0.0:4000"}
	wg := sync.WaitGroup{}
	for _,url := range urls{
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			mockSendToServer(url)
		}(url)

	}
	wg.Wait()
}
go tool trace trace.out
https://mp.weixin.qq.com/s/hIs318h6iJW2O9--QVqh6w#

