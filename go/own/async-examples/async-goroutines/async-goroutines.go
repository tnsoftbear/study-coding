package main

import (
	"fmt"
	"sort"
	"strings"
)

func _main() {
	done := make(chan struct{})
	dataChan := make(chan []byte)
	filteredDataChan := make(chan []byte)
	sortedDataChan := make(chan []byte)
	errChan := make(chan error)

	// горутина для загрузки данных с сервера
	go func() {
		data, err := fetchDataFromServer()
		if err != nil {
			errChan <- err
			return
		}
		dataChan <- data
	}()

	// горутина для фильтрации данных
	go func() {
		data := <-dataChan
		filteredData, err := filterData(data)
		if err != nil {
			errChan <- err
			return
		}
		filteredDataChan <- filteredData
	}()

	// горутина для сортировки данных
	go func() {
		filteredData := <-filteredDataChan
		sortedData, err := sortData(filteredData)
		if err != nil {
			errChan <- err
			return
		}
		sortedDataChan <- sortedData
	}()

	// горутина для вывода отсортированных данных на экран
	go func() {
		sortedData := <-sortedDataChan
		err := printData(sortedData)
		if err != nil {
			errChan <- err
			return
		}
		done <- struct{}{}
	}()

	// обработка ошибок
	select {
	case err := <-errChan:
		fmt.Println(err.Error())
	case <-done:
		return
	}
}

func fetchDataFromServer() ([]byte, error) {
	return []byte("From the Muddy Banks Of The Wishkah"), nil
}

func filterData(data []byte) ([]byte, error) {
	return filterVowels(data), nil
}

func sortData(data []byte) ([]byte, error) {
	sort.Slice(data, func(i, j int) bool {
		return data[i] < data[j]
	})
	return data, nil
}

func printData(data []byte) error {
	fmt.Println(string(data))
	return nil
}

func filterVowels(bytes []byte) []byte {
	var filteredBytes []byte
	vowels := "aeiouAEIOU "
	for _, b := range bytes {
		if !strings.ContainsAny(string(b), vowels) {
			filteredBytes = append(filteredBytes, b)
		}
	}
	return filteredBytes
}
