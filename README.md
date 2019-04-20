# web-temp

This script reads an MCP9808 temperature sensor from a Raspberry Pi and displays it on a webpage.

## Prerequisites

### Using with a MCP9808 temperature sensor
You'll need to wire it up

## Building

For a Raspberry Pi Zero
```
make
```

For local environment
```
make build-local
```

## Execute the script
#### Normal mode with sensor
```
./web-temp
```

#### Dev mode
```
./web-temp -dev true
```

#### Alternate port
```
./web-temp -dev=true -port=8080
```
