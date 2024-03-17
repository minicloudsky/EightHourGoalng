package main

import (
	"fmt"
	"sync"
)

// Job 表示要执行的任务
type Job struct {
	ID int
}

// Worker 表示工作池中的工作者
type Worker struct {
	ID         int
	JobCh      chan Job
	WorkerPool chan chan Job
	Quit       chan bool
}

// NewWorker 创建一个新的工作者
func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		ID:         id,
		JobCh:      make(chan Job),
		WorkerPool: workerPool,
		Quit:       make(chan bool),
	}
}

// Start 启动工作者
func (w *Worker) Start() {
	go func() {
		for {
			// 注册当前工作者到工作池
			w.WorkerPool <- w.JobCh

			select {
			case job := <-w.JobCh:
				// 执行任务
				fmt.Printf("Worker %d 执行任务 %d\n", w.ID, job.ID)

			case <-w.Quit:
				// 收到退出信号
				return
			}
		}
	}()
}

// Stop 停止工作者
func (w *Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

// Pool 表示工作池
type Pool struct {
	WorkerNum  int
	JobCh      chan Job
	WorkerPool chan chan Job
	wg         sync.WaitGroup
}

// NewPool 创建一个新的工作池
func NewPool(workerNum int) *Pool {
	return &Pool{
		WorkerNum:  workerNum,
		JobCh:      make(chan Job),
		WorkerPool: make(chan chan Job, workerNum),
	}
}

// Start 启动工作池
func (p *Pool) Start() {
	// 创建工作者并启动它们
	for i := 0; i < p.WorkerNum; i++ {
		worker := NewWorker(i+1, p.WorkerPool)
		worker.Start()
	}

	// 分配任务给工作者
	go func() {
		for {
			select {
			case job := <-p.JobCh:
				// 从任务通道中接收任务并分配给空闲的工作者
				go func(job Job) {
					jobCh := <-p.WorkerPool
					jobCh <- job
				}(job)
			}
		}
	}()
}

// Wait 等待所有任务完成
func (p *Pool) Wait() {
	p.wg.Wait()
}

func main() {
	// 创建工作池并启动
	pool := NewPool(3)
	pool.Start()

	// 添加任务到工作池
	for i := 1; i <= 100; i++ {
		job := Job{ID: i}
		pool.JobCh <- job
	}

	// 关闭任务通道并等待所有任务完成
	close(pool.JobCh)
	pool.Wait()

	// 停止工作池中的工作者
	for i := 0; i < pool.WorkerNum; i++ {
		worker := <-pool.WorkerPool
		worker <- Job{} // 发送空任务以触发工作者退出
	}

	fmt.Println("工作池已关闭")
}
