# День 01 - Go Bootcamp: Comparing Incomparable

## Оглавление

* [Описание проекта](#описание-проекта)
* [Функциональность](#функциональность)

  * [Упражнение 00: Чтение](#упражнение-00-чтение)
  * [Упражнение 01: Оценка повреждений](#упражнение-01-оценка-повреждений)
  * [Упражнение 02: После вечеринки](#упражнение-02-после-вечеринки)
* [Технологии](#технологии)
* [Структура проекта](#структура-проекта)
* [Примеры использования](#примеры-использования)

---

## Описание проекта

Проект реализует набор утилит для работы с различными форматами данных и файловыми системами. Основные задачи включают:

* чтение и преобразование между XML и JSON
* сравнение структур данных рецептов
* эффективное сравнение снимков файловых систем

Проект разработан для решения реальных задач сравнения данных в различных форматах.

---

## Функциональность

### Упражнение 00: Чтение

**Конвертация между XML и JSON форматами:**

* Автоматическое определение формата по расширению файла
* Преобразование XML → JSON с форматированием
* Преобразование JSON → XML с форматированием
* Единое представление данных независимо от формата

**Поддерживаемые команды:**

```bash
go run readDB.go -f original_database.xml
go run readDB.go -f stolen_database.json
```

---

### Упражнение 01: Оценка повреждений

**Сравнение баз данных рецептов:**

* Обнаружение добавленных/удалённых рецептов
* Сравнение времени приготовления
* Анализ изменений в ингредиентах:

  * Добавление/удаление ингредиентов
  * Изменение количества
  * Изменение единиц измерения

**Поддержка XML и JSON форматов**

**Поддерживаемые команды:**

```bash
go run compareDB.go --old original_database.xml --new stolen_database.json
```

---

### Упражнение 02: После вечеринки

**Сравнение снимков файловых систем:**

* Обнаружение добавленных/удалённых файлов
* Оптимизированная работа с большими файлами
* Потоковая обработка без загрузки в память
* Простой текстовый вывод изменений

**Поддерживаемые команды:**

```bash
go run compareFS.go --old snapshot1.txt --new snapshot2.txt
```

---

## Технологии

**Язык программирования:** Go

**Ключевые пакеты:**

* `encoding/xml` — работа с XML
* `encoding/json` — работа с JSON
* `flag` — обработка аргументов командной строки
* `bufio` — потоковое чтение файлов
* `os` — работа с файловой системой

**Алгоритмы:**

* Рекурсивное сравнение структур
* Хеширование для эффективного сравнения
* Потоковая обработка больших данных

**Форматирование:**

* Pretty-print для XML/JSON
* Отступы в 4 пробела

---

## Структура проекта

```
Go_Day01/
├── README_RU.md
└── src/
    ├── ex00/                          # Конвертер XML/JSON
    │   ├── original_database.xml      # Пример XML-файла
    │   ├── readDB.go                  # Основная программа
    │   └── stolen_database.json       # Пример JSON-файла
    ├── ex01/                          # Сравнение рецептов
    │   ├── compareDB.go               # Основная программа
    │   ├── original_database.xml
    │   └── stolen_database.json
    └── ex02/                          # Сравнение файловых систем
        ├── compareFS.go               # Основная программа
        ├── snapshot1.txt              # Пример снимка ФС
        └── snapshot2.txt              # Пример снимка ФС
```

---

## Примеры использования

### Упражнение 00: Чтение

```bash
# Конвертация XML в JSON
cd src/ex00
go run readDB.go -f original_database.xml

# Конвертация JSON в XML
go run readDB.go -f stolen_database.json
```

**Пример вывода (XML → JSON):**

```json
{
    "cake": [
        {
            "name": "Red Velvet Strawberry Cake",
            "time": "40 min",
            "ingredients": [
                {
                    "name": "Flour",
                    "count": "3",
                    "unit": "cups"
                },
                ...
            ]
        }
    ]
}
```

---

### Упражнение 01: Оценка повреждений

```bash
cd src/ex01
go run compareDB.go --old original_database.xml --new stolen_database.json
```

**Пример вывода:**

```
ADDED cake "Moonshine Muffin"
REMOVED cake "Blueberry Muffin Cake"
CHANGED cooking time for cake "Red Velvet Strawberry Cake" - "45 min" instead of "40 min"
ADDED ingredient "Coffee beans" for cake "Red Velvet Strawberry Cake"
REMOVED ingredient "Vanilla extract" for cake "Red Velvet Strawberry Cake"
CHANGED unit for ingredient "Flour" for cake "Red Velvet Strawberry Cake" - "mugs" instead of "cups"
```

---

### Упражнение 02: После вечеринки

```bash
cd src/ex02
go run compareFS.go --old snapshot1.txt --new snapshot2.txt
```

**Пример вывода:**

```
ADDED /etc/systemd/system/very_important/stash_location.jpg
REMOVED /var/log/browser_history.txt
ADDED /Users/baker/secret_recipes/
```
