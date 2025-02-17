/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

//go:generate protoc -I ../helloworld --go_out=plugins=grpc:../helloworld ../helloworld/helloworld.proto

package impl

import (
	"context"
	pb "github.com/duiniwukenaihe/king-k8s/grpc/proto"
	"log"
)

// server is used to implement helloworld.GreeterServer.
type Say struct{}

// SayHello implements helloworld.GreeterServer
func (s *Say) SayHello(ctx context.Context, in *pb.SayRequest) (*pb.SayReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.SayReply{Message: "Hello" + in.Name}, nil
}
