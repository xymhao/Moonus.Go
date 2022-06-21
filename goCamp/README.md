# go 训练营毕业总结

## 微服务架构（BFF、Service、Admin、Job、Task 分模块）
使用Envoy或者kong作为网关，可以用作七层负载均衡和路由规划，以及设置超时、限频、熔断策略。
BFF层，针对不同的设备，开发人员可以更加专注业务逻辑交付，以及性能优化。


## API 设计（包括 API 定义、错误码规范、Error 的使用）
- API的定义可以采用Restful的架构风格，一切皆资源。rpc框架：优先可以考虑gRPC来实现我们的api。
- 常用网络框架Gin
目录结构：api-服务名称-示例。
考虑API的向后兼容。
- 统一做业务错误码。
- 命名方式：
/<package_name>.<version>.<service_name>/{method}

- 错误处理
https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto
    - 统一的错误编码

- Go 项目工程化（项目结构、DI、代码分层、ORM 框架）
项目结构
- api：API协议定义目录
- internal：
    - biz：业务逻辑，DDD domain层
    - data： 数据层的封装，repo
    - service： DDD application层，处理DTO之间的转换，将API层的DTO->DO

通过wire使用依赖倒置，我们可以再biz层定义接口，data层实现逻辑，通过能力供应商模式，保证了领域层的绝对稳定。

- 并发的使用（errgroup 的并行链路请求）
goCamp/week13/cmd/services/main.go 通过errgroup 管理goroutine， 快速初始化和实现优雅终止。

- 微服务中间件的使用（ELK、Opentracing、Prometheus、Kafka）
    - ELK进行日志采集，和日志观测。
    - Opentracing：（https://opentracing.io/docs/overview/what-is-tracing/） :微服务下离不开全链路的追踪，便于故障排查，以及导致性能低下的原因。
        - Jaeger:Uber 开源的分布式跟踪系统
    - Prometheus:实现服务的基础指标的采集，通过grafana实现指标的展示。
    - Kafka：消息队列，用于日志采集、服务解耦、流量削峰，最终一致性的分布式事务。


-缓存的使用优化（一致性处理、Pipeline优化）
    - 一致性处理
        - 写操作使用Set操作，覆盖缓存；
        - 读操作，使用SetNX，操作回写MISS数据。
    - 通过Pipeline批量处理命令，降低网络IO的时延


# 毕业总结
    通过Go进阶训练营的学习，深入的学习了Go语言的相关知识，同时不仅仅接触到了Go，还学习微服务体系下的知识，扩宽了自己的知识面。
    从基础知识，Go的异常处理，并发编程，随着计算机的硬件的发展，了解到Go语言设置为什么被称为现在语言；再到Go语言的工程化实践，学习开源的框架，
轻松的从0到1开发一个Go语言的Web程序；再到架构设计，学习了DDD在Go工程下的实践思路，学习了kafka、redis设计我们的服务架构；总后还有DevOps相关的
学习，包含ELK的日志采集，Opentracing的全链路追踪，Prometheus的指标采集，以及grafana的可观测面板。系统性的学习了微服务体系下的相关知识，收获颇丰。
