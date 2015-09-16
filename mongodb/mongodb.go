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

func (m *MongoDB) PostDocument(database string, collection string, document interface{}) error {
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

func (m *MongoDB) GetDocuments(database string, collection string, document interface{}, skip int, limit int) (interface{}, error) {
    if m == nil {
        return nil, errors.New("The mongodb session does not exist.")
    }
    
    // database
    d := m.session.DB(database)

    // collection
    c := d.C(collection)

    var result []interface{}
    
    err := c.Find(document).Skip(skip).Limit(limit).Iter().All(&result)

    // error
    if err != nil && err.Error() != "not found" {
        return nil, err
    }

    if result != nil {
        return result, nil
    }

    return nil, nil
}

func (m *MongoDB) DeleteDocuments(database string, collection string, document interface{}) error {
    if m == nil {
        return errors.New("The mongodb session does not exist.")
    }
    
    // database
    d := m.session.DB(database)

    // collection
    c := d.C(collection)
    
    _, err := c.RemoveAll(document)

    // error
    if err != nil && err.Error() != "not found" {
        return err
    }

    return nil
}
