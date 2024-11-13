# Architecture

## Init

### Folder Structure

```mermaid
classDiagram
    GotMainFolder --> Meta
    GotMainFolder --> Blob
    Blob --> Commit
    Commit --> FileDelta
    Blob --> FTree
    GotMainFolder --> Config
    class GotMainFolder {
        name: .got
    }
    class Meta {
        name: meta
    }
    class Commit {
        +string Hash
        +FileDelta
    }
    class FTree {
        Abstraction of files
    }
    class FileDelta {
        Lines Added
        Lines Deleted
    }
    class Blob {
        intermediate folder
    }
    class Config {
        Possible configs
    }
```