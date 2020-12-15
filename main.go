package main

import (
  "context"
  "fmt"
  "time"
)

type DemoOne struct {
  id  int8
  age int32
  add int8
}
type DemoTwo struct {
  age int32
  id  int8
  add int8
}

func main() {
  //r:=gin.Default()
  ////func
  //r.GET("get",getting)
  /**
  指针对齐 位置不同占用内存大小不同
  空结构体的宽度是0，他占用了0字节的内存空间  由空结构体组成的也不占用内存空间
   */
  //fmt.Println(unsafe.Sizeof(DemoOne{}))//size为12
  //fmt.Println(unsafe.Sizeof(DemoTwo{}))//size为8

  //u1:=uuid.Must(uuid.NewV4())
  //fmt.Printf("%s",u1)
  //u2,err:=uuid.NewV4()
  //if err!=nil {
  //  fmt.Printf("something wrong %s",u2)
  //  return
  //}
  //fmt.Printf("%s",u2)
  //u2, _ = uuid.FromString("476698d0-b0b2-4990-9216-8d15f59d61c9")
  //fmt.Printf("successfully passed:%s",u2)

  //超时怎么做 模拟控制函数
  /**
    一个routine启动以后无法控制它 如果他不会自己结束 那么将会导致严重的内存泄漏
   1.控制并发 group. context
   2.群组控制 group 单个控制context
   3.优雅解决单个协程退出 select+channle  其中channel换成context 追踪协程
   context.Background() 根节点context
   指示函数在cancel中
   */
  //ctx,cancel:=context.WithTimeout(context.Background(),1*time.Second)
  //defer cancel()
  //select {
  //case <-time.After(1*time.Second):fmt.Println("overslept")
  //case <-ctx.Done():fmt.Println(ctx.Err())
  //       }

      //stop:=make(chan bool)
      //go func() {
      //  for  {
      //    select {
      //    case <-stop: fmt.Println("监控退出，停止了 退出协程...")
      //      return
      //    default:
      //      fmt.Println("监控中。。。")
      //      time.Sleep(2*time.Second)
      //    }
      //  }
      //}()
      //time.Sleep(10*time.Second)
      //stop<-true
      //time.Sleep(5*time.Second)//程序还将运行5s


      //一个改造
  ctx, cancel := context.WithCancel(context.Background())
  go func(ctx context.Context) {
    for {
      select {
      case <-ctx.Done():
        fmt.Println("监控退出，停止了...")
        return
      default:
        fmt.Println("goroutine监控中...")
        time.Sleep(2 * time.Second)
      }
    }
  }(ctx)

  time.Sleep(10 * time.Second)
  fmt.Println("可以了，通知监控停止")
  cancel()
  //为了检测监控过是否停止，如果没有监控输出，就表示停止了
  time.Sleep(5 * time.Second)
}

//context的作用
/**
  1.当父context被取消 所有的都被取消
  2.with 返回新的context 和函数 CancelFunc 取消子代
  3.相同的context并发是安全的
  4.包含结构体 结构体方法
  5.防止泄漏 需要结束
原理：
当我们调用cancel()时，相当与给Done发送了一个终止信号，在每一个使用Context方法里都会select监听这个Done，所以，调用cancel()时，会将所有使用这个content的方法全部停掉

 */
//func getting(context *gin.Context) {
//
//}
