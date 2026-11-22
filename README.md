# Engineering Lab

Das Engineering Lab ist unser gemeinsamer Arbeitsbereich für kleine, praxisorientierte Softwareprojekte zum Erlernen von Go und Python. Diese Datei hält den bisherigen Projektplan, unsere Arbeitsweise und die nächsten Schritte fest.

**Leitgedanke: Start small. Build incrementally. Keep it useful.**

## Ziel

Wir bauen eigenständig funktionierende Anwendungen und erweitern sie schrittweise. Dabei sollen sowohl die Umsetzung als auch die zugrunde liegenden Konzepte verstanden werden.

Langfristig schaffen die Projekte ein praktisches Fundament für:

- Platform Engineering
- Cloud Engineering & Architecture
- Machine Learning Engineering
- Site Reliability Engineering (SRE)
- AI Engineering
- Cybersecurity & DevSecOps

## Die drei Lernprojekte

| Projekt | Sprache | Schwerpunkt |
| --- | --- | --- |
| [`go-monitor`](go-monitor/) | Go | Netzwerk, Monitoring und Concurrency |
| [`py-log-analyzer`](py-log-analyzer/) | Python | Datenverarbeitung und Automatisierung |
| [`py-ml-pipeline`](py-ml-pipeline/) | Python | Machine Learning und MLOps |

Die Projekte beginnen unabhängig voneinander. Alle drei Projektordner sind bereits mit einer kurzen englischen README zum jeweiligen Vorhaben angelegt. Implementierung und Einrichtung der Entwicklungsumgebungen folgen erst beim Einstieg in das jeweilige Projekt.

## Arbeitsverzeichnis und Git

Unser lokaler Arbeitsbereich liegt unter:

```text
/Users/tillcornely/Desktop/github/engineering-lab/
```

Die geplante Struktur ist:

```text
engineering-lab/
├── README.md
├── go-monitor/
├── py-log-analyzer/
└── py-ml-pipeline/
```

Jedes Projekt erhält ein eigenes Git-Repository, eine eigene Dokumentation und eine unabhängige Entwicklungsumgebung. Das übergeordnete Verzeichnis `engineering-lab/` wird nicht als Git-Repository initialisiert. Diese README dient als gemeinsame Projektübersicht.

Für das Bewerbungsportfolio soll später ein tatsächlich umgesetztes Einzelprojekt wie `go-monitor` auf dem GitHub-Profil hervorgehoben werden. Der gemeinsame Lernplan ist keine fertige Arbeitsprobe.

## Unsere Arbeitsweise

- In kleinen, nachvollziehbaren Schritten entwickeln.
- Neues anhand konkreter Aufgaben erklären und Code gemeinsam nachvollziehen.
- Eigenes Verständnis und selbstständiges Entwickeln in den Mittelpunkt stellen.
- Mit wenigen Dateien beginnen und Standardbibliotheken bevorzugen.
- Frameworks, Abhängigkeiten und zusätzliche Architektur bei konkretem Bedarf einführen.
- Verständliche Funktionen, klare Verantwortlichkeiten und aussagekräftige Namen verwenden.
- Fehler explizit behandeln und automatisierte Tests schrittweise ergänzen.
- Kleine, nachvollziehbare Commits erstellen und die Dokumentation aktuell halten.
- Eingaben validieren, Timeouts berücksichtigen und keine Secrets im Repository speichern.

## Aktiver Fokus: Go Monitoring Agent

Ein schlanker Monitoring-Agent soll Websites, APIs und die Integrität von Dateien überwachen. Wir beginnen mit einem HTTP-Status-Checker.

### Kleine Startstruktur

```text
go-monitor/
├── main.go
├── go.mod
├── README.md
└── .gitignore
```

`main.go` enthält zunächst den Einstiegspunkt und die gesamte Programmlogik. `go.mod` definiert das Go-Modul und seine Abhängigkeiten. Der vorgesehene Modulpfad lautet `github.com/crnly/go-monitor`.

Die geplante Initialisierung aus dem Hauptverzeichnis:

```bash
mkdir go-monitor
cd go-monitor
go mod init github.com/crnly/go-monitor
git init
touch main.go README.md .gitignore
```

### Entwicklungsphasen

| Phase | Ergebnis | Lerninhalte |
| --- | --- | --- |
| 1. HTTP-Status-Checker | Eine URL prüfen und deren Antwortzeit messen | Packages, Variablen, Datentypen, Funktionen, HTTP-Anfragen, Fehlerbehandlung, Zeitmessung und CLI-Argumente |
| 2. Refactoring & Testing | Die Implementierung strukturieren und automatisiert testen | Rückgabewerte, Packages, Unit-Tests, `go test` und Table-driven Tests |
| 3. SHA-256 File Checker | Die Integrität ausgewählter Dateien prüfen | Dateisystem, Datei-I/O, SHA-256 und Hash-Vergleiche |
| 4. Mini-HTTP-Server | Monitoring-Ergebnisse über eine HTTP-API bereitstellen | Handler, Structs, JSON und REST-API-Grundlagen |
| 5. Monitoring Agent | Die Komponenten zusammenführen und Checks regelmäßig ausführen | Modularisierung, Goroutines, Channels, Concurrency, Timeouts und Synchronisierung |
| 6. Deployment & Observability | Die Anwendung für einen realistischen Betrieb vorbereiten | Konfiguration, strukturierte Logs, Docker, GitHub Actions, Prometheus und Security |

