namespace Logging.Abstractions
{
    public readonly record struct LogEntry(
        DateTime TimeStamp, 
        LogLevel Level, 
        string Category, 
        string Message
    );
}
