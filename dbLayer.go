package main

import (
  "errors"
  "fmt"
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
)

const CONNECTIONSTRING = "mongodb://127.0.0.1"

type mongoDocument struct {
  Id bson.ObjectId `bson:"_id"`
  ShortUrl string `bson:"shorturl"`
  LongUrl string `bson:"longurl"`
}

type MongoConnection struct {
  originalSession *mgo.Session
}

func NewDBConnection() (conn *MongoConnection) {
  conn = new(MongoConnection)
  conn.createLocalConnection()
  return
}

func (c *MongoConnection) createLocalConnection() (err error) {
  fmt.Println("Connecting to local mongo database...")
  c.originalSession, err = mgo.Dial(CONNECTIONSTRING)
  if err == nil {
    fmt.Println("Connection established to mongo server")
    urlcollection := c.originalSession.DB("LinkShortenerDB").C("UrlCollection")
    if urlcollection == nil {
      err = errors.New("Collection could not be created, might to manually instantiate/connect MongoDB")
    }
    index := mgo.Index {
      Key: []string{"$text:shorturl"},
      Unique: true,
      DropDups: true,
    }
    urlcollection.EnsureIndex(index)
  } else {
    fmt.Printf("Error occured while creating mongodb connection: %s", err.Error())
  }
  return
}

func (c *MongoConnection) getSessionAndCollection() (session *mgo.Session, urlCollection *mgo.Collection, err error) {
  if c.originalSession != nil {
    session = c.originalSession.Copy()
    urlCollection = session.DB("LinkShortenerDB").C("UrlCollection")
  } else {
    err = errors.New("No original session found")
  }
  return
}

func (c *MongoConnection) FindshortUrl(longurl string) (sUrl string, err error) {
  // Create empty document struct
  result := mongoDocument{}
  // Get copy of the original session and a collection
  session, urlCollection, err := c.getSessionAndCollection()
  if err != nil {
    return
  }
  defer session.Close()
  err = urlCollection.Find(bson.M{"longurl": longurl}).One(&result)
  if err != nil {
    return
  }
  return result.ShortUrl, nil
}

func (c *MongoConnection) FindlongUrl(shortUrl string) (sUrl string, err error) {
  // Create empty document struct
  result := mongoDocument{}
  // Get copy of the original session and a collection
  session, urlCollection, err := c.getSessionAndCollection()
  if err == nil {
    return
  }
  defer session.Close()
  err = urlCollection.Find(bson.M{"shorturl": shortUrl}).One(&result)
  if err != nil {
    return
  }
  return result.LongUrl, nil
}

func (c *MongoConnection) AddUrls(longUrl string, shortUrl string) (err error) {
  // Get copy of current session
  session, urlCollection, err := c.getSessionAndCollection()
  if err == nil {
    defer session.Close()
    // Insert a document with the provided function arguments
    err = urlCollection.Insert(
      &mongoDocument {
        Id: bson.NewObjectId(),
        ShortUrl: shortUrl,
        LongUrl: longUrl,
      },
    )
    if err != nil {
      // Check: see if there is an error due to duplicate shorturl
      if mgo.isDup(err) {
        err = errors.New("That shorturl already exists")
      }
    }
  }
  return
}
