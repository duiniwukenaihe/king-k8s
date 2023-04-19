package impl

import (
	"context"
	"encoding/json"
	pb "github.com/duiniwukenaihe/king-k8s/grpc/proto"
	"github.com/duiniwukenaihe/king-k8s/util"
	"github.com/duiniwukenaihe/king-utils/common/log"
)

type Deployment struct{}

func (s *Deployment) GetByLabels(ctx context.Context, in *pb.DeploymentRequest) (*pb.DeploymentResponse, error) {
	// 生成通用参数
	commonParams, err := GenerateCommonParams(in.Cluster, in.Namespace, in.Name)
	if err != nil {
		log.Errorf("Generate Common Params Error: %v", err)
		return nil, err
	}
	deployment, err := util.GetDeploymentBySelectorLabel(in.Labels, in.Namespace, commonParams.ClientSet)
	if err != nil {
		log.Errorf("Get Deployment Error: %v", err)
		return nil, err
	}
	// 序列化
	b, err := json.Marshal(deployment)
	if err != nil {
		log.Errorf("Marshal Deployment Error: %v", err)
		return nil, err
	}
	return &pb.DeploymentResponse{Data: b}, nil
}
