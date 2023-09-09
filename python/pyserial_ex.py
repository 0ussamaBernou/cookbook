# pip install pyserial
import serial
from sys import argv

dev = serial.Serial(argv[1])

if len(argv) > 2:
    # binary_str = "".join(
    #     format(byte, "08b") for byte in bytearray(argv[2], encoding="utf-8")
    # )
    encoded = argv[2].encode("utf-8")
    dev.write(encoded)

while True:
    print(dev.readline())
