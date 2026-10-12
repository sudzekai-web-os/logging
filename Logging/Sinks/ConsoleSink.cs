using Logging.Abstractions;
using System.Text;

namespace Logging.Sinks
{
    internal sealed class ConsoleSink(Stream stream) : ILogSink
    {
        private Stream _stream = stream;

        private StreamWriter? _writer;
        private bool IsOpen;

        private void Close()
        {
            IsOpen = false;
            
            if (_writer == null)
                return;

            _writer.Flush();
            _writer.Dispose();
        }

        private void Open()
        {
            _writer = new(_stream,
                new UTF8Encoding(encoderShouldEmitUTF8Identifier: false),
                bufferSize: 4096,
                leaveOpen: true)
            {
                AutoFlush = false
            };

            IsOpen = true;
        }

        public void Write(LogEntry entry)
        {
            if (_writer == null || !IsOpen)
                Open();

            if (_writer == null || !IsOpen)
                return;

            _writer.Write(entry.TimeStamp.ToString("HH:mm:ss.fff"));
            _writer.Write(" [");
            _writer.Write(entry.Level);
            _writer.Write("] ");
            _writer.Write(entry.Category);
            _writer.Write(": ");
            _writer.WriteLine(entry.Message);

            Close();
        }
    }
}