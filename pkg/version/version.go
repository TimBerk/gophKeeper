package version

var (
	Version   = "dev"
	BuildTime = "unknown"
)

// String - метод для построения версии приложения
func String() string { return Version + " (" + BuildTime + ")" }
