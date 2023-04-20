package main

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/chebyrash/promise"
)

/**
Это пример вызова цепочки промисов с использованием Then() и Catch(),
Я не понял, а как реджектить цепочку если ошибка произошла внутри колбека переданного в Then()? Напр. в filterData()
*/

func main() {
	// Создание Promise для загрузки данных с сервера
	promiseFetchData := promise.New(
		func(
			resolve func([]byte),
			reject func(error)) {
			data, err := fetchDataFromServer()
			if err != nil {
				reject(err)
				return
			}
			resolve(data)
		})

	promiseFilterData := promise.Then(
		promiseFetchData,
		context.Background(),
		func(data []byte) []byte {
			filteredData, _ := filterData(data)
			return filteredData
		})

	// Создание Promise для сортировки данных
	promiseSortData := promise.Then(
		promiseFilterData,
		context.Background(),
		func(data []byte) []byte {
			sortedData, _ := sortData(data)
			return sortedData
		})

	// Создание Promise для вывода отсортированных данных на экран
	promisePrintData := promise.Then(
		promiseSortData,
		context.Background(),
		func(data []byte) error {
			return printData(data)
		})

	promiseError := promise.Catch(
		promisePrintData,
		context.Background(),
		func(err error) error {
			fmt.Println(err.Error())
			return err
		})

	promiseError.Await(context.Background())
}

func fetchDataFromServer() ([]byte, error) {
	if true {
		return []byte("From the Muddy Banks Of The Wishkah"), nil
	}

	return nil, errors.New("Cannot load data")
}

func filterData(data []byte) ([]byte, error) {
	if true {
		return filterVowels(data), nil
	}

	return nil, errors.New("Cannot filter data")
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
