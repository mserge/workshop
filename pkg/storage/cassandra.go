package storage

import "github.com/gocql/gocql"

func InitStorage(hostport, keyspace string) (*gocql.Session, error) {
	clusterConfig := gocql.NewCluster(hostport)
	clusterConfig.Keyspace = keyspace

	session, err := clusterConfig.CreateSession()
	if err != nil {
		return nil, err
	}

	return session, nil
}
