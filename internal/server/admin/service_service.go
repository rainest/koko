package admin

import (
	"context"
	"fmt"
	"net/http"

	pbModel "github.com/kong/koko/internal/gen/grpc/kong/admin/model/v1"
	v1 "github.com/kong/koko/internal/gen/grpc/kong/admin/service/v1"
	"github.com/kong/koko/internal/model"
	"github.com/kong/koko/internal/resource"
	"github.com/kong/koko/internal/store"
	"go.uber.org/zap"
)

type ServiceService struct {
	v1.UnimplementedServiceServiceServer
	CommonOpts
}

func (s *ServiceService) GetService(ctx context.Context,
	req *v1.GetServiceRequest) (*v1.GetServiceResponse, error) {
	result := resource.NewService()
	s.logger.With(zap.String("id", req.Id)).Debug("reading service by id")
	err := s.store.Read(ctx, result, store.GetByID(req.Id))
	if err != nil {
		return nil, s.err(err)
	}
	return &v1.GetServiceResponse{
		Item: result.Service,
	}, nil
}

func (s *ServiceService) CreateService(ctx context.Context,
	req *v1.CreateServiceRequest) (*v1.CreateServiceResponse, error) {
	res := resource.NewService()
	res.Service = req.Item
	err := s.store.Create(ctx, res)
	if err != nil {
		return nil, s.err(err)
	}
	setHeader(ctx, http.StatusCreated)
	return &v1.CreateServiceResponse{
		Item: res.Service,
	}, nil
}

func (s *ServiceService) DeleteService(ctx context.Context,
	request *v1.DeleteServiceRequest) (*v1.DeleteServiceResponse, error) {
	err := s.store.Delete(ctx, store.DeleteByID(request.Id),
		store.DeleteByType(resource.TypeService))
	if err != nil {
		return nil, s.err(err)
	}
	setHeader(ctx, http.StatusNoContent)
	return &v1.DeleteServiceResponse{}, nil
}

func (s *ServiceService) ListServices(ctx context.Context,
	_ *v1.ListServicesRequest) (*v1.ListServicesResponse, error) {
	list := resource.NewList(resource.TypeService)
	if err := s.store.List(ctx, list); err != nil {
		return nil, s.err(err)
	}
	return &v1.ListServicesResponse{
		Items: servicesFromObjects(list.GetAll()),
	}, nil
}

func (s *ServiceService) err(err error) error {
	return handleErr(s.logger, err)
}

func servicesFromObjects(objects []model.Object) []*pbModel.Service {
	res := make([]*pbModel.Service, 0, len(objects))
	for _, object := range objects {
		service, ok := object.Resource().(*pbModel.Service)
		if !ok {
			panic(fmt.Sprintf("expected type '%T' but got '%T'",
				&pbModel.Service{}, object.Resource()))
		}
		res = append(res, service)
	}
	return res
}
