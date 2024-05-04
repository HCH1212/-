package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type response struct {
	status  int
	message string
}

var (
	success = response{
		status:  10000,
		message: "success",
	}
	/*
		receivedSuccess = response{
			status:  10001,
			message: "received",
		}
	*/
	param = response{
		status:  40001,
		message: "参数有误",
	}

	internal = response{
		status:  50001,
		message: "Internal Server Error",
	}
)

func OK(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"staues": success.status,
		"info":   success.message,
	})
}

func OKWithData(c *gin.Context, data any) {
	c.JSON(http.StatusOK, gin.H{
		"status":  10000,
		"message": "success",
		"data":    data,
	})
}

func ParamError(c *gin.Context) {
	c.JSON(http.StatusBadRequest, gin.H{
		"staues": param.status,
		"info":   param.message,
	})
}

func InternalError(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"staues": internal.status,
		"info":   internal.message,
	})
}
