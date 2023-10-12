from sys import argv
import requests
from urllib.parse import urlencode


# def make_tiny(url):
#     request_url = "http://tinyurl.com/api-create.php?" + urlencode({"url": url})
#     with contextlib.closing(urlopen(request_url)) as response:
#         return response.read().decode("utf-8")


def make_tiny(url):
    request_url = "http://tinyurl.com/api-create.php?" + urlencode({"url": url})
    response = requests.get(request_url)
    return response.text


if __name__ == "__main__":
    # url = argv[1]
    urls = argv[1::]
    for i, url in enumerate(urls):
        tiny = make_tiny(url)
        print(tiny)
