import sys
import re

f = open("input.txt", "r")
input = f.read()
lines = [re.split('[\\s=;,]+', x) for x in input.splitlines()]
# adjacency list
G = {x[1]: set(x[10:]) for x in lines}
# dict of flowrates for all valves with nonzero flowrate
F = {x[1]: int(x[5]) for x in lines if int(x[5]) != 0}
# dict of valve names to bitmap representing only that it is turned on
I = {x: 1 << i for i, x in enumerate(F)}
# map of distances from valve to valve
T = {x: {y: 1 if y in G[x] else float('+inf') for y in G} for x in G}

# populate distance matrix
for k in T:
    for i in T:
        for j in T:
            T[i][j] = min(T[i][j], T[i][k] + T[k][j])

# distance is a proxy to time
# there are n^2 states, where n = number of valves
# answer is a map with key bitmap to maximal flowrate so far (with regard to time left)
def visit(v, budget, state, value, answer):
    # print(answer)
    answer[state] = max(answer.get(state, 0), value)
    for u in F:
        newbudget = budget - T[v][u] - 1
        if I[u] & state or newbudget < 0:
            continue
        # state | I[u] produces a bitmap with valve u turned on
        visit(u, newbudget, state | I[u], value + newbudget * F[u], answer)
    return answer


total1 = max(visit('AA', 30, 0, 0, {}).values())
visited2 = visit('AA', 26, 0, 0, {})
total2 = max(
    v1 + v2 for k1, v1 in visited2.items() for k2, v2 in visited2.items() if not k1 & k2
)
print(total1, total2)
