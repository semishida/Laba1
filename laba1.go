package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"
)

// Node представляет узел для стека и очереди.
type Node struct {
	data string
	next *Node
}

// Stack представляет стек.
type Stack struct {
	head *Node
}

// Queue представляет очередь.
type Queue struct {
	head *Node
	tail *Node
}

// Set представляет множество.
type Set struct {
	data []string
}

// HashTable представляет хеш-таблицу.
type HashTable struct {
	data []hashTableEntry
}

type hashTableEntry struct {
	key   string
	value string
}

// NewSet создает новое множество.
func NewSet() *Set {
	return &Set{}
}

// Add добавляет элемент в множество.
func (set *Set) Add(value string) {
	// Проверяем, есть ли элемент уже в множестве
	if !set.Contains(value) {
		set.data = append(set.data, value)
	}
}

// Contains проверяет, содержит ли множество указанный элемент.
func (set *Set) Contains(value string) bool {
	for _, v := range set.data {
		if v == value {
			return true
		}
	}
	return false
}

// Remove удаляет элемент из множества.
func (set *Set) Remove(value string) {
	for i, v := range set.data {
		if v == value {
			set.data = append(set.data[:i], set.data[i+1:]...)
			return
		}
	}
}

// Push добавляет элемент на вершину стека.
func (stack *Stack) Push(value string) {
	node := &Node{data: value}
	if stack.head == nil {
		stack.head = node
	} else {
		node.next = stack.head
		stack.head = node
	}
}

// Pop удаляет и возвращает элемент с вершины стека.
func (stack *Stack) Pop() (string, error) {
	if stack.head == nil {
		return "", errors.New("стек пуст")
	}
	value := stack.head.data
	stack.head = stack.head.next
	return value, nil
}

// Enqueue добавляет элемент в конец очереди.
func (queue *Queue) Enqueue(value string) {
	node := &Node{data: value}
	if queue.head == nil {
		queue.head = node
		queue.tail = node
	} else {
		queue.tail.next = node
		queue.tail = node
	}
}

// Dequeue извлекает элемент из начала очереди и возвращает его значение.
func (queue *Queue) Dequeue() (string, error) {
	if queue.head == nil {
		return "", errors.New("очередь пуста")
	}
	value := queue.head.data
	queue.head = queue.head.next
	// Если после извлечения элемента очередь осталась пустой, обновляем указатель на хвост.
	if queue.head == nil {
		queue.tail = nil
	}
	return value, nil
}

// NewHashTable создает новую хеш-таблицу.
func NewHashTable() *HashTable {
	return &HashTable{}
}

// Put добавляет пару ключ:значение в хеш-таблицу.
func (ht *HashTable) Put(key, value string) {
	for i, entry := range ht.data {
		if entry.key == key {
			ht.data[i].value = value
			return
		}
	}
	ht.data = append(ht.data, hashTableEntry{key: key, value: value})
}

// Get возвращает значение по ключу из хеш-таблицы.
func (ht *HashTable) Get(key string) (string, bool) {
	for _, entry := range ht.data {
		if entry.key == key {
			return entry.value, true
		}
	}
	return "", false
}

// Delete удаляет запись из хеш-таблицы по ключу.
func (ht *HashTable) Delete(key string) {
	for i, entry := range ht.data {
		if entry.key == key {
			ht.data = append(ht.data[:i], ht.data[i+1:]...)
			return
		}
	}
}

func main() {
	// Определение флагов для аргументов командной строки
	stackFile := flag.String("stack", "stack.txt", "Файл для стека")
	queueFile := flag.String("queue", "queue.txt", "Файл для очереди")
	setFile := flag.String("set", "set.txt", "Файл для множества")
	tableFile := flag.String("table", "hash_table.txt", "Файл для хеш-таблицы")

	// Анализ аргументов командной строки
	flag.Parse()

	// Создание объектов стека, очереди, множества и хеш-таблицы
	stack := &Stack{}
	queue := &Queue{}
	set := NewSet()
	hashTable := NewHashTable()

	// Загрузка данных из файлов, используя значения из аргументов командной строки
	if err := loadStackFromFile(stack, *stackFile); err != nil {
		fmt.Println("Ошибка загрузки данных стека:", err)
	}

	if err := loadQueueFromFile(queue, *queueFile); err != nil {
		fmt.Println("Ошибка загрузки данных очереди:", err)
	}

	if err := loadSetFromFile(set, *setFile); err != nil {
		fmt.Println("Ошибка загрузки данных множества:", err)
	}

	if err := loadHashTableFromFile(hashTable, "hash_table.txt"); err != nil {
		fmt.Println("Ошибка загрузки данных хеш-таблицы:", err)
	}

	fmt.Println("Программа для работы с данными (стек, очередь, множество, хеш-таблица)")
	for {
		fmt.Println("\nМеню:")
		fmt.Println("1. Работа со стеком")
		fmt.Println("2. Работа с очередью")
		fmt.Println("3. Работа с множеством")
		fmt.Println("4. Работа с хеш-таблицей")
		fmt.Println("5. Выход")

		fmt.Print("Выберите опцию: ")

		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch choice {
		case 1:
			handleStackMenu(stack)
		case 2:
			handleQueueMenu(queue)
		case 3:
			handleSetMenu(set)

			// Сохранение данных множества в файл после его модификации
			if err := saveSetToFile(set, "set.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных множества:", err)
			}

		case 4:
			handleHashTableMenu(hashTable, tableFile)

			// Сохранение данных хеш-таблицы в файл после его модификации
			if err := saveHashTableToFile(hashTable, "hash_table.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных хеш-таблицы:", err)
			}
		case 5:
			// Сохранение данных в файл перед выходом из программы
			if err := saveStackToFile(stack, "stack.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных стека:", err)
			}

			if err := saveQueueToFile(queue, "queue.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных очереди:", err)
			}

			// Сохранение данных множества в файл перед выходом из программы
			if err := saveSetToFile(set, "set.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных множества:", err)
			}

			// Сохранение данных хеш-таблицы в файл перед выходом из программы
			if err := saveHashTableToFile(hashTable, "hash_table.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных хеш-таблицы:", err)
			}

			fmt.Println("Выход из программы.")
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

