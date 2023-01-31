IPN - IP变更通知工具
---
IPN在运行时会尝试获取您的IP，若与上次运行时IP不一致，则会向您发送邮件通知。

## 特性

- 获取公网IP并发送至指定邮箱
- 能够运行在包括路由器在内的大多数设备中
- 避免因使用DDNS而被运营商约谈
- 持续更新更多特性

## 使用说明

1. [下载](https://github.com/devzhi/ipn/releases)对应平台的可执行文件和配置文件`config.toml`
2. 配置定时任务，Linux下使用crond，Windows下使用计划任务，Padavan等路由器系统可直接在后台设置，弱上述条件均不具备，您也可以修改配置文件，启用内置定时任务。
3. Enjoy it.

## 配置文件说明

请查看注释。

```toml
# 设备名称
Device = "test device"
# SMTP服务器地址
SmtpHost = "smtp.qq.com"
# SMTP服务器端口
SmtpPort = 587
# 发件人地址
From = "xxxx@qq.com"
# 发件人密码（部分供应商使用授权码，如QQ邮箱）
Password = "xxxxx"
# 收件人地址
To = ["xxx@qq.com"]
# 启用内置cron
EnableCron = false
# Cron表达式
CronExp = "*/5 * * * *"
```

## 贡献&反馈

请直接发起PR或提交Issue即可。

## 开源协议

[MIT License](https://github.com/devzhi/ipn/blob/main/LICENSE)