package main

import (
	"fmt"
	"sort"
)

func GetName(m map[string]string, id string) (string, bool) {
	value, ok := m[id]
	return value, ok
}
func HasFloor(m map[int]string, floor int) bool {
	_, ok := m[floor]
	return ok
}
func ReadPrice(m map[string]float64, course string) (float64, bool) {
	value, ok := m[course]
	return value, ok
}

func main() {
	//Создай пустую карту map[string]int для посещаемости буткемпа. Добавь 3 пары (дни недели - число). Выведи всю карту и len.
	attendance := map[string]int{}
	attendance["вторник"] = 1
	attendance["среда"] = 2
	attendance["четверг"] = 3
	fmt.Println(len(attendance))

	//Создай карту map[string]string с ID - имя сотрудника: добавь akhmad - Кудузов Ахмад, rasul - Назиров Расул. Выведи по ключу akhmad.
	ID := map[string]string{}
	ID["akhmad"] = "Кудузов Ахмад"
	ID["rasul"] = "Назиров Расул"
	fmt.Println(ID["akhmad"])

	//Создай карту map[int]string с номерами этажей буткемпа: 1 - кухня, 2 - коворкинг. Прочитай несуществующий ключ 3.
	butcemp := map[int]string{}
	butcemp[1] = "кухня"
	butcemp[2] = "коворкинг"
	fmt.Println(butcemp[3])

	//Создай карту map[string]float64 с названием курса - цена. Измени цену существующего курса.
	exchangeRate := map[string]float64{
		"dollar":  80.12,
		"bitcoin": 30000.21,
	}
	exchangeRate["dollar"] = 80.31
	fmt.Println(exchangeRate)

	//Создай и выведи карту-литерал map[string]string с тремя любимыми локациями в Грозном.
	Grozny := map[string]string{
		"Мечеть «Сердце Чечни»":   "Крупнейшая мечеть Европы, архитектурный символ Грозного.",
		"Грозный‑Сити":            "Современный деловой квартал с высотными зданиями и панорамными видами.",
		"Парк Хусейна Бен Талала": "Зелёная зона для прогулок с фонтанами и зонами отдыха.",
	}
	fmt.Println(Grozny)

	//Присвой карту переменной другой переменной (например, b := a) и измени через новую переменную один элемент. Выведи обе.
	location := Grozny
	location["Мечеть «Сердце Чечни»"] = " мечеть с золотыми куполами и минаретами."
	fmt.Println(Grozny, "\n")
	fmt.Println(location)

	//Создай карту map[string]string с 2 элементами, затем перезапиши один ключ новым значением. Убедись, что длина не изменилась.
	locations := map[string]string{
		"Мечеть «Сердце Чечни»": "Крупнейшая мечеть Европы, архитектурный символ Грозного.",
		"Грозный‑Сити":          "Современный деловой квартал с высотными зданиями и панорамными видами.",
	}
	fmt.Println(len(locations))
	locations["Мечеть «Сердце Чечни»"] = "мечеть с золотыми куполами и минаретами."
	fmt.Println(len(locations), "\n")

	//В map[string]string добавь ключ со строкой пустого значения "". Выведи разницу между «нет ключа» и «пустая строка».
	l := map[string]string{
		"":   "gg",
		"12": "",
	}
	fmt.Println("\n", l, "\n", "\n")

	//Создай карту с 4 студентами (ID - имя). Удали одного существующего студента и выведи len до/после.
	//Удали несуществующий ID — программа не должна падать, длина не меняется.
	//Удали все элементы (по одному в цикле), проверь, что len == 0.
	//Попробуй удалить ключ ещё раз. Обрати внимание что происходит при повторном удалении одного и того же ключа.
	names := map[int]string{
		1: "usup",
		2: "muslim",
		3: "hamzat",
		4: "imran",
	}
	fmt.Println(len(names))
	delete(names, 1)
	fmt.Println(len(names), "\n")
	delete(names, 5)
	fmt.Println(len(names))

	for k, _ := range names {
		delete(names, k)
	}

	fmt.Println(len(names))

	delete(names, 1)

	//Создай карту «адрес - название локации». Сделай цикл for range и выведи пары.
	loc := map[string]string{
		"Махмуда А. Эсамбаева, 1":             "Мечеть «Сердце Чечни»",
		"В. В. Путина, 30":                  "Грозный‑Сити",
		"Шейха Али‑Хаджи Акушинского, 11":     "Парк Хусейна Бен Талала",
	}
	for address, name := range loc {
		fmt.Printf("Адрес: %s\nНазвание: %s\n\n", address, name)
	}

	//Выведи только ключи, затем только значения (два отдельных цикла).
	for key,_ := range loc {
		fmt.Printf("Ключи:%s\n",key)
	}
	fmt.Println("\n")
		for _,value := range loc {
		fmt.Printf("Значение:%s\n",value)
	}
	fmt.Println("\n")
	//*Реализуй печать в алфавитном порядке ключей: собери ключи в срез, отсортируй, пройдись по нему.
	keys := []string{}
	for k ,_ := range loc {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, v := range keys{
		fmt.Println(v)
	}
	fmt.Println("\n")

	//Сформируй строку отчёта: "<адрес>: <название>" на каждой строке.
	 reportLines := []string{}
	for _, address := range keys { 
		name := loc[address]
		reportLine := fmt.Sprintf("%s: %s", address, name)
		reportLines = append(reportLines, reportLine)
	}
	for _, line := range reportLines {
		fmt.Println(line)
	}

	//вывод функций
	fmt.Println(GetName(ID, "akhmad"))
	fmt.Println(HasFloor(butcemp, 1))
	fmt.Println(HasFloor(butcemp, 3))
	fmt.Println(ReadPrice(exchangeRate, "bitcoi"))

}
