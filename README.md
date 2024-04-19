
Тестируем API с использованием testify
Рассмотрим сервис из предыдущего урока. 
Напомним логику работы. 
Он возвращает список кафе по запросу. В запросе указано сколько вернуть кафе и из какого города. 
Обработчик принимает этот запрос и формирует ответ.
Если какие-то параметры указаны некорректно (нет такого города, неправильно указано количество), обработчик вернёт ошибку.
Сервер будет ожидать обращение по пути /cafe. В GET-параметрах ожидается:
count — количество кафе, которые нужно вернуть
city — город, в котором нужно найти кафе
Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе. В сервисе будет только один город moscow, в котором будет всего 4 кафе.

Нужно реализовать три теста:
Запрос сформирован корректно, сервис возвращает код ответа 200 и тело ответа не пустое.
Город, который передаётся в параметре city, не поддерживается. Сервис возвращает код ответа 400 и ошибку wrong city value в теле ответа.
Если в параметре count указано больше, чем есть всего, должны вернуться все доступные кафе.
Проверки должны осуществляться с помощью пакета testify.