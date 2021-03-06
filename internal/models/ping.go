package models

import "time"

type Ping struct {
	NodeId    string `storm:"id"`
	Timestamp time.Time
}

type PingRepository interface {
	FindByNodeID(nodeId string) (*Ping, error)
	Save(ping *Ping) error
	GetAll() (*[]Ping, error)
}
