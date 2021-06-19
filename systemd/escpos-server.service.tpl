[Unit]
Description=ESC-POS printer service
After=network.target
StartLimitIntervalSec=5

[Service]
Type=simple
Restart=always
RestartSec=5
User=printer
WorkingDirectory={{ installDir }}/server
ExecStart={{ installDir }}/.venv/bin/python3 {{ installDir }}/server/escpos-server

[Install]
WantedBy=multi-user.target
