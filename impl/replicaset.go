package impl

import (
	"github.com/gin-gonic/gin"
	"github.com/duiniwukenaihe/king-k8s/resource"
	"github.com/duiniwukenaihe/king-utils/common"
	"github.com/duiniwukenaihe/king-utils/common/access"
	"github.com/duiniwukenaihe/king-utils/common/handle"
	"github.com/duiniwukenaihe/king-utils/common/log"
	"net/http"
)

func GetReplicaSet(c *gin.Context) {
	responseData := HandleReplicaSet(common.Get, c)
	c.JSON(responseData.Code, responseData)
}

func DeleteReplicaSet(c *gin.Context) {
	responseData := HandleReplicaSet(common.Delete, c)
	c.JSON(responseData.Code, responseData)
}

func ListReplicaSet(c *gin.Context) {
	responseData := HandleReplicaSet(common.List, c)
	c.JSON(responseData.Code, responseData)
}

func HandleReplicaSet(action common.ActionType, c *gin.Context) (responseData *common.ResponseData) {
	// 获取clientSet，如果失败直接返回错误
	clientSet, err := access.Access(c.Query("cluster"))
	responseData = handle.HandlerResponse(nil, err)
	if responseData.Code != http.StatusOK {
		log.Errorf("%s%s", common.K8SClientSetError, err)
		return
	}
	// 获取HTTP的参数，存到handle.Resources结构体中
	commonParams := handle.GenerateCommonParams(c, clientSet)
	r := resource.ReplicaSetResource{Params: commonParams}
	// 调用结构体方法
	switch action {
	case common.Get:
		response, err := r.Get()
		responseData = handle.HandlerResponse(response, err)
	case common.Delete:
		err := r.Delete()
		responseData = handle.HandlerResponse(nil, err)
	case common.List:
		response, err := r.List()
		if err != nil {
			responseData = handle.HandlerResponse(nil, err)
		} else {
			responseData = handle.HandlerResponse(response.Items, err)
		}
	}
	return
}
