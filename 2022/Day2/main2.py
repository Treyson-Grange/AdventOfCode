totalPoints2 = 0
with open ("input.txt",'rt') as myFile:
    for round in myFile:
        rounds = round.split()
        if rounds[1] == "Y":
            if rounds[0] =="A":
                totalPoints2+= 4
            if rounds[0] == "B":
                totalPoints2 += 5
            if rounds[0] == "C":
                totalPoints2 += 6
        if rounds[1] == "X":
            if rounds[0] == "A":
                totalPoints2 += 3
            if rounds[0] == "B":
                totalPoints2 += 1
            if rounds[0] == "C":
                totalPoints2 += 2
        if rounds[1] == "Z":
            if rounds[0] == "A":
                totalPoints2 += 8
            if rounds[0] == "B":
                totalPoints2 += 9
            if rounds[0] == "C":
                totalPoints2 += 7

print(totalPoints2)
            #totalPoints2 += 6