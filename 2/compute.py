import os
import re
import functools

RED = 12
GREEN = 13
BLUE = 14
if __name__ == "__main__":
    power_sets = []
    with open("input.txt") as f:
        # for each game
        while (game_text := f.readline()):
            game = re.split(r':', game_text)[0]
            BLUE = 0
            RED = 0
            GREEN = 0

            reds = re.findall(r'\d+ red', game_text)
            for i in reds:
                inst = int(re.search(r'\d+', i).group(0))
                if inst > RED:
                    RED = inst

            greens = re.findall(r'\d+ green', game_text)
            for i in greens:
                inst = int(re.search(r'\d+', i).group(0))
                if inst > GREEN:
                    GREEN = inst

            blues = re.findall(r'\d+ blue', game_text)
            for i in blues:
                inst = int(re.search(r'\d+', i).group(0))
                if inst > BLUE:
                    BLUE = inst

            power = BLUE * RED * GREEN
            power_sets.append(power)

    print((functools.reduce(lambda a, b: a + b, power_sets)))

