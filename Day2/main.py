#ile = open("input.txt",'rt')
enPaper = "B"
enRock = "A"
enSci = "C"
myRock = "X"
mySci = "Z"
myPaper = "Y"
print("Hi")
totalPoints = 0
with open ("input.txt",'rt') as myFile:
    for round in myFile:
        rounds = round.split()
        
        enChoice = ""
        myChoice = ""
        if rounds[0] == enPaper:
            enChoice = "paper"
        if rounds[0] == enRock:
            enChoice = "rock"
        if rounds[0] == enSci:
            enChoice = "sci"
        if rounds[1] == myRock:
            myChoice = "rock"
        if rounds[1] == mySci:
            myChoice = "sci"
        if rounds[1] == myPaper:
            myChoice = "paper"
        if(myChoice == "paper"):
            totalPoints+= 2
        if(myChoice == "sci"):
            totalPoints+= 3
        if(myChoice == "rock"):
            totalPoints+= 1
        if(myChoice == enChoice):
            totalPoints += 3
        if(myChoice == "paper" and enChoice == "rock"):
            totalPoints += 6
        if(myChoice == "rock" and enChoice == "sci"):
            totalPoints += 6
        if(myChoice == "sci" and enChoice == "paper"):
            totalPoints += 6
        print(totalPoints)