# Site Alarm Service

# Install as service

Create `/usr/lib/systemd/user/notifier.service` file with content.

```
[Unit]
Description=Site Notifier
[Service]
Type=simple
ExecStart=/data/notifier/alarmservice --c=/data/notifier/config.yml
[Install]
WantedBy=multi-user.target
```

Run to enable service
`systemctl enable /usr/lib/systemd/user/notifier.service`

Run to start service
`systemctl start /usr/lib/systemd/user/notifier.service`
