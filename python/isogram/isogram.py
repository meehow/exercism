def is_isogram(string):
    string = string.lower()
    for i, letter in enumerate(string):
        if letter < 'a' or letter > 'z':
            continue
        for l2 in string[i+1:]:
            if letter == l2:
                return False
    return True
