using Sudzekai.WebOs.Logging.Abstractions;
using Sudzekai.WebOs.Logging.Internals;
using Sudzekai.WebOs.Logging.Sinks;
using Sudzekai.WebOs.Logging.Sinks.Multiple;

namespace Sudzekai.WebOs.Logging
{
    public class LoggerFactory : ILoggerFactory
    {
        private readonly QueueWorker _worker;

        public LogLevel MinLevel { get; private set; }
        
        internal LogQueue Queue { get; private set; }


        private LoggerFactory(ILogSink sink)
        {
            Queue = new();
            _worker = new(Queue, sink);
        }

        public static ILoggerFactory CreateConsole()
            => CreateLoggerFactory(new ConsoleSink());

        public static ILoggerFactory CreateFile(string fileName)
            => CreateLoggerFactory(new FileSink(fileName));

        public static ILoggerFactory CreateCustom(params ILogSink[] sinks)
            => CreateLoggerFactory(sinks);

        public ILoggerFactory Configure(Action<LoggerFactoryOptions> configure)
        {
            LoggerFactoryOptions options = new();
            configure(options);

            MinLevel = options.MinLevel;

            return this;
        }

        public ILogger NewLogger(string category)
        {
            return new Logger(category, this);
        }

        private static ILogSink CreateSink(ILogSink[] sinks)
        {
            MultipleSinkBuilder builder = new();

            foreach (var sink in sinks)
                builder.AddSink(sink);

            return builder.Build();
        }

        private static ILoggerFactory CreateLoggerFactory(params ILogSink[] sinks)
            => new LoggerFactory(CreateSink(sinks));

        public void Dispose()
        {
            _worker.Dispose();
        }

        public ValueTask DisposeAsync()
        {
            return _worker.DisposeAsync();
        }
    }
}
