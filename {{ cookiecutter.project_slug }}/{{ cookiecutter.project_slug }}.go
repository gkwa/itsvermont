package {{ cookiecutter.project_slug }}

import (
	"log/slog"
)

func Main() int {
	slog.Debug("{{ cookiecutter.project_slug }}", "test", true)
	run()

	return 0
}
