#!/usr/bin/env python3
import numpy as np

LETTERS = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ'
l2n = {c: i for i, c in enumerate(LETTERS)}


def react(svec):
    while True:
        diff = np.diff(svec)
        polar = np.abs(diff)==26
        dbl_polar = np.abs(np.diff(diff))==52
        drop = np.full(len(svec), False)
        drop[:-1] |= polar
        drop[1:] |= polar
        drop[:-2] ^= (polar[:-1] & dbl_polar)
        if not np.any(drop):
            return svec
        svec = svec[~drop]


if __name__ == '__main__':
    with open('../05.txt', 'r') as f:
        s = f.read().strip()

    svec = np.array(list(map(l2n.get, s)))

    # Output part 1 result
    print(len(react(svec)))

    cnt = [
        len(react(svec[~((svec==i) | (svec==i+26))]))
        for i in range(26)
    ]

    # Output part 2 result
    print(min(cnt))
