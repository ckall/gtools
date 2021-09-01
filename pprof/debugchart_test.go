package pprof

import (
	_ "github.com/mkevac/debugcharts"
	_ "net/http/pprof"
	"testing"
)

/**
 * @auth: kuncheng
 * @Date: 2021/8/31
 */
func TestChart(t *testing.T)  {
	InitChart(":9090")
}