#!/usr/bin/env python3
def add_node(lic):
    n, m = next(lic), next(lic)
    return {
        'nodes': {i: add_node(lic) for i in range(n)},
        'metadata': [next(lic) for i in range(m)]
    }


def sum_metadata(node):
    return sum(node['metadata']) + \
        sum(sum_metadata(v) for k, v in node['nodes'].items())


def sum_nodes(node):
    if not node['nodes']:
        return sum(node['metadata'])
    else:
        child_sums = {k: sum_nodes(v) for k, v in node['nodes'].items()}
        return sum(child_sums.get(i-1, 0) for i in node['metadata'])


if __name__ == '__main__':
    with open('../08.txt', 'r') as f:
        lic = map(int, f.read().split())

    tree = {0: add_node(lic)}

    # Output part 1 answer
    print(sum_metadata(tree[0]))

    # Output part 2 answer
    print(sum_nodes(tree[0]))
