import requests
from bs4 import BeautifulSoup

# Function to search for Patagonia jackets and return a list of URLs for websites that sell them
def find_jacket_urls(search_term):
    # Send a GET request to Google using the search term
    response = requests.get('https://www.google.com/search?q=' + search_term)
    
    # Parse the HTML of the search results page using BeautifulSoup
    soup = BeautifulSoup(response.text, 'html.parser')
    
    # Find all the links on the page that contain the search term
    links = soup.find_all('a', href=True, text=search_term)
    
    # Return the list of URLs for the websites that sell Patagonia jackets
    print(links)
    return [link['href'] for link in links]

# Function to extract the details of Patagonia jackets from a website
def find_jacket_details(url):
    # Send a GET request to the website using the URL
    response = requests.get(url)
    
    # Parse the HTML of the page using BeautifulSoup
    soup = BeautifulSoup(response.text, 'html.parser')
    
    # Find all the elements on the page that contain the details of the jackets
    jacket_elements = soup.find_all('div', class_='jacket-details')
    
    # Extract the details of the jackets from the elements and return them as a list of dictionaries
    jackets = []
    for jacket_element in jacket_elements:
        jacket = {}
        jacket['price'] = jacket_element.find('a')
    print(jackets)


find_jacket_urls("gun")
