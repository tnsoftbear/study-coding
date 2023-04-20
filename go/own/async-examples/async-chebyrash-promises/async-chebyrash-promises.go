package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/chebyrash/promise"
)

func main() {
	// Создание Promise для загрузки данных с сервера
	promiseFetchData := promise.New(func(resolve func([]byte), reject func(error)) {
		data, err := fetchDataFromServer()
		if err != nil {
			reject(err)
			return
		}
		resolve(data)
	})
	dataFetched, err := promiseFetchData.Await(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(string(*data))

	// Создание Promise для фильтрации данных
	promiseFilterData := promise.New(func(resolve func([]byte), reject func(error)) {
		filteredData, err := filterData(*dataFetched)
		if err != nil {
			reject(err)
			return
		}
		resolve(filteredData)
	})
	dataFiltered, err := promiseFilterData.Await(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(string(*dataFiltered))

	// Создание Promise для сортировки данных
	promiseSortData := promise.New(func(resolve func([]byte), reject func(error)) {
		sortedData, err := sortData(*dataFiltered)
		if err != nil {
			reject(err)
			return
		}
		resolve(sortedData)
	})
	dataSorted, err := promiseSortData.Await(context.Background())
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// fmt.Println(string(*dataSorted))

	// Создание Promise для вывода отсортированных данных на экран
	promisePrintData := promise.New(func(resolve func([]byte), reject func(error)) {
		err := printData(*dataSorted)
		if err != nil {
			reject(err)
			return
		}
		resolve(nil)
	})
	_, err = promisePrintData.Await(context.Background())
	if err != nil {
		fmt.Println(err.Error())
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
