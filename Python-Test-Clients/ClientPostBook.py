import requests

url = "http://localhost:8080"

load = {'id': 4, 'title': 'My GO Adventures', 'author': 'Jake Ressler', 'genre': 'Educational', 'shelf': "C", 'pages': 100}

headers = {"Content-Type": "application/json"}

r = requests.post(url+"/book", json = load, headers=headers)

print(r.content)

print(r)