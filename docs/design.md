# SwiftlyGo Design Document
## SwiftlyGo Overview
#### Purpose
An Installer and Toolchain Manager for the Swift Programming Language on Linux.

Initially SwiftlyGo will be available for Debian / Ubuntu based distributions but can be easily extended to support other platforms.
#### Functions
* Check the host OS system for Swift compatibilty
* Detect the host Linux OS distribution and version
* Check Versions of Swift available for detected distribution
* Download and install Swift toolchains
* Install the system dependencies needed for Swift
* Activate and use different locally installed Swift versions
* Delete installed Swift toolchains
* Remove the SwiftlyGo app

#### Development Priorities
 
* Ease of development
* Portability
* Functionality
* Ease of distribution

When considering these priorities the decision was made to develop in the [Go](https://go.dev/) language and use the [Cobra Cli](https://pkg.go.dev/github.com/spf13/cobra) package.

#### Swift Toolchain Installation Location.
When choosing the installation location of the SwiftlyGo app and the Swift toolchains it was decided to use the standard installation location of `/usr/libexec/[name]/[version]/`.  
Each toolchain version is contained in its own directory.
```
/usr/libexec/swift
|-- 5.10.1
|   `-- usr
|       |-- bin
|       |-- include
|       |-- lib
|       |-- libexec
|       |-- local
|       `-- share
|-- 6.0
|   `-- usr
|       |-- bin
|       |-- include
|       |-- lib
|       |-- libexec
|       |-- local
|       `-- share
`-- 6.0.3
    `-- usr
        |-- bin
        |-- include
        |-- lib
        |-- libexec
        |-- local
        `-- share

```
The **SwiftlyGo** App is installed in the same location.
```
/usr/libexec/swiftlygo/
`-- bin
    |-- swiftlygo
    `-- swiftlygo.env
```
## SwiftlyGo Operational Flow Diagrams

### Installation of SwiftlyGo
```mermaid
graph TD
A[Run install script] --> B
B[Get system information] --> C
C[Check Swift compatibility ]
C --> |Yes| E
E[Create swiftlygo.env file] --> F
F[Install SwiftlyGo] --> I
I[Installation Complete]   
C --> |No| D[System not compatible] 
```

### Installation of a Swift Toolchain
```mermaid
graph TD;
A[swiftlygo install version] --> B;
B[Get Swift Version] --> C;
C[Get available versions] --> D;
D[Version available?];
D --> |Yes| E;
E[Version installed?];
E --> |Yes| H;
E --> |No| F;
F[Download Swift] --> G;
G[Extract Swift] --> H;
H[Create symlinks] --> I;
I[Dependencies installed?];
I --> |Yes| K;
I --> |No| J;
J[Install Dependencies] --> K;
K[Installation Complete];
D --> |No| M;
M[Version not available] --> N
N[Show available versions];

```

### Activation of a Swift Toolchain
```mermaid
graph TD
X[swiftlygo activate version] --> A
A[Get Swift Version] --> B
B[Get installed versions] --> C
C[Version available?]
C --> |Yes| F    
F[Update symlinks] --> G
G[Activation Complete]
C --> |No| D[Version not available] --> E[Show available versions]
```

### Deletion of a Swift Toolchain
```mermaid
graph TD
X[swiftlygo delete version] --> A
A[Get Swift Version] --> B
B[Get installed versions] --> C
C[Version available?]
C --> |Yes| F    
F[Remove symlinks] --> G
G[Remove version] --> H
H[Deletion Complete]
C --> |No| D[Version not available] --> E[Show available versions]
```
