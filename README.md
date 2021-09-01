# gck


日志系统  
```
  日志分割
  格式定义
```
推送系统
```
  信息框警告：
     钉钉   
  邮箱警告
     待完成    
  电话警告 
     待完成      
```
### 链路跟踪:
```
   * zipkin 
```
### 性能监控

##### 访问http://HOST:PORT/debug/pprof

##### 内存火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/heap

##### CPU火焰图本地运行go tool pprof -http=:8081 http://HOST:PORT/debug/pprof/profile?seconds=10

##### 实时监控访问http://HOST:PORT/debug/charts