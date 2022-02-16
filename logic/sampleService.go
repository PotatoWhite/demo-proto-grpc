package logic

import (
	"context"
	"log"
	pb "potato/demo-protocolbuffer-grpc/grpc/proto"
)

type Service struct {
}

func (s Service) GetAllSamples(ctx context.Context, request *pb.GetAllSamplesRequest) (*pb.GetAllSamplesResponse, error) {
	log.Println(request)
	samples := make([]*pb.Sample, request.Max)
	for index := 0; int32(index) < request.Max; index++ {
		samples[index] = &pb.Sample{
			Id:   int32(index),
			Data: request.Data,
		}
	}

	return &pb.GetAllSamplesResponse{
		Samples: samples,
	}, nil
}
