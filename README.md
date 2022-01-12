## Задача 1

###  При запуске программы стартует веб-сервер, готовый к принятнию и обработке http-запросов.
#### Для запуска программы необходимо запустить /api/main.go

#### Описание эндпойнтов, реализованных в программе и примеры запросов к ним:


- ###### POST api/savestate - обновить информацию клиента
        curl -v -X POST -H "Content-Type: application/json" --data '{"application": "first-client", "param1": 2, "param2": "0.05"}' 'localhost:5000/api/savestate'

- ###### GET api/getstate - запросить информацию клиента
        curl -v -X GET -H "Content-Type: application/json" --data '{"application": "first-client"}' 'localhost:5000/api/getstate'
