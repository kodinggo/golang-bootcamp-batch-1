package service

import (
	"context"
	pb "kodinggo/pb/story"
)

type Service struct {
	pb.UnimplementedStoryServiceServer
}

func NewService() *Service {
	return new(Service)
}

func (s *Service) FindByID(ctx context.Context, in *pb.FindByIDRequest) (out *pb.Story, err error) {
	return &pb.Story{
		Id:    1,
		Title: "Example title",
	}, nil
}
func (s *Service) FindAll(ctx context.Context, in *pb.FindAllRequest) (out *pb.Stories, err error) {
	stories := []*pb.Story{
		{
			Id:    1,
			Title: "Example title 1",
		},
		{
			Id:    2,
			Title: "Example title 2",
		},
	}

	return &pb.Stories{Stories: stories}, nil
}
