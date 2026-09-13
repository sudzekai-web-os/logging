# logging

Потокобезопасная реализация логирования с уровнями, категориями и настраиваемым `io.Writer`.

## Установка

```bash
go get github.com/sudzekai-web-os/logging
```

## Использование

```go
factory := logging.NewLoggerFactory(os.Stdout)
factory.SetMinLevel(types.Debug)

logger := factory.NewLogger("worker")
logger.LogInformation("запущена обработка: %s", jobID)
logger.LogError("ошибка обработки: %v", err)
```

Доступны методы `LogDebug`, `LogInformation`, `LogWarning`, `LogError` и `LogCritical`, а также универсальный `Log`.

По умолчанию минимальный уровень — `types.Information`. Уровень можно задать типом `types.LogLevel` или строкой через `SetMinLevelStr`: поддерживаются `debug`, `info`, `warn`, `error` и `critical`.

Фабрика синхронизирует изменение writer и минимального уровня, поэтому ее можно использовать из нескольких горутин.