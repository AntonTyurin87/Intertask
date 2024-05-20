![Go Report](https://goreportcard.com/badge/github.com/AntonTyurin87/Intertask) ![Repository Top Language](https://img.shields.io/github/languages/top/AntonTyurin87/Intertask) ![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/AntonTyurin87/Intertask)

<a href="https://codeclimate.com/github/AntonTyurin87/Intertask/maintainability"><img src="https://api.codeclimate.com/v1/badges/01c7db710db54263326d/maintainability" /></a>


# Intertask
Систему для добавления и чтения постов и комментариев
(разработано в рамках задачи перед собеседованием)

# Задание

Реализовать систему для добавления и чтения постов и комментариев с использованием GraphQL, аналогичную комментариям к постам на популярных платформах, таких как Хабр или Reddit.

<c> [подробное задание и критерии оценки](https://github.com/AntonTyurin87/Intertask/blob/main/docs/task_text.md) </c>

<c> [анализ заданяи](https://github.com/AntonTyurin87/Intertask/blob/main/docs/analisis.md) </c>

<c> [план выполнения](https://github.com/AntonTyurin87/Intertask/blob/main/docs/work_plan.md) </c>

<c> [план тестирования](https://github.com/AntonTyurin87/Intertask/blob/main/docs/test_plan.md) </c>

# Описание реализации Intertask

Приложение Intertask на вход принимает строку с POST запросом. На выходе предоставляет данные, в соответствии со структурой запроса. Формирование запросов, ответов и обновлений состояния подписки осуществляется по средствам GraphQL. При получении корректного запроса приложение обращается к базе данных и возвращает ответ. Параметры пагинации принимаются в запросе, а так же ограничены значениями по умолчанию.
При создании подписки уведомление о новом комментарии приходят в том случае, если комментировать выбранный пост возможно. Уведомления приходят в асинхронном режиме.
Возможна как работа приложения с базой данных PostgresQL, так и из памяти. Работа с хранилищем реализована через интерфейс "Blog".
Приложение Intertask может быть развернуто по средствам Docker.

# Docker параметры

* Настройка выбора пустой базы данных или с предзаписанными данными для ручного тестирования.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/DB.jpg">
</p>

* Настройки для выбора режима работы из памияти или из базы данных PostgresQL. По умолчанию переменная "IN_MEMORY=false". При задании значения "IN_MEMORY=true" приложение не подключается к PostgresQL.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/Chenge.jpg">
</p>

# Работа с Intertask по средствам GraphQL Playground for Chrome

#### <c> [Схема приложения из Playground](https://github.com/AntonTyurin87/Intertask/blob/main/docs/introspectionSchema.json) </c>

При настройках по умолчанию подключение на http://localhost:8080/graphql

### * Запрос на все посты, с ограничением по количеству и по номеру первого в выводе.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/query_posts.jpg">
</p>

### * Запрос на один пост с комментариями, с ограничением по количеству и по номеру первого в выводе.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/query_post.jpg">
</p>

### * Запрос на создание поста.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/mutatuin_createpost.jpg">
</p>

### * Запрос на изменение возможности комментирования поста.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/mutatuin_commentstatus.jpg">
</p>

### * Запрос на создание комментария к посту.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/mutatuin_createcomment.jpg">
</p>

### * Запрос на создание потписки к посту.
<p align="left">
  <img src="https://github.com/AntonTyurin87/Intertask/blob/main/docs/imeges/subscription_post.jpg">
</p>
