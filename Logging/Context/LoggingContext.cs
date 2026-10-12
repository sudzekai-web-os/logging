using Sudzekai.WebOs.Logging.Internals;

namespace Sudzekai.WebOs.Logging.Context
{
    internal static class LoggingContext
    {
        public static readonly LogQueue Queue = new();
    }
}
