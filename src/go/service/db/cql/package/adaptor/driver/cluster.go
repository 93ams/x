package driver

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"time"
)

type (
	Keyspace string
	Hosts    []string
	Options  struct {
		Consistency gocql.Consistency
		Keyspace    Keyspace
		Hosts       Hosts
	}
)

func NewCluster(opts Options) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Max:        10 * time.Second,
		Min:        time.Second,
		NumRetries: 5,
	}
	if len(opts.Hosts) == 0 {
		opts.Hosts = []string{"localhost:9042"}
	}
	cluster := gocql.NewCluster(opts.Hosts...)
	if opts.Keyspace == "" {
		opts.Keyspace = "system_schema"
	}
	cluster.Keyspace = string(opts.Keyspace)
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = opts.Consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

func NewSession(cluster *gocql.ClusterConfig) (gocqlx.Session, func(), error) {
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return gocqlx.Session{}, nil, err
	}
	return session, func() { session.Close() }, nil
}
