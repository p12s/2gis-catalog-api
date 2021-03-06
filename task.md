# Задание

Реализовать упрощенную версию 2GIS справочника.  
Предполагаемая аудитория - 1 000 000 пользователей в месяц.  
Размер базы фирм составляет порядка 100 000 записей.
Взаимодействие с пользователем происходит посредством HTTP запросов к API серверу.
Все ответы представляют собой JSON объекты.

## Наш справочник будет иметь 3 вида объектов  
1) *Фирма*
2) *Здание*
3) *Рубрика*
    
**Фирма** представляет собой карточку организации в справочнике и должна содержать в себе следующую информацию:
- название (Например, ООО “Рога и Копыта”)
- несколько номеров телефонов (2-222-222, 3-333-333, 8-923-666-13-13)
- находится в одном  конкретном здании (Например, Блюхера, 32/1)
- может относиться к нескольким рубрикам (Например, “Полуфабрикаты оптом”, “Мясная продукция”)

**Здание** содержит в себе как минимум информацию о конкретном здании, а именно:
- адрес здания
- географические координаты его местоположения (в виде долгота/широта, или в виде линейных координат на декартовой плоскости x/y)

**Рубрика**
- позволяет классифицировать род деятельности фирм в каталоге
- Имеют название и могут в древовидном виде вкладываться друг в друга. 
- В простом случае реализации уровень вложенности будет 2-3 рубрики, в сложном - произвольный.
  
Вот пример возможных рубрик и их иерархии:
- Еда
    - Полуфабрикаты оптом
    - Мясная продукция
- Автомобили
    - Грузовые
    - Легковые
        - Запчасти для подвески
        - Шины/Диски
- Спорт
    - ...

## Реализовать функциональные возможности
- выдачу всех организаций находящихся в конкретном здании
- выдачу списка всех организаций, которые относятся к указанной рубрике
- выдачу информации об организациях по их идентификаторам
- добавление в справочник здания вместе с организациями, которые в ней находятся

## Реализовать технические моменты
- Формат маршрутов для доступа к методам, а также формат ответа и запросов
- Спроектировать базу данных + реализовать скрипт ее формирования
- Реализовать API согласно ТЗ
- Подготовить тестовые данные (дамп базы, скрипт для генерации тестового набора данных)
- Инструкцию развертывания на локальной машине

## Используемые технологии
- Язык:  Go (желательно) или PHP
- БД: Любая реляционная (например, PostgreSQL или MySQL)