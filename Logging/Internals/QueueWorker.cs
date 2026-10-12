using Sudzekai.WebOs.Logging.Abstractions;

namespace Sudzekai.WebOs.Logging.Internals
{
    internal sealed class QueueWorker : IAsyncDisposable, IDisposable
    {
        private readonly LogQueue _queue;
        private readonly ILogSink _sink;
        private readonly Task _worker;
        private readonly CancellationTokenSource _cts = new();
        private bool _disposed;

        public QueueWorker(LogQueue queue, ILogSink sink)
        {
            _queue = queue;
            _sink = sink;
            _worker = Task.Factory.StartNew(
                () => RunAsync(_cts.Token),
                CancellationToken.None,
                TaskCreationOptions.LongRunning,
                TaskScheduler.Default).Unwrap();
        }

        private async Task RunAsync(CancellationToken ct)
        {
            var reader = _queue.GetReader();

            try
            {
                while (await reader.WaitToReadAsync(ct))
                {
                    while (reader.TryRead(out var entry))
                        _sink.Write(entry);
                }
            }
            catch (OperationCanceledException) when (ct.IsCancellationRequested)
            {
                // Нормальная отмена
            }
            catch (Exception ex)
            {
                Console.Error.WriteLine($"Log worker crashed: {ex}");
            }
        }

        public async ValueTask DisposeAsync()
        {
            if (_disposed) return;
            _disposed = true;

            _queue.Complete();     // завершит WaitToReadAsync → false
            await _worker;         // дождёмся завершения
            _cts.Cancel();         // на случай зависания
            _cts.Dispose();
        }

        public void Dispose()
        {
            if (_disposed) return;
            _disposed = true;

            _queue.Complete();     // завершит WaitToReadAsync → false
            _worker.Wait();        // дождёмся завершения
            _cts.Cancel();         // на случай зависания
            _cts.Dispose();
        }
    }
}