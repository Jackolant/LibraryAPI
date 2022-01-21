import requests

url = "http://localhost:8080"

r = requests.get(url+"/book/3")

print(r.content)

print(r)