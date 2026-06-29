# Cato
> works for both  **Linux (Hyprland)** and **Windows**.


Cato is a tiny desktop cat that lives on your screen and follows your mouse cursor. It walks, runs, idles with different animations, and stays pinned above all your windows without stealing focus.


considering more animations and interactions .



| Arch | idk_lol |
|:-----|:--------|
| <img src="Screenshots/arch.png" alt="arch" /> | <img src="Screenshots/idk_lol.png" alt="idk_lol" /> |
| Licking | Running |
| <img src="Screenshots/licking.png" alt="licking" /> | <img src="Screenshots/running.png" alt="running" /> |


## Requirements

### Linux
- Go
- Hyprland ( for other DE's future maybe)

### Windows
- Go
- Windows 10/11





## Hyprland Configuration

Add the following window rules to your Hyprland config:

```lua
hl.window_rule({match = {title = "^(cato)$"}, no_focus = true})
hl.window_rule({match = {title = "^(cato)$"}, float = true})
hl.window_rule({match = {title = "^(cato)$"}, pin = true})
hl.window_rule({match = {title = "^(cato)$"}, no_shadow = true})
hl.window_rule({match = {title = "^(cato)$"}, no_blur = true})
hl.window_rule({match = {title = "^(cato)$"}, no_initial_focus = true})
hl.window_rule({match = {title = "^(cato)$"}, no_anim = true})
hl.window_rule({match = {title = "^(cato)$"}, move = {0, 0}})
```

Reload Hyprland after adding the rules.

```
hyprctl reload
```



## Windows Configuration

no configs required 


## Running

### Clone the repository

```
git clone https://github.com/HASSAAN-AMIN/cato
cd cato
```

### Linux

```
chmod +x run.sh
./run.sh
```

or :

```
go run .
```

or  :

```
go build -o cato 
./cato
```


### Windows

```
run.bat
```

Or:

```
go run .
```





### Automatic OS selection


it uses Go build tags to identify the OS and build according to that

- `linux.go` → Linux (Hyprland)
- `windows.go` → Windows


## Contribution

Contributions are always welcome!!!!!!!!!!!

If you'd like to add new animations, behaviors, features, improve the codebase, or fix bugs, feel free to open a pull request.

I'd especially love contributions that make **Cato work on more Linux distributions and desktop environments**, not just the current Arch Linux + Hyprland setup.

also contribuitions for MacOS support