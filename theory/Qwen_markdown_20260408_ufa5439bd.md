# 🚀 Go Web Development Roadmap (без основ языка)
> Фокус: от `net/http` до production-ready gRPC-сервера. Предполагается, что базовый Go (типы, goroutines, channels, interfaces) уже знаком.

---

## 📦 Модуль 1: Ядро `net/http` и первый сервер
**🎯 Цель:** Понимать жизненный цикл HTTP-запроса в Go, уметь поднимать сервер без сторонних зависимостей.
**📚 Ключевые темы:**
- `http.ListenAndServe`, `http.Server`, настройка таймаутов (`ReadTimeout`, `WriteTimeout`, `IdleTimeout`)
- Интерфейсы `http.Handler` и `http.HandlerFunc`
- Стандартный маршрутизатор `http.ServeMux` (Go 1.22+: паттерны `/method/path`, `{param}`)
- Работа с `http.Request` и `http.ResponseWriter`
- Graceful shutdown через `context` и `signal.Notify`
- Базовая обработка ошибок и возврат статус-кодов

**🛠 Практика:**
- Написать сервер с 3 маршрутами: `GET /health`, `GET /users/{id}`, `POST /users`
- Реализовать чтение тела запроса и возврат JSON
- Добавить graceful shutdown по `SIGINT`/`SIGTERM`

