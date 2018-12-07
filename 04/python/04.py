#!/usr/bin/env python3
import re
from collections import defaultdict
from datetime import datetime

import numpy as np


def get_minute(s):
    s_dt = re.findall(r'\[(.*)\].*', s)[0]
    return datetime.strptime(s_dt, '%Y-%m-%d %H:%M').minute


if __name__ == '__main__':
    with open('../04.txt', 'r') as f:
        lines = [s.strip() for s in f.readlines()]
    lines = sorted(lines, key=lambda s: datetime.strptime(
        re.findall(r'\[(.*)\].*', s)[0], '%Y-%m-%d %H:%M'
    ))

    guards = defaultdict(lambda: np.zeros(60))

    li = iter(lines)
    while True:
        try:
            s = next(li)
        except StopIteration:
            break
        gid_match = re.findall(r'Guard #(\d+)', s)
        if len(gid_match):
            gid = int(gid_match[0])
        else:
            asleep = get_minute(s)
            awake = get_minute(next(li))
            guards[gid][asleep:awake] += 1

    sleepy_guard = sorted(guards.items(), key=lambda x: x[1].sum())[-1]

    # Output part 1 result
    print(sleepy_guard[0] * np.argmax(sleepy_guard[1]))

    most_freq = sorted(guards.items(), key=lambda x: x[1].max())[-1]

    # Output part 2 result
    print(most_freq[0] * np.argmax(most_freq[1]))
