package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type ExchangeRatesApiResponse struct {
	Success bool               `json:"success"`
	Rates   map[string]float64 `json:"rates"`
}

type ExchangeRatesSymbolsApiResponse struct {
	Success bool                   `json:"success"`
	Symbols map[string]interface{} `json:"symbols"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Currency Converter")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter amount in USD")

	from := widget.NewSelect(nil, func(s string) {})
	to := widget.NewSelect(nil, func(s string) {})
	output := widget.NewLabel("")

	convertBtn := widget.NewButton("Convert", func() {
		amount, err := strconv.ParseFloat(input.Text, 64)
		if err != nil {
			output.SetText("Invalid input")
			return
		}

		fromCurrency := from.Selected
		toCurrency := to.Selected
		convertedAmount, err := convert(amount, fromCurrency, toCurrency)
		if err != nil {
			output.SetText("Conversion error")
			return
		}

		output.SetText(fmt.Sprintf("%.2f %s = %.2f %s", amount, fromCurrency, convertedAmount, toCurrency))
	})

	content := container.NewVBox(
		input,
		container.NewHBox(from, to),
		convertBtn,
		output,
	)

	w.SetContent(content)

	// Fetch available currencies from API
	currencies, err := getAvailableCurrencies()
	if err != nil {
		output.SetText(err.Error())
	}

	// Sort the currencies alphabetically
	sort.Strings(currencies)

	// Populate the select widgets with the available currencies
	from.Options = currencies
	to.Options = currencies

	w.ShowAndRun()
}

func convert(amount float64, fromCurrency string, toCurrency string) (float64, error) {
	url := fmt.Sprintf("https://api.exchangerate.host/latest?base=%s&symbols=%s", fromCurrency, toCurrency)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var apiResponse ExchangeRatesApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return 0, err
	}

	if !apiResponse.Success {
		return 0, fmt.Errorf("API request failed")
	}

	if rate, ok := apiResponse.Rates[toCurrency]; ok {
		return amount * rate, nil
	}

	return 0, fmt.Errorf("Invalid currency code")
}

func getAvailableCurrencies() ([]string, error) {
	url := "https://api.exchangerate.host/symbols"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var apiResponse ExchangeRatesSymbolsApiResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return nil, err
	}

	if !apiResponse.Success {
		return nil, fmt.Errorf("API request failed")
	}
	currencies := make([]string, 0, len(apiResponse.Symbols))
	for currency := range apiResponse.Symbols {
		currencies = append(currencies, currency)
	}

	return currencies, nil
}
