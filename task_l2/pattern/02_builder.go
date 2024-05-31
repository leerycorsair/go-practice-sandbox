package pattern

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Применимость:
// 1) Когда для объектов получается слишком большой конструктор.
// 2) Когда требуется создавать вариации объектов.
// 3) Когда объект создается пошагово.

// Плюсы:
// 1) Создание объекта пошагово.
// 2) Изоляция создания объекта.

// Минусы:
// 1) Разростание числа интерфейсов/типов.

type Database struct {
	schemes     []string
	db          []string
	tables      []string
	constraints []string
	indexes     []string
}

type DatabaseBuilder interface {
	CreateScheme()
	CreateDatabase()
	CreateTables()
	CreateConstraints()
	CreateIndexes()
	GetResult() *Database
}

type firstBuilder struct {
	db *Database
}

func (b *firstBuilder) CreateScheme() {
	b.db.schemes = []string{"f1", "f2", "f3"}
}

func (b *firstBuilder) CreateDatabase() {
	b.db.db = []string{"f1", "f2", "f3"}
}

func (b *firstBuilder) CreateTables() {
	b.db.tables = []string{"f1", "f2", "f3"}
}

func (b *firstBuilder) CreateConstraints() {
	b.db.constraints = []string{"f1", "f2", "f3"}
}

func (b *firstBuilder) CreateIndexes() {
	b.db.indexes = []string{"f1", "f2", "f3"}
}

func (b *firstBuilder) GetResult() *Database {
	return b.db
}

type secondBuilder struct {
	db *Database
}

func (b *secondBuilder) CreateScheme() {
	b.db.schemes = []string{"s1", "s2"}
}

func (b *secondBuilder) CreateDatabase() {
	b.db.db = []string{"s1", "s2"}
}

func (b *secondBuilder) CreateTables() {
	b.db.tables = []string{"s1", "s2"}
}

func (b *secondBuilder) CreateConstraints() {
	b.db.constraints = []string{"s1", "s2"}
}

func (b *secondBuilder) CreateIndexes() {
	b.db.indexes = []string{"s1", "s2"}
}

func (b *secondBuilder) GetResult() *Database {
	return b.db
}

type DatabaseDirector struct {
}

func (d *DatabaseDirector) Init(builder DatabaseBuilder) {
	builder.CreateScheme()
	builder.CreateDatabase()
	builder.CreateTables()
	builder.CreateConstraints()
	builder.CreateIndexes()
}
