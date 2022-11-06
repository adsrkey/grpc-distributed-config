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
![](C:/Users/ealim/Downloads/image_2022-11-06_12-48-30.png)
![](C:/Users/ealim/Downloads/image_2022-11-06_12-50-37.png)
#### Path тот же, что указан в VAULT_MOUNT_PATH
```distributed_config```

![](C:/Users/ealim/Downloads/image_2022-11-06_12-51-53.png)

### Далее будет использован клиент BloomRPC
#### GRPC запрос на создание конфига
![](C:/Users/ealim/Downloads/image_2022-11-06_13-00-18.png)
#### Появился новый secret с путём - названием сервиса
![](C:/Users/ealim/Downloads/image_2022-11-06_13-01-31.png)
#### P.S "process" нужен для того, чтобы помечать, что приложение использует конфигурацию
![](C:/Users/ealim/Downloads/image_2022-11-06_13-03-05.png)
#### GRPC запрос на получение конфига
![](C:/Users/ealim/Downloads/image_2022-11-06_13-05-37.png)
#### Приложение использует конфиг, "process" переходит в true
![](C:/Users/ealim/Downloads/image_2022-11-06_13-07-17.png)
#### GRPC запрос на обновление конфига
![](C:/Users/ealim/Downloads/image_2022-11-06_13-08-29.png)
#### Приложение пока что не использует новую версию конфига, поэтому process false
![](C:/Users/ealim/Downloads/image_2022-11-06_13-10-54.png)
#### Т.к запросов на получение конфига не приходило и приложение не использует новую версию конфигурации, эту версию можно удалить
![](C:/Users/ealim/Downloads/image_2022-11-06_13-18-13.png)
#### Конфиг 3 версии удалён
![](C:/Users/ealim/Downloads/image_2022-11-06_13-19-02.png)
