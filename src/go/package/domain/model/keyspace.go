package model

import "strings"

type (
	Replication         map[string]any
	KeySpaceKey         string
	ReplicationStrategy string
	KeySpace            struct {
		KeySpaceKey
		Durable     bool
		Replication Replication
		Tables      []Table
	}
)

const (
	SimpleReplicationStrategy  ReplicationStrategy = "SimpleStrategy"
	NetworkReplicationStrategy ReplicationStrategy = "NetworkTopologyStrategy"
)

func NewNetworkReplication(datacenters map[string]any) Replication {
	ret := Replication(datacenters)
	ret["class"] = string(NetworkReplicationStrategy)
	return ret
}
func NewSimpleReplication(factor int) Replication {
	return Replication{
		"class":              string(SimpleReplicationStrategy),
		"replication_factor": factor,
	}
}
func (r Replication) Strategy() ReplicationStrategy { return ReplicationStrategy(r["class"].(string)) }
func (k KeySpaceKey) String() string                { return string(k) }
func (k KeySpaceKey) Raw() any                      { return k }

func FilterTest(in []KeySpace) (out []KeySpace) {
	for _, ks := range in {
		if strings.HasPrefix(ks.KeySpaceKey.String(), "test") {
			out = append(out, ks)
		}
	}
	return
}
