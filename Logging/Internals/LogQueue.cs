using Logging.Abstractions;
using System.Threading.Channels;

namespace Logging.Internals
{
    internal sealed class LogQueue
    {
        private readonly Channel<LogEntry> _queue = Channel.CreateBounded<LogEntry>(
            new BoundedChannelOptions(10_000)
            {
                FullMode = BoundedChannelFullMode.DropOldest,
                SingleReader = true,
                SingleWriter = false
            });

        public bool TryEnqueue(LogEntry entry)
            => _queue.Writer.TryWrite(entry);

        public void Complete() 
            => _queue.Writer.Complete();

        public ChannelReader<LogEntry> GetReader()
            => _queue.Reader;
    }
}
