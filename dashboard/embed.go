package dashboard

import (
    "embed"
)

var (
    //go:embed dashboard/*
    Dashboard embed.FS

    //go:embed equity/*
    Equity embed.FS
)
