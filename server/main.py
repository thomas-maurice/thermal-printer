from concurrent import futures
import time
import math
import os
import logging
import functools
from escpos.printer import Usb
from usb.core import USBError

import grpc

import api_pb2_grpc
import api_pb2

from google.protobuf import empty_pb2

def printer_guard(function):
    @functools.wraps(function)
    def wrapper(*args, **kwargs):
        try:
            return function(*args, **kwargs)
        except USBError:
            # probably failed because the cable was disconnected or something
            # then we just suicide the server altogether
            logging.error("Shit broke, suiciding the server")
            os._exit(1)
    return wrapper

class PrintServicer(api_pb2_grpc.PrintServiceServicer):
    def __init__(self):
        self.printer = Usb(0x0416, 0x5011, profile="POS-5890")

    @printer_guard
    def Print(self, request, context):
        self.printer.set(font=request.font)
        self.printer.text(request.line)
        return empty_pb2.Empty()
    
    @printer_guard
    def Blank(self, request, context):
        self.printer.set(font=request.font)
        for i in range(0, request.number):
            self.printer.text("\n")
        return empty_pb2.Empty()
    
    @printer_guard
    def QR(self, request, context):
        self.printer.set(font=1)
        self.printer.qr(request.code, size=request.pixel_size, native=False, center=True)
        return empty_pb2.Empty()
    
    @printer_guard
    def Bar(self, request, context):
        self.printer.soft_barcode(
            barcode_type='ean13',
            data=request.code,
            center=request.center,
        )
        self.printer.set(font=1)
        for i in range(0, request.blanks):
            self.printer.text("\n")

        return empty_pb2.Empty()

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=1))
    api_pb2_grpc.add_PrintServiceServicer_to_server(
        PrintServicer(), server)
    server.add_insecure_port('[::]:8069')
    server.start()
    server.wait_for_termination()


if __name__ == '__main__':
    logging.basicConfig()
    serve()