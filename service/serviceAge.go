package service

import (
	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"context"
	"quiz-nds-be/controller"
	"quiz-nds-be/generated/proto"
)

type ServiceAgeStruct struct {
	logger interfaces.Logger
	ctr    controller.ControllerAgeInterface
}

func NewServiceAge(logger interfaces.Logger, ctr controller.ControllerAgeInterface) proto.AgeServer {
	return &ServiceAgeStruct{
		logger: logger,
		ctr:    ctr,
	}
}

func (s *ServiceAgeStruct) HitungUmur(ctx context.Context, req *proto.RequestTanggalLahir) (*proto.ResponseTanggalLahirSekarang, error) {
	data, err := s.ctr.HitungUmur(req.GetBirthYear())
	if err != nil {
		return nil, err
	}

	return &proto.ResponseTanggalLahirSekarang{
		CurrentAge: data,
	}, nil

}
