package main

import "fmt"

func OnceInStr(str string) bool {
	m := make(map[rune]struct{})
	for _, v := range str {
		if v >= 97 {
			v -= 32
		}
		if _, ok := m[v]; ok {
			return false
		}
		m[v] = struct{}{}
	}
	return true
}

func main() {
	testOnceInStr()

}
func testOnceInStr() {
	tests := []struct {
		input    string
		expected bool
		name     string
	}{
		{
			name:     "Уникальные символы (разный регистр)",
			input:    "abCdef",
			expected: true,
		},
		{
			name:     "Повторяющиеся символы (разный регистр)",
			input:    "aAbB",
			expected: false,
		},
		{
			name:     "Повторяющиеся символы (один регистр)",
			input:    "hello",
			expected: false,
		},
		{
			name:     "Пустая строка",
			input:    "",
			expected: true,
		},
		{
			name:     "Unicode-символы (уникальные)",
			input:    "абвгАБВГ",
			expected: true,
		},
		{
			name:     "Unicode-символы (повторяющиеся)",
			input:    "абвгА",
			expected: false,
		},
		{
			name:     "Символы с пробелами и знаками",
			input:    "a b!c@",
			expected: true,
		},
		{
			name:     "Повторяющиеся пробелы",
			input:    "a  b",
			expected: false,
		},
	}

	for _, test := range tests {
		fmt.Printf("Тест: %s\n", test.name)
		fmt.Printf("Входная строка: %q\n", test.input)
		result := OnceInStr(test.input)
		fmt.Printf("Результат: %v\n", result)
		if result == test.expected {
			fmt.Println("Статус: Успех")
		} else {
			fmt.Printf("Статус: Ошибка, ожидалось %v, получено %v\n", test.expected, result)
		}
		fmt.Println()
	}
}
