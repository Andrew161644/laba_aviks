Этот проект содержит 5 разделов:

docs - раздел с документами 

Build - папка со всеми скриптами 
    - microservices - скрипт, который потребуется для запуска нескольких микросервисов (grpc)
    - data_base_only_start - запускает базу данных в контейнере 
    - rest_server_start - скрипт для запуска простого сервера 
    - stop_clear_containers - необходим для остановки и отчистки всех контейнеров
    - BuildEnv - кастомный образ для сборки приложения (сюда можно не заходить)
    - local_scripts - папка с docker-compose для базы данных (сюда можно не заходить)

migrations - папка, где хранятся все скрипты, выполняемые при старте базы данных

grpc_service - клиент и сервер, которые используют grpc, потребуется в будущем
    - clent пример удаленного вызова процедуры (grpc)
    - database код, взаимодействия с базой данных(модели/провайдеры)
    - database_test ?? надо добавить
    - main запуск grpc сервера
    - grpc_server логика сервера(например вытащить пользователя из базы данных и отдать его клиенту)
    - user_mapper - приведения типа UseRole(база данных) к типу UserResponse(ответ клиенту)
    - etc. файлы сгенерированный код обуспечивающий клиент серверное взаимодействие по протоколу HTTP/2

rest - обыкновенный сервер для выполнения первой лабораторной работы, может возвращать html + css (http://localhost:8080/page)


Общие пояснения: 

Для запуска приложений не используя докер на своем собственном компе см. комментарии
На каждый написанный провайдер нужен тест
В static директориях лежит html и css код, все директории с html/css/js будут static
В rest/api/handlers создается то, что увидит пользователь в браузере
    - view папка с моделями старицы(соответствует html, отображение html в виде класса, содержит динамические поля html)
    - application - простой механизм внедрения зависимостей, работает через указатели, можно один раз инициализировать
        какую либо сущность (базу данных например), затем в нескольких местах использовать injection
    - mainpage содержит функции, которые будут выполнены, при переходе пользователя на главную страницу

go.mod корень проекта
 - go sum файл с зависимостями