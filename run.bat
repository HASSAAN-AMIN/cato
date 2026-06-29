@echo off
setlocal enabledelayedexpansion

echo   Building Cato 
echo.

:: Check  Go
where go >nul 2>&1
if errorlevel 1 (
    echo  Go is not installed!
    echo.
    echo Download it from: https://go.dev/dl/
    pause
    exit /b 1
)

echo  Go version:
go version
echo.



echo  Building Cato...
go build -buildvcs=false -o cato.exe

if errorlevel 1 (
    echo.
    echo  Build failed! Check the error above.
    pause
    exit /b 1
)

echo.
echo  Build successful!
echo  Starting Cato...

start "" cato.exe

echo.
echo   Cato is running now ty for using it
pause