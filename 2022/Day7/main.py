s = 0
m = 1000000000000
dirs = {'/':0}
cd = ['/']
# calc dir sizes
with open("Day7/input.txt", "r") as file:
    for line in file.readlines():
        lsplit = line[:-1].split(" ")
        if lsplit[0] == '$':
            if lsplit[1] == 'cd':
                if lsplit[2] == '..':
                    cd.pop(-1)
                elif lsplit[2] == '/':
                    cd = ['/']
                else:
                    cd.append(lsplit[2])
        elif lsplit[0] == 'dir':
            dirs["".join(cd) + lsplit[1]] = 0
        else:
            dirs["".join(cd)] += int(lsplit[0])
            for i in range(1, len(cd)):
                dirs["".join(cd[:-i])] += int(lsplit[0])
for d in dirs: # calc p1
    if (dirs[d] <= 100000):
        s += dirs[d]
for d in dirs: # calc p2
    if (dirs[d] >= dirs['/'] - 40000000 and dirs[d] < m):
        m = dirs[d]

print(s, m)