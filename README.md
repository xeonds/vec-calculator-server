## 分布式计算作业-支持多客户端并发访问的向量计算服务端程序

## Feature

- 向量点乘
- 向量叉乘
- 多用户并发访问

## Benchmark

对点乘和叉乘API各进行了`64`组测试，测试用例随机生成，每个用例的向量长度在`1-10000`之间，每组测试进行1000000000次以保证结果稳定性。设置最大并发数为`1000000000`，测试结果如下：

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