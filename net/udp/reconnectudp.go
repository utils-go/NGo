package udp

import (
	"fmt"
	"github.com/utils-go/ngo/collections/concurrent"
	"net"
	"sync"
	"time"
)

/*
长连接：将连接存储起来，不主动关闭
*/

type ReconnectUdp struct {
	DstAddr *net.UDPAddr
	//当连接或者写入报错时，则重新初始化conn
	reconnectChan chan struct{}
	//关闭通知
	closeChan chan struct{}
	//接收到的字节
	RecvBuffer *concurrent.ConcurrentListT[[]byte]
	//tcp连接
	conn *net.UDPConn
	mu   sync.Mutex
}

func (u *ReconnectUdp) SendMsg(data []byte) error {
	if u.conn == nil {
		u.reconnect()
		time.Sleep(time.Millisecond * 100)
		return fmt.Errorf("%s not connect yet", u.DstAddr)
	}

	u.mu.Lock()
	_, err := u.conn.WriteToUDP(data, u.DstAddr)
	u.mu.Unlock()

	if err != nil {
		u.reconnect()
		time.Sleep(time.Millisecond * 100)
		return err
	}
	return nil
}
func NewRUdpConnection(serverAddr string) (*ReconnectUdp, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return nil, err
	}
	u := &ReconnectUdp{
		DstAddr:       udpAddr,
		reconnectChan: make(chan struct{}),
		closeChan:     make(chan struct{}),
		RecvBuffer:    concurrent.NewListT[[]byte](),
	}
	go u.handleReconnect()
	go u.handRead()
	return u, nil
}
func (u *ReconnectUdp) connect() {
	if u.conn != nil {
		u.conn.Close()
	}

	con, err := net.ListenUDP("udp", nil)
	if err != nil {
		return
	}

	fmt.Printf("listen to [%s] success\n", con.LocalAddr())
	u.conn = con
}
func (u *ReconnectUdp) handleReconnect() {
	for {
		select {
		case <-u.reconnectChan:
			u.connect()
		case <-u.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", u.DstAddr)
			return
		}
	}
}

// add signal without block
func (t *ReconnectUdp) reconnect() {
	select {
	case t.reconnectChan <- struct{}{}:
	default:
	}
}

// 处理读取数据
func (u *ReconnectUdp) handRead() {
	buffer := make([]byte, 1024)
	for {
		select {
		case <-u.closeChan:
			fmt.Printf("remote connection:%s has been closed,exit reconnect goroutine\n", u.DstAddr)
			return
		default:
		}
		if u.conn == nil {
			u.reconnect()
			time.Sleep(time.Millisecond * 200)
			continue
		}

		n, addr, err := u.conn.ReadFromUDP(buffer)
		if err == nil {
			if addr.String() == u.DstAddr.String() {
				u.RecvBuffer.Add(buffer[:n])
			} else {
				fmt.Printf("recv from [%s] but listen on [%s]\n", addr.String(), u.DstAddr.String())
			}
		} else {
			u.reconnect()
			time.Sleep(time.Millisecond * 200)
		}
	}
}
func (u *ReconnectUdp) Close() {
	close(u.closeChan)
}
