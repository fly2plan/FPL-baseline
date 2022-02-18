with open('DecToASCII.txt') as f:
    chars = []
    for i in range(128):
        chars.append(chr(i))
    for i, c in enumerate(chars):
        print(f'const u32 \'{c}\' = {i}')