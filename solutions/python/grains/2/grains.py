def square(number):
    """
    This function takes a number and return the number of grains expected on the square of a chessboard.
    """
    if number < 1 or number > 64:
        raise ValueError("square must be between 1 and 64")
    return 2 ** (number - 1)
        
def total():
    """
    This function calculate the total number of grains on the chessboard.
    """
    total_num = 0
    for num in range(64):
        total_num += square(num + 1)
    return total_num