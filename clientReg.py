import requests

response = requests.get('http://proxy:8888/api/reg?port=8000')
if response.status_code == 200:
	print "client reggistered"