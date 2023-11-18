
def getPresents():
    with open('input.txt', 'r') as file:
        x, y = 0, 0
        visited_houses = {(0, 0)}
        for char in file.read():
            if char == '>':
                x += 1
            elif char == '<':
                x -= 1
            elif char == '^':
                y += 1
            elif char == 'v':
                y -= 1
            visited_houses.add((x, y))
        return len(visited_houses)

#Part two not done yet

if __name__ == '__main__':
    print(getPresents())

