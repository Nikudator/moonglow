module moonglow

go 1.24.3

require (
	github.com/gofiber/fiber/v2 v2.52.8
	moonglow/database v0.0.0
	moonglow/routes v0.0.0
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/gofiber/fiber v1.14.6 // indirect
	github.com/gofiber/utils v0.0.10 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/schema v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.7.5 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/philhofer/fwd v1.1.3-0.20240916144458-20a13a1f6b7c // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/tinylib/msgp v1.2.5 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.62.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.38.0 // indirect
	golang.org/x/sync v0.14.0 // indirect
	golang.org/x/sys v0.33.0 // indirect
	golang.org/x/text v0.25.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	moonglow/handlers v0.0.0 // indirect
	moonglow/models v0.0.0 // indirect
)

replace moonglow/database => /database

replace moonglow/models => /models

replace moonglow/handlers => /handlers

replace moonglow/routes => /routes
