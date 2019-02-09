CREATE KEYSPACE tracking WITH replication = {
  'class':'SimpleStrategy',
  'replication_factor':1
};

USE tracking;

CREATE TABLE tracking (
  messageid text,
  userid text,
  status text,
  timestamp timestamp,
  PRIMARY KEY ((messageid), userid)
);