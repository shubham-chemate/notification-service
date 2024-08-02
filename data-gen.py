import json


with open("data.json", "w") as of:
    with open("raw-data.txt", "r") as f:
        cnt=0
        for l in f:
            l = l.strip()
            l = l[l.find(' '):]
            receiver = ["Shubham", "Rushikesh", "Kaiwalya", "Pranav", "Siddharth"][cnt%5]
            msg = {
                'sendTo': receiver,
                'text': l
            }
            of.write(json.dumps(msg)+'\n')
            cnt += 1
        