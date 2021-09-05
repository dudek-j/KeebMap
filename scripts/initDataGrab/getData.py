from bs4 import BeautifulSoup
import requests
import re
import json

response = requests.get("https://www.keebtalk.com/t/list-of-keyboard-retailers-shops-stores-vendors/9022")
contents = response.text
soup = BeautifulSoup(contents, 'lxml')

rows = soup.find_all("tr")

file = open("output.json", "w")

file.write("[")

for row in rows:
    output = "{"
    cols = row.find_all("td")
    if len(cols) == 5:
        text = cols[0].text
        country = cols[2]
        if len(text) > 1:
            data = re.split("http", text)
            if len(data) == 2:
                output = output + "name: " + '"' + data[0] + '"' + "," + " url: " + '"' + "http" + data[1].strip() + '"' + ", country: "+ '"' + country.find("img")["title"][1:-1] + '"' + ", region: "+ '"' + country.find("img")["title"][1:-1] + '"' + " },"
                file.write(output)
                
file.write("]")