#!/usr/bin/env bash

name=twistlock_authz_broker
cp ./twistlock_authz_broker /usr/bin/

cat <<SERVICE > "/lib/systemd/system/twistlock-authz.service"
[Unit]
Description=Twistlock docker authorization plugin
After=syslog.target
[Service]
Type=simple
ExecStart=/usr/bin/authz-broker
[Install]
WantedBy=multi-user.target
SERVICE

# sudo systemctl enable my-app
# sudo systemctl start my-app
