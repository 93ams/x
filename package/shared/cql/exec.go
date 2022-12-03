package cql

import (
	"context"
	"github.com/gocql/gocql"
	"github.com/scylladb/gocqlx/v2"
	"log"
)

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
