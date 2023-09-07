package tcp

import (
	"fmt"
	"github.com/utils-go/ngo/collections/concurrent"
	"net"
	"sync"
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
	RecvBuffer *concurrent.ConcurrentListT[byte]
	// tcp连接
	conn net.Conn
	mu   sync.Mutex
}

func (t *ReconnectTcp) SendMsg(data []byte) error {
	if t.conn == nil {
		t.reconnect()
		return fmt.Errorf("%s not connect yet", t.DstAddr)
	}

	t.mu.Lock()
	_, err := t.conn.Write(data)
	t.mu.Unlock()

	if err != nil {
		//出错，重连
		t.reconnect()
		time.Sleep(time.Millisecond * 100)
		return err
	}
	return nil
}

// new reconnect tcp connection
func NewRTcpConnection(addr string) (*ReconnectTcp, error) {
	if _, err := net.ResolveTCPAddr("tcp", addr); err != nil {
		return nil, err
	}
	t := &ReconnectTcp{
		DstAddr:       addr,
		reconnectChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		RecvBuffer:    concurrent.NewListT[byte](),
	}
	go t.handleReconnect()
	go t.handleRead()
	return t, nil
}
func (t *ReconnectTcp) connect() {
	if t.conn != nil {
		t.conn.Close()
	}
	con, err := net.DialTimeout("tcp", t.DstAddr, time.Second*500)
	if err != nil {
		t.conn = nil
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
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", t.DstAddr)
			return
		}
	}
}

// add signal without block
func (t *ReconnectTcp) reconnect() {
	select {
	case t.reconnectChan <- struct{}{}:
	default:
	}
}

// 连接
func (t *ReconnectTcp) handleRead() {
	buf := make([]byte, 1024)
	for {
		//处理关闭
		select {
		case <-t.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit handleRead goroutine\n", t.DstAddr)
			return
		default:
		}
		if t.conn == nil {
			t.reconnect()
			time.Sleep(time.Millisecond * 200)
			continue
		}
		n, err := t.conn.Read(buf)
		if err == nil {
			t.RecvBuffer.AddRange(buf[:n])
		} else {
			t.reconnect()
			time.Sleep(time.Millisecond * 200)
		}
	}
}

// 关闭
func (t *ReconnectTcp) Close() {
	close(t.closeChan)
}
