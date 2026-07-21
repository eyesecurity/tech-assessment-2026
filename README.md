# Eye Security - Engineering Assessment

This repository is the starting point for the Eye Security take-home assessment.
It boots a Go backend and a React frontend that already talk to each other. Your
assignment brief (sent separately) describes the SecOps detection-triage task you
will build on top of it.

---

## Architecture

```
backend/   Go + Gin API server  -> http://localhost:8080
frontend/  React + Vite app     -> http://localhost:5173
```

The frontend calls `GET /api/health` on startup to prove connectivity (see the
badge in the top-right of the page).

---

## Getting started

### Backend

Requires Go 1.22+.

```bash
cd backend
go mod tidy
go run .
```

### Frontend

Requires Node 20+ and pnpm.

```bash
cd frontend
pnpm install
pnpm dev
```

Open http://localhost:5173. The connectivity badge should read **backend: ok**.

---

## What's provided

The boilerplate is deliberately minimal — it sets up the plumbing and leaves the
design decisions to you.

- **Backend**
  - Gin server on `:8080` with CORS configured for the frontend.
  - `GET /api/health` — the connectivity example the frontend calls.
  - `data.ReadCSV(path)` (`backend/data/reader.go`) — reads a CSV into raw,
    untyped records (`[]map[string]string`). How you model the data is up to you.
  - Your detection CSV is provided separately by Eye Security and belongs in
    `backend/data/` (the folder where `reader.go` lives).
- **Frontend**
  - React 19 + Vite + Tailwind, styled and booting cleanly.
  - A single detections screen skeleton in `src/App.tsx`.
  - `useHealth()` (`src/api/useHealth.ts`) — an example React Query hook. Reuse
    the pattern for your own data fetching.

Nothing here dictates your domain model, your API contract, or how you integrate
the Enrichment Service — those choices are yours.

---

## Are we not worried that this is in the public domain?

Short answer: no.

The heart of this assessment is the interview, not the code in isolation. Two
senior engineers will sit down and walk through your work with you — how you
modelled the problem, the trade-offs you made, how you handled the flaky
Enrichment Service, and what you'd do differently. That conversation is where the
signal is.

Publicly available solutions from previous candidates may well exist, and their
quality varies. You're welcome to look around — but a solution you didn't reason
through yourself won't help you when we ask you to explain and defend your
decisions live. A thoughtful, partial solution you understand deeply is worth far
more to us than a polished one you can't talk through.

So build in a way that lets you own every line. That's what we're here to see.

---

## Your notes

Use this section (or expand the README) to record your key decisions,
assumptions, and trade-offs — the interview will walk through them.

```
- _Assumptions:_
- _Decisions & trade-offs:_
- _If I had more time:_
```
