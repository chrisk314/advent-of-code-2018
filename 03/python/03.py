#!/usr/bin/env python3
import re

import numpy as np


def extract_dims(s):
    pttrn = r'#\d+\s@\s(\d+),(\d+):\s(\d+)x(\d+)'
    return list(map(int, re.match(pttrn, s).groups()))


if __name__ == '__main__':
    with open('../03.txt', 'r') as f:
        dims = np.array([extract_dims(s) for s in f.readlines()])

    dims[:,(2,3)] += dims[:,(0,1)]

    fabric = np.zeros(np.max(dims[:,(2,3)], axis=0))

    for cut in dims:
        fabric[cut[0]:cut[2], cut[1]:cut[3]] += 1

    # Output part 1 result
    print(len(np.where(fabric.flatten() > 1)[0]))

    for idx, cut in enumerate(dims):
        if np.all(fabric[cut[0]:cut[2], cut[1]:cut[3]] == 1):
            break

    # Output part 2 result
    print('#%d' % (idx+1))
