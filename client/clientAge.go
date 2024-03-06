package client

import (
	"bitbucket.bri.co.id/nds/nds-go-module-logger/interfaces"
	"strconv"
	"time"
)

type ClientAgeInterFace interface {
	HitungUmur(birthYear string) int32
}

type ClientAgeStruct struct {
	logger interfaces.Logger
}

func NewClientAgeStruct(logger interfaces.Logger) ClientAgeInterFace {
	return &ClientAgeStruct{
		logger: logger,
	}
}

func (c *ClientAgeStruct) HitungUmur(birthYear string) int32 {
	// parse date string to time.Time
	birth, err := time.Parse("2006-01-02", birthYear)
	if err != nil {
		c.logger.Error("Error parsing hari ulang tahun", err)
		return 0
	}

	currentTime := time.Now()
	age := currentTime.Year() - birth.Year()
	if currentTime.YearDay() < birth.YearDay() {
		age--
	}

	c.logger.Info("BIRTH DATE", strconv.Itoa(birth.Year()))
	c.logger.Info("AGE", strconv.Itoa(age))

	return int32(age)
}
