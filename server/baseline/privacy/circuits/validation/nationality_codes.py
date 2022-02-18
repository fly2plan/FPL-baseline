with open('nationality_codes.txt') as f:
    lines = f.readlines()
    lines = [i for i in lines if i != '\n']
    lines = [i.strip() for i in lines]
    lines = [i for i in lines if i != '']
    lines = [i for i in lines if len(i) in [2, 3, 4] and not i[-1].islower()]
    lines = [int(''.join([str(ord(j)) for j in i])) for i in lines]
    lines_2 = [i for i in lines if len(str(i)) == 4]
    lines_3 = [i for i in lines if len(str(i)) == 6]
    lines_4 = [i for i in lines if len(str(i)) == 8]
    print(sorted(lines_2))
    print(sorted(lines_3))
    print(sorted(lines_4))
    print(len(lines_2))
    print(len(lines_3))
    print(len(lines_4))
    print(len(lines))
    print(len(lines_2) + len(lines_3) + len(lines_4))

def letterAsASCII(digits) -> int:
    res = 0
    for i in range(len(digits)):
        res = res + digits[i] * 10 ** (len(digits) - i - 1)
    return res

print(letterAsASCII([1, 2, 3, 4]))