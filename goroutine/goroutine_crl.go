/**
 @author:way
 @date:2021/12/25
 @note
**/

package main

//type Glimit struct {
//	n int
//	c chan struct{}
//}
//
//
////New  initialization Glimit struct
//func New(n int) *Glimit {
//	return &Glimit{
//		n: n,
//		c: make(chan struct{}, n),
//	}
//}
//
//
//// Run f in a new goroutine but with limit.
//func (g *Glimit) Run(f func()) {
//	g.c <- struct{}{}
//	go func() {
//		f()
//		<-g.c
//	}()
//}
//
//var wg = sync.WaitGroup{}
//func main() {
//	number := 10
//	g := New(2)
//	for i := 0; i < number; i++ {
//		wg.Add(1)
//		value := i
//		goFunc := func() {
//			// 做一些业务逻辑处理
//			fmt.Printf("go func: %d\n", value)
//			time.Sleep(time.Second)
//			wg.Done()
//		}
//		g.Run(goFunc)
//	}
//	wg.Wait()
//}