using Logging.Internals;

namespace Logging.Context
{
    internal static class LoggingContext
    {
        public static readonly LogQueue Queue = new();
    }
}
