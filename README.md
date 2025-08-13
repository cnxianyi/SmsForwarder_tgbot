# SmsForwarder_tgbot

## 效果预览

![md](https://tncache1-f1.v3mh.com/image/2025/08/13/7b3eab9d9a128b846843a2feefe2e7d3.png)

## 使用前提

1. 已经创建TG bot 并获取到 token
2. 已经开启 SmsForwarder 的 http 服务
3. 已经开启 frps 服务.(调试可以不开)

## 使用方法

1. 下载对应的二进制文件到服务器上
2. 在二进制文件夹所在位置创建一个 .env 文件

```env
TELEGRAM_BOT_TOKEN="填入token"

SIGN="填入签名" # 只支持校验签名和不签名

BASE_URL="填入frps后的url. 如http://192.168.1.2:5000"
```

3. 部署

### 方法一：直接运行
```bash
# 给二进制文件添加执行权限
chmod +x SmsForwarder_tgbot-linux-amd64

# 直接运行
./SmsForwarder_tgbot-linux-amd64
```

### 方法二：使用 systemctl 服务部署（推荐）

#### 步骤1：创建服务文件
```bash
sudo nano /etc/systemd/system/smsforwarder-tgbot.service
```

#### 步骤2：添加服务配置
```ini
[Unit]
Description=SMS Forwarder Telegram Bot
After=network.target
Wants=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/smsforwarder-tgbot
ExecStart=/opt/smsforwarder-tgbot/SmsForwarder_tgbot-linux-amd64
Restart=always
RestartSec=10
StandardOutput=journal
StandardError=journal
SyslogIdentifier=smsforwarder-tgbot

[Install]
WantedBy=multi-user.target
```

#### 步骤3：创建应用目录并部署
```bash
# 创建应用目录
sudo mkdir -p /opt/smsforwarder-tgbot

# 复制二进制文件
sudo cp SmsForwarder_tgbot-linux-amd64 /opt/smsforwarder-tgbot/

# 环境配置文件
sudo vim /opt/smsforwarder-tgbot/.env

# 设置权限
sudo chmod +x /opt/smsforwarder-tgbot/SmsForwarder_tgbot-linux-amd64
sudo chown -R root:root /opt/smsforwarder-tgbot
```

#### 步骤4：启动和管理服务
```bash
# 重新加载systemd配置
sudo systemctl daemon-reload

# 启用服务（开机自启）
sudo systemctl enable smsforwarder-tgbot

# 启动服务
sudo systemctl start smsforwarder-tgbot

# 查看服务状态
sudo systemctl status smsforwarder-tgbot

# 查看日志
sudo journalctl -u smsforwarder-tgbot -f

# 停止服务
sudo systemctl stop smsforwarder-tgbot

# 重启服务
sudo systemctl restart smsforwarder-tgbot
```