import requests

base_url = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates"
params = {"f":"pjson", "SingleLine": "Abilene, TX"}

print(requests.get(base_url, params=params).json())
