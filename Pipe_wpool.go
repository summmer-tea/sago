package sago

import (
	"fmt"
	utils "gitee.com/xiawucha365/sago/internal/tool"
	"sync"
	"time"
)

//非共享协程池
type WPool struct {
	//任务传递
	work chan WorkerInterface
	//并发数控制
	limitChan   chan interface{}
	wg          sync.WaitGroup
	TotalNum    int
	Counter     int
	CounterFail int
	mutexFail   sync.Mutex
	mutex       sync.Mutex
	TimeStart   int64
	TimeOut     int
}

// 协程池
func NewWPool(maxGoroutines int, totalNum int, timeOut int) *WPool {

	p := WPool{
		TotalNum:  totalNum,
		work:      make(chan WorkerInterface),
		limitChan: make(chan interface{}, maxGoroutines),
	}

	p.wg.Add(totalNum)
	p.TimeOut = timeOut

	return &p
}

// 提交任务
func (p *WPool) Commit(w WorkerInterface) {
	p.limitChan <- "ok"
	p.work <- w

}

// 控制最大并发数
func (p *WPool) Run() {
	//任务开始时间记录
	p.TimeStart = time.Now().Unix()

	//新起一个协程
	go func() {
		for w := range p.work {

			go func(wr WorkerInterface) {

				// 收尾工作 容灾
				defer func() {
					p.wg.Done()
					<-p.limitChan
					if err := recover(); err != nil {
						Log.Error("task error", err)
					}

					//p.runtimelog()
				}()

				//超时机制start
				if p.TimeOut > 0 {
					timeout_ch := make(chan interface{})

					go func() {
						p.runtask(wr, timeout_ch)
					}()

					for {
						select {
						case <-timeout_ch:
							//logger.Info(wr.GetTaskID(), ok)
							goto ForEnd
						case <-time.After(time.Duration(p.TimeOut) * time.Second):
							//打印超时的任务id
							//p.CountFail()
							Log.Warn(wr.GetTaskID(), "timeout")
							goto ForEnd
						}

					}
				ForEnd:
					//超时机制end
				} else {
					//没有设置超时时间情况下
					//执行job里的任务
					if err := wr.Task(); err != nil {
						p.CountFail()
						panic(err)
					} else {
						//成功
						p.CountOk()
					}
				}

			}(w)

		}
	}()

}

func (p *WPool) runtask(wr WorkerInterface, timeout_ch chan interface{}) {
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

// 等待组 关闭channel
func (p *WPool) Shutdown() {
	close(p.work)
	p.wg.Wait()

}

//计数器
func (p *WPool) CountOk() {
	p.mutex.Lock()
	//runtime.Gosched()
	p.Counter++
	p.mutex.Unlock()

}

//计数器-失败
func (p *WPool) CountFail() {
	p.mutexFail.Lock()
	//runtime.Gosched()
	p.CounterFail++
	p.mutexFail.Unlock()

}

// log
func (p *WPool) Runtimelog() {
	//计数器
	ttime := utils.MathDecimal(float64(time.Now().Unix() - p.TimeStart))
	trange := utils.MathDecimal(float64(p.TotalNum) / ttime)
	if p.Counter > 0 || p.CounterFail > 0 {
		Log.Sucess(fmt.Sprintln("runtime:总数|成功|失败:", p.TotalNum, "|", p.Counter, "|", p.CounterFail, "", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)"))
	}
}

////并发池实例
//var wool *pipe.Wool
//
////任务实例
//type worker struct {
//	name string
//}
//
////要执行的任务列表
//var name_slices = []string{"001", "002", "003", "004", "005", "006", "007", "008", "009"}
//
//
//func (m *worker) Task() error {
//
//	fmt.Println("job:" + m.name + " start")
//	time.Sleep(time.Second * 3)
//	fmt.Println("job:" + m.name + " end")
//	return nil
//}
//
////例子演示
//func main() {
//
//	//初始化框架
//	app.Init()
//	_ = app.G_log.Warn("并发开始")
//	wool = pipe.NewWool(4, len(name_slices))
//	wool.Run()
//
//	for _, name := range name_slices {
//		np := worker{
//			name: name,
//		}
//		wool.Commit(&np)
//	}
//
//	wool.Shutdown()
//
//}
