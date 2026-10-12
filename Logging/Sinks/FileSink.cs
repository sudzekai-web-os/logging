using Logging.Abstractions;
using System.Text;

namespace Logging.Sinks
{
    internal sealed class FileSink(string fileName) : ILogSink
    {
        private readonly string _fileName = fileName;
        
        private FileStream? _writer;
        private bool IsOpen;

        private void Close()
        {
            IsOpen = false;

            if (_writer is null)
                return;
        }

        private void Open()
        {
            if (string.IsNullOrWhiteSpace(_fileName))
                return;

            _writer = new(_fileName!,
                FileMode.Append,
                FileAccess.Write,
                FileShare.Read,
                bufferSize: 4096,
                useAsync: false);

            IsOpen = true;
        }

        public void Write(LogEntry entry)
        {
            if (_writer is null || !IsOpen)
                Open();

            if (_writer is null || !IsOpen)
                return;

            var sb = new StringBuilder(64);

            sb.Append(entry.TimeStamp.ToString("HH:mm:ss.fff"))
              .Append(" [").Append(entry.Level).Append("] ")
              .Append(entry.Category).Append(" : ")
              .Append(entry.Message)
              .Append(Environment.NewLine);

            byte[] buffer = Encoding.UTF8.GetBytes(sb.ToString());
            _writer.Write(buffer, 0, buffer.Length);

            Close();
        }
    }
}