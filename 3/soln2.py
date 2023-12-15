import os
import re
import functools

# given a cell with a digit, find the indices of the maximal number corresponding
def maximal_num_coord(row, col, array):
    if not array[row][col].isdigit():
        raise Exception
    
    left_max = col
    while (left_max >= 0 and array[row][left_max].isdigit()):
        left_max -= 1
    left_max += 1

    right_max = col
    while (right_max < len(array[0]) and array[row][right_max].isdigit()):
        right_max += 1
    right_max -= 1
    return (left_max, right_max)

# given two tuples each indicating a range, return if they are intersecting
def are_intersecting(t1, t2):
    a, b = t1
    c, d = t2
    if a < c:
        if c > b:
            return False
        else:
            return True
    elif c == a:
        return False
    else:
        if a > b:
            return False
        else:
            return True

def is_cell_coord_inside(row, col, array):
    return row >= 0 and row < len(array) and col >= 0 and col < len(array[0])

if __name__ == '__main__':
    array = []
    with open('input.txt', 'r') as file:
        while line := file.readline():
            array.append(line.strip())

    stars_positions = []
    for line_num, line in enumerate(array):
        stars = re.finditer(r'\*', line)
        for star in stars:
            stars_positions.append((line_num, star.start()))

    # given the position of a star, either return the corresponding gear ratio or None
    def find_gear_ratio(row, col):
        cogs = []
        try:
            if array[row][col - 1].isdigit():
                left, right = maximal_num_coord(row, col - 1, array)
                cogs.append(int(array[row][left:right + 1]))
        except IndexError:
            pass
        try:
            if array[row][col + 1].isdigit():
                left, right = maximal_num_coord(row, col + 1, array)
                cogs.append(int(array[row][left:right + 1]))
        except IndexError:
            pass

        top = [None, None, None]
        if is_cell_coord_inside(row - 1, col - 1, array) and array[row - 1][col - 1].isdigit():
            top[0] = maximal_num_coord(row - 1, col - 1, array)
        if is_cell_coord_inside(row - 1, col, array) and array[row - 1][col].isdigit():
            top[1] = maximal_num_coord(row - 1, col, array)
        if is_cell_coord_inside(row - 1, col + 1, array) and array[row - 1][col + 1].isdigit():
            top[2] = maximal_num_coord(row - 1, col + 1, array)
        top = list(filter(lambda a: a, top))
        unique_top = set(top)
        for l, r in unique_top:
            cogs.append(int(array[row - 1][l:r+1]))


        bottom = [None, None, None]
        if is_cell_coord_inside(row + 1, col - 1, array) and array[row + 1][col - 1].isdigit():
            bottom[0] = maximal_num_coord(row + 1, col - 1, array)
        if is_cell_coord_inside(row + 1, col, array) and array[row + 1][col].isdigit():
            bottom[1] = maximal_num_coord(row + 1, col, array)
        if is_cell_coord_inside(row + 1, col + 1, array) and array[row + 1][col + 1].isdigit():
            bottom[2] = maximal_num_coord(row + 1, col + 1, array)
        bottom = list(filter(lambda a: a, bottom))
        unique_bottom = set(bottom)
        for l, r in unique_bottom:
            cogs.append(int(array[row + 1][l:r+1]))

        if len(cogs) == 2:
            print('gear ratio star is at', row, col)
            return cogs[0] * cogs[1]
        else:
            print('non gear ratio star is at', row, col, "its cogs are", cogs, "unique_bottom", unique_bottom, "unique_top", unique_top)
            return None

    gear_ratios = []
    for star_position in stars_positions:
        if gr := find_gear_ratio(star_position[0], star_position[1]):
            gear_ratios.append(gr)

    print(gear_ratios)
    print(functools.reduce(lambda a, b: a + b, gear_ratios))