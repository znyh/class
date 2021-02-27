package config

import (
	"time"
)

//type Config struct {
//	Server *Server
//	TLS    *TLS
//}
//
//type Server struct {
//	Addr    string
//	Timeout string
//}
//
//type TLS struct {
//	Pem string
//	Key string
//}

type Comet struct {
	Sndbuf       int
	Rcvbuf       int
	KeepAlive    bool
	Reader       int
	ReadBuf      int
	ReadBufSize  int
	Writer       int
	WriteBuf     int
	WriteBufSize int
}

// TCP is tcp config.
type TCP struct {
	Bind []string
}

// Websocket is websocket config.
type Websocket struct {
	Bind        []string
	TLSOpen     bool
	TLSBind     []string
	CertFile    string
	PrivateFile string
}

// Protocol is proto config.
type Protocol struct {
	Proxy            bool
	Timer            int
	TimerSize        int
	SvrProto         int
	CliProto         int
	HandshakeTimeout time.Duration
	HeartbeatTimeout time.Duration
}

type Auth struct {
	Open  bool
	AppID string
}

// Bucket is bucket config.
type Bucket struct {
	Size    int
	Channel int
}

//ChanSize
type ChanSize struct {
	Push       int
	Close      int
	Disconnect int
}

// Config is comet config.
type ServerConfig struct {
	Comet     *Comet
	TCP       *TCP
	Websocket *Websocket
	Protocol  *Protocol
	Auth      *Auth
	Bucket    *Bucket
	ChanSize  *ChanSize
}
