import requests
import csv
import re

CITY_COL = 0
ID_COL = 1
POP_COL = 2

def write_to_csv(cities, csv_name):
    with open(csv_name, 'w', newline='') as csv_file:
        writer = csv.writer(csv_file, delimiter=',')
        for c in cities:
            writer.writerow([c['city'], c['id'], c['population'], c['lat'], c['lng']])

# geocoding via esri rest api
def geocode_addresses(cities):
    base_url = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer/findAddressCandidates"
    params = {"f":"pjson", "SingleLine": "Abilene, TX"}

    for c in cities:
        params["SingleLine"] = c['city']
        resp = requests.get(base_url, params=params).json()
        print(resp['candidates'][0]['address'])
        c['lat'] = resp['candidates'][0]['location']['y']
        c['lng'] = resp['candidates'][0]['location']['x']

# needed for proper formatting
def clean_cities(cities):
    for c in cities:
        # get elements of rough string
        split_name = re.split("--|, ", c['city'])
        city = split_name[0]
        state = ""
        for s in split_name:
            # state identifier
            if len(s) == 2:
                state = s
                break
        c['city'] = city + ", " + state

# Return w/ Name, ID, Population, and blank Lat/Lng
def city_factory(row):
    return {"city": row[CITY_COL], "id": row[ID_COL], 'population': row[POP_COL], 'lat': 0.0, 'lng': 0.0}

# get metros from csv
def get_cities_list(csv_name):
    # read cities
    cities = []
    with open(csv_name, newline='') as csv_file:
        reader = csv.reader(csv_file, delimiter=',')
        next(reader, None)
        for row in reader:
            cities.append(city_factory(row))
    return cities

def main():
    cities = get_cities_list('urbanareas.csv')
    clean_cities(cities)
    geocode_addresses(cities)
    write_to_csv(cities, 'cleancities.csv')

main()
