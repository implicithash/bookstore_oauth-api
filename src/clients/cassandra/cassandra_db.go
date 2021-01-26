package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	cluster *gocql.ClusterConfig
)

func init() {
	// connect to Cassandra cluster
	cluster := gocql.NewCluster("192.168.99.100")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum
}

func GetSession() (*gocql.Session, error) {
	return cluster.CreateSession()
}
