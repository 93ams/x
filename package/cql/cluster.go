package cql

import (
	"context"
	"github.com/gocql/gocql"
	"log"
	"time"
)

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

func SafeExec(ctx context.Context, s *gocql.Session, stmt string) error {
	if err := s.AwaitSchemaAgreement(ctx); err != nil {
		log.Printf("error waiting for schema agreement pre running stmt=%q err=%v\n", stmt, err)
		return err
	}
	if err := s.Query(stmt).RetryPolicy(&gocql.SimpleRetryPolicy{}).Exec(); err != nil {
		log.Printf("error running stmt stmt=%q err=%v\n", stmt, err)
		return err
	}
	if err := s.AwaitSchemaAgreement(ctx); err != nil {
		log.Printf("error waiting for schema agreement running stmt=%q err=%v\n", stmt, err)
		return err
	}
	return nil
}
