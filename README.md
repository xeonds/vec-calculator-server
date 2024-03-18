## 分布式计算作业-支持多客户端并发访问的向量计算服务端程序

## Feature

- 向量点乘
- 向量叉乘
- 多用户并发访问

### 设计文档

本程序使用Gin作为Web框架实现整个Web服务，借助Golang语言级别的轻量级协程Goroutine实现多用户并发访问，并且Gin内部实现了基于Goroutine的线程池，降低了线程创建的开销。

我还编写了一个简单的前端页面，用于测试向量计算服务端程序。

另外在Makefile中，我添加了`make bench`命令，用于进行性能测试。Payload在`bench.lua`中，我使用了一组简单的数值，可以根据情况自行修改。

使用方法如下：首先使用`make init && make`命令编译程序，然后使用`make run`命令启动服务端程序。初次启动时，程序会创建`config.yaml`文件并退出，此时可以修改配置文件中的端口号等信息。然后再次使用`make run`命令启动服务端程序。

最后可以使用`make bench`命令进行性能测试。通过修改`bench.lua`和`Makefile`中的参数，可以进行不同的性能测试。

## Benchmark

测试平台配置为`12th Gen Intel(R) Core(TM) i7-12700H`，`40GB`内存，`512GB`固态硬盘。

首先对语言实现的原生点乘和叉乘函数直接进行性能基准测试。

对点乘和叉乘各进行了`64`组测试，测试用例随机生成，每个用例的向量长度在`1-10000`之间，每组测试进行1000000000次以保证结果稳定性。设置最大并发数为`1000000000`，测试结果如下：

```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^BenchmarkAPI$ vec-calc-server/test
goos: linux
goarch: amd64
pkg: vec-calc-server/test
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
1000000000	         0.0000099 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	vec-calc-server/test	1.800s
```

根据`TPS = Total Transactions / Time`的计算公式，可得：

```
TPS = 1,000,000,000 * 64 / 1.800 = 355,5555,5555
```

其次对Web服务进行性能测试。借助`wrk`工具，在10000并发数下，叉乘接口的API压力测试结果如下：

```
xeonds@ark-station:~/code/vec-calculator-server$ make bench 
cd build && ./vec-calc-web-linux-amd64-1.0.0 & sleep 1 && \
wrk -t 20 -c 10000 -d 180s -s bench.lua --latency "http://localhost:8080/api/calc/mul"
Running 3m test @ http://localhost:8080/api/calc/mul
  20 threads and 10000 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency    61.27ms   62.15ms   1.46s    93.06%
    Req/Sec     9.74k     1.33k   22.11k    71.75%
  Latency Distribution
     50%   48.47ms
     75%   66.32ms
     90%   90.10ms
     99%  364.05ms
  34884275 requests in 3.00m, 4.35GB read
Requests/sec: 193693.29
Transfer/sec:     24.75MB
```

可以看到，将服务作为一个Web服务后，性能有所下降，但是依然可以达到`193693.29`的TPS。