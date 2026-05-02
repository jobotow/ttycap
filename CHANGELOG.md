# Changelog
All notable changes to termcolor will be documented in this file

## [0.2.0] - 2026-05-02

### Added
- Color level detection is now made robust with additions such as (Force-color level detection, CI system detection and checkTermProgram)

### Fixed
- Color flag parser now treats values (negative numbers and "-") as values instead of flags.
- Dropped `SupportsColor()`, `Disabled()`, `Enabled()`, and `IsTerminal()` funtions,  for one function `Capability()`


---

## [0.1.1] - 2026-04-30

### Fixed
- Color flags without values now return true, which implies flag was present

---

## [0.1.0] - 2026-04-30

### Added
- Detects whether a terminal supports color and at what level.
- Check if color has being explicitly enabled or disabled via flags or environment variables.
- Support for `--color` and `--no-color` flags in both GNU and POSIX styles.
- Respects NO_COLOR and FORCE_COLOR environment variables.
- Falls back to auto-detection when no color flags are provided
- Detects if output is a TTY.
- Automatically handles when user uses `--color=auto` or `--color=tty` or `--color tty`
- Simple API: Enabled(), Disabled(), Auto(), Capability(), SupportsColor(), IsTerminal().

---
