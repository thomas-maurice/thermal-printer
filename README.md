# Thermal printer controller

Small gRPC service to be able to control a thermal printer over the network

You can make it install itself and create a systemd unit file & udev rules like so
```bash
$ sudo python3 ./install.py --installSystemd=True --installUdev=True --user=printer
$ sudo systemctl daemon-reload
$ sudo systemctl restart udev
# optionally
$ sudo systemctl start escpos-server
$ sudo systemctl enable escpos-server
```

udev rules:
```
    SUBSYSTEM=="usb", ATTR{idVendor}=="0416", ATTR{idProduct}=="5011", MODE="666"
```
