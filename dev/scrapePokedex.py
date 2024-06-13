from bs4 import BeautifulSoup
import csv

# Open the HTML file and parse it with BeautifulSoup
with open('pokedex.html', 'r') as f:
    soup = BeautifulSoup(f, 'html.parser')

# Open a new CSV file in write mode
with open('pokedex.csv', 'w', newline='') as f:
    # Create a CSV writer
    writer = csv.writer(f)

    # Write the headers to the CSV file
    writer.writerow(['Number', 'Name', 'Type1', 'Type2', 'Image URL'])

    # Find the table rows
    rows = soup.find_all('tr')

    # Loop through the rows
    for row in rows:
        # Find the cells
        cells = row.find_all('td')

        # Check if this row contains a Pokemon
        if len(cells) > 0:
            # Extract the number, name, types, and image URL
            number = cells[0].find('span').text.strip()
            name = cells[1].text.strip()
            types = cells[2].find_all('a')
            type1 = types[0].text.strip()
            type2 = types[1].text.strip() if len(types) > 1 else ''
            image_url = cells[0].find('img')['src']

            # Write the data to the CSV file
            writer.writerow([number, name, type1, type2, image_url])