**📖 Ресурсы:**
- `net/http` official docs
- [Let's Go (Alex Edwards)](https://lets-go.alexedwards.net/) (главы 1-3)
- Go 1.22 ServeMux release notes

⏱️ ~1-2 недели

---

## 📦 Модуль 2: Middleware, маршрутизация и обработка данных
**🎯 Цель:** Выстроить архитектуру обработки запросов, добавить cross-cutting concerns.
**📚 Ключевые темы:**
- Паттерн Middleware (`http.Handler` wrapper)
- Логирование, recovery (panic → 500), CORS, request ID
- Парсинг данных: `r.FormValue`, `r.ParseMultipartForm`, `json.NewDecoder(r.Body).Decode()`
- Валидация входящих данных (структурные теги, ручная проверка)
- Стандартизация ответов: обёртка `ResponseEnvelope`, обработка ошибок через `errors.Is/As`
- Работа с контекстом: `context.WithValue`, `context.WithTimeout`, проброс в БД/внешние вызовы

**🛠 Практика:**
- Написать middleware: логгер, CORS, recovery, авторизация (заглушка)
- Реализовать парсинг JSON + валидацию через `go-playground/validator`
- Централизовать обработку ошибок (кастомный `ErrorHandler`)

**📖 Ресурсы:**
- `net/http` middleware patterns
- [Go Web Examples](https://gowebexamples.com/)
- `github.com/go-chi/chi` (для понимания продвинутой маршрутизации)

⏱️ ~2 недели

---

## 📦 Модуль 3: Работа с данными и слой хранения
**🎯 Цель:** Интегрировать реляционную БД в веб-приложение по best practices.
**📚 Ключевые темы:**
- `database/sql`, пул соединений, `sql.Conn`
- Prepared statements, транзакции (`Begin`, `Commit`, `Rollback`)
- Паттерн Repository/DAO в контексте веб-сервиса
- Миграции: `pressly/goose` или `golang-migrate/migrate`
- N+1 проблема, батчинг, индексация (базово)
- Mocking БД для тестов (`sqlmock`)

**🛠 Практика:**
- Подключить PostgreSQL, создать таблицу `users`
- Реализовать CRUD через репозиторий с транзакциями
- Настроить миграции и откат
- Написать тесты репозитория с `sqlmock`

**📖 Ресурсы:**
- `database/sql` docs
- [Go SQL Best Practices](https://github.com/golang-standards/project-layout/tree/master#sql)
- `pressly/goose` README

⏱️ ~2 недели

---

## 📦 Модуль 4: Фреймворки и экосистема
**🎯 Цель:** Научиться выбирать и эффективно использовать веб-фреймворки.
**📚 Ключевые темы:**
- Когда фреймворк нужен, а когда `net/http` + `chi` достаточно
- Сравнение: `chi` (stdlib-like), `echo` (легкий, быстрый), `gin` (популярный, но тяжелее)
- Роутинг с параметрами, группы маршрутов, версии API
- Binding & Validation (автоматическое маппинг в структуры)
- Конфигурация: `spf13/viper`, env-файлы, секреты
- Шаблонизация (опционально, если SSR не нужен → пропустить)

**🛠 Практика:**
- Переписать сервер из Модуля 2 на `chi` или `echo`
- Добавить группы `/api/v1/`, `/api/v2/`
- Вынести конфиги в `viper`, добавить `.env`
- Реализовать healthcheck с метриками (заглушка)

**📖 Ресурсы:**
- Официальные docs выбранного фреймворка
- [Go Web Frameworks Benchmark](https://github.com/the-benchmarker/web-frameworks)
- `spf13/viper` examples

⏱️ ~1-2 недели

---

## 📦 Модуль 5: API Design, документация и тестирование
**🎯 Цель:** Довести API до production-стандартов.
**📚 Ключевые темы:**
- RESTful принципы, именование ресурсов, идемпотентность
- Пагинация (`limit/offset` или cursor-based), фильтрация, сортировка
- Rate limiting (`golang.org/x/time/rate` или middleware)
- Кеширование ответов (Redis или in-memory)
- Автодокументация: `swaggo/swag`, OpenAPI 3.0
- Тестирование HTTP: `httptest.NewServer`, `httptest.ResponseRecorder`, моки зависимостей

**🛠 Практика:**
- Добавить пагинацию и фильтрацию в `GET /users`
- Настроить rate limiter (100 req/min на IP)
- Сгенерировать Swagger UI через `swag init`
- Покрыть хендлеры тестами через `httptest`

**📖 Ресурсы:**
- [Microsoft REST API Guidelines](https://github.com/microsoft/api-guidelines)
- `swaggo/swag` README
- `net/http/httptest` docs

⏱️ ~2 недели

---

## 📦 Модуль 6: Безопасность и аутентификация
**🎯 Цель:** Защитить API и реализовать production-ready auth flow.
**📚 Ключевые темы:**
- JWT: структура, подпись, валидация, refresh tokens
- OAuth2 / OpenID Connect (базово, через `golang.org/x/oauth2`)
- Хеширование паролей: `golang.org/x/crypto/bcrypt`
- Security headers, CORS, CSRF (для браузерных клиентов)
- HTTPS/TLS: самоподписанные и Let's Encrypt (certbot)
- Защита от инъекций, XSS, перебора (fail2ban / lockout)

**🛠 Практика:**
- Реализовать `/auth/register` + `/auth/login` → возврат JWT
- Middleware для проверки токена, извлечение `user_id` в контекст
- Добавить refresh token flow
- Настроить HTTPS через `tls.Config` (dev-режим)

**📖 Ресурсы:**
- `golang-jwt/jwt/v5`
- [OWASP API Security Top 10](https://owasp.org/API-Security/)
- `golang.org/x/crypto/bcrypt`

⏱️ ~2 недели

---

## 📦 Модуль 7: gRPC Server с нуля
**🎯 Цель:** Разработать высокопроизводительный RPC-сервер, интегрировать с HTTP.
**📚 Ключевые темы:**
- Protocol Buffers: синтаксис `proto3`, типы, сообщения, сервисы
- Генерация кода: `protoc`, `protoc-gen-go`, `protoc-gen-go-grpc`
- gRPC Server: `grpc.NewServer`, регистрация сервисов, запуск
- Unary RPCs vs Server/Client/Bi-directional Streaming
- Интерцепторы: аутентификация, логирование, метрики, recovery
- Интеграция с HTTP/JSON: `grpc-ecosystem/grpc-gateway/v2`
- Статусы и ошибки: `google.golang.org/grpc/status`, `codes`

**🛠 Практика:**
- Описать сервис `UserService` в `.proto`
- Сгенерировать Go-код, реализовать интерфейс
- Поднять gRPC-сервер, добавить unary `CreateUser`, `GetUser`
- Написать interceptor для логирования и auth
- Подключить `grpc-gateway`, проверить HTTP/JSON маппинг
- Написать клиент для тестирования (можно через `grpcurl`)

**📖 Ресурсы:**
- [gRPC Go Quickstart](https://grpc.io/docs/languages/go/quickstart/)
- `protoc` official docs
- `grpc-gateway` documentation
- `github.com/grpc-ecosystem/go-grpc-middleware`

⏱️ ~3 недели

---

## 📦 Модуль 8: Production Readiness & Observability
**🎯 Цель:** Подготовить сервис к деплою в реальные условия.
**📚 Ключевые темы:**
- Структурированное логирование: `log/slog` (Go 1.21+) или `uber-go/zap`
- Метрики: `prometheus/client_golang`, `/metrics` endpoint
- Трейсинг: OpenTelemetry SDK, экспортеры (Jaeger/Tempo)
- Healthchecks: `/health` (liveness), `/ready` (readiness)
- Docker: multi-stage build, минимальные образы (`distroless`/`alpine`)
- Graceful shutdown в связке `net/http` + `gRPC`
- CI/CD basics: линтеры (`golangci-lint`), тесты, сборка, деплой

**🛠 Практика:**
- Добавить `slog` + request ID в цепочку
- Экспортировать метрики HTTP и gRPC в Prometheus
- Настроить OpenTelemetry tracer
- Написать `Dockerfile` (multi-stage), `docker-compose` с БД
- Реализовать единый graceful shutdown для HTTP+gRPC
- Запустить `golangci-lint` и исправить все ворнинги

**📖 Ресурсы:**
- `log/slog` official guide
- [OpenTelemetry Go](https://opentelemetry.io/docs/languages/go/)
- `golangci-lint` config examples
- [Production Go (Uber)](https://github.com/uber-go/guide/blob/master/style.md)

⏱️ ~2-3 недели

---

## 🧩 Финальный проект (Capstone)
Собери всё воедино:
- gRPC-сервис + HTTP gateway
- Авторизация через JWT
- PostgreSQL с миграциями и репозиторием
- Prometheus + OpenTelemetry
- Docker + `docker-compose`
- Покрытие тестами >70%, `golangci-lint` в CI
- Документация: OpenAPI + gRPC `.proto`

---

## 📁 Как использовать
1. Сохрани этот файл как `go-web-dev-roadmap.md`
2. Отмечай выполненные задачи `[x]`
3. Не пропускай практику: теория без кода в Go = иллюзия знаний
4. Фокусируйся на одном модуле, пока не соберёшь работающий прототип

> 💡 Совет: Go 1.21+ изменил стандартную библиотеку (`log/slog`, улучшенный `ServeMux`, `context` в `http.Client`). Убедись, что используешь актуальные версии пакетов.

Удачи в разработке! 🚀