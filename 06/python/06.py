#!/usr/bin/env python3
from itertools import product

import numpy as np


if __name__ == '__main__':
    # Load data
    data = np.loadtxt('../06.txt', delimiter=',', dtype=int)
    data -= np.min(data, axis=0)

    # Initialise grid
    dims = np.max(data, axis=0) + 1
    grid = np.full(np.prod(dims), -1)

    coords = np.array(list(product(range(dims[0]), range(dims[1]))))

    dist = np.sum(np.abs(coords[:,np.newaxis,:] - data), axis=2)
    for i, row in enumerate(dist):
        if len(np.where(row==np.min(row))[0]) == 1:
            grid[i] = np.argmin(row)

    grid = grid.reshape(dims)

    finite = set(range(len(data))) ^ (
        set(grid[0,:]) | set(grid[:,-1]) | set(grid[-1,:]) | set(grid[:,0])
    )

    areas = np.array([len(np.where(grid==x)[0]) for x in finite])

    # Output part 1 result
    print(np.max(areas))

    # Output part 2 result
    print(len(np.where(np.sum(dist, axis=1) < 10000)[0]))
