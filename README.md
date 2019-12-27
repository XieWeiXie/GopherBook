# GopherBook


[![Build Status](https://travis-ci.com/wuxiaoxiaoshen/GopherBook.svg?token=NJwtDqGPUSoHiysBfFqE&branch=master)](https://travis-ci.com/wuxiaoxiaoshen/GopherBook)

:fire:  :fire:  :fire:  :fire:  :fire:  :fire:



> 定位： 以示例式的编程指南书。

---
:fire:  :fire:  :fire:  :fire:  :fire:  :fire:

章节|内容|进度
:---|:---|:---|
第一章|环境搭建|100&permil;|
第二章|语言特性|100&permil;|
第三章|数据操作|100&permil;|
第四章|内置库常用操作|100&permil;|
第五章|编写图表库|1000&permil;|
第六章| 编写测试|100&permil;|
第七章| Go 爬虫|100&permil;|
第八章|自己动手实现命令行工具|100&permil;|
第九章|自己动手实现一个库|100&permil;|
第十章|web服务|100&permil;|
第十一章| GoWeb 开发手册|100&permil;|
第十二章| GoGraphQl 开发手册|100&permil;|
第十三章| 面向接口编程|100&permil;|
附录| 学习路径 |100&permil;|


:fire:  :fire:  :fire:  :fire:  :fire:  :fire:

**编程清单**
1. 明确需求
2. 合理组织项目结构
3. 全部使用 Go Module 进行代码管理
3. 优先使用内置库
4. 其次选择优秀第三方库
5. Git 分布式代码管理：其一分支、其二Commit Message 需规范
5. 单元测试保障代码质量，尤其注意关键核心代码
6. 项目必须有构建文件: Makefile
7. 项目容器化：Dockerfile 
8. 项目必须有 CI/CD 工具，比如：GithubActions, Travis
9. 简单项目：docker 容器启动，稍复杂 docker-compose 多容器部署，多节点一律使用 k8s，且部署脚本（yaml) 需托管
10. 项目必须有监控：prometheus 配合 grafana

**编程范式**

- 一律使用面向对象编程思想，时刻谨记面向对象四大特性：抽象、封装、多态、继承，谨慎使用面向过程编程思想
- 面向接口、组合编程，维护系统可拓展性
- 接口命名需抽象化、接口实现者命名需具体化
- 对外暴露信息需尽可能的少
- 合理使用： SOLID 思想：S: 职责单一、O: 对扩展开放，对修改关闭、L: 里式替换 I: 实现者不管理接口 D: 控制反转

**面向提升编程**

- 规范每个项目，即使是 Side Project
- 根据个人水平，攻克优秀第三方库源码阅读，学习实现思想、设计思想、分层架构
- 总结：定期汇总所学所得，文章组织出来，尽可能提升自己的处理问题的难度和复杂度，提升竞争力

**系统工程**

- 编程不是孤立的使用编程语言实现项目需求，靠的是合理的使用多种技术，支持业务需求，仍然需要持续不断的攻克后端相关技术

```text
SQL
MySQL/PostgresSQL
ElasticSearch
Redis
Docker/Docker-compose
Kubernetes
Linux
Kafka/RocketMQ
```

- 技术选型：1. 明确场景和需求 2. 根据场景和需求选择待选方案 3. 根据社区活跃度、参与度 从待选方案中选择其一

---

&copy;XieWei&trade;




