import requests

ex_info = requests.get('https://api.binance.com/api/v3/exchangeInfo')
result = ex_info.json()['symbols']
for i in result:
    print(i)