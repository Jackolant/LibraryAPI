import requests

url = "http://localhost:8080"

id = "2"

print("Deleting book id: " + id)

r = requests.delete(url+"/book/" + id)

print(r.content)

print(r)