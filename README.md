# cs-lab (Monorepo)

Learning-first monorepo for CS × Go × Algorithms × System Design.

## Structure

- `cs/` — CS experiments (concurrency, network, database)
- `lab/` — distributed/messaging/cache experiments
- `systems/` — small systems (url-shortener, job-queue, chat)
- `algo/` — AtCoder/LeetCode + data structures
- `shared/` — shared utilities
- `docs/` — ADRs, weekly notes, templates

## Quick start

```bash
# clone the repo, then:
go version
go work sync
make test
```

## Weekly loop

- Tue: English × CS reading (1.5h)
- Thu: Algorithms (1.5h)
- Sat: Implementation (3h)
- Sun: English summary (1–2h)

See `.github/ISSUE_TEMPLATE/weekly.yml` and scheduled KPI/Weekly workflows.
