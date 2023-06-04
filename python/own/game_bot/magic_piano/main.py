from pyautogui import *
import pyautogui
import time
import keyboard
import random
import win32api, win32con


def click(x, y):
    win32api.SetCursorPos((x, y))
    win32api.mouse_event(win32con.MOUSEEVENTF_LEFTDOWN, 0, 0)
    time.sleep(0.1)  # This pauses the script for 0.1 seconds
    win32api.mouse_event(win32con.MOUSEEVENTF_LEFTUP, 0, 0)
    print("Clicked ", x, y)


while not keyboard.is_pressed('q'):

    if pyautogui.pixel(750, 400)[0] == 0:
        click(750, 400)
    if pyautogui.pixel(890, 400)[0] == 0:
        click(890, 400)
    if pyautogui.pixel(1050, 400)[0] == 0:
        click(1050, 400)
    if pyautogui.pixel(1140, 400)[0] == 0:
        click(1140, 400)