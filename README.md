<h1><img alt="logo" src="docs/img/logo.png" height="30px"> DailyCards</h1>

> 自建每日打卡小工具 [Demo](https://card.linkinstars.com/)

## Feature
> 核心只有两个功能 Less Is More 

### 打卡
通过书写 markdown 来打卡每日记录

### 分享
将打卡内容以 **卡片图片** 或者 **链接** 的方式分享到任意地方

## Preview
![card](/docs/img/card.png)

## Why?
之前就一直在打卡，只不过一直都是最简单文字描述，然后打开在微信群里面，会有几个问题：
1. 不利于分享到其他地方（不好看）
2. 不利于归档和回顾

使用其他小程序或者 APP 的话，也会有几个问题：
1. 分享之后会带有使用成本，必须用小程序才能查看
2. 很多工具能实现但是有很多其他额外的功能，负担有点重
3. 数据不在自己这里（隐私党不爽）

之前看到一个网站 [poet.so](https://poet.so/) 可以将推特的内容分享成一张好看的卡片，于是就有了做这个的想法。

好看样式的卡片会极大的增加打卡和分享的欲望，更容易使人坚持。

## Install
下载对应操作系统的二进制文件

- 使用 `-init` 初始化配置文件
- 使用 `-c` 指定目录，如果不指定则默认会在当前用户目录下创建如 `/root/daily-cards/`

```bash
# 步骤 1：初始化配置文件
./dc -init -c /tmp/dc/

# 步骤 2：启动
./dc -c /tmp/dc/
```

## Usage
- 管理员访问 http://127.0.0.1:8080/card/uc/login
- 用户访问 http://127.0.0.1:8080/card/page

## Config
> 如果没有指定目录，默认配置文件生成在 `/root/daily-cards/config.yaml`，修改配置后重启才生效

- WebPort: 端口
- Username: 用户名，用于管理员登录，默认 admin
- Password: 密码，用于管理员登录，默认 admin
- Nickname: 显示昵称
- Avatar: 显示头像
- Secret: 秘钥，用于加密jwt
- SiteName: 网站名称
- ShowQrcode: 是否展示二维码

## Notice
- 数据存储在 `daily-cards.db` 文件中，所以你只需要备份这个文件数据就不会丢
- 没有适配图片，虽然你可以写，但可能卡片展示会有问题，个人尝试了一下，不建议在卡片内放图片，对于打卡来说文字足以

## Dependence
- Golang + Vue3 
- sqlite3
