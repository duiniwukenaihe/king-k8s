package impl

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/duiniwukenaihe/king-k8s/resource"
	"github.com/duiniwukenaihe/king-utils/common"
	"github.com/duiniwukenaihe/king-utils/common/access"
	"github.com/duiniwukenaihe/king-utils/common/handle"
	"github.com/duiniwukenaihe/king-utils/common/log"
	"net/http"
)

func GetIngress(c *gin.Context) {
	responseData := HandleIngress(common.Get, c)
	c.JSON(responseData.Code, responseData)
}

func ListIngress(c *gin.Context) {
	responseData := HandleIngress(common.List, c)
	c.JSON(responseData.Code, responseData)
}

func DeleteIngress(c *gin.Context) {
	responseData := HandleIngress(common.Delete, c)
	c.JSON(responseData.Code, responseData)
}

func PatchIngress(c *gin.Context) {
	responseData := HandleIngress(common.Patch, c)
	c.JSON(responseData.Code, responseData)
}

func UpdateIngress(c *gin.Context) {
	responseData := HandleIngress(common.Update, c)
	c.JSON(responseData.Code, responseData)
}

func CreateIngress(c *gin.Context) {
	responseData := HandleIngress(common.Create, c)
	c.JSON(responseData.Code, responseData)
}

func GetIngressChart(c *gin.Context) {
	responseData := HandleIngress(common.GetChart, c)
	c.JSON(responseData.Code, responseData)
}

func GetIngressByDeployment(c *gin.Context) {
	responseData := HandleIngress("IngressByDeployment", c)
	c.JSON(responseData.Code, responseData)
}

func HandleIngress(action common.ActionType, c *gin.Context) (responseData *common.ResponseData) {
	// 获取clientSet，如果失败直接返回错误
	clientSet, err := access.Access(c.Query("cluster"))
	responseData = handle.HandlerResponse(nil, err)
	if responseData.Code != http.StatusOK {
		log.Errorf("%s%s", common.K8SClientSetError, err)
		return
	}
	// 获取HTTP的参数，存到handle.Resources结构体中
	commonParams := handle.GenerateCommonParams(c, clientSet)
	r := resource.IngressResource{Params: commonParams}
	// 调用结构体方法
	switch action {
	case common.Get:
		response, err := r.Get()
		responseData = handle.HandlerResponse(response, err)
	case common.List:
		response, err := r.List()
		if err != nil {
			responseData = handle.HandlerResponse(nil, err)
		} else {
			responseData = handle.HandlerResponse(response.Items, err)
		}
	case common.Delete:
		err := r.Delete()
		responseData = handle.HandlerResponse(nil, err)
	case common.Patch:
		if err := c.BindJSON(&r.Params.PatchData); err == nil {
			response, err := r.Patch()
			responseData = handle.HandlerResponse(response, err)
		} else {
			responseData = handle.HandlerResponse(nil, err)
		}
	case common.Update:
		if err := c.BindJSON(&r.PostData); err == nil {
			response, err := r.Update()
			responseData = handle.HandlerResponse(response, err)
		} else {
			responseData = handle.HandlerResponse(nil, err)
		}
	case common.Create:
		if err := r.GenerateCreateData(c); err == nil {
			if r.PostData != nil {
				response, err := r.Create()
				responseData = handle.HandlerResponse(response, err)
			} else {
				responseData = handle.HandlerResponse(nil, errors.New("the post data does not match the type"))
			}
		} else {
			responseData = handle.HandlerResponse(nil, err)
		}
	case common.GetChart:
		response, err := r.GetChart()
		responseData = handle.HandlerResponse(response, err)
	case "IngressByDeployment":
		response, err := r.GetIngressByDeployment()
		responseData = handle.HandlerResponse(response, err)
	}
	return
}
