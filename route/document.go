package route

import (
    "strconv"
    "net/http"
    "encoding/json"
    
    "github.com/ghmlee/mongodb-rest-api/context"
)

func PostDocument(
    c context.Context,
    w http.ResponseWriter,
    r *http.Request) (int, interface{}) {
    var status int
    var res interface{}

    // collection
    db := c.Params["database"]
    col := c.Params["collection"]

    // doc
    var doc interface{}
    err := json.Unmarshal(c.Body, &doc)
    if err != nil {
        status = http.StatusBadRequest
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    err = c.MongoDB.PostDocument(db, col, doc)
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    status = http.StatusOK
    
    return status, res
}

func PutDocument(
    c context.Context,
    w http.ResponseWriter,
    r *http.Request) (int, interface{}) {
    var status int
    var res interface{}

    // collection
    db := c.Params["database"]
    col := c.Params["collection"]

    // doc
    var doc interface{}
    err := json.Unmarshal(c.Body, &doc)
    if err != nil {
        status = http.StatusBadRequest
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    var query = make(map[string]string)
    for k := range c.QueryParams {
        query[k] = c.GetQueryParam(k, "")
    }

    err = c.MongoDB.PutDocument(db, col, query, doc)
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    status = http.StatusOK
    
    return status, res
}

func GetDocuments(
    c context.Context,
    w http.ResponseWriter,
    r *http.Request) (int, interface{}) {
    var status int
    var res interface{}

    // collection
    db := c.Params["database"]
    col := c.Params["collection"]

    skip := c.GetQueryParam("skip", "0")
    limit := c.GetQueryParam("limit", "1")
    sort := c.GetQueryParam("sort", "-_id")

    // skip
    s, err := strconv.ParseUint(skip, 10, 64)
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    // limit
    l, err := strconv.ParseUint(limit, 10, 64)
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    delete(c.QueryParams, "limit")
    delete(c.QueryParams, "skip")
    delete(c.QueryParams, "sort")

    var query = make(map[string]interface{})
    for k := range c.QueryParams {
        query[k] = c.GetQueryParam(k, "")

        v := c.GetQueryParam(k, "")
        i, err := strconv.ParseUint(v, 10, 64)
        if err == nil {
            query[k] = i
        } else {
            query[k] = v
        }
    }

    result, err := c.MongoDB.GetDocuments(db, col, query, sort, int(s), int(l))
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    status = http.StatusOK

    return status, result
}

func DeleteDocuments(
    c context.Context,
    w http.ResponseWriter,
    r *http.Request) (int, interface{}) {
    var status int
    var res interface{}

    // collection
    db := c.Params["database"]
    col := c.Params["collection"]

    var query = make(map[string]string)
    for k := range c.QueryParams {
        query[k] = c.GetQueryParam(k, "")
    }

    err := c.MongoDB.DeleteDocuments(db, col, query)
    if err != nil {
        status = http.StatusInternalServerError
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    status = http.StatusOK

    return status, res
}
