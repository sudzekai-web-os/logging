using Sudzekai.WebOs.Logging.Abstractions;
using System.Text;

namespace Sudzekai.WebOs.Logging.Sinks
{
    public sealed class ConsoleSink : ILogSink
    {
        private const string Reset = "\u001b[0m";
        private const string DarkGray = "\u001b[90m";

        private StreamWriter? _writer;
        private bool IsOpen;

        public ConsoleSink()
        {
        }

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

            var (code, color) = MapLevel(entry.Level);

            _writer.Write(DarkGray);
            _writer.Write('[');
            _writer.Write(entry.TimeStamp.ToString("HH:mm:ss"));
            _writer.Write(']');
            _writer.Write(Reset);

            _writer.Write(' ');
            _writer.Write(color);
            _writer.Write(code.PadRight(4));
            _writer.Write(Reset);

            _writer.Write(' ');
            _writer.Write(DarkGray);
            _writer.Write(entry.Category);
            _writer.Write(Reset);

            _writer.Write("\n           ");
            _writer.Write(entry.Message);
            _writer.Write('\n');

            _writer.Flush();
            Close();
        }

        private static (string code, string color) MapLevel(LogLevel level) => level switch
        {
            LogLevel.Trace => ("TRCE", "\u001b[30;47m"),
            LogLevel.Debug => ("DBUG", "\u001b[39;49m"),
            LogLevel.Information => ("INFO", "\u001b[34m"),
            LogLevel.Warning => ("WARN", "\u001b[33m"),
            LogLevel.Error => ("ERR", "\u001b[31m"),
            LogLevel.Critical => ("CRIT", "\u001b[37;41m"),
            _ => ("----", Reset),
        };
    }
}