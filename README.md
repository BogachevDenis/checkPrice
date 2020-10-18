# checkPrice
# Запуск проекта:
Для запуска проекта выполните следующие команды:
<br>
$ git clone https://github.com/BogachevDenis/checkPrice.git
<br>
$ cd checkPrice
<br>
$ docker-compose build
<br>
$ docker-compose up
<br>
# Работа с приложением:
Приложение запустится на http://localhost:8080/
<br>
Для удобства приложение имеет UI, данные можно вводить через браузер
<br><br>
Также данные можно отправить через curl:
<br>
$ curl -X POST http://localhost:8080/create -d '{"email": "test@test.ru","url":"https://www.avito.ru/ad"}'
