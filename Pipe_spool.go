package sago

import (
	"gitee.com/xiawucha365/sago/internal/logger"
	utils "gitee.com/xiawucha365/sago/internal/tool"
	"sync"
	"time"
)

///共享协程池
type SPool struct {
	work      chan WorkerInterface
	wg        sync.WaitGroup
	Counter   int
	mutex     sync.Mutex
	TimeStart int64
}

// 协程池
func NewSPool(maxGoroutines int) *SPool {
	p := SPool{
		work: make(chan WorkerInterface),
	}
	//任务开始时间记录
	p.TimeStart = time.Now().Unix()

	p.wg.Add(maxGoroutines)

	for i := 0; i < maxGoroutines; i++ {
		//协程池
		go func() {

			for w := range p.work {
				//消费 并 执行job里的任务
				if err := w.Task(); err != nil {
					logger.Error("task error", err)
				} else {
					//计数器
					p.Count()

					ttime := utils.MathDecimal(float64(time.Now().Unix() - p.TimeStart))
					trange := utils.MathDecimal(float64(p.Counter) / ttime)
					logger.Info("runtime:总数(", p.Counter, ")", "消耗时间:(", ttime, "s)", "平均:(", trange, "次/s)")
				}
			}

			defer p.wg.Done()

		}()
	}

	return &p
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
func (p *SPool) Count() {
	p.mutex.Lock()
	//runtime.Gosched()
	p.Counter++
	p.mutex.Unlock()

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