Das angestrebte Verhalten für Phase 1:

```bash
go run . https://example.com
```

Beispielausgabe; die Antwortzeit ist nur ein Beispiel:

```text
https://example.com: 200 OK (142ms)
```

### Architektur bei wachsendem Bedarf

Wenn zusätzliche Komponenten entstehen, kann sich die Struktur so entwickeln:

```text
go-monitor/
├── cmd/monitor/main.go
├── internal/
│   ├── checks/
│   │   ├── http.go
│   │   ├── http_test.go
│   │   └── file.go
│   ├── monitor/runner.go
│   └── api/server.go
├── configs/example.json
├── .github/workflows/test.yml
├── Dockerfile
├── go.mod
├── go.sum
├── README.md
└── .gitignore
```

`cmd/monitor/` enthält den Einstiegspunkt, `internal/checks/` die Prüfungen, `internal/monitor/` deren Planung und Ausführung und `internal/api/` die HTTP-API. Das Verzeichnis `internal/` beschränkt externe Imports gemäß den Go-Paketregeln. `go.sum` entsteht bei Bedarf durch Abhängigkeiten.

Weitere geplante Funktionen sind parallele Checks, konfigurierbare Intervalle, strukturierte Logs, eine Status-API und Prometheus-Metriken. Diese Struktur und Funktionen werden schrittweise eingeführt.

## Später: Python Log Analyzer

Ein CLI-Tool liest Logdateien ein, erkennt Fehler und Warnungen, zählt deren Häufigkeit und gibt Statistiken aus. Wiederkehrende Fehlermuster sollen sichtbar werden.

Als spätere Erweiterungen sind mehrere Logformate, JSON- und CSV-Export, zeitbasierte Auswertungen, Anomalieerkennung und eine Integration mit dem Go Monitoring Agent vorgesehen.

Geplante Startstruktur:

```text
py-log-analyzer/
├── src/log_analyzer/
│   ├── __init__.py
│   └── analyzer.py
├── tests/test_analyzer.py
├── pyproject.toml
├── README.md
└── .gitignore
```

## Später: Python ML Pipeline

Eine kleine Pipeline bildet den vollständigen ML-Workflow ab: Daten laden, vorbereiten und in Trainings- und Testdaten aufteilen, ein Modell trainieren, evaluieren und die Ergebnisse dokumentieren.

Spätere Erweiterungen können Experiment Tracking mit MLflow, Modellversionierung, automatisierte Tests, Model Serving über eine API, Containerisierung und Monitoring der Modellqualität sein.

Geplante Startstruktur:

```text
py-ml-pipeline/
├── src/ml_pipeline/
│   ├── __init__.py
│   ├── data.py
│   ├── train.py
│   └── evaluate.py
├── tests/
├── data/sample/
├── pyproject.toml
├── README.md
└── .gitignore
```

## Langfristige Perspektive

Die Projekte können später zu einer kleinen Monitoring- und Analyseplattform zusammenwachsen:

```text
Go Monitoring Agent
        │ Metriken und Logs
        ▼
Python Log Analyzer
        │ Aufbereitete Daten
        ▼
Python ML Pipeline
        │ Anomalieerkennung
        ▼
Monitoring API / Dashboard
```

Docker, Kubernetes, Terraform, GitHub Actions, Prometheus, Grafana, OpenTelemetry und MLflow kommen infrage, sobald die jeweiligen Grundlagen verstanden wurden und ein konkreter Anwendungsfall besteht.

## Aktueller Stand und nächste Schritte

Stand: 22. November 2026. Das Hauptverzeichnis und die drei Projektordner existieren. Jedes Projekt hat eine kurze englische README; Projektcode, Go-Modul und Git-Repositories wurden noch nicht angelegt.

**Aktives Projekt:** `go-monitor` · **Aktuelle Phase:** 1 – HTTP-Status-Checker

- [x] Arbeitsverzeichnis `engineering-lab` erstellen
- [x] Projektplan und Arbeitsweise in dieser README dokumentieren
- [x] Die drei Projektverzeichnisse mit englischer README anlegen
- [ ] Go-Modul initialisieren
- [ ] Git-Repository in `go-monitor` initialisieren
- [ ] Erste `main.go` und `.gitignore` erstellen
- [ ] HTTP-Status-Checker entwickeln
- [ ] Programm lokal ausführen und überprüfen
- [ ] Erste Änderungen committen

Bei weiteren Arbeitsschritten halten wir den aktuellen Stand und die nächsten Aufgaben hier nach.
