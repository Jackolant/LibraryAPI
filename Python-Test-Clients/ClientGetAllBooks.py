import requests

url = "http://localhost:8080"

r = requests.get(url+"/books")

print(r.content)

print(r)