package test

import (
	setting2 "github.com/BigerWANG/my-gin/pkg/setting"
	"testing"
)

func TestSetting(t *testing.T) {
	setting2.LoadServer()
	setting2.LoadApp()
	setting2.LoadBase()

	t.Log(setting2.RunMode)
	t.Log(setting2.HTTPPort)
	t.Log(setting2.ReadTimeout)
	t.Log(setting2.WriteTimeout)
	t.Log(setting2.PageSize)
	t.Log(setting2.JwtSecret)
}

