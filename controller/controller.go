package controller

import (
	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"quiz-nds-be/client"
)

type ControllerInterface interface {
	HitungLuasPersegi(panjang int32, lebar int32) (int32, error)
	HitungKelilingPersegi(sisi int32) (int32, error)
	HitungLuasLingkaran(diameter float64) (float64, error)
	HitungKelilingLingkaran(diameter float64) (float64, error)
	HitungLuasSegitiga(alas int32, tinggi int32) (int32, error)
	HitungKelilingSegitiga(sisiA int32, sisiB int32, sisiC int32) (int32, error)
}

type ControllerStruct struct {
	logger interfaces.Logger
	cl     client.ClientInterFace
}

type ControlleStruct struct {
	logger interfaces.Logger
}

func NewController(logger interfaces.Logger, cl client.ClientInterFace) ControllerInterface {
	return &ControllerStruct{
		logger: logger,
		cl:     cl,
	}
}

func (c *ControllerStruct) HitungLuasPersegi(panjang int32, lebar int32) (int32, error) {
	if panjang < 1 {
		c.logger.Error("Error", "Panjang Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Panjang harus Lebih Dari 0")
	}
	if lebar < 1 {
		c.logger.Error("Error", "Lebar Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Lebar harus Lebih Dari 0")
	}
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)
	luas := c.cl.HitungLuasPersegi(panjang, lebar)
	c.logger.Info("LUAS", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungKelilingPersegi(sisi int32) (int32, error) {
	if sisi < 1 {
		c.logger.Error("Error", "Sisi harus lebih dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi harus lebih dari 0")
	}

	c.logger.Info("PANJANG SISI", sisi)

	keliling := c.cl.HitungKelilingPersegi(sisi)
	c.logger.Info("KELILING", keliling)

	return keliling, nil
}

func (c *ControllerStruct) HitungLuasLingkaran(diameter float64) (float64, error) {
	if diameter < 1 {
		c.logger.Error("Error", "Diameter harus lebih dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Diameter harus lebih dari 0")
	}

	c.logger.Info("PANJANG DIAMETER", diameter)

	luas := c.cl.HitungLuasLingkaran(diameter)
	c.logger.Info("LUAS", luas)

	return luas, nil
}

func (c *ControllerStruct) HitungKelilingLingkaran(diameter float64) (float64, error) {
	if diameter < 1 {
		c.logger.Error("Error", "Diameter harus lebih dari 0")
		return 0, status.Error(codes.FailedPrecondition, "Diameter harus lebih dari 0")
	}

	c.logger.Info("PANJANG DIAMETER", diameter)

	keliling := c.cl.HitungKelilingLingkaran(diameter)
	c.logger.Info("KELILING", keliling)

	return keliling, nil
}

func (c *ControllerStruct) HitungLuasSegitiga(alas int32, tinggi int32) (int32, error) {
	if alas < 1 {
		c.logger.Error("Error", "Alas Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Alas harus Lebih Dari 0")
	}
	if tinggi < 1 {
		c.logger.Error("Error", "Tinggi Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Tinggi harus Lebih Dari 0")
	}

	c.logger.Info("ALAS:", alas)
	c.logger.Info("TINGGI", tinggi)

	luas := c.cl.HitungLuasSegitiga(alas, tinggi)
	c.logger.Info("LUAS:", luas)
	return luas, nil
}

func (c *ControllerStruct) HitungKelilingSegitiga(sisiA int32, sisiB int32, sisiC int32) (int32, error) {
	if sisiA < 1 {
		c.logger.Error("Error", "Sisi A Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi A harus Lebih Dari 0")
	}
	if sisiB < 1 {
		c.logger.Error("Error", "Sisi B Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi B harus Lebih Dari 0")
	}
	if sisiC < 1 {
		c.logger.Error("Error", "Sisi C Harus Lebih Dri 0")
		return 0, status.Error(codes.FailedPrecondition, "Sisi C harus Lebih Dari 0")
	}

	c.logger.Info("SISI A:", sisiA)
	c.logger.Info("SISI B:", sisiB)
	c.logger.Info("SISI C:", sisiC)

	keliling := c.cl.HitungKelilingSegitiga(sisiA, sisiB, sisiC)
	c.logger.Info("KELILING:", keliling)
	return keliling, nil
}
