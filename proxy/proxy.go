package proxy

import "net"

// Proxy is struct contains conn client,
// and client pipe full-duplex.
type Proxy struct {
	conn      net.Conn
	clientIn  net.Conn
	clientOut net.Conn
	handler   *Handler
}

// type CertPair struct {
// 	certPath string
// 	keyPath  string
// }

// Config is struct for config values for proxy
type Config struct {
	// sizeBufSend    int64
	// sizeBufRecv    int64
	// sizeBufInitial int64
	// certPair       CertPair
}

// NewProxy return a struct proxy created.
func NewProxy(conn net.Conn) (*Proxy, error) {
	return newProxyConfig(conn, Config{
		// 	sizeBufSend: int64(32*1024),
		// 	sizeBufRecv: int64(16*1024),
		// 	sizeBufInitial: int64(16*1024),
		// 	certPair: CertPair{
		// 		certPath: "cert.pem",
		// 		keyPath: "key.pem",
		// 	}
	})
}

// NewProxyConfig return a struct proxy created by config.
func NewProxyConfig(conn net.Conn, config Config) (*Proxy, error) {
	return newProxyConfig(conn, config)
}

func newProxyConfig(conn net.Conn, config Config) (*Proxy, error) {
	p := &Proxy{
		conn: conn,
	}
	if e := p.configure(config); e != nil {
		panic(e)
	}
	return p, nil
}

func (p *Proxy) configure(config Config) error {

	return nil
}
