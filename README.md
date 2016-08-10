# blog
框架版本:  
[tinygo 1.0 beta](https://github.com/kdada/tinygo/tree/dev1.0)  
  
目录结构:  
```
main.go             main
web.cfg             tinygo web配置文件
|-app
  |-favicon.ico     图标文件
  |-robots.txt      搜索引擎文件
  |-css             css文件目录
  |-js              js文件目录
  |-src             ts源码文件,使用tsc编译到js目录
  |-tmpl            angular2的模板文件目录
  |-views           tinygo框架视图模板文件目录
    |-layout.json   tinygo框架视图布局文件
    |-layout        布局模板文件目录
      |-layout.html
    |-home          视图模板文件目录
      |-index.html
|-controllers       控制器源码
  |-home.go
|-models            数据模型源码
  |-model.go
|-services          服务源码
  |-service.go
```
  

database:   PostgreSQL  
golang pkg: [github.com/lib/pq](https://github.com/lib/pq)

数据库创建:[./doc/once.sql](doc/once.sql),[./doc/blog.sql](doc/blog.sql)  
接口文档:[./doc/api.md](doc/api.md)

### blog项目用途:  
1. 测试tinygo 1.0框架的基础功能  
    * 依赖注入(同时可按类型和名称进行注入)
    * 路由功能
    * Session功能
    * 基础SQL功能
    * 模板多重布局功能
    * 日志功能
    * 基于表达式的字符串验证功能
    * Web事件功能
2. 测试与PostgreSQL的兼容性
3. 测试与Angular2的兼容性
4. 实现一个基于Markdown的博客