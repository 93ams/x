// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Replication interface {
	IsReplication()
	GetClass() *ReplicationStrategy
}

type Datacenter struct {
	Name              string `json:"name"`
	ReplicationFactor int    `json:"replication_factor"`
}

type KeySpace struct {
	Name        string      `json:"name"`
	Tables      []*Table    `json:"tables"`
	Durable     *bool       `json:"durable"`
	Replication Replication `json:"replication"`
}

type NetworkTopologyReplication struct {
	Class       *ReplicationStrategy `json:"class"`
	Datacenters []*Datacenter        `json:"datacenters"`
}

func (NetworkTopologyReplication) IsReplication()                      {}
func (this NetworkTopologyReplication) GetClass() *ReplicationStrategy { return this.Class }

type NewKeyspace struct {
	Name        string          `json:"name"`
	Replication *NewReplication `json:"replication"`
}

type NewReplication struct {
	Strategy *ReplicationStrategy `json:"strategy"`
}

type NewTable struct {
	Name   string    `json:"name"`
	Params []*string `json:"params"`
}

type SimpleReplication struct {
	Class             *ReplicationStrategy `json:"class"`
	ReplicationFactor int                  `json:"replication_factor"`
}

func (SimpleReplication) IsReplication()                      {}
func (this SimpleReplication) GetClass() *ReplicationStrategy { return this.Class }

type Table struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ReplicationStrategy string

const (
	ReplicationStrategySimpleStrategy          ReplicationStrategy = "SimpleStrategy"
	ReplicationStrategyNetworkTopologyStrategy ReplicationStrategy = "NetworkTopologyStrategy"
)

var AllReplicationStrategy = []ReplicationStrategy{
	ReplicationStrategySimpleStrategy,
	ReplicationStrategyNetworkTopologyStrategy,
}

func (e ReplicationStrategy) IsValid() bool {
	switch e {
	case ReplicationStrategySimpleStrategy, ReplicationStrategyNetworkTopologyStrategy:
		return true
	}
	return false
}

func (e ReplicationStrategy) String() string {
	return string(e)
}

func (e *ReplicationStrategy) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ReplicationStrategy(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ReplicationStrategy", str)
	}
	return nil
}

func (e ReplicationStrategy) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}