func handleStackMenu(stack *Stack) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню стека:")
		fmt.Println("1. Добавить элемент")
		fmt.Println("2. Извлечь элемент")
		fmt.Println("3. Вернуться в главное меню")

		fmt.Print("Выберите опцию: ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Введите элемент для добавления: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			value = strings.TrimSuffix(value, "\n") // Удаление символа новой строки
			stack.Push(value)
			fmt.Println("Элемент добавлен в стек.")

			// Сохранение стека в файл после добавления элемента
			if err := saveStackToFile(stack, "stack.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных стека:", err)
			}
		case 2:
			value, err := stack.Pop()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Извлеченный элемент:", value)

				// После извлечения элемента из стека, обновите файл с данными стека
				if err := saveStackToFile(stack, "stack.txt"); err != nil {
					fmt.Println("Ошибка сохранения данных стека:", err)
				}
			}
		case 3:
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

func handleQueueMenu(queue *Queue) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню очереди:")
		fmt.Println("1. Добавить элемент")
		fmt.Println("2. Извлечь элемент")
		fmt.Println("3. Вернуться в главное меню")

		fmt.Print("Выберите опцию: ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}
		switch choice {
		case 1:
			fmt.Print("Введите элемент для добавления: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			queue.Enqueue(value)
			fmt.Println("Элемент добавлен в очередь.")

			// Сохранение очереди в файл после добавления элемента
			if err := saveQueueToFile(queue, "queue.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных очереди:", err)
			}
		case 2:
			value, err := queue.Dequeue()
			if err != nil {
				fmt.Println("Ошибка:", err)
			} else {
				fmt.Println("Извлеченный элемент:", value)

				// После извлечения элемента из очереди, обновите файл с данными очереди
				if err := saveQueueToFile(queue, "queue.txt"); err != nil {
					fmt.Println("Ошибка сохранения данных очереди:", err)
				}
			}
		case 3:
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

func handleSetMenu(set *Set) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню множества:")
		fmt.Println("1. Добавить элемент")
		fmt.Println("2. Проверить наличие элемента")
		fmt.Println("3. Удалить элемент")
		fmt.Println("4. Вернуться в главное меню")

		fmt.Print("Выберите опцию: ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Введите элемент для добавления: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			if set.Contains(value) {
				fmt.Println("Ошибка: Вы указали существующий элемент.")
			} else {
				set.Add(value)
				fmt.Println("Элемент добавлен в множество.")

				// Сохранение данных множества в файл после добавления элемента
				if err := saveSetToFile(set, "set.txt"); err != nil {
					fmt.Println("Ошибка сохранения данных множества:", err)
				}
			}
		case 2:
			fmt.Print("Введите элемент для проверки: ")
			valueToCheck, _ := reader.ReadString('\n')
			valueToCheck = strings.TrimSpace(valueToCheck)

			if set.Contains(valueToCheck) {
				fmt.Println("Элемент найден в множестве.")
			} else {
				fmt.Println("Элемент не найден в множестве.")
			}

		case 3:
			fmt.Print("Введите элемент для удаления: ")
			valueToDelete, _ := reader.ReadString('\n')
			valueToDelete = strings.TrimSpace(valueToDelete)
			if set.Contains(valueToDelete) {
				set.Remove(valueToDelete)
				fmt.Println("Элемент удален из множества.")

				// Удаление элемента из файла множества после удаления из множества
				if err := deleteElementFromFile("set.txt", valueToDelete); err != nil {
					fmt.Println("Ошибка удаления элемента из файла множества:", err)
				}
			} else {
				fmt.Println("Ошибка: Элемент не найден в множестве.")
			}
			// Ваш код удаления элемента из множества

		case 4:
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

func handleHashTableMenu(hashTable *HashTable, tableFile *string) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nМеню хеш-таблицы:")
		fmt.Println("1. Добавить элемент")
		fmt.Println("2. Удалить элемент")
		fmt.Println("3. Прочитать элемент")
		fmt.Println("4. Вернуться в главное меню")

		fmt.Print("Выберите опцию: ")
		var choice int
		_, err := fmt.Scanln(&choice)
		if err != nil {
			fmt.Println("Ошибка ввода:", err)
			continue
		}

		switch choice {
		case 1:
			fmt.Print("Введите ключ для добавления: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			if _, found := hashTable.Get(key); found {
				fmt.Println("Ошибка: Вы указали существующий ключ.")
			} else {
				fmt.Print("Введите значение для добавления: ")
				value, _ := reader.ReadString('\n')
				value = strings.TrimSpace(value)
				hashTable.Put(key, value)
				fmt.Println("Элемент добавлен в хеш-таблицу.")

				// Сохранение данных хеш-таблицы в файл после добавления элемента
				if err := saveHashTableToFile(hashTable, *tableFile); err != nil {
					fmt.Println("Ошибка сохранения данных хеш-таблицы:", err)
				}
			}
		case 2:
			fmt.Print("Введите ключ для удаления: ")
			keyToDelete, _ := reader.ReadString('\n')
			keyToDelete = strings.TrimSpace(keyToDelete)
			if _, found := hashTable.Get(keyToDelete); found {
				hashTable.Delete(keyToDelete)
				fmt.Println("Элемент удален из хеш-таблицы.")

				// Удаление элемента из файла хеш-таблицы после удаления из хеш-таблицы
				if err := deleteElementFromFile(*tableFile, keyToDelete); err != nil {
					fmt.Println("Ошибка удаления элемента из файла хеш-таблицы:", err)
				}
			} else {
				fmt.Println("Ошибка: Элемент не найден в хеш-таблице.")
			}

		case 3:
			fmt.Print("Введите ключ для чтения: ")
			keyToRead, _ := reader.ReadString('\n')
			keyToRead = strings.TrimSpace(keyToRead)

			value, found := hashTable.Get(keyToRead)
			if found {
				fmt.Printf("Значение для ключа %s: %s\n", keyToRead, value)
			} else {
				fmt.Println("Элемент не найден в хеш-таблице.")
			}

		case 4:
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

// Функция для сохранения данных стека в файл
func saveStackToFile(stack *Stack, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := stack.head
	for current != nil {
		_, err := fmt.Fprintln(file, current.data)
		if err != nil {
			return err
		}
		current = current.next
	}

	return nil
}

// Функция для сохранения данных очереди в файл
func saveQueueToFile(queue *Queue, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	current := queue.head
	for current != nil {
		_, err := fmt.Fprintln(file, current.data)
		if err != nil {
			return err
		}
		current = current.next
	}

	return nil
}

// Функция для сохранения данных множества в файл
func saveSetToFile(set *Set, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range set.data {
		_, err := fmt.Fprintln(file, value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Функция для сохранения данных хеш-таблицы в файл
func saveHashTableToFile(hashTable *HashTable, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, entry := range hashTable.data {
		_, err := fmt.Fprintf(file, "%s:%s\n", entry.key, entry.value)
		if err != nil {
			return err
		}
	}

	return nil
}

// Функция для загрузки данных стека из файла
func loadStackFromFile(stack *Stack, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		stack.Push(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Функция для загрузки данных множества из файла
func loadSetFromFile(set *Set, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		set.Add(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Функция для загрузки данных очереди из файла
func loadQueueFromFile(queue *Queue, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		queue.Enqueue(line)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

// Функция для загрузки данных хеш-таблицы из файла
func loadHashTableFromFile(hashTable *HashTable, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ":")
		if len(parts) == 2 {
			key := parts[0]
			value := parts[1]
			hashTable.Put(key, value)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
func deleteElementFromFile(filename, elementToDelete string) error {
	// Чтение содержимого файла в массив строк
	lines, err := readLinesFromFile(filename)
	if err != nil {
		return err
	}

	// Поиск элемента в массиве строк и удаление его, если он существует
	for i, line := range lines {
		if line == elementToDelete {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}

	// Запись обновленных данных обратно в файл
	err = writeLinesToFile(filename, lines)
	if err != nil {
		return err
	}

	return nil
}
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}

// Функция для записи массива строк в файл
func writeLinesToFile(filename string, lines []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range lines {
		_, err := fmt.Fprintln(file, line)
		if err != nil {
			return err
		}
	}

	return nil
}

// Функция для чтения строк из файла и возврата их в виде массива
