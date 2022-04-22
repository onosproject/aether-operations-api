package swagger

import (
	"embed"
)

//go:embed dist/*
var OpenAPI embed.FS
