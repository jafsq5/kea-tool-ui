package web

import "embed"

//go:embed templates/*.html
//go:embed static/**
var Assets embed.FS
