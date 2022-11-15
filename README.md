## 介绍

<div align=center>
<img src="https://kevn.gitee.io/gsadmindoc/img/gslogo.png" width=20%><br/>
<div style="padding: 20px 0"><a href="https://kevn.gitee.io/gsadmindoc">文档</a> | <a href="https://gsadmin.suiyidian.cn/">演示</a></div>

   <a href='https://gitee.com/kevn/gsadmin' target="_blank">
				<img src='https://gitee.com/kevn/gsadmin/badge/star.svg?theme=dark' alt='star' style="vertical-align: middle">
			</a>
			<a href='https://github.com/sonhineboy/gsadmin' target="_blank">
				<img src='https://img.shields.io/github/stars/sonhineboy/gsadmin.svg?logo=GitHub' alt='star' style="vertical-align: middle">
			</a>
<a href='https://gin-gonic.com/' target="_blank">
				<img src='https://img.shields.io/badge/go-gin-blue?logo=go' alt='star' style="vertical-align: middle">
			</a>
   <a href='https://lolicode.gitee.io/scui-doc/' target="_blank">
				<img src='https://img.shields.io/badge/vue-scui-yellow' alt='star' style="vertical-align: middle">
			</a>
</div>


GS Admin=gin+scui 它是golang 开发的一个企业级后台。遵循MIT开源协议。前端框架是scui,SCUI基于 Vue3、elementPlus 持续性的提供独家组件和丰富的业务模板帮助你快速搭建企业级中后台前端任务。后端框架是gin,Gin是一个golang的微框架，封装比较优雅，具有快速灵活，容错方便等特点。内置了权限管理、用户管理等基础模块儿，还支持了事件服务，方便业务解耦。后续会根据用户的反馈更新内容！

## 快速开始

<b>一共分三步</b>：   
1、 拉代码  
2、 部署后端服务  
3、 部署前端代码

```sh
#第一步
git clone https://gitee.com/kevn/gsadmin.git

#第二步
#服务端

cd {项目目录}/service && go mod tidy

#配置配置文件，config.yaml
#初始化数据 {项目目录}/service/databases/*.sql 执行里面的sql
go run main.go

#第三步
#web端
cd {项目目录}/web/scui && npm install

# 启动项目(开发模式)
npm run serve

```

> 演示账号密码：test/123456  
> 注意：goland 的环境必须配置好，数据不要忘记初始化

## 加入贡献

GS Admin 是一个开源项目，一个开源项目的发展离不开开源社区的力量支持，如果您希望参与 GS Admin 的开发，可以先从 [issues](https://gitee.com/kevn/gsadmin/issues) 开始，通常来说会有以下的一些步骤：

- 1.关注 [issues](https://gitee.com/kevn/gsadmin/issues) 的动态，评论回复帮助提出疑问的用户；
- 2.根据 [issues](https://gitee.com/kevn/gsadmin/issues) 的内容，找寻根据自己当前对项目的了解程度，去修复力所能及的 BUG 或实现功能，并以 Pull Request 的形式提交至 [kevn/gsadmin](https://gitee.com/kevn/gsadmin) 仓库；
- 3.关注自己提交 Pull Request 的进度和状态，以推动您的 Pull Request 尽快合入主仓库；
- 4.对其他人提交的 Pull Request 进行 Code Review，并给出您的建议和看法。
- 5.坚持并持续进行上述步骤。  

## 交流方式
<b>QQ交流群</b>：584407821 (1群)

<b>微信群</b>：加个人微信后，拉入  
备注：<u>gsadmin 无备注不通过</u>

<img src="https://kevn.gitee.io/gsadmindoc/img/wx.jpg" width="50%">

