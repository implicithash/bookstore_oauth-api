package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	//cluster *gocql.ClusterConfig
	session *gocql.Session
)

func init() {
	// connect to Cassandra cluster
	cluster := gocql.NewCluster("192.168.99.100")
	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
