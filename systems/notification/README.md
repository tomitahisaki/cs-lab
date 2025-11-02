# Notification System (Mini Design)

A small, extensible notification system focusing on clean responsibilities and
future channel expansion.

MVP ships Slack only; interfaces allow adding LINE/Email and etc later.

## Goals

- A minimal but realistic system design (1-week scope).
- Clear separation: domain, application, infrastructure.
- Async dispatch with simple retry and backoff.
- Easy to run locally.

## Non-goals

- Full-blown distributed system or external message broker.
- Advanced delivery guarantees beyond at-least-once.
- Multi-tenant ACL/permissions.

## Architecture Overview

- **Domain**: `Message`, `Notifier` interface.
- **Application**: `Dispatcher` consumes a queue and calls a `Notifier`.
- **Infrastructure**: Slack implementation, in-memory queue, backoff.

See `/diagram/architecture.mmd` and `/diagram/sequence-send.mmd`.

## Quickstart

```bash
cd systems/notification/src
cp .env.example .env    # set SLACK_BOT_TOKEN and SLACK_CHANNEL
go mod tidy
go run ./...
```

## Configration

- SLACK_BOT_TOKEN: Bot token for slack API
- SLACK_CHANNEL: Channel ID to send message
- DISPATCHER_WORKERS: Number of dispatcher workers (default: 2)
- RETRY_MAX_ATTEMPTS: Max retry attempts for failed messages (default: 3)

## Project Architecture

- internal/domain: Core entities and interfaces.
- internal/app: Use cases (Dispather)
- internal/infra: External systems (Slack, Queue)
- pkg: Shared utilities (Backoff)

## Roadmap

- Base Slack Notifier
- In-Memory Queue and Dispatcher
- Retry with exponential backoff
- Minimal metrics (counts, stdeer logs)
- enable to add LINE or Email adapter(stretch)
