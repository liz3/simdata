from ctypes import c_float, byref
import time
import sdl2
import sdl2.ext
import sdl2.sdlimage
from datetime import datetime
import random
import time
import json
import sys

def parseData(raw):
    split = raw.strip().split('----')
    dataRaw = split[2].strip()
    game = ""
    lines = split[1].split("\n")
    for x in lines:
        if x.startswith("game "):
            game = x[5:]
            break
    data = json.loads('[' + dataRaw[:len(dataRaw)-1] + ']')
    return (split[1], data,game)

# rendering helper for
def draw_steering_state(renderer, x,y,thinkness, value):
    pixVal = int(value * 200)
    if pixVal > 1000:
        return
    for targetY in range(y-thinkness//2,y+(thinkness//2)):
        for targetX in range(x-204,x-200):
            renderer.draw_point([targetX,targetY], sdl2.ext.Color(70,70,70))
    for targetY in range(y-thinkness//2,y+(thinkness//2)):
        for targetX in range(x+200,x+204):
            renderer.draw_point([targetX,targetY], sdl2.ext.Color(70,70,70))
    for targetY in range(y-thinkness//2,y+(thinkness//2)):
        if pixVal < 0:
            for targetX in range(x+pixVal,x):
                renderer.draw_point([targetX,targetY], sdl2.ext.Color(255,0,0))
        else:
            for targetX in range(x,x+pixVal):
                renderer.draw_point([targetX,targetY], sdl2.ext.Color(255,0,0))


def draw_line_aggrate(renderer, x,y,height,thinkness, value):
    pixValue = int((((value) * (height // 2)) - height//2)) * -1
    for targetY in range(y+(height-pixValue), y+height):
        for targetX in range(x-(thinkness //2), x+(thinkness//2)):
            renderer.draw_point([targetX,targetY], sdl2.ext.Color(0,0,255))
def plot_entry(entry, start, baseOffset, renderer,b):
    renderer.draw_point([start, baseOffset + int(entry["steering"] * 50)],  sdl2.ext.Color(255,255,255) if b else sdl2.ext.Color(255,0,0))
    renderer.draw_point([start, baseOffset + (250) + -int(entry["speed"] * 0.5)],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(0,120,255))
    renderer.draw_point([start, baseOffset + (400) +  -int(entry["throttle"] * 50)],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(120,0,255))
    renderer.draw_point([start, baseOffset + (600) +  -int(entry["brake"] * 50)],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(0,0,255))
    renderer.draw_point([start, baseOffset + (800) +  -( entry["gear"] ) * 10],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(0,255,0))

def plot_entry_pc2(entry, start, baseOffset, renderer, b):
    renderer.draw_point([start, baseOffset + entry["steering"]],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(255,0,0))
    renderer.draw_point([start, baseOffset + (250) + -int(entry["speed"] * 1.6)],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(0,120,255))
    renderer.draw_point([start, baseOffset + (400) +  -int(entry["throttle"] * 0.5)],  sdl2.ext.Color(255,255,255) if b else sdl2.ext.Color(120,0,255))
    renderer.draw_point([start, baseOffset + (600) +  -int(entry["brake"] * 0.5)],  sdl2.ext.Color(255,255,255) if b else sdl2.ext.Color(0,0,255))
    renderer.draw_point([start, baseOffset + (800) +  -( entry["gear"] - 90) * 10],  sdl2.ext.Color(255,255,255) if b else  sdl2.ext.Color(0,255,0))

def compute_display(index, game, zoom):
    if game == "PC2":
        return index % zoom == 0
    elif game == "F12020":
        return index % zoom == 0
    return False

def main():
    content = open(sys.argv[1]).read()
    obj = parseData(content)
    parsed = obj[1]
    game = obj[2]
    sdl2.ext.init()
    window = sdl2.ext.Window("Graphs: " + game, size=(1920, 1080))
    renderer = sdl2.ext.Renderer(window)
    factory = sdl2.ext.SpriteFactory(renderer=renderer) #
    font_manager2 = sdl2.ext.FontManager(font_path = "Roboto-Regular.ttf", size = 18) # Creating Font Manager
    running = True
    state = "tt"
    lastMouseX = 0
    window.show()
    offset = 0
    zoom = 0
    if game == "PC2":
        zoom = 45
    elif game == "F12020":
        zoom = 2
    lap1 = 1
    lap2 = 2

    while running:
        events = sdl2.ext.get_events()
        for event in events:
            if event.type == sdl2.SDL_QUIT:
                running = False
                break
            if event.type == sdl2.SDL_KEYDOWN:
                code = event.key.keysym.scancode
                print(code)
                if code == 30:
                    lap1 += 1
                if code == 31:
                    lap1 -= 1
                if code == 32:
                    lap2 += 1
                if code == 33:
                    lap2 -= 1
                if code == 82:
                    zoom -= 2
                if code == 81:
                    zoom += 2
                if code == 80 and offset > 0:
                    offset -= 100
                if code == 79:
                    offset += 100
            if event.type == sdl2.SDL_MOUSEMOTION:
                motion = event.motion
                lastMouseX = motion.x

        renderer.clear()
        start = 20
        start2 = 20
        lapOffset1 = 0
        lapOffset2 = 0
        for index, entry in enumerate(parsed):
            if entry["current_lap"] == lap1:
                if lapOffset1 == 0:
                    lapOffset1 = index
                if compute_display(index, game, zoom) and index - lapOffset1 > offset:
                    if game == "F12020":
                        plot_entry(entry, start, 150, renderer, False)
                    elif game == "PC2":
                        plot_entry_pc2(entry, start, 150, renderer, False)
                    start += 1
            if entry["current_lap"] == lap2:
                if lapOffset2 == 0:
                    lapOffset2 = index
                if compute_display(index, game, zoom) and index - lapOffset2 > offset:
                    if game == "F12020":
                        plot_entry(entry, start2, 150, renderer, True)
                    elif game == "PC2":
                        plot_entry_pc2(entry, start2, 150, renderer, True)
                    start2 += 1
#        texture = factory.from_text(str(state), fontmanager=font_manager2)
        for y in range(1080):
            renderer.draw_point([lastMouseX, y], sdl2.ext.Color(255,0,0))
            renderer.draw_point([lastMouseX + 1, y], sdl2.ext.Color(255,0,0))
        renderer.present()

if __name__ == "__main__":
    main()
