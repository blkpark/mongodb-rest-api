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

    // param
    var param interface{}
    err := json.Unmarshal(c.Body, &param)
    if err != nil {
        status = http.StatusBadRequest
        res = map[string]interface{}{
            "error": err.Error(),
        }
        return status, res
    }

    err = c.MongoDB.PostDocument(db, col, param)
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

    var param = make(map[string]string)
    for k := range c.QueryParams {
        param[k] = c.GetQueryParam(k, "")
    }

    result, err := c.MongoDB.GetDocuments(db, col, param, int(s), int(l))
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

    var param = make(map[string]string)
    for k := range c.QueryParams {
        param[k] = c.GetQueryParam(k, "")
    }

    err := c.MongoDB.DeleteDocuments(db, col, param)
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
