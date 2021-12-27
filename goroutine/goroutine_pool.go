/**
 @author:way
 @date:2021/12/25
 @note
**/

package main
//
//import "fmt"
//
//type Pool struct {
//	work chan func() //工作协程的chan，无缓冲区（同步）
//	sem chan struct{} //控制并发数，带缓冲区
//}
//
//func NewPool(size int)  *Pool{
//	return &Pool{
//		work: make(chan func()),
//		sem: make(chan struct{}, size),
//	}
//}
//
//func (pool *Pool)NewTask(task func())  {
//	select {
//	case pool.work <-task:
//		fmt.Println("pool.work sends success.")
//	case pool.sem  <- struct{}{}:
//		go pool.worker(task)
//	}
//}
//
//func (pool *Pool)worker(task func())  {
//	defer func() {
//		<- pool.sem //理论上是不会走这个流程
//	}()
//	//重复利用开启的goroutine
//	for  {
//		task()
//		//消费者 (如果消费者没准备好，同步的channel就不会发送成功，也就是pool.work <-task 事件不会被触发
//		task = <-pool.work
//	}
//}
//
//func main() {
//	NewPool(5)
//}