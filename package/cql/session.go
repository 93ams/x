package cql

import (
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
)

func NewSession(cluster *gocql.ClusterConfig) (gocqlx.Session, error) {
	session, err := gocqlx.WrapSession(cluster.CreateSession())
	if err != nil {
		return gocqlx.Session{}, err
	}
	return session, nil
}
