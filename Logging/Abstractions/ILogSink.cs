namespace Logging.Abstractions
{
    public interface ILogSink
    {
        void Write(LogEntry entry);
    }
}
