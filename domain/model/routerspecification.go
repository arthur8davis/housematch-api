package model

import (
	"database/sql"
	"github.com/gin-gonic/gin"
)

type RouterSpecification struct {
	Api *gin.Engine
	DB  *sql.DB
}
