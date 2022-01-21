import requests

url = "http://localhost:8080"

load = {'id': 7, 'title': 'Updated Go Adventures', 'author': 'Updated to someone else', 'genre': 'Science Fiction', 'shelf': "B", 'pages': 800}

headers = {"Content-Type": "application/json"}

r = requests.put(url+"/books/4", json = load, headers=headers)

print(r.content)

print(r)