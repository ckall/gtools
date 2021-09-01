## 框架开启pprof



#### 访问http://HOST:PORT/debug/pprof

#### 内存火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/heap

#### CPU火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/profile?seconds=10

#### 实时监控访问http://HOST:PORT/debug/charts