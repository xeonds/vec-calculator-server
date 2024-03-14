## 分布式计算作业-支持多客户端并发访问的向量计算服务端程序

## Feature

- 向量点乘
- 向量叉乘
- 多用户并发访问

## Benchmark

对点乘和叉乘API各进行了`64`组测试，测试用例随机生成，每个用例的向量长度在`1-10000`之间，每组测试进行`10亿`次以保证结果稳定性。

```
Running tool: /usr/bin/go test -benchmem -run=^$ -bench ^BenchmarkAPI$ vec-calc-server/test
goos: linux
goarch: amd64
pkg: vec-calc-server/test
cpu: 12th Gen Intel(R) Core(TM) i7-12700H
BenchmarkAPI/case-0-20         	[GIN] 2024/03/15 - 01:55:59 | 200 |     305.271µs |                 | POST     "/api/calc/mul"
1000000000	         0.0003481 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-0#01-20      	[GIN] 2024/03/15 - 01:55:59 | 200 |     369.961µs |                 | POST     "/api/calc/dot"
1000000000	         0.0002685 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-1-20         	[GIN] 2024/03/15 - 01:55:59 | 200 |    8.171779ms |                 | POST     "/api/calc/mul"
1000000000	         0.007987 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-1#01-20      	[GIN] 2024/03/15 - 01:55:59 | 200 |    8.972781ms |                 | POST     "/api/calc/dot"
1000000000	         0.006544 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-2-20         	[GIN] 2024/03/15 - 01:55:59 | 200 |    8.258421ms |                 | POST     "/api/calc/mul"
1000000000	         0.005618 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-2#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.226949ms |                 | POST     "/api/calc/dot"
1000000000	         0.004117 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-3-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |    6.583122ms |                 | POST     "/api/calc/mul"
1000000000	         0.007272 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-3#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.989597ms |                 | POST     "/api/calc/dot"
1000000000	         0.005468 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-4-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |     4.02502ms |                 | POST     "/api/calc/mul"
1000000000	         0.003161 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-4#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    2.534545ms |                 | POST     "/api/calc/dot"
1000000000	         0.002538 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-5-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |    6.459178ms |                 | POST     "/api/calc/mul"
1000000000	         0.006161 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-5#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    5.141992ms |                 | POST     "/api/calc/dot"
1000000000	         0.005096 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-6-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.286692ms |                 | POST     "/api/calc/mul"
1000000000	         0.004367 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-6#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    3.584979ms |                 | POST     "/api/calc/dot"
1000000000	         0.003555 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-7-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |    5.851163ms |                 | POST     "/api/calc/mul"
1000000000	         0.008546 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-7#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |     4.95232ms |                 | POST     "/api/calc/dot"
1000000000	         0.004975 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-8-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |     4.84215ms |                 | POST     "/api/calc/mul"
1000000000	         0.004890 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-8#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.716738ms |                 | POST     "/api/calc/dot"
1000000000	         0.004230 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-9-20         	[GIN] 2024/03/15 - 01:56:00 | 200 |    3.091025ms |                 | POST     "/api/calc/mul"
1000000000	         0.003160 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-9#01-20      	[GIN] 2024/03/15 - 01:56:00 | 200 |    2.536411ms |                 | POST     "/api/calc/dot"
1000000000	         0.002584 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-10-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |     819.336µs |                 | POST     "/api/calc/mul"
1000000000	         0.001495 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-10#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |     1.05354ms |                 | POST     "/api/calc/dot"
1000000000	         0.001134 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-11-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    3.391931ms |                 | POST     "/api/calc/mul"
1000000000	         0.004053 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-11#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |    3.577756ms |                 | POST     "/api/calc/dot"
1000000000	         0.004738 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-12-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    7.762491ms |                 | POST     "/api/calc/mul"
1000000000	         0.008380 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-12#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |    6.677343ms |                 | POST     "/api/calc/dot"
1000000000	         0.006261 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-13-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    2.585977ms |                 | POST     "/api/calc/mul"
1000000000	         0.004449 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-13#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |    3.525651ms |                 | POST     "/api/calc/dot"
1000000000	         0.002182 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-14-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    6.343701ms |                 | POST     "/api/calc/mul"
1000000000	         0.006290 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-14#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |    5.081927ms |                 | POST     "/api/calc/dot"
1000000000	         0.008662 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-15-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |     620.505µs |                 | POST     "/api/calc/mul"
1000000000	         0.001042 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-15#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |     517.447µs |                 | POST     "/api/calc/dot"
1000000000	         0.0008632 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-16-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |     451.801µs |                 | POST     "/api/calc/mul"
1000000000	         0.0004752 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-16#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |     575.073µs |                 | POST     "/api/calc/dot"
1000000000	         0.0003881 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-17-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    5.672073ms |                 | POST     "/api/calc/mul"
1000000000	         0.005057 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-17#01-20     	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.574285ms |                 | POST     "/api/calc/dot"
1000000000	         0.004392 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-18-20        	[GIN] 2024/03/15 - 01:56:00 | 200 |    4.111659ms |                 | POST     "/api/calc/mul"
1000000000	         0.004604 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-18#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    3.249605ms |                 | POST     "/api/calc/dot"
1000000000	         0.003427 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-19-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |     105.563µs |                 | POST     "/api/calc/mul"
1000000000	         0.0000883 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-19#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |      91.937µs |                 | POST     "/api/calc/dot"
1000000000	         0.0000664 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-20-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    1.886873ms |                 | POST     "/api/calc/mul"
1000000000	         0.002819 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-20#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |      1.4644ms |                 | POST     "/api/calc/dot"
1000000000	         0.001512 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-21-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |   10.661227ms |                 | POST     "/api/calc/mul"
1000000000	         0.005871 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-21#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    4.845913ms |                 | POST     "/api/calc/dot"
1000000000	         0.006819 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-22-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    3.036452ms |                 | POST     "/api/calc/mul"
1000000000	         0.002984 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-22#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    2.497289ms |                 | POST     "/api/calc/dot"
1000000000	         0.002559 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-23-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |     2.61399ms |                 | POST     "/api/calc/mul"
1000000000	         0.002665 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-23#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    2.140861ms |                 | POST     "/api/calc/dot"
1000000000	         0.002431 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-24-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |     294.904µs |                 | POST     "/api/calc/mul"
1000000000	         0.0003337 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-24#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |     378.358µs |                 | POST     "/api/calc/dot"
1000000000	         0.0004446 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-25-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    5.186256ms |                 | POST     "/api/calc/mul"
1000000000	         0.005334 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-25#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    4.258031ms |                 | POST     "/api/calc/dot"
1000000000	         0.004368 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-26-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |     852.613µs |                 | POST     "/api/calc/mul"
1000000000	         0.0008717 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-26#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |     671.223µs |                 | POST     "/api/calc/dot"
1000000000	         0.0007425 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-27-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    4.081159ms |                 | POST     "/api/calc/mul"
1000000000	         0.003817 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-27#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |     3.09936ms |                 | POST     "/api/calc/dot"
1000000000	         0.003213 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-28-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    6.964162ms |                 | POST     "/api/calc/mul"
1000000000	         0.006777 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-28#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    5.534345ms |                 | POST     "/api/calc/dot"
1000000000	         0.005500 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-29-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    3.971073ms |                 | POST     "/api/calc/mul"
1000000000	         0.004676 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-29#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    3.295262ms |                 | POST     "/api/calc/dot"
1000000000	         0.003391 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-30-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    5.019665ms |                 | POST     "/api/calc/mul"
1000000000	         0.004886 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-30#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    4.915974ms |                 | POST     "/api/calc/dot"
1000000000	         0.004006 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-31-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    8.126532ms |                 | POST     "/api/calc/mul"
1000000000	         0.007241 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-31#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    6.042992ms |                 | POST     "/api/calc/dot"
1000000000	         0.005958 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-32-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    6.790825ms |                 | POST     "/api/calc/mul"
1000000000	         0.006482 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-32#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    5.339318ms |                 | POST     "/api/calc/dot"
1000000000	         0.008189 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-33-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    4.063764ms |                 | POST     "/api/calc/mul"
1000000000	         0.003397 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-33#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    2.664655ms |                 | POST     "/api/calc/dot"
1000000000	         0.002682 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-34-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    8.661778ms |                 | POST     "/api/calc/mul"
1000000000	         0.006724 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-34#01-20     	[GIN] 2024/03/15 - 01:56:01 | 200 |    6.329015ms |                 | POST     "/api/calc/dot"
1000000000	         0.005575 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-35-20        	[GIN] 2024/03/15 - 01:56:01 | 200 |    6.258684ms |                 | POST     "/api/calc/mul"
1000000000	         0.006241 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-35#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.213366ms |                 | POST     "/api/calc/dot"
1000000000	         0.005022 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-36-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    6.414061ms |                 | POST     "/api/calc/mul"
1000000000	         0.006417 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-36#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.330255ms |                 | POST     "/api/calc/dot"
1000000000	         0.005303 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-37-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    7.806493ms |                 | POST     "/api/calc/mul"
1000000000	         0.007500 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-37#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    6.328612ms |                 | POST     "/api/calc/dot"
1000000000	         0.006332 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-38-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |     1.94775ms |                 | POST     "/api/calc/mul"
1000000000	         0.002072 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-38#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    1.616648ms |                 | POST     "/api/calc/dot"
1000000000	         0.001644 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-39-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.981857ms |                 | POST     "/api/calc/mul"
1000000000	         0.003759 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-39#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    4.815235ms |                 | POST     "/api/calc/dot"
1000000000	         0.004864 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-40-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.200808ms |                 | POST     "/api/calc/mul"
1000000000	         0.005157 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-40#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    4.306657ms |                 | POST     "/api/calc/dot"
1000000000	         0.004230 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-41-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    1.112673ms |                 | POST     "/api/calc/mul"
1000000000	         0.001133 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-41#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |     942.001µs |                 | POST     "/api/calc/dot"
1000000000	         0.001503 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-42-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    7.607761ms |                 | POST     "/api/calc/mul"
1000000000	         0.007381 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-42#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    7.293077ms |                 | POST     "/api/calc/dot"
1000000000	         0.008549 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-43-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    2.861521ms |                 | POST     "/api/calc/mul"
1000000000	         0.002727 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-43#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    3.597893ms |                 | POST     "/api/calc/dot"
1000000000	         0.002259 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-44-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |     952.221µs |                 | POST     "/api/calc/mul"
1000000000	         0.001010 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-44#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    1.302362ms |                 | POST     "/api/calc/dot"
1000000000	         0.0008380 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-45-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    6.927543ms |                 | POST     "/api/calc/mul"
1000000000	         0.006860 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-45#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.605648ms |                 | POST     "/api/calc/dot"
1000000000	         0.005600 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-46-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    6.611454ms |                 | POST     "/api/calc/mul"
1000000000	         0.006975 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-46#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.385245ms |                 | POST     "/api/calc/dot"
1000000000	         0.005431 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-47-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    6.208781ms |                 | POST     "/api/calc/mul"
1000000000	         0.005992 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-47#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    5.023884ms |                 | POST     "/api/calc/dot"
1000000000	         0.005214 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-48-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    2.282077ms |                 | POST     "/api/calc/mul"
1000000000	         0.002312 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-48#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    1.923941ms |                 | POST     "/api/calc/dot"
1000000000	         0.001945 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-49-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |     2.69863ms |                 | POST     "/api/calc/mul"
1000000000	         0.002698 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-49#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |     2.25671ms |                 | POST     "/api/calc/dot"
1000000000	         0.002262 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-50-20        	[GIN] 2024/03/15 - 01:56:02 | 200 |    4.649356ms |                 | POST     "/api/calc/mul"
1000000000	         0.003652 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-50#01-20     	[GIN] 2024/03/15 - 01:56:02 | 200 |    3.021397ms |                 | POST     "/api/calc/dot"
1000000000	         0.003059 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-51-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |     979.026µs |                 | POST     "/api/calc/mul"
1000000000	         0.0009799 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-51#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |     808.608µs |                 | POST     "/api/calc/dot"
1000000000	         0.0008350 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-52-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    7.789595ms |                 | POST     "/api/calc/mul"
1000000000	         0.008536 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-52#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    6.266907ms |                 | POST     "/api/calc/dot"
1000000000	         0.01020 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-53-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.040781ms |                 | POST     "/api/calc/mul"
1000000000	         0.002076 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-53#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    1.654624ms |                 | POST     "/api/calc/dot"
1000000000	         0.001930 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-54-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    7.015085ms |                 | POST     "/api/calc/mul"
1000000000	         0.006548 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-54#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    5.377917ms |                 | POST     "/api/calc/dot"
1000000000	         0.005548 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-55-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    7.355943ms |                 | POST     "/api/calc/mul"
1000000000	         0.007064 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-55#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    5.989724ms |                 | POST     "/api/calc/dot"
1000000000	         0.005961 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-56-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    7.882375ms |                 | POST     "/api/calc/mul"
1000000000	         0.007828 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-56#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    8.499707ms |                 | POST     "/api/calc/dot"
1000000000	         0.006374 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-57-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.136416ms |                 | POST     "/api/calc/mul"
1000000000	         0.001330 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-57#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |     1.08315ms |                 | POST     "/api/calc/dot"
1000000000	         0.001114 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-58-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |     3.82611ms |                 | POST     "/api/calc/mul"
1000000000	         0.002663 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-58#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.180235ms |                 | POST     "/api/calc/dot"
1000000000	         0.002251 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-59-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    5.928201ms |                 | POST     "/api/calc/mul"
1000000000	         0.005873 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-59#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |      5.0036ms |                 | POST     "/api/calc/dot"
1000000000	         0.004766 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-60-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    4.084131ms |                 | POST     "/api/calc/mul"
1000000000	         0.005826 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-60#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    4.087756ms |                 | POST     "/api/calc/dot"
1000000000	         0.003386 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-61-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.074459ms |                 | POST     "/api/calc/mul"
1000000000	         0.002925 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-61#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.252282ms |                 | POST     "/api/calc/dot"
1000000000	         0.001733 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-62-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.612712ms |                 | POST     "/api/calc/mul"
1000000000	         0.002660 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-62#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.136132ms |                 | POST     "/api/calc/dot"
1000000000	         0.002672 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-63-20        	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.597274ms |                 | POST     "/api/calc/mul"
1000000000	         0.002549 ns/op	       0 B/op	       0 allocs/op
BenchmarkAPI/case-63#01-20     	[GIN] 2024/03/15 - 01:56:03 | 200 |    2.048287ms |                 | POST     "/api/calc/dot"
1000000000	         0.002186 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	vec-calc-server/test	4.042s
```

## TODO

编写并发访问测试代码