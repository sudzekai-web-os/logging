namespace Sudzekai.WebOs.Logging.Abstractions
{
    public interface ILoggerFactory : IAsyncDisposable, IDisposable
    {
        ILogger NewLogger(string category);

        ILoggerFactory Configure(Action<LoggerFactoryOptions> configure);
    }
}
