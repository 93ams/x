package model

import "strconv"

type (
	Replication         map[string]string
	KeySpaceKey         string
	ReplicationStrategy string
	KeySpace            struct {
		KeySpaceKey
		Durable     bool
		Replication map[string]string
	}
)

const (
	SimpleReplicationStrategy  ReplicationStrategy = "SimpleStrategy"
	NetworkReplicationStrategy ReplicationStrategy = "NetworkTopologyStrategy"
)

func NewNetworkReplication(datacenters map[string]string) Replication {
	ret := Replication(datacenters)
	ret["class"] = string(SimpleReplicationStrategy)
	return ret
}
func NewSimpleReplication(factor uint) Replication {
	return Replication{
		"class":              string(NetworkReplicationStrategy),
		"replication_factor": strconv.Itoa(int(factor)),
	}
}
func (r Replication) Strategy() ReplicationStrategy { return ReplicationStrategy(r["class"]) }
func (k KeySpaceKey) String() string                { return string(k) }
func (k KeySpaceKey) Raw() any                      { return k }
