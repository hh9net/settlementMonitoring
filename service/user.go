package service

import (
	"github.com/sirupsen/logrus"
	"settlementMonitoring/dto"
)

func Login(req dto.ReqLogin) (int, error) {
	logrus.Print(req)
	return 200, nil
}
