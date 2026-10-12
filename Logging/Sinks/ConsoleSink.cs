using Sudzekai.WebOs.Logging.Abstractions;
using System.Text;

namespace Sudzekai.WebOs.Logging.Sinks
{
    internal sealed class ConsoleSink : ILogSink
    {
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
            _writer = new(Console.OpenStandardOutput(),
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