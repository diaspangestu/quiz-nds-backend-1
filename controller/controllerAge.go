package controller

import (
	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"quiz-nds-be/client"
)

type ControllerAgeInterface interface {
	HitungUmur(birthDay string) (int32, error)
}

type ControllerAgeStruct struct {
	loggger interfaces.Logger
	cl      client.ClientAgeInterFace
}

type ControlleAgeStruct struct {
	logger interfaces.Logger
}

func NewControllerAge(logger interfaces.Logger, cl client.ClientAgeInterFace) ControllerAgeInterface {
	return &ControllerAgeStruct{
		loggger: logger,
		cl:      cl,
	}
}

func (c *ControllerAgeStruct) HitungUmur(birthDay string) (int32, error) {
	if birthDay == "" {
		c.loggger.Error("Error", "Tanggal lahir tidak boleh kosong")
		return 0, status.Error(codes.FailedPrecondition, "Tanggal lahir tidak boleh kosong")
	}

	c.loggger.Info("BIRTH DATE", birthDay)
	age := c.cl.HitungUmur(birthDay)
	c.loggger.Info("AGE", age)

	return age, nil
}
