# Game Slots Service

Простой сервис для работы с игровыми слотами. Поддерживает загрузку конфигурации для катушек, линий и выплат, а 
также расчет комбинаций.

# Установка и Запуск

Клонировать репозиторий выполнив команду:

`git clone https://github.com/DrrBurger/game-slots`

Для запуска сервиса:
- В docker: выполнить команду `make dc`
- Локально: выполнить команду `make run`  

Для запуска тестов необходимо выполнить одну из команд:
1. `make test` запуск всех тестов
2. `make cover` для запуска тестов с покрытием
3. `make cover-html` для запуска тестов с покрытием и получения отчёта в html формате

Для запуска линтера необходимо выполнить команду `make lint`

Список всех команд можно получить выполнив команду `make help`

# Примеры запросов

### Загрузка конфигурации для катушек:
```
curl -X POST 
     -H "Content-Type: application/json" \
     -d '[
       ["A", "B", "C", "D", "E"],
       ["F", "A", "F", "B", "C"],
       ["D", "E", "A", "G", "A"]
     ]' \
     "http://localhost:8080/upload?type=reels&name=default"
```
### Загрузка конфигурации для линий:
```
curl -X POST \
     -H "Content-Type: application/json" \
     -d '[
    {
        "line": 1, 
        "positions": [
            {"row": 0, "col": 0},
            {"row": 1, "col": 1},
            {"row": 2, "col": 2},
            {"row": 1, "col": 3},
            {"row": 0, "col": 4}
        ]
    },
    {
        "line": 2, 
        "positions": [
            {"row": 2, "col": 0},
            {"row": 1, "col": 1},
            {"row": 0, "col": 2},
            {"row": 1, "col": 3},
            {"row": 2, "col": 4}
        ]
    },
    {
        "line": 3, 
        "positions": [
            {"row": 1, "col": 0},
            {"row": 2, "col": 1},
            {"row": 1, "col": 2},
            {"row": 0, "col": 3},
            {"row": 1, "col": 4}
        ]
    }
    ]' \
     "http://localhost:8080/upload?type=lines&name=default"
```

### Загрузка конфигурации для выплат:
```
curl -X POST \
     -H "Content-Type: application/json" \
     -d '[
    {
        "symbol": "A",
        "payout": [0, 0, 50, 100, 200]
    },
    {
        "symbol": "B",
        "payout": [0, 0, 40, 80, 160]
    },
    {
        "symbol": "C",
        "payout": [0, 0, 30, 60, 120]
    },
    {
        "symbol": "D",
        "payout": [0, 0, 20, 40, 80]
    },
    {
        "symbol": "E",
        "payout": [0, 0, 10, 20, 40]
    },
    {
        "symbol": "F",
        "payout": [0, 0, 5, 10, 20]
    },
    {
        "symbol": "G",
        "payout": [0, 0, 2, 5, 10]
    }
    ]' \
     "http://localhost:8080/upload?type=payouts&name=default"
```

### Расчет комбинаций:
После загрузки всех конфигураций (катушки, линии, выплаты)
```
curl "http://localhost:8080/calculate?reels=default&lines=default&payouts=default"
```

# Метрики:
Вы можете просмотреть метрики Prometheus, отправив GET запрос на:
```
http://localhost:8080/metrics
```

# Примечание
- В качестве хранилища использовано inmemory хранилище. Для реального проекта заменить на Postgres или иное.
- Конфигурации не вынесены в отдельный файл. 
- Не реализован Graceful shutdown