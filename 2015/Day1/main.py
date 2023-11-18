def floorGetter():
    try:
        with open('input.txt', 'r') as file:
            floor = 0
            for char in file.read():
                if char == '(':
                    floor += 1
                elif char == ')':
                    floor -= 1

            return floor



    except FileNotFoundError:
        print(f"The file input.txt was not found.")
    except Exception as e:
        print(f"An error occurred: {e}")


if __name__ == '__main__':
    print(floorGetter())
