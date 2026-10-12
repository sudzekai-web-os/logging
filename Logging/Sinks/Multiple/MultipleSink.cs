using Logging.Abstractions;

namespace Logging.Sinks.Multiple
{
    internal class MultipleSink : ILogSink
    {
        private readonly IList<ILogSink> sinks = [];

        internal MultipleSink(IList<ILogSink> sinks)
        {
            this.sinks = sinks;
        }

        public void Write(LogEntry entry)
        {
            foreach (var sink in sinks)
                sink.Write(entry);
        }
    }
}
