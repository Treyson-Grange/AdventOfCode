def getWrappingPaper():
    totalSqFootage = 0
    with open('input.txt', 'r') as file:
        for line in file.readlines():
            asdf = line.split('x')
            input1 = 2 * int(asdf[0]) * int(asdf[1])
            input2 = 2 * int(asdf[1]) * int(asdf[2])
            input3 = 2 * int(asdf[2]) * int(asdf[0])
            SA = input1 + input2 + input3
            smallestSideSA = min(input1,input2,input3) / 2
            totalSqFootage += SA + smallestSideSA
        return totalSqFootage

def getRibbon():
    totalRibbon = 0
    with open('input.txt', 'r') as file:
        for line in file.readlines():
            asdf = line.split('x')
            l = int(asdf[0])
            w = int(asdf[1])
            h = int(asdf[2])
            bowPart = l*w*h

            values = [l,w,h]
            sortValues = sorted(values)
            wrapPart = (sortValues[0] * 2) + (sortValues[1] * 2)
            totalRibbon += bowPart + wrapPart
        return totalRibbon

if __name__ == '__main__':
    print(getWrappingPaper())
    print(getRibbon())
