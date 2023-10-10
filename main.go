package main

import (
	"bufio"
	"errors"
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
	data map[string]bool
}

// HashTable представляет хеш-таблицу.
type HashTable struct {
	data map[string]string
}

// NewSet создает новое множество.
func NewSet() *Set {
	return &Set{data: make(map[string]bool)}
}

// Add добавляет элемент в множество.
func (set *Set) Add(value string) {
	set.data[value] = true
}

// Contains проверяет, содержит ли множество указанный элемент.
func (set *Set) Contains(value string) bool {
	return set.data[value]
}

// Remove удаляет элемент из множества.
func (set *Set) Remove(value string) {
	delete(set.data, value)
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
	return &HashTable{data: make(map[string]string)}
}

// Put добавляет пару ключ:значение в хеш-таблицу.
func (ht *HashTable) Put(key, value string) {
	ht.data[key] = value
}

// Get возвращает значение по ключу из хеш-таблицы.
func (ht *HashTable) Get(key string) (string, bool) {
	value, ok := ht.data[key]
	return value, ok
}

// Delete удаляет запись из хеш-таблицы по ключу.
func (ht *HashTable) Delete(key string) {
	delete(ht.data, key)
}

func main() {
	// Создание объектов стека, очереди, множества и хеш-таблицы
	stack := &Stack{}
	queue := &Queue{}
	set := NewSet()
	hashTable := NewHashTable()

	// Загрузка данных из файлов
	if err := loadStackFromFile(stack, "stack.txt"); err != nil {
		fmt.Println("Ошибка загрузки данных стека:", err)
	}

	if err := loadQueueFromFile(queue, "queue.txt"); err != nil {
		fmt.Println("Ошибка загрузки данных очереди:", err)
	}

	// Загрузка данных множества из файла
	if err := loadSetFromFile(set, "set.txt"); err != nil {
		fmt.Println("Ошибка загрузки данных множества:", err)
	}

	// Загрузка данных хеш-таблицы из файла
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
			handleHashTableMenu(hashTable)

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
			set.Add(value)
			fmt.Println("Элемент добавлен в множество.")

			// Сохранение данных множества в файл после добавления элемента
			if err := saveSetToFile(set, "set.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных множества:", err)
			}
		case 2:
			fmt.Print("Введите элемент для проверки: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			if set.Contains(value) {
				fmt.Println("Элемент найден в множестве.")
			} else {
				fmt.Println("Элемент не найден в множестве.")
			}
		case 3:
			fmt.Print("Введите элемент для удаления: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			set.Remove(value)
			fmt.Println("Элемент удален из множества.")

			// Сохранение данных множества в файл после удаления элемента
			if err := saveSetToFile(set, "set.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных множества:", err)
			}
		case 4:
			return
		default:
			fmt.Println("Некорректный выбор. Попробуйте ещё раз.")
		}
	}
}

func handleHashTableMenu(hashTable *HashTable) {
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
			fmt.Print("Введите значение для добавления: ")
			value, _ := reader.ReadString('\n')
			value = strings.TrimSpace(value)
			hashTable.Put(key, value)
			fmt.Println("Элемент добавлен в хеш-таблицу.")

			// Сохранение данных хеш-таблицы в файл после добавления элемента
			if err := saveHashTableToFile(hashTable, "hash_table.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных хеш-таблицы:", err)
			}
		case 2:
			fmt.Print("Введите ключ для удаления: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			hashTable.Delete(key) // Вызывайте Delete для удаления элемента
			fmt.Println("Элемент удален из хеш-таблицы.")

			// Сохранение данных хеш-таблицы в файл после удаления элемента
			if err := saveHashTableToFile(hashTable, "hash_table.txt"); err != nil {
				fmt.Println("Ошибка сохранения данных хеш-таблицы:", err)
			}
		case 3:
			fmt.Print("Введите ключ для чтения: ")
			key, _ := reader.ReadString('\n')
			key = strings.TrimSpace(key)
			value, found := hashTable.Get(key)
			if found {
				fmt.Printf("Значение для ключа %s: %s\n", key, value)
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

	for key := range set.data {
		_, err := fmt.Fprintln(file, key)
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

	for key, value := range hashTable.data {
		_, err := fmt.Fprintf(file, "%s=%s\n", key, value)
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
		parts := strings.Split(line, "=")
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
