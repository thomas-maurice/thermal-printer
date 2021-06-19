#!/usr/bin/env python3

import subprocess
import logging
import argparse
import time
import sys
import os

parser = argparse.ArgumentParser(description='Installs the esc-pos server')
parser.add_argument('--installSystemd', type=bool, help="installs the systemd unit file", default=False)
parser.add_argument('--installUdev', type=bool, help="installs the udev rules file", default=False)
parser.add_argument('--user', type=str, help="which user to run as", default="printer")
args = parser.parse_args()

logging.basicConfig(level=logging.INFO)

installDir = os.path.dirname(os.path.realpath(__file__))

logging.info("current directory {}".format(installDir))
logging.info("install systemd config: {}".format(args.installSystemd))
logging.info("create the virtual env at {}/.venv/".format(installDir))

process = subprocess.Popen(["python3", "-m", "venv", ".venv"], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
for c in iter(lambda: process.stdout.read(1), b''):
    sys.stdout.buffer.write(c)

process.wait()

sys.stdout.flush()

if process.returncode != 0:
    logging.error("failed to create the virtual env")
    sys.exit(process.returncode)

process = subprocess.Popen(["{}/.venv/bin/pip3".format(installDir), "install", "-r", "{}/requirements.txt".format(installDir)], stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
for c in iter(lambda: process.stdout.read(1), b''):
    sys.stdout.buffer.write(c)

process.wait()

if process.returncode != 0:
    logging.error("failed to install the packages")
    sys.exit(process.returncode)

sys.stdout.flush()

with open(installDir + "/systemd/escpos-server.service.tpl", "r") as tmplFile:
    tmpl = tmplFile.read()
    rendered = tmpl.replace("{{ installDir }}", installDir)
    rendered = rendered.replace("{{ user }}", args.user)
    print("\n")
    logging.info("generated systemd config file to install to /lib/systemd/system/escpos-server.service\n")
    print("8<-" * 30, "\n")
    print(rendered)
    print("8<-" * 30)
    if args.installSystemd:
        with open("/lib/systemd/system/escpos-server.service", "wb") as unitFile:
            unitFile.write(rendered)
            logging.info("written unit file successfully")

if args.installUdev:
    with open("/etc/udev/rules.d/80-escpos-server-printer.rules", "wb") as unitFile:
        unitFile.write("SUBSYSTEM==\"usb\", ATTR{idVendor}==\"0416\", ATTR{idProduct}==\"5011\", MODE=\"666\"\n")
        logging.info("written udev file successfully")
