package migrations

const (
	CreateKeyspace = `CREATE KEYSPACE %s WITH replication = {
        'class' : 'SimpleStrategy',
        'replication_factor' : 1
    };`

	CreateTable = `CREATE TABLE %s (
    	messageid text,
    	userid text,
    	status text,
    	timestamp timestamp,
    	PRIMARY KEY ((messageid),userid)
	);`
)
