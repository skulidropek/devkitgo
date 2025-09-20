# DevKitGo

Набор утилит, упрощающий подключение статического анализа и инструментов трассировки в Go‑проектах. Репозиторий включает три ключевые команды:

- `cmd/analyze` — агрегированный статический анализатор на базе `multichecker`, [Staticcheck](https://github.com/dominikh/go-tools/tree/master/cmd/staticcheck) и [GoSuggestMembersAnalyzer (SMB)](https://github.com/skulidropek/GoSuggestMembersAnalyzer).
- `cmd/lint` — обёртка вокруг [GoLint](https://github.com/skulidropek/GoLint), умеющая подтягивать зависимости без глобальной установки.
- `cmd/instrument` — запускатор [gotrace-instrument](https://github.com/skulidropek/gotrace/tree/main/cmd/gotrace-instrument) для автоматического добавления трассировок (проект [gotrace](https://github.com/skulidropek/gotrace)).

## Быстрый старт

1. Зафиксируйте инструменты один раз:
   ```bash
   go generate -tags tools ./tools
   ```

2. Запустите нужную команду напрямую через `go run`:
   ```bash
   # Статический анализ всего проекта
   go run github.com/skulidropek/devkitgo/cmd/analyze@latest ./...

   # Линтер GoLint (использует go-lint из devkit)
   go run github.com/skulidropek/devkitgo/cmd/lint@latest ./...

   # Инструментирование кода трассировками
   go run github.com/skulidropek/devkitgo/cmd/instrument@latest --src ./internal
   ```

   Каждая обёртка сначала ищет локально установленный бинарь (`go-lint`, `gotrace-instrument`). Если он отсутствует, команда автоматически выполнит `go run` закреплённой версии.

## Интеграция в Makefile

Пример целей для линтинга и инструментирования:

```make
GO ?= go
TARGET ?= ./...
GOCACHE_DIR ?= $(PWD)/.gocache
GOLANGCI_LINT_CACHE_DIR ?= $(PWD)/.golangci-cache
GO_LINT_MODULE := github.com/skulidropek/devkitgo/cmd/lint@latest
GO_LINT_BIN ?= $(shell command -v go-lint 2>/dev/null)
GO_LINT_CMD := $(if $(GO_LINT_BIN),$(GO_LINT_BIN),$(GO) run $(GO_LINT_MODULE))

.PHONY: lint

lint:
	@mkdir -p $(GOCACHE_DIR) $(GOLANGCI_LINT_CACHE_DIR)
	GOCACHE=$(GOCACHE_DIR) GOLANGCI_LINT_CACHE=$(GOLANGCI_LINT_CACHE_DIR) $(GO_LINT_CMD) $(TARGET)
```

## Переменные окружения

- `DEVKIT_GO_LINT_MODULE` — переопределяет версию модуля, используемую обёрткой `lint`.
- `DEVKIT_GOTRACE_MODULE` — аналогично для `instrument`.

Это полезно, если требуется протестировать другую ревизию инструмента, не меняя исходники devkit.

## Обновление инструментов

Команда `go generate -tags tools ./tools` устанавливает зафиксированные версии [go-lint](https://github.com/skulidropek/GoLint), [smbgo](https://github.com/skulidropek/GoSuggestMembersAnalyzer/tree/main/cmd/smbgo) и [gotrace-instrument](https://github.com/skulidropek/gotrace/tree/main/cmd/gotrace-instrument) в локальный `$GOBIN`. Запускайте её после обновления `tools/tools.go` или при первой настройке окружения.

## Требования

- Go 1.24.4+
- Доступ к модульному прокси либо зеркалам, обеспечивающий загрузку зависимостей Staticcheck и прочих пакетов.
