package main

import (
	"errors"
	"fmt"
	"math"
)

// Book описывает книгу.
type Book struct {
	id     int
	title  string
	author string
	year   int
	size   int
	rate   float64
}

// NewBook создает новую книгу.
func NewBook(id int, title, author string, year, size int, rate float64) *Book {
	return &Book{
		id:     id,
		title:  title,
		author: author,
		year:   year,
		size:   size,
		rate:   rate,
	}
}

// GetID возвращает идентификатор книги.
func (b *Book) GetID() int {
	return b.id
}

// SetID устанавливает новый идентификатор книги.
func (b *Book) SetID(newID int) {
	b.id = newID
}

// GetTitle возвращает название книги.
func (b *Book) GetTitle() string {
	return b.title
}

// SetTitle устанавливает новое название книги.
func (b *Book) SetTitle(newTitle string) {
	b.title = newTitle
}

// GetAuthor возвращает автора книги.
func (b *Book) GetAuthor() string {
	return b.author
}

// SetAuthor устанавливает нового автора книги.
func (b *Book) SetAuthor(newAuthor string) {
	b.author = newAuthor
}

// GetYear возвращает год издания книги.
func (b *Book) GetYear() int {
	return b.year
}

// SetYear устанавливает новый год издания книги.
func (b *Book) SetYear(newYear int) {
	b.year = newYear
}

// GetSize возвращает размер книги.
func (b *Book) GetSize() int {
	return b.size
}

// SetSize устанавливает новый размер книги.
func (b *Book) SetSize(newSize int) {
	b.size = newSize
}

// GetRate возвращает рейтинг книги.
func (b *Book) GetRate() float64 {
	return b.rate
}

// SetRate устанавливает новый рейтинг книги.
func (b *Book) SetRate(newRate float64) {
	b.rate = newRate
}

// ComparisonMode описывает режимы сравнения книг.
type ComparisonMode int

const (
	CompareByYear ComparisonMode = iota
	CompareBySize
	CompareByRate
)

var comparisonModes = [...]string{
	"Год",
	"Размер",
	"Рейтинг",
}

// String возвращает текстовое представление режима сравнения.
func (cm ComparisonMode) String() string {
	return comparisonModes[cm]
}

// BookComparer реализует сравнение книг по различным критериям.
type BookComparer struct {
	mode ComparisonMode
}

// NewBookComparer создает новый объект для сравнения книг.
func NewBookComparer(mode ComparisonMode) *BookComparer {
	return &BookComparer{mode: mode}
}

// Compare сравнивает две книги согласно установленному режиму.
func (bc *BookComparer) Compare(book1, book2 *Book) bool {
	switch bc.mode {
	case CompareByYear:
		return book1.GetYear() > book2.GetYear()
	case CompareBySize:
		return book1.GetSize() > book2.GetSize()
	case CompareByRate:
		return math.Round(book1.GetRate()*100)/100 > math.Round(book2.GetRate()*100)/100
	default:
		panic(errors.New("неизвестный режим сравнения"))
	}
}

func main() {
	book1 := NewBook(1, "Книга 1", "Автор 1", 2020, 200, 4.7)
	book2 := NewBook(2, "Книга 2", "Автор 2", 2019, 300, 4.5)

	comparer := NewBookComparer(CompareByYear)
	result := comparer.Compare(book1, book2)
	fmt.Printf("Сравниваем книги по году: %t\n", result)

	comparer = NewBookComparer(CompareBySize)
	result = comparer.Compare(book1, book2)
	fmt.Printf("Сравниваем книги по размеру: %t\n", result)

	comparer = NewBookComparer(CompareByRate)
	result = comparer.Compare(book1, book2)
	fmt.Printf("Сравниваем книги по рейтингу: %t\n", result)
}
