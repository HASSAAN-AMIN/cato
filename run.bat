@echo off

where go >nul 2>&1
if errorlevel 1 (
    echo oh hoo go isnt installed 
    echo.
    echo download go first 
    echo https://go.dev/dl/
    pause
    exit /b 1
)

echo Building Cato...
go build -o cato.exe

if errorlevel 1 (
    echo smth baddd happened
    pause
    exit /b 1
)

echo Running Cato nowwwwwwwww
cato.exe