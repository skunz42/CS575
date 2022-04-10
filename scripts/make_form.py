arr = []

with open("../data/cities2.txt") as f:
    lines = f.readlines()
    for l in lines:
        var = l.rstrip()
        print(f"    <option value =\"{var}\">{var}</option>")

