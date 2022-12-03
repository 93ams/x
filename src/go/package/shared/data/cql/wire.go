package cql

import "github.com/google/wire"

var Set = wire.NewSet(NewCluster, NewSession)
