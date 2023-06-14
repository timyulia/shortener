# API по созданию сокращённых ссылок
## Запуск:
Выбрать режим работы можно в файле .env:
* MODE=db
* MODE=inMemory
1. Запуск контейнеров с базой данных и приложением. Для первого запуска
нужен --build.
```bash
docker-compose up --build
```
2. Миграция для базы данных из папки migrations.
```bash
make migrate
```

## Использование
1. Запрос: http://localhost:8000/  POST
```json
{
    "URL":"https://www.google.com"
}
```
Ответ:
```json
{
    "shortURL": "CL_rVxjFkR"
}
```

2. Запрос: http://localhost:8000/CL_rVxjFkR GET__
Ответ:
```json
{
    "longURL": "https://www.google.com"
}
```

Proto файл был создан, но grpc не реализован.

## Тестирование с выводом покрытия кода тестами

```bash
make test-cover
```

## Алгоритм сокращения

Детерминированный алгоритм на основе хеш-функции и base-63. По заданной ссылке считается хеш-сумма sha224. По ней формируется 10 символов путем перевода
некоторых бит в base-63. Так как хеш-функция имеет равномерное распределение сокращение ее значений увеличивает веротяность коллизий только за счет
уменьшения количество возможнох вариантов (вместо 2^224 - 63^10<=64^10=(2^6)^10).

Также проводится проверка на коллизии, если полученная короткая ссылка уже существует, но привязана к другой длинной ссылке, генерируется новый варинат, а на вход функции
подается полученная ранее короткая ссылка.

## Задание
Ссылка должна быть:
* Уникальной; на один оригинальный URL должна ссылаться только одна сокращенная ссылка;
* Длиной 10 символов;
* Из символов латинского алфавита в нижнем и верхнем регистре, цифр нижнего подчеркивания.

Сервис должен быть написан на Go и принимать следующие запросы по http:
1. **Метод Post**, который будет сохранять оригинальный URL в базе и возвращать сокращённый.
2. **Метод Get**, который будет принимать сокращённый URL и возвращать оригинальный.

**Условие со звёздочкой:**
Сделать работу сервиса через GRPC, то есть составить proto и реализовать сервис с двумя соответствующими эндпойнтами.

Решение должно соответствовать условиям:
* Сервис распространён в виде Docker-образа;
* В качестве хранилища ожидаем in-memory решение и PostgreSQL. Какое хранилище использовать, указывается параметром при запуске сервиса;
* Реализованный функционал покрыт Unit-тестами.
