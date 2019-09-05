package sago

import (
	"fmt"
	utils "gitee.com/xiawucha365/sago/internal/tool"
	"sync"
	"time"
)

///共享协程池
type SPool struct {
	work          chan WorkerInterface
	wg            sync.WaitGroup
	TotalNum      int
	Counter       int
	CounterFail   int
	mutexFail     sync.Mutex
	mutexOk       sync.Mutex
	TimeStart     int64
	TimeOut       int
	MaxGoroutines int
}

// 协程池
func NewSPool(maxGoroutines int) *SPool {
	p := SPool{
		MaxGoroutines: maxGoroutines,
		work:          make(chan WorkerInterface),
	}
	//任务开始时间记录
	p.TimeStart = time.Now().Unix()
	return &p
}

func (p *SPool) Run() {
	p.wg.Add(p.MaxGoroutines)

	for i := 0; i < p.MaxGoroutines; i++ {
		//协程池
		go func() {

			for w := range p.work {
				timeout_ch := make(chan interface{})

				for {
					select {
					case <-timeout_ch:
						//logger.Info(wr.GetTaskID(), ok)
						goto ForEnd
					case <-time.After(time.Duration(p.TimeOut) * time.Second):
						//打印超时的任务id
						//p.CountFail()
						Log.Warn(w.GetTaskID(), "timeout")
						goto ForEnd
					}

				}
			ForEnd:

				go func() {
					p.runtask(w, timeout_ch)
				}()
			}

			// 收尾工作 容灾
			//defer func() {
			//	p.wg.Done()
			//	if err := recover(); err != nil {
			//		Log.Error("task error", err)
			//	}
			//}()
			//
			//for w := range p.work {
			//	//消费 并 执行job里的任务
			//	if err := w.Task(); err != nil {
			//		p.CountFail()
			//		panic(err)
			//	} else {
			//		//计数器
			//		p.CountOk()
			//	}
			//}
		}()
	}
}

func (p *SPool) runtask(wr WorkerInterface, timeout_ch chan interface{}) {
	defer func() {
		if err := recover(); err != nil {
			Log.Error("task error", err)
		}
	}()

	//执行job里的任务
	err := wr.Task()
	if err != nil {
		p.CountFail()
		panic(err)
	} else {
		p.CountOk()
	}
	timeout_ch <- "ok"
}

// 提交任务
func (p *SPool) Commit(w WorkerInterface) {
	p.work <- w
}

// 等待组 关闭channel
func (p *SPool) Shutdown() {
	close(p.work)
	p.wg.Wait()

}

//计数器
func (p *SPool) CountOk() {
	p.mutexOk.Lock()
	p.Counter++
	p.mutexOk.Unlock()

}

//计数器-失败
func (p *SPool) CountFail() {
	p.mutexFail.Lock()
	p.CounterFail++
	p.mutexFail.Unlock()

}

// log
func (p *SPool) Runtimelog() {
	//计数器
	ttime := utils.MathDecimal(float64(time.Now().Unix() - p.TimeStart))
	trange := utils.MathDecimal(float64(p.TotalNum) / ttime)
	if p.Counter > 0 || p.CounterFail > 0 {
		Log.Sucess(fmt.Sprintln("runtime:总数|成功|失败:", p.TotalNum, "|", p.Counter, "|", p.CounterFail, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
	}
}

//
////共享固定goroutine 例子
//var woker_queues = []string{"001","002","003","004","005","006","007"}
//type worker2 struct {
//	name string
//}
//
////业务代码
//func(m *worker2) Task() error {
//
//	// 收尾工作 容灾
//	defer func(){
//		if err:= recover();err!=nil{
//			fmt.Println(err)
//		}
//	}()
//
//	fmt.Println("job:"+m.name+" start")
//	time.Sleep(time.Second*3)
//	fmt.Println("job:"+m.name+" end")
//
//	return nil
//}
//func main()  {
//	//创建并发池
//	p := pipe.NewSPool(4)
//
//	for _,name := range woker_queues{
//		np := worker2{name:name,}
//		//提交任务
//		p.Commit(&np)
//	}
//
//	//等待组 关闭channel
//	p.Shutdown()
//
//}
