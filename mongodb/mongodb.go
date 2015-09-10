package mongodb

import (
    "errors"

    "gopkg.in/mgo.v2"
)

type MongoDB struct {
    session *mgo.Session
    database *mgo.Database
}

func NewMongoDB(host string, database string) (*MongoDB, error) {
    s, err := mgo.Dial(host)
    if err != nil {
        return nil, err
    }

    d := s.DB(database)

    m := MongoDB{
        session: s,
        database: d,
    }

    return &m, nil
}

func (m *MongoDB) NewDocument(collection string, document interface{}) error {
    if m == nil || m.database == nil {
        return errors.New("The database does not exist.")
    }

    // collection
    c := m.database.C(collection)

    var result interface{}
    err := c.Find(document).One(&result)

    // error
    if err != nil && err.Error() != "not found" {
        return err
    }

    if result != nil {
        return nil
    }

    err = c.Insert(document)
    if err != nil {
        return err
    }

    return nil
}
