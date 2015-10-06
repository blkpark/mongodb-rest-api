package route

import (
    "net/http"
    "encoding/json"
    
    "github.com/ghmlee/mongodb-rest-api/context"
)

func PostDocument(
    c context.Context,
    w http.ResponseWriter,
    r *http.Request) (int, map[string]interface{}) {
    var status int
    var res map[string]interface{}

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

    err = c.MongoDB.NewDocument(db, col, param)
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
