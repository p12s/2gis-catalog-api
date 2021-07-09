![License](https://img.shields.io/github/license/p12s/ispring-todo-list-api)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/p12s/2gis-catalog-api?style=plastic)
[![Coverage Status](https://codecov.io/gh/p12s/2gis-catalog-api/branch/master/graph/badge.svg?token=sTWAW1J7hW)](https://codecov.io/gh/p12s/2gis-catalog-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/p12s/2gis-catalog-api)](https://goreportcard.com/report/github.com/p12s/2gis-catalog-api)
<img src="https://github.com/p12s/2gis-catalog-api/workflows/lint-build/badge.svg?branch=master">

**Внимание:** *Тестовое задание найдено на просторах github-а. Для обучения и тренировки, попробовал решить ее в меру своего понимания. На ревью не отправлял, за оптимальность не ручаюсь.*

# Упрощенная версия 2ГИС справочника

## Задача
Создать REST-API сервис, предполагаемая аудитория которого - 1 000 000 пользователей в месяц.  
Размер базы фирм составляет порядка 100 000 записей.     
Подробнее [здесь](task.md)

## Реализовать функциональные возможности
- ❌ выдачу всех организаций находящихся в конкретном здании
- ❌ выдачу списка всех организаций, которые относятся к указанной рубрике
- ❌ выдачу информации об организациях по их идентификаторам
- ❌ добавление в справочник здания вместе с организациями, которые в ней находятся

## Реализовать технические моменты
- ❌ формат маршрутов для доступа к методам, а также формат ответа и запросов
- ❌ спроектировать базу данных + реализовать скрипт ее формирования
- ❌ реализовать API согласно ТЗ
- ❌ подготовить тестовые данные (дамп базы, скрипт для генерации тестового набора данных)
- ❌ инструкцию развертывания на локальной машине

## Используемые инструменты
- ✅ Язык:  Go (желательно) или PHP
  **Go**
- ❌ БД: Любая реляционная (например, PostgreSQL или MySQL)
- ❇️ [Go-kit](https://github.com/go-kit/kit) как библиотека для микросервисов
- ❇️ [Go-chi](https://github.com/go-chi/chi) для роутинга
mongo??


## Добавил от себя
- бейдж с %-ом покрытия unit-тестами (сейчас 0%)
- бейдж с результатом запуска тестов и пробной сборки в Github Actions для master-ветки
