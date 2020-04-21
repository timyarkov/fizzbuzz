import sys

def fizzbuzz(num, divs_ls):
    """

    Performs fizzbuzz function. Takes list of defined multiples and associated word. Starts printing numbers from 1 to the specified end (num), and if a number is a multiple of one of the numbers in divs_ls, the associated word will be added to result

    """
    i = 0
    while i < num:
        result = ""
        j = 0
        while j < len(divs_ls):
            if (i + 1) % divs_ls[j][0] == 0:
                result += divs_ls[j][1]
            j += 1
        if result == "":
            print(i + 1)
        else:
            print(result)
        i += 1



print("number of iterations for fizzbuzz????")
while True:
    try:
        num = int(input("> "))
    except TypeError:
        print("bad input :( try again.")
        continue
    except ValueError:
        print("bad input :( try again.")
        continue
    if num <= 0:
        print("bad input :( try again.")
        continue
    else:
        break

print("how many divs do you want")
while True:
    try:
        divs_num = int(input("> "))
    except TypeError:
        print("bad input :( try again.")
        continue
    except ValueError:
        print("bad input :( try again.")
        continue
    if divs_num <= 0:
        print("bad input :( try again.")
        continue
    else:
        break

print("input a number, followed by the phrase to be said")
#get div_ls
div_ls = []
i = 0
while i < divs_num:
    try:
        div_str, word = input("{}/{}> ".format(i + 1, divs_num)).split(" ")
    except ValueError:
        print("bad input :( try again.")
        continue

    try:
        div = int(div_str)
    except TypeError:
        print("bad input :( try again.")
        continue
    except ValueError:
        print("bad input :( try again.")
        continue
    
    div_word_tuple = (div, word)
    div_ls.append(div_word_tuple)
    i += 1

fizzbuzz(num, div_ls)