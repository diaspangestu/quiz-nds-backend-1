package client

import (
	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"math"
)

type ClientInterFace interface {
	HitungLuasPersegi(panjang int32, lebar int32) int32
	HitungKelilingPersegi(sisi int32) int32
	HitungLuasLingkaran(diameter float64) float64
	HitungKelilingLingkaran(diameter float64) float64
	HitungLuasSegitiga(alas int32, tinggi int32) int32
	HitungKelilingSegitiga(sisiA int32, sisiB int32, sisiC int32) int32
}

type ClientStruct struct {
	logger interfaces.Logger
}

func NewClientStruct(logger interfaces.Logger) ClientInterFace {
	return &ClientStruct{
		logger: logger,
	}
}

func (c *ClientStruct) HitungLuasPersegi(panjang int32, lebar int32) int32 {
	c.logger.Info("PANJANG SISI", panjang)
	c.logger.Info("LEBAR SISI", lebar)

	luas := panjang * lebar
	c.logger.Info("LUAS", luas)

	return luas
}

func (c *ClientStruct) HitungKelilingPersegi(sisi int32) int32 {
	c.logger.Info("PANJANG SISI", sisi)

	keliling := 4 * sisi
	c.logger.Info("KELILING", keliling)

	return keliling
}

func (c *ClientStruct) HitungLuasLingkaran(diameter float64) float64 {
	c.logger.Info("PANJANG DIAMETER", diameter)

	phi := math.Pi
	luas := phi * math.Pow(diameter, 2)
	c.logger.Info("LUAS", luas)

	return luas
}

func (c *ClientStruct) HitungKelilingLingkaran(diameter float64) float64 {
	c.logger.Info("PANJANG DIAMETER", diameter)

	phi := math.Pi
	keliling := phi * diameter
	c.logger.Info("KELILING", keliling)

	return keliling
}

func (c *ClientStruct) HitungLuasSegitiga(alas int32, tinggi int32) int32 {
	c.logger.Info("ALAS:", alas)
	c.logger.Info("TINGGI:", tinggi)

	luas := (alas * tinggi) / 2
	c.logger.Info("LUAS:", luas)

	return luas
}

func (c *ClientStruct) HitungKelilingSegitiga(sisiA int32, sisiB int32, sisiC int32) int32 {
	c.logger.Info("SISI A:", sisiA)
	c.logger.Info("SISI B:", sisiB)
	c.logger.Info("SISI C:", sisiC)

	keliling := sisiA + sisiB + sisiC
	c.logger.Info("KELILING:", keliling)

	return keliling
}
