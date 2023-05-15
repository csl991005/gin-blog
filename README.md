# 介绍
gin+vue 全栈制作一个博客。学习以及巩固关于 Vue 和 golang 的知识
# 目录结构
```
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  
│        
├─api         
├─config // 项目配置入口   
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面      
│  └─front  // 前端展示页面 
├─upload   
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码
    ├─admin             
    └─front
```

# 实现功能
* 简单的用户管理权限设置
* 用户密码加密存储
* 文章分类自定义
* 列表分页
* 图片上传七牛云
* JWT 认证
* 自定义日志功能
* 跨域 cors 设置
* 文章评论功能

# 技术栈
* golang
  * Gin web framework
  * gorm(v1 && v2)
  * jwt-go
  * scrypt
  * gin-contrib/cors
  * go-playground/validator/v10
  * go-ini
* JavaScript
  * vue
  * vue cli
  * Vite
  * vue router
  * ant design vue
  * vuetify
  * axios
  * tinymce
* MySQL version:8.0.21

# 目前进度
* 后端api ✅
* 后台管理页面 ✅
* 前台展示页面 ✅
