# Сервис вычисления время подачи машины

Реализован сервис вычисления время подачи машины с точностью до минуты.

Описание протокола находится в файле [./docs/eta-swagger.yml](./docs/eta-swagger.yml).

Возвращает ID машины и Количество минут до ее прибытия.

Запуск сервиса:
```bash
   $ make run
```

Пример запроса к сервису:
```bash
   $ curl -i 'http://127.0.0.1:8082/eta?lat=55.752992&lng=37.618333&limit=3'
```

## Примечание
Решил сделать клиентов для сервисов из swagger-спецификации. Использовал комманды:
* `swagger generate client -f docs/car-swagger.yml -t api-cars`
* `swagger generate client -f docs/predict-swagger.yml -t api-predict`

Нагрузку пробовал генерировать ab командой:
`ab -n500 -c 20 '127.0.0.1:8082/eta?lat=55.752992&lng=37.618333&limit=3'`

## Долго думал как реализовать требование: "Высокая скорость ответа". Из того что пришло на ум, но не стал реализовывать:
* Делать несколько запросов с клиента - по первому запросу создавать фоновую задачу и возвращать Id задачи и статус "Ищем", а далее либо web-socket, либо повторными запросами запрашивать/получать результат. Реализовывать не стал - подумал вряд ли вы имели ввиду именно этот кейс. Отдельного кэширования через Redis и Memcache делать тоже не стал. По желанию могу реализовать.

