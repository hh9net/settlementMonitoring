package utils

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestStrTimeTotime(t *testing.T) {
	StrTimeTotime("1212-12-12 12:12:12")
}

func TestStrTimeToNowtime(t *testing.T) {
	StrTimeToNowtime()
}

//KuaizhaoTimeNowFormat
func TestKuaizhaoTimeNowFormat(t *testing.T) {
	logrus.Print(KuaizhaoTimeNowFormat())
}
