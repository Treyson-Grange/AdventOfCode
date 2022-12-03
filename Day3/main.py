dict = {
    'A' : 27,
    'B' : 28,
    'C' : 29,
    'D' : 30,
    'E' : 31,
    'F' : 32,
    'G' : 33,
    'H' : 34,
    'I' : 35,
    'J' : 36,
    'K' : 37,
    'L' : 38,
    'M' : 39,
    'N' : 40,
    'O' : 41,
    'P' : 42,
    'Q' : 43,
    'R' : 44,
    'S' : 45,
    'T' : 46,
    'U' : 47,
    'V' : 48,
    'W' : 49,
    'X' : 50,
    'Y' : 51,
    'Z' : 52,
    'a' : 1,
    'b' : 2,
    'c' : 3,
    'd' : 4,
    'e' : 5,
    'f' : 6,
    'g' : 7,
    'h' : 8,
    'i' : 9,
    'j' : 10,
    'k' : 11,
    'l' : 12,
    'm' : 13,
    'n' : 14,
    'o' : 15,
    'p' : 16,
    'q' : 17,
    'r' : 18,
    's' : 19,
    't' : 20,
    'u' : 21,
    'v' : 22,
    'w' : 23,
    'x' : 24,
    'y' : 25,
    'z' : 26,


} 
with open ("Day3/input.txt",'rt') as myFile:
    # print(ord('A'))
    # print(ord('B'))# subtract 38 for uper
    # print(ord('a'))
    # print(ord('b'))#subtract 96 for lower
    totalPriority = 0
    for rucksack in myFile:
        length = len(rucksack)
        halfLength = int(length/2)
        firstComp = rucksack[0:halfLength]
        secondComp = rucksack[halfLength:]
        #print(firstComp)
        #print(secondComp)
        for char in firstComp:
            if char in secondComp:
                totalPriority += dict[char]
                break

print(totalPriority)

