from forex_python.converter import CurrencyRates

def curExchange(_value, _currentCurrencyName, _newCurrencyName):
    c = CurrencyRates()
    _newValue = _value * c.get_rate(_currentCurrencyName, _newCurrencyName)
    print(_newValue, c.get_rate(_currentCurrencyName, _newCurrencyName))
    return (_newValue, _newCurrencyName)