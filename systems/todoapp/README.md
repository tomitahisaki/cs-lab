# TodoApp

A simple command-line ToDo management application written in Go.  
This project is designed to practice **Go application design**, **layered architecture**, and **JSON-based persistence**.

## Overview

TodoApp lets you create, list, and complete tasks directly from the terminal.  
Tasks are stored in a local **JSON file**, so your data is saved even after you close the app — no database setup required.

## Features

- Add a new task  
- List all tasks  
- Mark a task as done  
- Persistent storage using a JSON file (`tasks.json`)  
- Fully decoupled layers (Domain / Application / Infrastructure)

## Architecture

+-------------------+
| cmd/todo | ← CLI entry point
+-------------------+
|
v
+-------------------+
| Application | ← TaskService (business use cases)
+-------------------+
|
v
+-------------------+
| Domain | ← Task entity & TaskRepository interface
+-------------------+
|
v
+-------------------+
| Infrastructure | ← FileTaskRepo (JSON file storage)
+-------------------+

## Project Structure

todoapp/
  ├── cmd/
  │   └── todo/
  │       └── main.go
  ├── internal/
  │   ├── domain/
  │   │   └── task.go
  │   ├── app/
  │   │   └── task_service.go
  │   └── infra/
  │       ├── memory_task_repo.go
  │       └── file_task_repo.go
  ├── db/
  │   └── tasks.json
  ├── go.mod
  ├── README.md
  └── design.md

## Goals

Practice Go language basics and struct organization

Learn responsibility separation via layered design

Understand how to persist data using JSON files as a lightweight database
