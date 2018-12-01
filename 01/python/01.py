#!/usr/bin/env python3
import numpy as np


if __name__ == '__main__':
    # Load data from file
    nums = np.loadtxt('../01.txt')

    # Output part 1 result
    print(int(np.sum(nums)))

    def find_repeat(nums):
        csum = np.zeros(len(nums))
        m = {}
        # Loop forever until same frequency observed twice
        while True:
            csum = csum[-1] + np.cumsum(nums)
            for n in csum:
                if n in m:
                    return n
                else:
                    m[n] = None

    # Output part 2 result
    print(int(find_repeat(nums)))
