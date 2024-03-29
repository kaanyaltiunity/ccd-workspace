package utils

const (
	Reset      = "\x1b[0m"
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)

type TextColorTags struct {
	Modifiers  TextColorModifiers
	Foreground TextForegroundColors
	Background TextBackgroundColors
}

type TextColorModifiers struct {
	Reset      string
	Bright     string
	Dim        string
	Underscore string
	Blink      string
	Reverse    string
	Hidden     string
}

type TextForegroundColors struct {
	Black   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
}

type TextBackgroundColors struct {
	Black   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string
}

var ColorTags = TextColorTags{
	Modifiers: TextColorModifiers{
		Reset:      Reset,
		Bright:     Bright,
		Dim:        Dim,
		Underscore: Underscore,
		Blink:      Blink,
		Reverse:    Reverse,
		Hidden:     Hidden,
	},
	Foreground: TextForegroundColors{
		Black:   FgBlack,
		Red:     FgRed,
		Green:   FgGreen,
		Yellow:  FgYellow,
		Blue:    FgBlue,
		Magenta: FgMagenta,
		Cyan:    FgCyan,
		White:   FgWhite,
	},
	Background: TextBackgroundColors{
		Black:   BgBlack,
		Red:     BgRed,
		Green:   BgGreen,
		Yellow:  BgYellow,
		Blue:    BgBlue,
		Magenta: BgMagenta,
		Cyan:    BgCyan,
		White:   BgWhite,
	},
}
