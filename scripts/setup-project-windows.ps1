param (
    [string]$RUN_ENV = "dev"
)

if ($RUN_ENV -eq "dev") {
    Write-Output "Running in dev mode"
    Copy-Item -Path "environments/.env.dev" -Destination ".env" -Force
    Copy-Item -Path "environments/.env.dev" -Destination "./cmd/api/application/.env" -Force
    Copy-Item -Path "environments/.env.dev" -Destination "./cmd/api/services/.env" -Force
} else {
    Write-Output "Running in prod mode"
    Copy-Item -Path "environments/.env.prod" -Destination ".env" -Force
    Copy-Item -Path "environments/.env.prod" -Destination "./cmd/api/application/.env" -Force
    Copy-Item -Path "environments/.env.prod" -Destination "./cmd/api/services/.env" -Force
}