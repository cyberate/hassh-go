package proxy

import (
	"bytes"
	"io"
	"log"
	"net"
)

// Handler middleware from client connection
type Handler struct {
	firstBuf  []byte
	encrypted bool
}

// Handle return the socket handled and error or nil
func Handle(conn net.Conn) (handle *Handler, err error) {
	s := &Handler{}
	if _, err = s.extractBuf(conn); err != nil {
		if err != nil {
			log.Fatalln(err)
		}
		panic("No minimal bufer size")
	}
	if ok := s.checkTLS(); ok {
		s.encrypted = ok
	}
	return s, err
}

// IsTLS return if conn is encrypted
func (h *Handler) IsTLS() bool {
	return h.encrypted
}

// func (h *Handler) configure(conn net.Conn) {

// 	// if h.encrypted = h.checkTLS(); h.encrypted {
// 	// 	cer, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
// 	// 	if err != nil {
// 	// 		log.Fatalln(err)
// 	// 		return
// 	// 	}
// 	// 	config := &tls.Config{Certificates: []tls.Certificate{cer}}
// 	// 	h.socketOut = tls.Server(h.socketOut, config)
// 	// }
// }

func (h *Handler) extractBuf(conn net.Conn) (readed int64, err error) {
	firstBuf := make([]byte, 16*1024)
	nr, er := conn.Read(firstBuf)
	if nr > 0 {
		h.firstBuf = firstBuf[0:nr]
		readed += int64(nr)
	}
	if er != nil {
		if er != io.EOF {
			err = er
		}
		return
	}
	return
}

func (h *Handler) checkTLS() bool {
	return bytes.Compare(h.firstBuf[0:3], []byte("\x16\x03\x01")) == 0
}

// func (h *Handler) input() {
// 	_, ew := h.socketIn.Write(h.firstBuf)
// 	if ew != nil {
// 		panic(ew)
// 	}

// 	_, ec := io.Copy(h.socketIn, h.client)
// 	if ec != nil {
// 		if ec != io.EOF {
// 			panic(ec)
// 		}
// 		return
// 	}

// 	return
// }

// func (h *Handler) output() {
// 	_, ec := io.Copy(h.client, h.socketIn)
// 	if ec != nil {
// 		if ec != io.EOF {
// 			panic(ec)
// 		}
// 		return
// 	}
// 	return
// }
