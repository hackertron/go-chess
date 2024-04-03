# go-chess

```
cmd/
  server/
    main.go <- initialize server command 
internal/
  views/
    public/
      htmx.js
    home.templ
    layout.templ
  handlers/
    home.go
  models/
    user.go <- backend agnostic user model
  data/
    user.go <- user data storage (e.g. postgres db access)
```