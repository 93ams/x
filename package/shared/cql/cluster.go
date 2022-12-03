package cql

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
	"time"
)

func NewSession(cluster *gocql.ClusterConfig) (gocqlx.Session, error, func()) {
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return gocqlx.Session{}, err, nil
	}
	return session, nil, func() { session.Close() }
}

func NewCluster(consistency gocql.Consistency, keyspace string, hosts ...string) *gocql.ClusterConfig {
	retryPolicy := &gocql.ExponentialBackoffRetryPolicy{
		Max:        10 * time.Second,
		Min:        time.Second,
		NumRetries: 5,
	}
	cluster := gocql.NewCluster(hosts...)
	cluster.Keyspace = keyspace
	cluster.Timeout = 5 * time.Second
	cluster.RetryPolicy = retryPolicy
	cluster.Consistency = consistency
	cluster.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(gocql.RoundRobinHostPolicy())
	return cluster
}

func SafeExec(ctx context.Context, s gocqlx.Session, stmt string, values ...any) error {
	if err := s.AwaitSchemaAgreement(ctx); err != nil {
		log.Printf("error waiting for schema agreement pre running stmt=%q err=%v\n", stmt, err)
		return err
	}
	if err := s.Session.Query(stmt, values...).RetryPolicy(&gocql.SimpleRetryPolicy{}).Exec(); err != nil {
		log.Printf("error running stmt stmt=%q err=%v\n", stmt, err)
		return err
	}
	if err := s.AwaitSchemaAgreement(ctx); err != nil {
		log.Printf("error waiting for schema agreement running stmt=%q err=%v\n", stmt, err)
		return err
	}
	return nil
}
