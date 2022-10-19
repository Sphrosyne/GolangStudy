package pool

var pool *WorkerPool

func GetPool() *WorkerPool {
	return pool
}

func init() {
	/*var (
		MaxWorker = os.Getenv("MAX_WORKERS")
		MaxQueue  = os.Getenv("MAX_QUEUE")
	)*/
	pool = NewWorkerPool(3)
	pool.Run()
}
