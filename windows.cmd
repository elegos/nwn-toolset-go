echo off
set arg=%1
shift

if "%arg%" == "/?" (
  echo "/c clean binaries"
  echo "/t run tests"
  echo "/a clean and rebuild binaries"
  exit /b
)

if "%arg%" == "/c" (
  ninja -f windows.ninja -t clean
  exit /b
)

if "%arg%" == "/t" (
  set "GOPATH=%cd%"
  set "GOBIN=%cd%\bin"

  go test -v -cover aurora/...
  exit /b
)

if "%arg%" == "/a" (
  ninja -f windows.ninja -t clean
)

ninja -f windows.ninja
