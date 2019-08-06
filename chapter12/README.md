# Chapter12

> 第18届国际泳联世界锦标赛

> HOST: https://www.fina-gwangju2019.com



## 开发思路

世锦赛关注点，最重要的是什么？

- 历史
    - 历届：时间、主办方
    - 简介
- 各国奖牌排名
    - 金牌
    - 银牌
    - 铜牌
- 项目
    - 简介
    - 历史
- 各项目获奖得主
    - 国家
    - 名称
    - 奖项
    - 项目

## 使用

pwd=GopherBook/chapter12/fina

- 启动容器
```
cd deployments && docker-compose up -d
```

- 导入数据

```text

cd data
docker cp fina-2019-08-06.sql ${imageID}:/

mysqldump -uroot -p < fina-2019-08-06.sql

```

- 启动服务

```text
make install

make run

```

- 调用接口

接口文档：schema.graphql