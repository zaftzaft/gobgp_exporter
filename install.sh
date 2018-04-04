go build

service=gobgp_exporter

systemctl is-enabled $service
if [ $? -eq 0 ]; then
  systemctl stop $service
fi


mv ./gobgp_exporter /usr/local/bin/gobgp_exporter

tee /usr/lib/systemd/system/${service}.service << EOS
[Unit]
Description=GoBGP Exporter

[Service]
Restart=always
ExecStart=/usr/local/bin/gobgp_exporter

[Install]
WantedBy=default.target
EOS


systemctl daemon-reload
systemctl enable $service
systemctl start $service



