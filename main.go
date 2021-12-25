package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"time"
)

type websocketClientManager struct {
	conn        *websocket.Conn
	addr        *string
	path        string
	sendMsgChan chan string
	recvMsgChan chan string
	isAlive     bool
	timeout     int
}

// 构造函数
func NewWsClientManager(addrIp, addrPort, path string, timeout int) *websocketClientManager {
	addrString := addrIp + ":" + addrPort
	var sendChan = make(chan string, 10)
	var recvChan = make(chan string, 10)
	var conn *websocket.Conn
	return &websocketClientManager{
		addr:        &addrString,
		path:        path,
		conn:        conn,
		sendMsgChan: sendChan,
		recvMsgChan: recvChan,
		isAlive:     false,
		timeout:     timeout,
	}
}

// 链接服务端
func (wsc *websocketClientManager) dail() {
	var err error
	u := url.URL{Scheme: "ws", Host: *wsc.addr, Path: wsc.path}
	log.Printf("connecting to %s", u.String())
	wsc.conn, _, err = websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return

	}
	wsc.isAlive = true
	log.Printf("connecting to %s 链接成功！！！", u.String())

}

// 发送消息
func (wsc *websocketClientManager) sendMsgThread() {
	go func() {
		for {
			msg := <-wsc.sendMsgChan
			err := wsc.conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("write:", err)
				continue
			}
		}
	}()
}

// 读取消息
func (wsc *websocketClientManager) readMsgThread() {
	go func() {
		for {
			if wsc.conn != nil {
				_, message, err := wsc.conn.ReadMessage()
				if err != nil {
					log.Println("read:", err)
					wsc.isAlive = false
					// 出现错误，退出读取，尝试重连
					break
				}
				log.Printf("recv: %s", message)
				// 需要读取数据，不然会阻塞
				wsc.recvMsgChan <- string(message)
			}

		}
	}()
}

// 开启服务并重连
func (wsc *websocketClientManager) start() {
	for {
		if wsc.isAlive == false {
			wsc.dail()
			wsc.sendMsgThread()
			wsc.readMsgThread()
		}
		time.Sleep(time.Second * time.Duration(wsc.timeout))
	}
}

func main() {
	for i:=0;i<1000;i++{

		go func() {
			wsc := NewWsClientManager("127.0.0.1", "8887", "/ws", 10)
			wsc.sendMsgChan <- "\"market.btcusdt.kline.5min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.etcusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.xrpusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.adausdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.ltcusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.xemusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.dashusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.xlmusdt.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.ethbtc.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.xrpbtc.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.eosbtc.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.dotbtc.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.eoseth.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.linketh.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.adaeth.kline.1min\""
			//time.Sleep(time.Second*1)
			//wsc.sendMsgChan <- "\"market.xmreth.kline.1min\""
			//time.Sleep(time.Second*1)
			msgChan := wsc.recvMsgChan
			wsc.start()
			println(msgChan)

		}()
		fmt.Println(i)
		if i == 999{
			time.Sleep(60*time.Second)
		}


	}






}