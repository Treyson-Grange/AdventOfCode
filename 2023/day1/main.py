def part1():
    with open('input.txt', 'r') as file:
        total = 0
        for line in file.readlines():
            first = ''
            second = ''
            for char in line:
                if char.isnumeric():
                    if first == '':
                        first = char
                    else:
                        second = char
            if(second == ''):#then only one num in the line
                stringNum =  first + '' + first
                numNum = int(stringNum)
                total += numNum
            elif (second != ''):
                stringNum = first + '' + second
                numNum = int(stringNum)
                total += numNum
        return total
def part2():
    with open('input.txt', 'r') as file:
        total = 0
        for line in file.readlines():
            modified_line = line.replace('one', '1')
            modified_line = modified_line.replace('two', '2')
            modified_line = modified_line.replace('three', '3')
            modified_line = modified_line.replace('four', '4')
            modified_line = modified_line.replace('five', '5')
            modified_line = modified_line.replace('six', '6')
            modified_line = modified_line.replace('seven', '7')
            modified_line = modified_line.replace('eight', '8')
            modified_line = modified_line.replace('nine', '9')
            print(modified_line)
            first = ''
            second = ''
            for char in modified_line:
                if char.isnumeric():
                    if first == '':
                        first = char
                    else:
                        second = char
            if(second == ''):
                stringNum =  first + '' + first
                numNum = int(stringNum)
                print(numNum)
                total += numNum
            elif (second != ''):
                stringNum = first + '' + second
                numNum = int(stringNum)
                print(numNum)
                total += numNum
        return str(total) + 'total'

if __name__ == '__main__':
    print(part1())
    print(part2())
