import requests

url = "http://localhost:8080"

r = requests.get(url+"/books/3")

print(r.content)

print(r)