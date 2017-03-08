# Todo

- `-fuzzy xxx` or `regex ` search
- `-output` for integration in scripts (needs to remove app log, and may
  be a `-select`):
  
```
cliper -fuzzy "https" -last 1 -output "{{ .clip }}" ls | open`
```

- CI and Tests
- Lint and refactoring
- Proper doc
- Fix Makefile tooling (see crosscompilation of sqlite)
