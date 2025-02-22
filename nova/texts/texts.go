package texts

const (
	ReadingError             = "Couldn't read input"
	StartFlag                = "start"
	StartDescription         = "Start Nova"
	CurrentDirectory         = "."
	RepositoryToClone        = "URL/Path of the repository to clone (Leave it blank to clone current repository):"
	Screenshot               = "screenshot"
	Files                    = "ls"
	Infos                    = "watcher"
	Logger                   = "log"
	Status                   = "status"
	StatusSuccessfully       = "%sStatus%s\n%s\n"
	Branches                 = "branches"
	CreateCommand            = `^nova \w+\.[0-9a-z]+$`
	SwitchBranchCommand      = `^go to \w+$`
	CommittedSuccessfully    = "Worktree saved\nhash -> %s ...msg -> %s\n"
	UnsavedChangesWarning    = "WARNING!!! File has unsaved changes. Press Ctrl-X %d more times to quit."
	FileOpenedSuccessfully   = "File opened"
	BeforeInput              = "~ "
	StopNova                 = "stop"
	ClearScreen              = "clear"
	StoreCreatedSuccessfully = "%sStore created: %s%s\n"
	QuitEditor               = "quit editor"
	UserCanceledInputPrompt  = "user canceled the input prompt"
	NameCannotBeEmpty        = "Name cannot be empty"
	RESET                    = "\033[0;0m"
	WHITE                    = "\033[1;37m"
	RED                      = "\033[1;41m"
	GREEN                    = "\033[1;32m"
	CYAN                     = "\033[1;36m"
	BACKGROUND_RED           = "\033[0;41m"
	BACKGROUND_ORANGE        = "\033[0;43m"
	BACKGROUND_CYAN          = "\033[0;46m"
	HtmlEmptyStringMsg       = `Cannot be empty`
)
