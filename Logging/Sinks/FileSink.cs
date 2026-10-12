using Sudzekai.WebOs.Logging.Abstractions;
using System.Text;

namespace Sudzekai.WebOs.Logging.Sinks
{
    public sealed class FileSink(string fileName) : ILogSink
    {
        private readonly string _fileName = fileName;

        private FileStream? _writer;
        private bool IsOpen;

        private void Close()
        {
            IsOpen = false;

            if (_writer is null)
                return;

            _writer.Flush();
            _writer.Dispose();
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

            var code = MapLevel(entry.Level);

            var sb = new StringBuilder(64);

            sb.Append('[')
              .Append(entry.TimeStamp.ToString("HH:mm:ss"))
              .Append("] ")
              .Append(code.PadRight(4))
              .Append(' ')
              .Append(entry.Category)
              .Append("\n           ")
              .Append(entry.Message)
              .Append('\n');

            byte[] buffer = Encoding.UTF8.GetBytes(sb.ToString());
            _writer.Write(buffer, 0, buffer.Length);

            Close();
        }

        private static string MapLevel(LogLevel level) => level switch
        {
            LogLevel.Trace => "TRCE",
            LogLevel.Debug => "DBUG",
            LogLevel.Information => "INFO",
            LogLevel.Warning => "WARN",
            LogLevel.Error => "ERR",
            LogLevel.Critical => "CRIT",
            _ => "----",
        };
    }
}