package cassandra

import (
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster("172.26.0.2")
	cluster.Keyspace = "oauthlocal"
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()

	if err != nil {
		panic(err)
	}

}

func GetSession() *gocql.Session {
	return session
}
