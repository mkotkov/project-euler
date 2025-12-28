#Check if a number is a palindrome
def is_palindrome(num):
    str_num = str(num)
    return str_num == str_num[::-1]