import re

# Fetch input data
with open('Day5/input.txt') as input_data:
    input_data = input_data.readlines()
    stacks = [stack.strip('\n') for stack in input_data[:8]]
    instructions = [line.strip() for line in input_data[10:]]

            # Create starting stacks as lists/saving lists to dict
    one = []
    two = []
    three = []
    four = []
    five = []
    six = []
    seven = []
    eight = []
    nine = []
    stacks_dict = {
        '1':one,
        '2':two,
        '3':three,
        '4':four,
        '5':five,
        '6':six,
        '7':seven,
        '8': eight,
        '9': nine,
        }
    
    for crate in stacks:
        if crate[:4] != '    ':
            one.append(crate[1])
        if crate[4:8] != '    ':
            two.append(crate[5])
        if crate[8:12] != '    ':
            three.append(crate[9])
        if crate[12:16] != '   ':
            four.append(crate[13])
        if crate[16:20] != '   ':
            five.append(crate[17])  
        if crate[12:16] != '   ':
            six.append(crate[21])  
        if crate[16:20] != '   ':
            seven.append(crate[25])
        if crate[20:24] != '   ':
            eight.append(crate[29])
        if crate[24:27] != '  ':
            nine.append(crate[33])    
               
            # Tidy up empty list elements
    for stack in stacks_dict.keys():
        stacks_dict[stack] = [crate for crate in stacks_dict[stack] if crate != ' ']
     # Follow instrcution to move crates across stacks
    for instruction in instructions:
        quanity, source, dest = re.findall(r'[0-9][0-9]*', instruction)
        move = stacks_dict[source][:int(quanity)]
        for x in move:
            stacks_dict[source].remove(x)
        stacks_dict[dest] = move + stacks_dict[dest]
                   
    
        # Concatente first element of list to find answer
answer = ''
for stack in stacks_dict.values():
    answer += stack[0]
    
print(answer)