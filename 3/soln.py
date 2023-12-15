import os
import re
import functools

def isSymbol(c : str):
    if c == '.':
        return False
    elif c.isdigit():
        return False
    else:
        return True
        

array = []
with open('input.txt', 'r') as file:
    while line := file.readline():
        array.append(line.strip())
    

array = array[:-1]
lines = len(array)

numbers_array = []
# store the position of symbols (non digit, non period)
for line_no, line in enumerate(array):
    numbers = re.finditer(r'\d+', line)
    for num in numbers:
        numbers_array.append((line_no, num))

print(numbers_array[:2])

def check_left(start, end, line_no, array):
    try:
        return isSymbol(array[line_no][start - 1])
    except IndexError:
        return False
def check_right(start, end, line_no, array):
    try:
        return isSymbol(array[line_no][end])
    except IndexError:
        return False
def check_above(start, end, line_no, array):
    try:
        for c in array[line_no - 1][start:end]:
            if isSymbol(c):
                return True
    except IndexError:
        return False
def check_below(start, end, line_no, array):
    try:
        for c in array[line_no + 1][start:end]:
            if isSymbol(c):
                return True
    except IndexError:
        return False

def check_topleft(start, end, line_no, array):
    try:
        return isSymbol(array[line_no - 1][start - 1])
    except IndexError:
        return False
def check_topright(start, end, line_no, array):
    try:
        return isSymbol(array[line_no - 1][end])
    except IndexError:
        return False
def check_bottomleft(start, end, line_no, array):
    try:
        return isSymbol(array[line_no + 1][start - 1])
    except IndexError:
        return False
def check_bottomright(start, end, line_no, array):
    try:
        return isSymbol(array[line_no + 1][end])
    except IndexError:
        return False

real_numbers_array = []
for i in numbers_array:
    line_no = i[0]
    start = i[1].start()
    end = i[1].end()

    number = i[1][0]

    # 8 edge cases divisible into 3 categories
    if check_left(start, end, line_no, array) \
    or check_right(start, end, line_no, array) \
    or check_above(start, end, line_no, array) \
    or check_below(start, end, line_no, array) \
    or check_topleft(start, end, line_no, array) \
    or check_topright(start, end, line_no, array) \
    or check_bottomleft(start, end, line_no, array) \
    or check_bottomright(start, end, line_no, array):
        real_numbers_array.append(number)


real_numbers_array = list(map(int, real_numbers_array))
print(real_numbers_array)
sum = 0
i = 0
while (i < len(real_numbers_array)):
    sum += real_numbers_array[i]
    i += 1
print(sum)
