# SudzekaiWebOS.Logging

Асинхронное потокобезопасное логирование для .NET: уровни, категории, очередь и сменные sink-и (консоль, файл, кастомные).

Состоит из двух пакетов:

- **SudzekaiWebOS.Logging.Abstractions** — интерфейсы и типы: `ILogger`, `ILoggerFactory`, `ILogSink`, `LogEntry`, `LogLevel`, `LoggerFactoryOptions`.
- **SudzekaiWebOS.Logging** — реализация поверх абстракций.

## Установка

```
dotnet add package SudzekaiWebOS.Logging
```

## Использование

```csharp
using Logging;

using var factory = LoggerFactory.CreateConsole(Console.OpenStandardOutput());

factory.Configure(o => o.MinLevel = Logging.Abstractions.LogLevel.Trace);

var logger = factory.NewLogger("Program");

logger.LogTrace("trace");
logger.LogDebug("debug");
logger.LogInformation("info");
logger.LogWarning("warn");
logger.LogError("error");
logger.LogCritical("critical");
```

Записи складываются в потокобезопасную очередь и записываются отдельным фоновым потоком, поэтому вызовы `Log*` не блокируют вызывающий поток.

## Конфигурация

- `factory.Configure(...)` — задаёт минимальный уровень (`o.MinLevel`). Уровни: `Trace`, `Debug`, `Information`, `Warning`, `Error`, `Critical`.
- `factory.NewLogger(category)` — создаёт логгер с категорией.

## Sink-и

- `LoggerFactory.CreateConsole(Stream stream)` — запись в консоль.
- `LoggerFactory.CreateFile(string fileName)` — запись в файл (append).
- `LoggerFactory.CreateCustom(params ILogSink[] sinks)` — произвольный набор sink-ов.

Собственный sink реализуется через `ILogSink.Write(LogEntry entry)`. Синк сам управляет открытием/закрытием ресурса на время записи.

## Лицензия

Apache-2.0 (см. [LICENSE](LICENSE)).