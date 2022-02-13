import csv

RADIUS = 0.904 # Roughly 100 km in degrees

# Write output to csv
def write_to_csv(edges, csv_name):
    with open(csv_name, 'w', newline='') as csv_file:
        writer = csv.writer(csv_file, delimiter=',')
        for e in edges:
            writer.writerow([e['start'], e['end'], e['distance']])

# Generate edge struct
def edge_factory(city_a, city_b):
    distance = ((city_a['lat'] - city_b['lat'])**2 + (city_a['lng'] - city_b['lng'])**2)**0.5
    return {'start': city_a['city'], 'end': city_b['city'], 'distance': distance}

# Determine if points are within a certain distance of each other
def within_circle(city_a, city_b):
    return (city_a['lat'] - city_b['lat'])**2 + (city_a['lng'] - city_b['lng'])**2 <= RADIUS**2

# Calculate edges and make list
def generate_edges(cities):
    edges = []
    for i in range(len(cities)):
        for j in range(i+1, len(cities)):
            if within_circle(cities[i], cities[j]):
                # add edge
                edges.append(edge_factory(cities[i], cities[j]))
                edges.append(edge_factory(cities[j], cities[i]))
    return edges

# Return w/ Name, ID, Population, and blank Lat/Lng
def city_factory(row):
    return {'city': row[0], 'id': row[1], 'population': int(row[2]), 'lat': float(row[4]), 'lng': float(row[3])}

# get metros from csv
def get_cities_list(csv_name):
    # read cities
    cities = []
    with open(csv_name, newline='') as csv_file:
        reader = csv.reader(csv_file, delimiter=',')
        for row in reader:
            cities.append(city_factory(row))
    return cities

def main():
    cities = get_cities_list('cleancities.csv')
    edges = generate_edges(cities)
    write_to_csv(edges, 'edges.csv')

main()
