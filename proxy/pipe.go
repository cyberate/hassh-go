package proxy

import (
	"net"
	"time"
)

type pipe struct {
	conn       net.Conn
	localAddr  net.Addr
	remoteAddr net.Addr
}

// Pipe return struct extended by net.pipe
func Pipe(lAddr, rAddr net.Addr) (net.Conn, net.Conn) {
	c1, c2 := net.Pipe()
	p1 := &pipe{
		conn:       c1,
		localAddr:  lAddr,
		remoteAddr: rAddr,
	}
	p2 := &pipe{
		conn:       c2,
		localAddr:  lAddr,
		remoteAddr: rAddr,
	}

	return p1, p2
}

func (p *pipe) LocalAddr() net.Addr {
	return p.localAddr
}
func (p *pipe) RemoteAddr() net.Addr {
	return p.remoteAddr
}

func (p *pipe) Read(b []byte) (int, error) {
	return p.conn.Read(b)
}

func (p *pipe) Write(b []byte) (int, error) {
	return p.conn.Write(b)
}

func (p *pipe) SetDeadline(t time.Time) error {
	return p.conn.SetDeadline(t)
}

func (p *pipe) SetReadDeadline(t time.Time) error {
	return p.conn.SetReadDeadline(t)
}

func (p *pipe) SetWriteDeadline(t time.Time) error {
	return p.conn.SetWriteDeadline(t)
}

func (p *pipe) Close() error {
	return p.conn.Close()
}
