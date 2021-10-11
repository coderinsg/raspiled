#!/usr/bin/env python

from gpiozero import LED
from time import sleep
import os
import atexit

greenled = LED(14)
amberled = LED(15)
redled = LED(18)
delay = float(os.environ['LED_DELAY'])

def exit_handler():
  print 'Exiting traffic light program.'
  print 'Clearing all lights ...'
  greenled.off()
  amberled.off()
  redled.off()

atexit.register(exit_handler)

while True:
  greenled.on()
  sleep(delay)
  greenled.off()

  amberled.on()
  sleep(delay)
  amberled.off()

  redled.on()
  sleep(delay)
  redled.off()
