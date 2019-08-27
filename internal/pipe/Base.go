package pipe

//协程池接口
//type PoolInterface interface {
//	Commit(w WorkerInterface)
//}

//任务接口
type WorkerInterface interface {
	Task() error
	GetTaskID() interface{}
}
