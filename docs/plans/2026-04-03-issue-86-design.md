# Issue 86 Design

## 1. Summary

Issue 86 requires this repository's checked-in protobuf-generated artifacts to be refreshed against the latest XMTP proto definitions on the `xmtp/proto` `main` branch. The change should be generator-driven, preserve this repository's local notification service contract, and result in a PR containing only the generated updates and any minimal dependency adjustments required to keep the codebase green.

## 2. Project Goals & Non-Goals

**Goals**

- Refresh generated protobuf outputs derived from `buf.build/xmtp/proto` so they match the latest upstream `main` schema available at execution time.
- Preserve generation of this repository's local notification service protos from `proto/notifications/v1/service.proto`.
- Keep all checked-in generated surfaces consistent across Go server code, Swift client artifacts, and integration test client code.
- Verify the regenerated code still builds and passes the repository's automated validation commands.

**Non-Goals**

- Changing the semantics of `proto/notifications/v1/service.proto` or the notification server API.
- Refactoring application logic outside the minimal adjustments needed to compile against regenerated protos.
- Introducing new code generation tooling beyond the Buf-based workflow already present in the repository.

## 3. Context

**Catalysts**

- GitHub issue: https://github.com/xmtp/example-notification-server-go/issues/86

**Codebase**

- `dev/gen-proto` regenerates local notification Go and TypeScript outputs, then regenerates upstream XMTP proto outputs, and clears `pkg/proto` before generation.
- `buf.gen.yaml` defines root-level generation targets, including `pkg/proto` and `swift/`.
- `proto/buf.gen.yaml` defines local notification service generation into `pkg/proto` and `integration/src/gen`.
- `pkg/proto/` contains checked-in Go artifacts generated from both local and upstream protos.
- `swift/` contains checked-in Swift artifacts generated from both local and upstream protos.
- `integration/src/gen/` contains the generated TypeScript integration client for the local notification service.
- `.github/workflows/test.yml`, `.github/workflows/lint.yml`, and `.github/workflows/buf.yml` define the repository's verification baseline.

**External docs**

- XMTP proto upstream: https://github.com/xmtp/proto
- Buf module source used by generation: `buf.build/xmtp/proto`

**Impact area**

- `pkg/proto/**`
- `swift/**`
- `integration/src/gen/**`
- `go.mod` / `go.sum` only if regeneration forces dependency updates

## 4. System Design

**Architecture overview**

The repository already has the required generation workflow, but it is split across two Buf templates. The implementation will use `proto/buf.gen.yaml` for local notification Go and TypeScript artifacts, `buf.gen.yaml` for local Swift notification artifacts, and the existing upstream generation path for XMTP proto artifacts from the latest available upstream schema. After regeneration, any compile or test failures caused by updated generated types will be fixed with the smallest application-side changes necessary to restore compatibility.

**New or modified interfaces**

No new interfaces are expected. Generated interfaces in `pkg/proto`, `swift`, and `integration/src/gen` may change to reflect upstream proto evolution.

**Key functions**

- `dev/gen-proto` SHALL remain the entrypoint for upstream proto regeneration.
- Root `buf.gen.yaml` SHALL remain the source of truth for cross-language generation targets.
- `proto/buf.gen.yaml` SHALL remain the source of truth for local notification service generation targets.

**Alternatives considered**

- Manual edits to checked-in generated files were rejected because they would drift from the established Buf workflow.
- Updating only Go outputs was rejected because the repository checks in generated Swift and integration client artifacts that must stay synchronized with the same schema set.

## 5. Libraries & Utilities Required

**External dependencies**

| Package | Version | Purpose |
|---------|---------|---------|
| `buf` | existing CLI in repo environment | Regenerate protobuf artifacts from local and upstream schemas |

**Internal modules**

| Module | Path | Purpose |
|--------|------|---------|
| proto generator | `dev/gen-proto` | Regenerates upstream and local proto outputs |
| root buf config | `buf.gen.yaml` | Defines Go and Swift generation targets |
| local buf config | `proto/buf.gen.yaml` | Defines Go and integration TypeScript generation targets |

## 6. Testing & Validation

### Acceptance Criteria

1. WHEN the repository proto refresh workflow is run against the latest `buf.build/xmtp/proto` module THE SYSTEM SHALL regenerate checked-in upstream XMTP proto artifacts under `pkg/proto` and `swift` to match the latest upstream schema.
2. WHEN proto regeneration completes THE SYSTEM SHALL preserve generation of the local notification service artifacts under `pkg/proto/notifications/v1`, `swift/notifications/v1`, and `integration/src/gen/notifications/v1`.
3. WHEN regenerated proto outputs introduce compile-time API changes THE SYSTEM SHALL update the minimal affected application code so the repository builds successfully without changing notification service behavior.
4. THE SYSTEM SHALL NOT include manual edits to generated protobuf files that cannot be reproduced by the repository's Buf generation configuration.
5. WHEN verification commands are run after regeneration THE SYSTEM SHALL pass the repository's unit test and lint-relevant checks required for a reviewable PR.

### Edge Cases

- Upstream proto changes may remove or rename generated fields or packages used by server code.
- Plugin version differences may change generated file headers or import layouts without semantic API changes.
- Local notification service generation must remain intact even though `dev/gen-proto` removes `pkg/proto` before regeneration.
- Generated output may require `go.mod` or `go.sum` updates if the generated code now references newer runtime packages.

### Verification Commands

- `./dev/gen-proto`
- `buf generate`
- `go test -p 1 ./...`
- `buf lint proto`
