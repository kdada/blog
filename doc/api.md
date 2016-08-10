# API SPEC


## API 总览:  
```
[可选的] 
(:参数名)   
```
 URL | 方法 | 仅管理员可用 | 描述
---|:-:|:-:|---
/ | GET | false | 首页
/index.html | GET | false | 首页,等同于/
/c(:Id)[/p(:Page).html] | GET | false | 分类页面
/a(:Id).html | GET | false | 文章页面
/img/(:Name).jpg | GET | false | 获取图片
/manager | GET | true | 管理页面,仅限管理员访问
/user/login | POST | false | 用户登录
/user/register | POST | false | 用户注册
/img/upload | POST | true | 上传图片
/category/new | POST | true | 创建分类
/category/rename | POST | true | 修改分类名称
/category/hide | POST | true | 隐藏分类
/category/show | POST | true | 显示分类
/category/delete | POST | true | 删除分类
/category/list | POST | true | 查看分类列表
/article/view | POST | true | 查看文章
/article/new | POST | true | 创建文章
/article/modify | POST | true | 修改文章
/article/hide | POST | true | 隐藏文章
/article/show | POST | true | 显示文章
/article/delete | POST | true | 删除文章
/article/top | POST | true | 置顶文章
/article/list | POST | true | 查看分类下的文章列表
/reply/new | POST | false | 回复
/reply/list | POST | false | 获取回复列表
  
  
  
  
### URL: / 或 /index.html  
方法: GET  
描述: 显示首页内容  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---

正确返回:  
```
首页html页面
```
错误返回:  
```
404 not found
```
  
  
  
### URL: /c(:Id)[/p(:Page).html]  
方法: GET  
描述: 分类页面  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 分类id
Page | int | false | 1 | 页码,如果不传该参数则默认为0

正确返回:  
```
显示id分类下的第page页
```
错误返回:  
```
404 not found
```
  
  
  
### URL: /a(:Id).html  
方法: GET  
描述: 文章页面  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
显示id对应的文章内容
```
错误返回:  
```
404 not found
```
  
  
  
### URL: /img/(:Name).jpg  
方法: GET  
描述: 图片  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Name | string | true | logo | 图片名称

正确返回:  
```
指定名称的图片
```
错误返回:  
```
404 not found
```
  
  
  
### URL: /manager  
方法: GET  
描述: 博客管理后台,只有处于登录状态的管理员才能访问  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---

正确返回:  
```
管理后台首页
```
错误返回:  
```
404 not found
```
  
  
  
### URL: /user/login  
方法: POST  
描述: 用户登录接口  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Email | string | true | test@test.com | 用户邮箱,必须符合邮箱格式
Password | string | true | 123456 | 用户密码,最短6个字符,最长20个字符

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //用户id
        Name:   string  //用户昵称
        Email:  string  //用户邮箱
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /user/register  
方法: POST  
描述: 用户注册接口  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Email | string | true | test@test.com | 用户邮箱,必须符合邮箱格式
Name | string | true | test | 用户昵称,最短2个字符,最长10个字符
Password | string | true | 123456 | 用户密码,最短6个字符,最长20个字符

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //用户id
        Name:   string  //用户昵称
        Email:  string  //用户邮箱
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /img/upload
方法: POST  
描述: 上传图片接口  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Name | string | true | test | 图片名称,只能包含英文字母和数字,最短2个字符,最长10个字符
Image | file | true | 文件 | 图片文件

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Name:   string  //图片名称
        Path:   string  //图片地址
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/new
方法: POST  
描述: 创建分类  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Name | string | true | 测试分类 | 分类名称,最短2个字符,最长10个字符

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //分类id
        Name:   string  //分类名称
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/rename
方法: POST  
描述: 重命名分类  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Name | string | true | 测试分类 | 新的分类名称,最短2个字符,最长10个字符

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //分类id
        Name:   string  //分类名称
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/hide
方法: POST  
描述: 隐藏分类,隐藏分类后,除管理员以外不可访问分类下所有文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 分类id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //分类id
        Name:   string  //分类名称
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/show
方法: POST  
描述: 显示分类  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 分类id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //分类id
        Name:   string  //分类名称
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/delete
方法: POST  
描述: 删除分类,删除后分类下的文章均不可访问  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 分类id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //分类id
        Name:   string  //分类名称
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /category/list
方法: POST  
描述: 查看分类列表  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: [{
        Id:     int     //分类id
        Name:   string  //分类名称
    }]
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/view
方法: POST  
描述: 创建文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:         int     //文章id
        Title:      string  //文章标题
        Content:    string  //文章内容
        CreateTime: string  //创建时间
        UpdateTime: string  //上次更新时间
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/new
方法: POST  
描述: 创建文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Title | string | true | 文章标题 | 文章标题,不可超过100个字符
Content | string | true | 文章内容 | 文章内容

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/modify
方法: POST  
描述: 修改文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id
Title | string | true | 文章标题 | 文章标题,不可超过100个字符
Content | string | false | 文章内容 | 文章内容,如果内容没有变更则可以不传该项

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/hide
方法: POST  
描述: 隐藏文章,隐藏后只有管理员能够查看  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/show
方法: POST  
描述: 显示文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/delete
方法: POST  
描述: 删除文章,删除后不可查看  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /article/top
方法: POST  
描述: 置顶文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Id | int | true | 1 | 文章id

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //文章id
        Title:  string  //文章标题
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```

  
  
  
### URL: /article/list
方法: POST  
描述: 文章列表  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Category | int | true | 1 | 分类id
Page | int | true | 0 | 页码

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        TotlePage: int  //总页数
        Articles:[{
            Id:         int     //文章id
            Title:      string  //文章标题
            UpdateTime: string  //上次更新时间
        }]
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
  
  
  
### URL: /reply/new
方法: POST  
描述: 创建回复  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Article | int | true | 1 | 文章id
Reply | int | false | 2 | 回复的回复id
Content | string | true | 回复内容 | 回复内容

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        Id:     int     //回复id
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```

  
  
  
### URL: /reply/list
方法: POST  
描述: 置顶文章  
传入参数:   

字段 | 类型 | 必须 | 例子 | 描述
---|---|---|---|---
Article | int | true | 1 | 文章id
Page | int | true | 1 | 页码,每页10条

正确返回:  
```
{
    Code:   int     //状态码,为0
    Data: {
        TotlePage: int //总页数
        Replies:[{
            Id:         int     //回复id
            Reply:      int     //回复的回复id
            Name:       string  //回复的回复人昵称
            CreateTime: string  //回复时间
        }]
    }
}
```
错误返回:  
```
{
    Code:   int     //状态码,不为0
    Message:string  //错误信息
}
```
