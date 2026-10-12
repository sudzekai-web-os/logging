using Logging.Abstractions;

namespace Logging
{
    internal class Logger : ILogger
    {
        internal Logger(string category, LoggerFactory factory)
        {
            Category = category;
            _factory = factory;
            _minLevel = factory.MinLevel;
        }

        public string Category { get; private set; }

        private readonly LoggerFactory _factory;

        private readonly LogLevel _minLevel;

        public bool IsEnabled(LogLevel logLevel)
            => logLevel != LogLevel.None && _minLevel <= logLevel;

        public void Log(LogEntry entry)
        {
            if (!IsEnabled(entry.Level))
                return;

            _factory.Queue.TryEnqueue(entry);
        }

        public void LogCritical(string message)
            => Log(new(DateTime.Now, LogLevel.Critical, Category, message));

        public void LogDebug(string message)
            => Log(new(DateTime.Now, LogLevel.Debug, Category, message));

        public void LogError(string message)
            => Log(new(DateTime.Now, LogLevel.Error, Category, message));

        public void LogInformation(string message)
            => Log(new(DateTime.Now, LogLevel.Information, Category, message));

        public void LogTrace(string message)
            => Log(new(DateTime.Now, LogLevel.Trace, Category, message));

        public void LogWarning(string message)
            => Log(new(DateTime.Now, LogLevel.Warning, Category, message));
    }
}
