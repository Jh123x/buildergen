# Refactor Plan: Layered Architecture with Google Wire

## Objective
Refactor the `buildergen` codebase into a strictly layered architecture using `google/wire` for dependency injection. The goal is to achieve high modularity, testability, and clear separation of concerns while ensuring no circular dependencies exist and following the Dependency Inversion Principle.

## Proposed Architecture (Layered DAG)

The application will be organized into layers that form a **Directed Acyclic Graph (DAG)**. Dependencies can only flow "downwards" through interfaces.

### 1. Handler Layer (`pkg/handler`)
* **Responsibility**: The entry point for the application logic. It parses input from various sources (CLI flags, YAML config files) and acts as a **translator**, mapping raw input into specific, minimal configuration objects required by the services.
* **Interactions**: 
    * Calls the **Service Layer** via a defined interface.
    * Depends on **Model** and **Common** layers.
* **Interface Ownership**: The `Handler` defines the interface it expects from the `Service`.

### 2. Service Layer (`pkg/service`)
* **Responsibility**: The core business logic orchestration. It receives granular configuration objects and coordinates the parsing, generation, and writing processes.
* **Interactions**:
    * Calls the **Parser Layer** and **Repository Layer** via defined interfaces.
    * Depends on **Model**, **Common**, and specialized component layers (like `pkg/parser`).
* **Interface Ownership**: The `Service` defines the interfaces it expects from its dependencies (e.g., `Parser` and `Repository`).

### 3. Component Layers (`pkg/parser`, `pkg/executor`)
* **Responsibility**: Specialized technical logic. 
    * `pkg/parser`: Handles AST manipulation and Go file parsing.
    * `pkg/executor`: Handles execution of external tools (e.g., `go fmt`).
* **Interactions**: Implements interfaces defined by the `Service` layer.

### 4. Repository / Writer Layer (`pkg/repository`)
* **Responsibility**: Handles all physical I/O operations, including filesystem side effects (creating directories, writing files) and STDOUT streaming.
* **Interactions**:
    * Depends on **Model** and **Common**.
* **Interface Ownership**: The `Service` defines the interface it expects from the `Repository`.

### 5. Model Layer (`pkg/model`)
* **Responsibility**: Contains pure data structures (DTOs) shared across all layers.
* **Constraints**: **Zero dependencies**. This layer must not import any other package from the project.

### 6. Common Layer (`pkg/common`)
* **Responsibility**: Provides cross-cutting concerns such as logging and custom error types.
* **Constraints**: **Strict Isolation**. `pkg/common` must have zero imports from the rest of the project to prevent circular dependencies.

---

## Dependency Rules & Constraints

1.  **Interface Inversion (The Consumer Defines)**: To avoid "Dependency Inversion Traps," interfaces must be defined by the layer that *consumes* them, not the layer that implements them. 
    *   If `Service` uses `Repository`, the interface lives in `pkg/service`.
2.  **Granular Configuration Mapping**: The `Handler` must not pass a single "God Object" (a massive configuration struct) through every layer. Instead, it should map raw input into specific, minimal configuration structs required by each service method. This improves testability and adheres to the Interface Segregation Principle.
3.  **Strict DAG**: A layer can only depend on layers "below" it. 
    *   `Handler` $\rightarrow$ `Service` $\rightarrow$ `Parser`/`Repository`
    *   All layers $\rightarrow$ `Model` / `Common`
4.  **Wire Integration**:
    *   Use `wire.NewSet` for each package to group providers (e.g., `service.ProviderSet`).
    *   The `wire.go` file will compose these sets to build the final application graph.
5.  **Minimal `main.go`**: The entry point should only initialize the Wire injector, handle top-level errors, and manage OS exit codes.

---

## Implementation Roadmap

### Phase 1: Foundation (Model & Common)
1.  Move all shared structs from `pkg/cmd` and `pkg/parser` into `pkg/model`.
2.  Extract logging and error definitions into `pkg/common`.
3.  **Verification**: Ensure `pkg/model` and `pkg/common` have no internal project imports.

### Phase 2: Component & Repository (Bottom-Up)
1.  Implement `pkg/parser` as a standalone technical component.
2.  Implement `pkg/repository` to handle all filesystem side effects.
3.  Define the `FileRepository` and `Parser` interfaces within the `Service` layer's scope.

### Phase 3: Service & Handler (Top-Down)
1.  Refactor business logic into `pkg/service`, using the interfaces defined in Step 2.
2.  Implement `pkg/handler` to parse CLI/YAML and map them to granular service configs.
3.  Create `wire.go` in the root directory, composing `ProviderSets` from each layer.
4.  Run `go generate ./...`.

### Phase 4: Final Integration & Cleanup
1.  Refactor `main.go` to be a thin wrapper around the Wire-generated injector.
2.  **Verification**: Run all existing tests. Ensure parser logic remains identical in behavior.

---

## Edge Cases & Risks

* **Circular Dependencies**: If an error in `common` needs to wrap a model, it could cause a cycle. **Mitigation**: Keep `common` strictly independent; perform error wrapping in the `Service` or `Handler` layers instead.
* **External Tooling**: Running commands like `go fmt`. **Mitigation**: Treat the Go CLI as an external dependency via an `Executor` interface injected into the services.
* **Filesystem Side Effects**: Ensuring directories exist. **Mitigation**: The `Repository` layer is responsible for all filesystem side effects to keep the `Service` layer pure and easily mockable.

## Clarifying Questions (Resolved)
1. **Logger Injection**: The `Logger` will be an injected dependency, potentially encapsulated within a specialized interface provided by the `common` layer.
2. **Context Propagation**: `context.Context` must be passed through all interface methods to ensure proper cancellation and timeout support for long-running generation tasks.