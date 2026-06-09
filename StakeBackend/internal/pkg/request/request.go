package request

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type StakeRequest struct {
	Amount string `json:"amount" binding:"required,gt=0"`
}

func ValidateStakeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req StakeRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			handleValidationError(c, err)
			c.Abort()
			return
		}
		c.Set("stake_request", req)
		c.Next()
	}
}

func handleValidationError(c *gin.Context, err error) {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			switch e.Tag() {
			case "required":
				c.JSON(400, gin.H{"code": 400, "msg": "参数 " + e.Field() + " 不能为空"})
				return
			case "gt":
				c.JSON(400, gin.H{"code": 400, "msg": "参数 " + e.Field() + " 必须大于 " + e.Tag()})
				return
			case "min":
				c.JSON(400, gin.H{"code": 400, "msg": "参数 " + e.Field() + " 最小值为 " + e.Param()})
				return
			case "max":
				c.JSON(400, gin.H{"code": 400, "msg": "参数 " + e.Field() + " 最大值为 " + e.Param()})
				return
			}
		}
	}
	c.JSON(400, gin.H{"code": 400, "msg": "参数错误: " + err.Error()})
}

func RegisterCustomValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("address", validateAddress)
	}
}

func validateAddress(fl validator.FieldLevel) bool {
	address := fl.Field().String()
	ok, _ := regexp.MatchString("^0x[0-9a-fA-F]{40}$", address)
	return ok
}