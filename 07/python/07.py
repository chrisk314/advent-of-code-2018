#!/usr/bin/env python3
import re
from collections import defaultdict
from copy import deepcopy

LETTERS = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'
l2n = {c: i+61 for i, c in enumerate(LETTERS)}


def get_in_out(s):
    return re.match(r'Step ([A-Z]).*step ([A-Z]).*', s).groups()


if __name__ == '__main__':
    with open('../07.txt', 'r') as f:
        in_out = [get_in_out(s) for s in f.readlines()]

    in2out = defaultdict(dict)
    out2in = defaultdict(dict)
    for i, o in in_out:
        in2out[i][o] = None
        out2in[o][i] = None
    universe = set(in2out) | set(out2in)

    # Backup these maps for part 2...
    tmp_io = deepcopy(in2out)
    tmp_oi = deepcopy(out2in)

    seq = []
    while len(in2out):
        choices = set(in2out) - set(out2in)
        choice = sorted(choices)[0]
        choices ^= {choice}
        seq += [choice]
        in2out.pop(choice)
        for k, v in list(out2in.items()):
            v.pop(choice, None)
            out2in[k] = v
            if not len(v):
                out2in.pop(k)
    seq += sorted(universe - set(seq))

    # Output part 1 result
    print(''.join(seq))

    # Reinitialise maps
    in2out = tmp_io
    out2in = tmp_oi

    n_workers = 5
    task = [None] * n_workers
    tick = [0] * n_workers
    done = set()

    for i, c in enumerate(sorted(set(in2out) - set(out2in))):
        task[i], tick[i] = c, l2n[c]

    cnt = 0
    while sum(tick) > 0:
        # print(cnt, tick, task)
        tick = [x-1 if y else x for x, y in zip(tick, task)]
        cnt += 1
        finished = [i for i, (x, y) in enumerate(zip(tick, task)) if y and x==0]
        if finished:
            for idx in finished:
                done |= {task[idx]}
                task[idx] = None

            active = set(x for x in task if x)

            # determine tasks which can now be started
            available = set()
            for c in done:
                for n in in2out[c]:
                    if not len(set(out2in[n]) - done):
                        available |= {n}
            available = sorted(available - done - active)
            # print(done, available)

            # Assign new tasks
            for i in range(n_workers):
                if not any([x==None for x in task]) or not len(available):
                    break
                if not task[i]:
                    c = available.pop(0)
                    task[i], tick[i] = c, l2n[c]

    # Output part 2 result
    print(cnt)
