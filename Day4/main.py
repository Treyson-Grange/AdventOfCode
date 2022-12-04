import sys
count = 0
part2count = 0
file = open("Day4\input.txt")
data = [line.strip().split(",") for line in file]
for pair in data:
    firstPair = [int(i) for i in pair[0].split("-")]
    secondPair = [int(i) for i in pair[1].split("-")]
    if (firstPair[0] <= secondPair[0] and secondPair[1] <= firstPair[1]):
        count += 1
    elif (secondPair[0] <= firstPair[0] and firstPair[1] <= secondPair[1]):
        count += 1
    if (secondPair[0] <= firstPair[1] and secondPair[1] >= firstPair[1]):
        part2count += 1
    elif (firstPair[0] <= secondPair[1] and firstPair[1] >= secondPair[1]):
        part2count += 1

print("Day 4 fuckers\n")
print("Amount of tasks that are completely overlapped: ",count,"\n")
print("Amount of tasks that are overlapped at all: ",part2count,"\n")



        