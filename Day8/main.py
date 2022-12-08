with open("Day8/input.txt", "r") as f:
    data = f.read().split("\n")[:-1]

data = [[int(y) for y in x] for x in data]
visible = 0
#print(data)
values = []

for x in range(len(data[0])):
    for y in range(len(data)):
        height = data[y][x]
        distances = []
        isVisible = False

        if x == 0 or y == 0 or x == len(data[0]) - 1 or y == len(data) - 1:
            isVisible = True
                   
        d = True
        for ix in range(1, x+1):
            if data[y][x-ix] >= height:
                d = False
                distances.append(ix)
                break

        if len(distances) < 1:
            distances.append(x)

        if d:
            isVisible = True

        d = True
        for iy in range(1, y+1):
            if data[y-iy][x] >= height:
                d = False
                distances.append(iy)
                break

        if len(distances) < 2:
            distances.append(y)

        if d:
            isVisible = True

        d = True
        for ix in range(1, len(data[0])-x):
            if data[y][x+ix] >= height:
                d = False
                distances.append(ix)
                break
        if d:
            isVisible = True

        if len(distances) < 3:
            distances.append(len(data[0])-x-1)

        d = True
        for iy in range(1, len(data)-y):
            if data[y+iy][x] >= height:
                d = False
                distances.append(iy)
                break

        if len(distances) < 4:
            distances.append(len(data[0])-y-1)

        if d:
            isVisible = True
        
        if isVisible:
            visible += 1

        s = 1
        for value in distances:
            s = s*value
        values.append(s)

print("Part 1: ",visible)
print("Part 2: ",max(values))
