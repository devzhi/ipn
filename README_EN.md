IPN - IP Change Notification Tool
---

When running, IPN attempts to obtain your IP. If it differs from the IP of the last run, it will send you an email notification.

## Features

- Obtains public IP and sends it to specified email address
- Can run on most devices, including routers
- Avoids potential warnings from ISPs due to DDNS usage
- Continuously updating with more features

## Usage Instructions

1. [Download](https://github.com/devzhi/ipn/releases) the executable file for your platform and the `config.toml` configuration file
2. Set up a scheduled task: use crond on Linux, Task Scheduler on Windows, or directly configure in the background for router systems like Padavan. If none of these options are available, you can modify the configuration file to enable the built-in scheduler.
3. Enjoy it.

## Configuration File Explanation

Please refer to the comments.

```toml
# Device name
Device = "test device"
# SMTP server address
SmtpHost = "smtp.qq.com"
# SMTP server port
SmtpPort = 587
# Sender's email address
From = "xxxx@qq.com"
# Sender's password (some providers use authorization codes, e.g., QQ Mail)
Password = "xxxxx"
# Recipient's email address
To = ["xxx@qq.com"]
# Always send email
AlwaysSend = false
# Enable built-in cron
EnableCron = false
# Cron expression
CronExp = "*/5 * * * *"
```

## Contribution & Feedback

Please directly initiate a PR or submit an Issue.

## License

[MIT License](https://github.com/devzhi/ipn/blob/main/LICENSE)

