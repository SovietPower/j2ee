# 后端

项目前端地址：https://github.com/SovietPower/j2ee-web

项目后端地址：https://github.com/SovietPower/j2ee

## 项目依赖

- Gin
- Gorm
- mysql
- godotenv
- jwt-go

## 目录结构
```
├─api 服务调用接口
├─config
│  └─locales i18n国际化文件（未使用）
├─constant 常量及相关信息
├─middleware 中间件
├─model 数据库模型
├─router 路由
├─runtime
│  └─logs 日志文件
├─serializer 返回给前端数据的序列化器
├─service 处理服务
└─util 工具
    └─logging 日志工具
```

## 运行

需先创建`.env`文件，填写数据库、密钥等信息，内容格式同`.env.example`。

```
go mod tidy
go run main.go
```

运行在`8080`端口。
