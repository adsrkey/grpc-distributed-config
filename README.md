# Distributed Config Grpc Service

## Инструкция по запуску

```
cd .\container\

make up_build
```
- поднятие и запуск docker контейнера

### http://localhost:8200
Token:
```myroot```
### Добавить новый Secrets Engine
![image_2022-11-06_12-48-30](https://user-images.githubusercontent.com/112159682/200163558-e7a75eba-776f-4078-bb80-e84b1b769390.png)
![image_2022-11-06_12-50-37](https://user-images.githubusercontent.com/112159682/200163568-5e895911-9ddd-4b8e-8dd5-24d1ed363106.png)

#### Path тот же, что указан в VAULT_MOUNT_PATH
```distributed_config```

![image_2022-11-06_12-51-53](https://user-images.githubusercontent.com/112159682/200163577-08e786ee-3df8-4a60-9de2-eb5c9ec2c331.png)

### Далее будет использован клиент BloomRPC
#### GRPC запрос на создание конфига
![image_2022-11-06_13-00-18](https://user-images.githubusercontent.com/112159682/200163583-c12904dd-cb80-4d35-97ea-a049e26458c2.png)

#### Появился новый secret с путём - названием сервиса
![image_2022-11-06_13-01-31](https://user-images.githubusercontent.com/112159682/200163588-8557070e-769c-4b4e-a455-0ea63ebbd985.png)

#### P.S "process" нужен для того, чтобы помечать, что приложение использует конфигурацию
![image_2022-11-06_13-03-05](https://user-images.githubusercontent.com/112159682/200163589-74cee513-7ef1-4e34-bc44-25d8e217f5c1.png)
#### GRPC запрос на получение конфига
![image_2022-11-06_12-50-37](https://user-images.githubusercontent.com/112159682/200163598-3feb78af-143c-4d6e-b322-0c340f2d4a71.png)
#### Приложение использует конфиг, "process" переходит в true
![image_2022-11-06_13-07-17](https://user-images.githubusercontent.com/112159682/200163604-9c6a559a-ef0f-4db8-8ced-2e08471dfa49.png)
#### GRPC запрос на обновление конфига
![image_2022-11-06_13-08-29](https://user-images.githubusercontent.com/112159682/200163615-0f6dde4d-e7f2-4e42-accb-54f602034dc4.png)
#### Приложение пока что не использует новую версию конфига, поэтому process false
![image_2022-11-06_13-10-54](https://user-images.githubusercontent.com/112159682/200163621-14eb47d2-f7f2-448d-87d4-b56f911911da.png)
#### Т.к запросов на получение конфига не приходило и приложение не использует новую версию конфигурации, эту версию можно удалить
![image_2022-11-06_13-18-13](https://user-images.githubusercontent.com/112159682/200163624-1238152f-114e-460b-a07e-f75c2cbf4797.png)
#### Конфиг 3 версии удалён
![image_2022-11-06_13-19-02](https://user-images.githubusercontent.com/112159682/200163630-4520ecdc-97f1-423b-88ef-6aa22c5cb11e.png)

