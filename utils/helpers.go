package utils

import "github.com/go-chi/render"

// GOCHI UTILS
func NewRenderList[T render.Renderer](items []T) []render.Renderer {
	list := []render.Renderer{}
	for _, item := range items {
		list = append(list, item)
	}
	return list
}

// SLICE UTILS
func Reverse[S ~[]E, E any](s S) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func Contains[T comparable](data []T, item T) bool {
	for _, da := range data {
		if da == item {
			return true
		}
	}
	return false
}

func Map[T any, Z any](data []T, callback func(T) Z) []Z {
	result := []Z{}
	for _, d := range data {
		result = append(result, callback(d))
	}
	return result
}

func Any[T any](data []T, callback func(T) bool) bool {
	for _, d := range data {
		if callback(d) {
			return true
		}
	}
	return false
}

func Take[T any](data *[]T, callback func(T) bool) []T {
	result := []T{}
	dd := *data
	for i := len(dd) - 1; i >= 0; i-- {
		if callback(dd[i]) {
			result = append(result, dd[i])
			*data = append((*data)[:i], (*data)[i+1:]...)
		}
	}
	return result
}

func Filter[T any](data []T, callback func(T) bool) []T {
	result := []T{}
	for _, d := range data {
		if callback(d) {
			result = append(result, d)
		}
	}
	return result
}

func Keys[T comparable, K any](data map[T]K) []T {
	result := []T{}
	for k := range data {
		result = append(result, k)
	}
	return result
}

func Values[T comparable, K any](data map[T]K) []K {
	result := []K{}
	for _, v := range data {
		result = append(result, v)
	}
	return result
}

func Last[T any](arr []T) T {
	return arr[len(arr)-1]
}

func RemoveAt[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}
