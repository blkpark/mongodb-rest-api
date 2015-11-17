package context

import (
    "github.com/blakepark/mongodb-rest-api/mongodb"
)

type Context struct {
    Params map[string]string
    Body []byte
    QueryParams map[string][]string
    MongoDB *mongodb.MongoDB
}

func (c *Context) GetQueryParam(key string, defaultValue string) string {
    values := c.QueryParams[key]
    if values == nil {
        return defaultValue
    }
    return values[0]
}
