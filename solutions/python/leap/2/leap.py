"""

This task is to determine whether a given year is a leap year.

"""

def leap_year(year):
    """
    This function take a parameter year and evaluate it to determine if it's a leap year. 
    The result is a Boolean value: True or False 
    
    """
    result = ""
    if year % 400 == 0:
        result = True
    elif year % 4 == 0 and year % 100 != 0:
        result = True
    else:    
        result = False
    return result 
