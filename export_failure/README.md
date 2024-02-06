# Dagger discards files :( 

I'm running gradle tests, and when they fail, I want to gather the generated junit test reports. I cannot do this because when dagger exits 0, the resulting container does not have the files I desire.

Each test simulates a successful test run, and a failed test run, and their corresponding outputs. 

In both cases, I expect to be able to retrieve files - even if the script fails.

## Run da tests.

```golang
go mod tidy
go test ./...
```