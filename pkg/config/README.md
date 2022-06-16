# Config PKG

The config pkg is used to manage the carbonaut configuration file. The package exposes two functions to do so `ReadFile` & `WriteFile`.
The Carbonaut configuration is stored as yaml file, in the structure of the struct `CarbonConfig`.
`WriteFile` overwrites any existing configuration.

Upon initialization this is how the runtime looks like.

```mermaid
sequenceDiagram
    actor User
    participant API
    participant Config
    User->>API: Init Carbonaut
    API->>Config: WriteFile(input) <br> apply defaults
    Config->>Config: Create Carbonaut <br> config file
    Config->>Config: Update carbonaut context <br> to find the file later
    Config-->>API: config created 
    API-->>User: config created
```

Interacting with the a data provider this is how the runtime looks like.

```mermaid
sequenceDiagram
    actor User
    participant API
    participant Config
    User->>API: Import Data
    API->>Config: ReadFile()
    Config->>Config: Lookup carbonaut context
    Config->>Config: Read carbonaut config file
    Config-->>API: Carbonaut Config 
    API->>Data: Import Data with config
    Data->>Data: ...
    Data-->>API:  
    API-->>User: Data imported
```
