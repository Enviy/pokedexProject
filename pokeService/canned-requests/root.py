#!/usr/local/bin/python

import requests

headers = {}
payload = ""
r = requests.post("http://localhost:8080/", data=payload, headers=headers)
print r