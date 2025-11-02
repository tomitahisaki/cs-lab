# Design Doc — Notification System (Mini)

## 1. Problem Statement

Build a small, extensible notification system that can send outbound messages reliable
through pluggable intefaces.

The MVP will focus on Slack as the initial channel, with a design that allows
easy addition of other channels (e.g., LINE, Email) in the future.

## 2. Requirements

- **Functional**
  - Enqueue a `Message{ channel, content, metadata }`.
  - Dispatch asynchronously to Slack.
  - Retry failed deliveries with exponential backoff (up to N attempts)
- **Non-Functional**
  - Clear layered architecture; easy to extend.
  - At-least-once delivery semantics.
  - Local dev friendly (no external broker).

## 3. Architecture

```mermaid
flowchart LR
  API[Producer (main.go)] --> Q[In-Memory Queue]
  Q --> D[Dispatcher (workers)]
  D -->|Notifier| S[Slack Adapter]
  S --> SlackAPI[(Slack API)]
```

```text
[cmd]                — Entry point (main)
   |
   v
[app] Dispatcher     — Use-case layer (workers / retry logic)
  |  \
  |   \__ uses --> [pkg/backoff]
  |                 [pkg/log] (optional)
  v
[domain]             — Core types & interfaces (Message, Notifier)
  ^
  |
[infra/queue]        — InMemoryQueue (depends on domain types)
[infra/slack]        — Slack Notifier (implements domain.Notifier) 
                       --> external: slack-go/slack
```

Rules

- cmd depends on app, infra, pkg, domain
- app depends on domain, pkg
- infra depends on domain
- domain and pkg have no dependencies
- No upward dependencies allowed and no cycles

## Components

- Domain Layer
  - Message: core entity representing a notification payload
    - Attributes: id, channel, text, metadata
  - Notifier: Abstract interface for sending messages
    - Method: Send(msg Message) error
- Application
  - Dispatcher
    - Consumes messages from the queue
    - Invokes the appropriate Notifier implenpementation
    - Handles retries, backoff scheduling, and worker concurrency
- Infrastructure
  - InMemoryQueue
    - Simple FIFO queue with delayed re-enqueue support
    - Used for async dispatch without external dependencies
  - SlackNotifier
    - Concrete Implemantation of the Notifier interface
    - Uses slack-go/slack to send messages to Slack channels
  - backoff
    - Helper utility providing exponential backoff with jitter
