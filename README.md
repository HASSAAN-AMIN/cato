# Stay Away -- underdevelopment rn --

# Cato



almost done for arch linux hyprland version






add this in config files please


```
hl.window_rule({match = {title = "^(cato)$"}, no_focus = true})
hl.window_rule({match = {title = "^(cato)$"}, float = true})
hl.window_rule({match = {title = "^(cato)$"}, pin = true})
hl.window_rule({match = {title = "^(cato)$"}, no_shadow = true})
hl.window_rule({match = {title = "^(cato)$"}, no_blur = true})
hl.window_rule({match = {title = "^(cato)$"}, no_initial_focus = true})
hl.window_rule({match = {title = "^(cato)$"}, no_anim = true})
hl.window_rule({match = {title = "^(cato)$"}, move = {0, 0}})
```



to run 

```
go mod tidy
```

and then

```
cd src
```

```
go run main.go
```