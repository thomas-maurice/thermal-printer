# Thermal printer controller

Small gRPC service to be able to control a thermal printer over the network

udev rules:
```
    SUBSYSTEM=="usb", ATTR{idVendor}=="0416", ATTR{idProduct}=="5011", MODE="666"
```
