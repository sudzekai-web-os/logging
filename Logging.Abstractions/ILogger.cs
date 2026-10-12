namespace Sudzekai.WebOs.Logging.Abstractions
{
    public interface ILogger
    {
        bool IsEnabled(LogLevel logLevel);

        void Log(LogEntry entry);

        void LogTrace(string message);

        void LogDebug(string message);

        void LogInformation(string message);

        void LogWarning(string message);

        void LogError(string message);

        void LogCritical(string message);
    }
}
