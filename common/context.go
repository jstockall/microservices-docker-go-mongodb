package common

import (
	"log"	
	"time"
	"gopkg.in/mgo.v2"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	MongoSession *mgo.Session
}

// Test database connectivity
func (c *Context) Ping() error {
	return c.MongoSession.Ping()
}

// Close mgo.Session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// Returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(AppConfig.Database).C(name)
}

// Create a new Context object for each HTTP request
func NewContext() *Context {
	session := GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}

// Session holds the mongodb session for database access
var session *mgo.Session

// Get database session
func GetSession() *mgo.Session {
	if session == nil {
		createDbSession()
	}
	return session
}

// Create database session
func createDbSession() {
	var err error
	log.Printf("Connecting to MongoDB host [%s]", AppConfig.MongoDBHost)
	session, err = mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    []string{AppConfig.MongoDBHost},
		Username: AppConfig.DBUser,
		Password: AppConfig.DBPwd,
		Timeout:  5 * time.Second,
	})
	if err != nil {
		log.Fatalf("[createDbSession]: %s\n", err)
	} else {
		log.Printf("Connected")
	}
}
