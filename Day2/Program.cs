Console.WriteLine("1. Example: {0}", FindNumberOfSafeReports(File.ReadAllText("example.txt"), false));
Console.WriteLine("1. Answer: {0}", FindNumberOfSafeReports(File.ReadAllText("input.txt"), false));

Console.WriteLine("2. Example: {0}", FindNumberOfSafeReports(File.ReadAllText("example.txt"), true));
Console.WriteLine("2. Answer: {0}", FindNumberOfSafeReports(File.ReadAllText("input.txt"), true));

static int FindNumberOfSafeReports(string input, bool allowSingleBadLevel)
{
    var reports = input.Split('\n').Where(l => l.Length != 0);

    return reports.Count(r => ReportLineIsSafe(r, allowSingleBadLevel));
}

static bool ReportLineIsSafe(string report, bool allowSingleBadLevel)
{
    var levels = report.Split(' ').Select(int.Parse).ToArray();

    if (!allowSingleBadLevel)
    {
        return ReportIsSafe(levels);
    }

    for (var i = 0; i < levels.Length; i++)
    {
        var reportWithRemovedLevel = levels.Take(i).Concat(levels.Skip(i + 1));
        if (ReportIsSafe(reportWithRemovedLevel))
        {
            return true;
        }
    }

    return false;
}

static bool ReportIsSafe(IEnumerable<int> report)
{
    using var enumerator = report.GetEnumerator();

    enumerator.MoveNext();
    var previousLevel = enumerator.Current;
    var previousDiff = 0;

    while (enumerator.MoveNext())
    {
        var level = enumerator.Current;
        var diff = level - previousLevel;

        if (!LevelChangeIsSafe(diff, previousDiff))
        {
            return false;
        }

        previousDiff = diff;
        previousLevel = level;
    }

    return true;
}

static bool LevelChangeIsSafe(int diff, int previousDiff)
{
    if ((diff < 0 && previousDiff > 0) || (diff > 0 && previousDiff < 0))
    {
        return false;
    }

    if (Math.Abs(diff) is < 1 or > 3)
    {
        return false;
    }

    return true;
}