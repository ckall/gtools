package pprof

import (
	_ "github.com/mkevac/debugcharts"
	"log"
	"net/http"
	_ "net/http/pprof"
)

/**
 * @auth: kuncheng
 * @Date: 2021/8/31
 */
func InitChart(addr string) {
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			log.Printf("debugcharts开启失败: %s", err.Error())
		}
	}()
}
