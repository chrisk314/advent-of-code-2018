#!/usr/bin/env python3
from collections import Counter
from itertools import combinations

import numpy as np


if __name__ == '__main__':
    with open('../02.txt', 'r') as f:
        l = [s.strip() for s in f.readlines()]

    counts = [0, 0]
    for s in l:
        c = Counter(s)
        counts[0] += 2 in c.values()
        counts[1] += 3 in c.values()
    chksum = counts[0] * counts[1]

    # Output part 1 result
    print(chksum)

    # Find strings with smallest Levenshtein distance
    l2n = {c: i for i, c in enumerate('abcdefghijklmnopqrstuvwxyz')}

    combs = np.array(list(combinations(range(len(l)), 2)))
    pairs = np.array([list(map(l2n.get, s)) for s in l])[combs]
    diff = pairs[:,1,:] - pairs[:,0,:]
    min_idx = np.argmin(np.count_nonzero(diff, axis=1))

    similar = ''.join(
        np.array(list(l[combs[min_idx, 0]]))[diff[min_idx] == 0]
    )

    # Output part 2 result
    print(similar)
