package mongodb

import (
    "errors"

    "gopkg.in/mgo.v2"
)

type MongoDB struct {
    session *mgo.Session
}

func NewMongoDB(host string) (*MongoDB, error) {
    s, err := mgo.Dial(host)
    if err != nil {
        return nil, err
    }

    m := MongoDB{
        session: s,
    }

    return &m, nil
}

func (m *MongoDB) NewDocument(database string, collection string, document interface{}) error {
    if m == nil {
        return errors.New("The mongodb session does not exist.")
    }
    
    // database
    d := m.session.DB(database)

    // collection
    c := d.C(collection)

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
