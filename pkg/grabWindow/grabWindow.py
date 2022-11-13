import sys
import time

import win32gui
import winsound
from PyQt5.QtWidgets import QApplication

# pip install PyQt5 -i https://pypi.tuna.tsinghua.edu.cn/simple/
# pip install pypiwin32 -i https://pypi.tuna.tsinghua.edu.cn/simple/

window_name_list = dict()  # 创建字典保存窗口的句柄与名称映射关系
app = QApplication(sys.argv)


# 获得所有窗口的名字
def get_all_hwnd(window_name, mouse):
    if win32gui.IsWindow(window_name) and win32gui.IsWindowEnabled(window_name) and win32gui.IsWindowVisible(
            window_name):
        window_name_list.update({window_name: win32gui.GetWindowText(window_name)})


# 可以异步播放声音
def play_music():
    winsound.PlaySound('alert', winsound.SND_ASYNC)


def alert(name, xPoint, xLength, yPoint, yLength, color: str):
    # 获得所有窗口的名字
    # win32gui.EnumWindows(get_all_hwnd, 0)
    # for h, t in window_name_list.items():
    #     if t != "":
    #         print(h, t)

    # 必要的代码

    # img.pixelColor(0,0).getRgb()
    hwnd1 = win32gui.FindWindow(None, name)
    screen = QApplication.primaryScreen()
    img = screen.grabWindow(hwnd1, xPoint, yPoint, xLength, yLength).toImage()
    for x in range(1, xLength):
        for y in range(1, yLength):
            rgb = img.pixelColor(x, y).getRgb()
            red = rgb[0]
            green = rgb[1]
            blue = rgb[2]
            if color == "red":
                if red > green and red > blue:
                    return True
            if color == "green":
                if green > red and green > blue:
                    return True
    return False


while 1:
    msg = alert('', 0, 0, 0, 0, "red")
    msg = msg or alert('', 0, 0, 0, 0, "red")
    if msg:
        play_music()
    time.sleep(1)
