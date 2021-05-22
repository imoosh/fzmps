package controllers

import (
    "centnet-fzmps/dao"
    "github.com/gin-gonic/gin"
)

func DB(c *gin.Context) *dao.Dao {
    d, ok := c.Get("db")
    if !ok {
        panic("not exists db")
    }
    return d.(*dao.Dao)
}
