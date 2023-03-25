package tcp

import (
	"fmt"
	"github.com/utils-go/ngo/collections/concurrent"
	"net"
	"time"
)

// tcp长连接
type ReconnectTcp struct {
	DstAddr string
	// 当有内容时，进行断线重连
	reconnectChan chan struct{}
	//通知关闭，结束协程，不进行断线重连
	closeChan chan struct{}
	// 将返回的内容放入这里
	RecvBuffer concurrent.ConcurrentListT[byte]
	// tcp连接
	conn net.Conn
}

func (t *ReconnectTcp) SendMsg(data []byte, timeout int) error {
	if t.conn == nil {
		t.reconnectChan <- struct{}{}
		time.Sleep(time.Millisecond * 100)
		return fmt.Errorf("%s not connect yet", t.DstAddr)
	}

	_, err := t.conn.Write(data)
	if err != nil {
		//出错，重连
		t.reconnectChan <- struct{}{}
		time.Sleep(time.Millisecond * 100)
		return err
	}
	return nil
}

// new reconnect tcp connection
func NewRTcpConnection(addr string) *ReconnectTcp {
	t := &ReconnectTcp{
		DstAddr:       addr,
		reconnectChan: make(chan struct{}, 100),
		closeChan:     make(chan struct{}, 100),
		RecvBuffer:    concurrent.ConcurrentListT[byte]{},
	}
	go t.handleReconnect()
	go t.handleRead()
	return t
}
func (t *ReconnectTcp) connect() {
	if t.conn != nil {
		t.conn.Close()
	}
	con, err := net.Dial("tcp", t.DstAddr)
	if err != nil {
		return
	}
	t.conn = con
	fmt.Printf("connect to %s success\n", con.RemoteAddr())
}

// 处理重连
func (t *ReconnectTcp) handleReconnect() {
	for {
		select {
		case <-t.reconnectChan:
			t.connect()
			//重连之后，清除掉t.reconnectChan中的内容
			t.reconnectChan = make(chan struct{}, 100)
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", t.DstAddr)
			return
		}
	}
}

// 连接
func (t *ReconnectTcp) handleRead() {
	//wait for connect
	time.Sleep(time.Second)
	buf := make([]byte, 1024*10)
	for {
		//处理关闭
		select {
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit handleRead goroutine\n", t.DstAddr)
			return
		default:
		}
		if t.conn == nil {
			t.reconnectChan <- struct{}{}
			time.Sleep(time.Millisecond * 100)
			continue
		}
		n, err := t.conn.Read(buf)
		if err == nil {
			t.RecvBuffer.AddRange(buf[:n])
		} else {
			t.reconnectChan <- struct{}{}
		}
	}
}

// 关闭
func (t *ReconnectTcp) Close() {
	t.closeChan <- struct{}{} //notify reconnect goroutine close
	t.closeChan <- struct{}{} //notify handleRead goroutine close
}
