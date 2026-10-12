using Sudzekai.WebOs.Logging.Abstractions;

namespace Sudzekai.WebOs.Logging.Sinks.Multiple
{
    internal class MultipleSinkBuilder
    {
        private readonly IList<ILogSink> _sinks = [];

        public void AddSink(ILogSink sink)
        {
            _sinks.Add(sink);
        }

        public ILogSink Build()
           => new MultipleSink(_sinks);
    }
}
