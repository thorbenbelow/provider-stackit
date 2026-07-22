# Agent Operating Manual & Repository Rules

This document establishes the strict operating procedures, command-first validation protocols, explicit boundaries, and escalation paths for any AI coding assistant working in this repository.

---

## Explicit Boundaries (Rigid Rules & "Never" List)

To prevent resource conflict, configuration drift, and security leaks, you must adhere to the following strict boundaries:

- **Never** use absolute paths (e.g., `file:///Users/foo/...` or `file:///home/foo/...`) inside any repository code, configurations, or documentation. Always use relative paths (e.g., `./readme.md`).
- **Never** auto-add your agent name as co-author.
- **Always** use the Conventional Commits specification for any changes committed to Git. The commit scope must always display the affected area.

## Maintaining this file

Keep this file for knowledge useful to almost every future agent session in this project. Do not repeat what the codebase already shows; point to the authoritative file or command instead. Prefer rewriting or pruning existing entries over appending new ones. When updating this file, preserve this bar for all agents and keep entries concise.
