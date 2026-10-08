def square(number):
    if number < 1 or number > 64:
        raise ValueError("square must be between 1 and 64")
    else:
        return 2 ** (number - 1)
        
def total():
    total_num = 0
    for num in range(64):
        total_num += square(num + 1)
    return total_num