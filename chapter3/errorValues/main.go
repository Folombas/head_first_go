package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

// 🎯 1. БАЗОВЫЕ ОШИБКИ С errors.New()
var (
	ErrInvalidInput  = errors.New("неверный ввод")
	ErrFileNotFound  = errors.New("файл не найден")
	ErrNetworkFailed = errors.New("сетевая ошибка")
)

// 🎯 2. КАСТОМНЫЕ ТИПЫ ОШИБОК
type ValidationError struct {
	Field   string
	Value   interface{}
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("ошибка валидации поля '%s': %s (значение: %v)",
		e.Field, e.Message, e.Value)
}

type APIError struct {
	StatusCode int
	URL        string
	Message    string
	Timestamp  time.Time
}

func (e APIError) Error() string {
	return fmt.Sprintf("API ошибка [%d] %s: %s (%s)",
		e.StatusCode, e.URL, e.Message, e.Timestamp.Format("2006-01-02 15:04:05"))
}

// 🎯 3. ФУНКЦИЯ, ВОЗВРАЩАЮЩАЯ ОШИБКИ
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("деление на ноль")
	}
	return a / b, nil
}

func validateUserAge(age int) error {
	if age < 0 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "возраст не может быть отрицательным",
		}
	}
	if age < 18 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "только для взрослых (18+)",
		}
	}
	if age > 150 {
		return ValidationError{
			Field:   "age",
			Value:   age,
			Message: "возраст слишком большой",
		}
	}
	return nil
}

func readConfigFile(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("%w: %s", ErrFileNotFound, filename)
	}
	// Симуляция чтения файла
	return nil
}

// 🎯 4. ОБЕРТЫВАНИЕ ОШИБОК (ERROR WRAPPING)
func processUserData(data map[string]string) error {
	ageStr, ok := data["age"]
	if !ok {
		return fmt.Errorf("processUserData: %w", ErrInvalidInput)
	}

	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return fmt.Errorf("processUserData: преобразование возраста: %w", err)
	}

	if err := validateUserAge(age); err != nil {
		return fmt.Errorf("processUserData: валидация: %w", err)
	}

	return nil
}

// 🎯 5. РАБОТА С МНОЖЕСТВЕННЫМИ ОШИБКАМИ
type MultiError struct {
	Errors []error
}

func (e MultiError) Error() string {
	var messages []string
	for _, err := range e.Errors {
		messages = append(messages, err.Error())
	}
	return fmt.Sprintf("несколько ошибок (%d): %v", len(e.Errors), messages)
}

func validateUserProfile(profile map[string]string) error {
	var errs []error

	if name, ok := profile["name"]; !ok || name == "" {
		errs = append(errs, ValidationError{Field: "name", Message: "имя обязательно"})
	}

	if ageStr, ok := profile["age"]; ok {
		if age, err := strconv.Atoi(ageStr); err == nil {
			if err := validateUserAge(age); err != nil {
				errs = append(errs, err)
			}
		} else {
			errs = append(errs, ValidationError{Field: "age", Message: "неверный формат"})
		}
	} else {
		errs = append(errs, ValidationError{Field: "age", Message: "возраст обязателен"})
	}

	if len(errs) > 0 {
		return MultiError{Errors: errs}
	}
	return nil
}

func main() {
	fmt.Println("🎯 ДЕМОНСТРАЦИЯ ОШИБОК В GO")
	fmt.Println("===========================")

	// 🎯 1. Базовые ошибки
	fmt.Println("\n1. БАЗОВЫЕ ОШИБКИ:")
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("❌ Ошибка деления: %v\n", err)
	} else {
		fmt.Printf("✅ Результат: %.2f\n", result)
	}

	// 🎯 2. Проверка типов ошибок
	fmt.Println("\n2. ПРОВЕРКА ТИПОВ ОШИБОК:")

	// errors.Is() - проверка конкретной ошибки
	err = readConfigFile("missing_file.txt")
	if errors.Is(err, ErrFileNotFound) {
		fmt.Printf("🔍 Найдена ошибка файла: %v\n", err)
	}

	// errors.As() - извлечение кастомного типа
	validationErr := validateUserAge(-5)
	if validationErr != nil {
		fmt.Printf("❌ Ошибка валидации: %v\n", validationErr)

		var valErr ValidationError
		if errors.As(validationErr, &valErr) {
			fmt.Printf("🔍 Детали ошибки: поле '%s', значение: %v\n",
				valErr.Field, valErr.Value)
		}
	}

	// 🎯 3. Обертывание ошибок
	fmt.Println("\n3. ОБЕРТЫВАНИЕ ОШИБОК:")
	userData := map[string]string{"age": "abc"}
	if err := processUserData(userData); err != nil {
		fmt.Printf("❌ Обработка данных: %v\n", err)

		// Проверка цепочки ошибок
		var numErr *strconv.NumError
		if errors.As(err, &numErr) {
			fmt.Printf("🔍 Числовая ошибка: %v\n", numErr.Err)
		}
	}

	// 🎯 4. Множественные ошибки
	fmt.Println("\n4. МНОЖЕСТВЕННЫЕ ОШИБКИ:")
	profile := map[string]string{"name": "", "age": "abc"}
	if err := validateUserProfile(profile); err != nil {
		fmt.Printf("❌ Ошибки профиля: %v\n", err)

		var multiErr MultiError
		if errors.As(err, &multiErr) {
			fmt.Printf("📊 Найдено ошибок: %d\n", len(multiErr.Errors))
			for i, e := range multiErr.Errors {
				fmt.Printf("   %d. %v\n", i+1, e)
			}
		}
	}

	// 🎯 5. PANIC и RECOVER
	fmt.Println("\n5. PANIC И RECOVER:")

	safeDivide := func(a, b float64) (result float64, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("паника перехвачена: %v", r)
			}
		}()

		if b == 0 {
			panic("критическая ошибка: деление на ноль!")
		}
		return a / b, nil
	}

	result, err = safeDivide(10, 0)
	if err != nil {
		fmt.Printf("🛡️ Безопасное деление: %v\n", err)
	}

	// 🎯 6. ПРАКТИЧЕСКИЕ СЦЕНАРИИ
	fmt.Println("\n6. ПРАКТИЧЕСКИЕ СЦЕНАРИИ:")

	// Симуляция API вызова
	apiError := APIError{
		StatusCode: 404,
		URL:        "https://api.example.com/user/123",
		Message:    "пользователь не найден",
		Timestamp:  time.Now(),
	}
	fmt.Printf("🌐 API ошибка: %v\n", apiError)

	// 🎯 ВЫВОДЫ
	fmt.Println("\n💡 КЛЮЧЕВЫЕ ПРИНЦИПЫ ОШИБОК В GO:")
	fmt.Println("✅ Ошибки - это значения, реализующие интерфейс error")
	fmt.Println("✅ Используйте errors.New() для простых ошибок")
	fmt.Println("✅ Создавайте кастомные типы для сложных сценариев")
	fmt.Println("✅ errors.Is() - для проверки конкретной ошибки")
	fmt.Println("✅ errors.As() - для извлечения кастомного типа")
	fmt.Println("✅ fmt.Errorf() с %w - для обертывания ошибок")
	fmt.Println("✅ defer/recover - для обработки паник")
	fmt.Println("❌ Не игнорируйте ошибки (никогда не используйте _ для ошибок!)")
}
