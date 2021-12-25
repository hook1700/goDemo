/**
 @author:way
 @date:2021/12/25
 @note
**/

package main

type Glimit struct {
	n int
	c chan struct{}
}


//New  initialization Glimit struct
func New(n int) *Glimit {
	return &Glimit{
		n: n,
		c: make(chan struct{}, n),
	}
}


// Run f in a new goroutine but with limit.
func (g *Glimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}