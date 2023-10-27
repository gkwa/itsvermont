package {{ cookiecutter.project_slug }}

import (
	"log/slog"
)

func Main() int {
	slog.Debug("{{ cookiecutter.project_slug }}", "test", true)

	return 0
}
