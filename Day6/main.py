

INPUT_FILE = "Day6/input.txt"


def readFile(filePath):
	with open(filePath) as file:
		return file.readlines()


def findAllMarkers():
	lines = readFile(INPUT_FILE)
	for line in lines:
		print("The answer for part 1 is:",findMarkerIndex(line, 4))
		print("The answer for part 2 is:",findMarkerIndex(line, 14))
		

def findMarkerIndex(line, packetSize):
	char = list(line.replace("\n", ""))
	index = 0
	while(index < len(char)):
		currentSet = char[index:index + packetSize]
		if(len(set(currentSet)) == len(currentSet)):
			return index + packetSize
		index += 1
	return -1
	
	

findAllMarkers()



	


	