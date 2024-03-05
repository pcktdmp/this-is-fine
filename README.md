# this-is-fine
This webservice always tells everything is fine.
Handy for mocking backends.

## Usage

`docker run -p 8080:8080 -e THIS_IS_FINE_LOG_LINES="INFO:info-line,ERROR:error-line,WARN:warning-line,DEBUG:debug-line ghcr.io/pcktdmp/this-is-fine:latest`
