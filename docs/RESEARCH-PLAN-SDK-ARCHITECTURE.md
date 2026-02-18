# Research Plan: go-api-sdk-jamfprotect Context & Preferred SDK Architecture

This document captures the current context and a research plan for aligning **go-api-sdk-jamfprotect** with your preferred SDK architecture (from **go-api-sdk-virustotal**) while reverse-engineering logic from **terraform-provider-jamfprotect-1**.

---

## 1. Current Context Summary

### 1.1 go-api-sdk-jamfprotect (this repo)

**Purpose**: Go SDK for the Jamf Protect GraphQL API.

**Structure**:
- **Root**: `jamfprotect/new.go` – `Client` aggregates services; `NewClient(clientID, clientSecret, options...)` / `NewClientFromEnv()`.
- **Client layer**: `jamfprotect/client/` – Transport, OAuth2 (client credentials), GraphQL execution, constants, errors, headers, OTEL.
- **Interface**: `jamfprotect/interfaces/client.go` – `GraphQLClient` with `DoGraphQL(ctx, endpoint, query, variables, target, headers...)` and `GetLogger()`.
- **Services** (per domain): e.g. `jamfprotect/services/plans/`:
  - `queries.go` – GraphQL fragments and query/mutation strings (constants).
  - `models.go` – Request/response and API types (Plan, CreatePlanRequest, etc.).
  - `crud.go` – Service struct holding `interfaces.GraphQLClient`; CRUD methods call `s.client.DoGraphQL(...)` with queries from `queries.go`; validation via `client.ErrInvalidInput`.
  - `validators.go` – Input validation (e.g. `ValidatePlanID`, `ValidateCreatePlanRequest`).

**API**: GraphQL-only; OAuth2 client credentials; base URL `https://apis.jamfprotect.cloud`; main endpoint `/app`. Pagination is cursor-based (`nextToken`, `pageInfo.next`). See `jamfprotect/docs/api_client_spec.md`.

**Notable**: Services depend on `interfaces.GraphQLClient`, not concrete transport. Return signature is `(result, *interfaces.Response, error)` so callers can access response metadata.

---

### 1.2 terraform-provider-jamfprotect-1 (source of logic)

**Purpose**: Terraform provider that manages Jamf Protect resources (plans, analytics, action configs, telemetry, unified logging filters, etc.).

**Structure**:
- **Provider**: `internal/provider/provider.go` – Configures `internal/client.Client` (URL, client ID, secret), passes it as `DataSourceData` / `ResourceData`.
- **Client**: `internal/client/` – Own OAuth2 + `DoGraphQL(ctx, path, query, variables, target)` (no `*Response` return; errors only). No interfaces; concrete implementation.
- **API wrapper**: `internal/jamfprotect/` – One `Service` struct wrapping `*client.Client`. Per-entity files (e.g. `plan.go`, `analytic.go`) contain:
  - **GraphQL**: Same query/mutation constants as the SDK (e.g. `planFields`, `createPlanMutation`, `getPlanQuery`).
  - **Types**: Request/response structs (e.g. `PlanInput`, `Plan`, `PlanCommsConfig`).
  - **Methods**: `CreatePlan`, `GetPlan`, `UpdatePlan`, `DeletePlan`, `ListPlans` (and equivalent for other entities). They call `s.client.DoGraphQL(...)` and return `(Plan, error)` or `(*Plan, error)`.

**Resources**: e.g. `internal/resources/plan/` – `PlanResource` holds `*jamfprotect.Service`; in `Configure` it does `r.service = jamfprotect.NewService(client)`. CRUD (create.go, etc.) maps Terraform state ↔ API via `buildVariables` / `apiToState` and calls `r.service.CreatePlan`, `r.service.GetPlan`, etc.

**Relationship**: The provider embeds the full API logic (queries, types, variable building) inside the provider repo. The SDK is the intended replacement: same operations and types should live in the SDK; the provider should eventually depend on the SDK and call e.g. `sdkClient.Plans.CreatePlan(...)` instead of `jamfprotect.Service.CreatePlan(...)`.

---

### 1.3 go-api-sdk-virustotal (preferred SDK architecture)

**Purpose**: Reference pattern for your Go API SDKs (REST, not GraphQL).

**Structure**:
- **Root**: `virustotal/new.go` – `Client` aggregates services; `NewClient(apiKey, options...)` / `NewClientFromEnv()`.
- **Client layer**: `virustotal/client/` – Transport (e.g. resty), request helpers (`Get`, `Post`, `PostWithQuery`, etc.), auth, headers, pagination, OTEL.
- **Interface**: `virustotal/interfaces/client.go` – `HTTPClient` with Get/Post/Put/Patch/Delete/GetBytes/GetPaginated/PostMultipart, plus `GetLogger()` and `QueryBuilder()`. `Response` is a struct (status, headers, body, duration, etc.) returned alongside results.
- **Services** (per API area, e.g. `ioc_reputation_and_enrichment/urls/`):
  - **constants.go** – Endpoint path and relationship/option constants (e.g. `EndpointURLs`, `RelationshipComments`). Single source of truth for paths and enum-like values.
  - **crud.go** – Service struct holding `interfaces.HTTPClient`; `NewService(client)`; methods implement a **service interface** (e.g. `URLsServiceInterface`) so the concrete `Service` is testable/mockable. Methods return `(result, *interfaces.Response, error)`.
  - **models** – Often in same package or adjacent; request/response types and option structs.
  - **crud_test.go** – Unit tests with mocked HTTP (e.g. httpmock), using the interface.

**Patterns**:
- **Constants**: All endpoint paths and shared string constants in `constants.go` (or shared_models for cross-service relationships).
- **Interfaces**: Each service implements an interface that describes its public API; `var _ URLsServiceInterface = (*Service)(nil)`.
- **Return signature**: `(T, *interfaces.Response, error)` so callers can inspect headers/status/duration.
- **Transport abstraction**: Services depend only on `interfaces.HTTPClient`; no dependency on concrete transport.

---

## 2. Gap Analysis: JamfProtect SDK vs Preferred (VirusTotal) Architecture

| Aspect | VirusTotal (preferred) | JamfProtect SDK (current) | Notes |
|--------|------------------------|----------------------------|--------|
| **Transport interface** | `HTTPClient` (Get, Post, …) | `GraphQLClient` (DoGraphQL) | Different protocol; JP is appropriate for GraphQL. |
| **Return signature** | `(T, *Response, error)` | `(T, *Response, error)` | Aligned. |
| **Constants** | `constants.go` per service (endpoints, relationships) | Endpoint in `client/constants.go`; queries in `queries.go` | JP could add a `constants.go` per service for endpoint path and any enum/relationship constants. |
| **Service interface** | Each service has `XxxServiceInterface`; `var _ XxxServiceInterface = (*Service)(nil)` | No service-level interface | Add interfaces for testability and consistency. |
| **Validators** | Not emphasized in VT snippet | `validators.go` per service | Keep; aligns with “validate before API call” in api_client_spec. |
| **Queries** | N/A (REST) | `queries.go` (GraphQL strings) | Keep; GraphQL-specific. |
| **File layout** | constants, crud, models, crud_test | queries, models, crud, validators | JP is fine; consider adding constants.go and service interface. |

---

## 3. Research Plan

### Phase 1 – Confirm current state and provider mapping

1. **Inventory terraform-provider-jamfprotect-1 entities**
   - List all `internal/jamfprotect/*.go` entity files and the corresponding SDK services (e.g. plan → Plans, analytic → Analytics, telemetry_v2 → TelemetryV2).
   - For each entity, note: GraphQL operations (queries/mutations), types (input/output), and any provider-specific helpers (e.g. `buildVariables`, special handling for commsConfig in plan).

2. **Map provider → SDK coverage**
   - Table: Provider entity | SDK service | Methods already in SDK | Methods/types still only in provider.
   - Identify any provider-only types or defaults (e.g. comms FQDN placeholder) that should stay in provider vs move to SDK.

3. **Document API surface differences**
   - Provider: `(Plan, error)` / `(*Plan, error)`; no `*Response`.
   - SDK: `(*Plan, *interfaces.Response, error)`.
   - List any provider error handling (e.g. `IsNotFoundError`) that the SDK already supports (e.g. `client.ErrNotFound`) so the provider can switch to the SDK and retain behavior.

### Phase 2 – Align SDK with preferred architecture

4. **Introduce per-service constants**
   - Add `constants.go` in each SDK service package where it adds value (e.g. endpoint path if it ever differs from `client.EndpointApp`, or enum/relationship constants used in variables or validation). Keep GraphQL strings in `queries.go`.

5. **Add service interfaces**
   - For each service (Plans, Analytics, ActionConfigs, etc.), define `XxxServiceInterface` in the same package with all public methods. Add `var _ XxxServiceInterface = (*Service)(nil)` in crud.go. This enables mocks and matches VT pattern.

6. **Optional: shared response/error helpers**
   - If the provider will use the SDK, ensure it can map SDK errors to its diagnostics (e.g. `errors.Is(err, client.ErrNotFound)`). No change needed if already present.

### Phase 3 – Reverse-engineering checklist (provider → SDK)

7. **Per-entity checklist**
   - For each entity in `internal/jamfprotect/`:
     - [ ] All GraphQL operations (queries/mutations) exist in SDK `queries.go` (or equivalent).
     - [ ] All request/response types exist in SDK `models.go` (naming can differ; align with SDK style e.g. CreateXxxRequest, UpdateXxxRequest).
     - [ ] Variable-building logic matches (create/update inputs → GraphQL variables). SDK may use a single `buildXxxVariables(req any)`; confirm it covers both create and update like the provider.
     - [ ] Pagination: list operations in SDK consume `pageInfo.next` and return full list (or expose cursor if needed); behavior matches provider’s list usage.
     - [ ] Validation: required fields and ID format validated in SDK (validators.go) and return `client.ErrInvalidInput` where appropriate.
     - [ ] Error handling: not-found and GraphQL errors mapped so provider can use `errors.Is(err, client.ErrNotFound)` etc.

8. **Provider migration path**
   - Plan steps to switch provider from `internal/client` + `internal/jamfprotect` to `go-api-sdk-jamfprotect`: replace `client.NewClient(...)` with `jamfprotect.NewClient(...)` (or FromEnv), pass SDK client into resources; replace `jamfprotect.NewService(providerClient)` with using SDK client’s services (e.g. `sdkClient.Plans`). Update each resource’s CRUD to use SDK method signatures (and `*interfaces.Response` if needed). Remove or thin out `internal/jamfprotect` and `internal/client` once all resources use the SDK.

### Phase 4 – Documentation and tests

9. **Document architecture**
   - Add or update `docs/` in the SDK repo: high-level architecture (transport vs services vs interfaces), how a new service should be added (queries, models, crud, validators, constants, interface), and how the SDK relates to the Terraform provider (who consumes it, migration status).

10. **Tests**
    - For each service: unit tests that mock `GraphQLClient` (or use a small in-memory GraphQL stub) to assert request variables and response handling. Optionally add an acceptance test with real credentials (similar to VT’s acceptance config) if desired.

---

## 4. Immediate Next Steps (Suggested)

1. **Inventory** – List all `internal/jamfprotect/*.go` files and the corresponding SDK `services/*` packages; create the mapping table (Phase 1.1–1.2).
2. **Pick one entity** – Choose one entity (e.g. Plan) and complete the per-entity checklist (Phase 3.7) in detail; fix any gaps in the SDK (missing types, validators, or variable building).
3. **Add interface + constants** – For the Plans service, add `PlansServiceInterface` and a `constants.go` if any constants are missing (Phase 2.4–2.5).
4. **Document** – Write or update `docs/ARCHITECTURE.md` (or similar) summarizing transport, services, interfaces, and provider migration (Phase 2 + 3.8 + 4.9).

---

## 5. References

- **go-api-sdk-jamfprotect**: `jamfprotect/new.go`, `jamfprotect/client/`, `jamfprotect/interfaces/client.go`, `jamfprotect/services/*/` (queries, models, crud, validators), `jamfprotect/docs/api_client_spec.md`.
- **terraform-provider-jamfprotect-1**: `internal/client/`, `internal/jamfprotect/service.go` + `*/*.go`, `internal/resources/*/` (resource.go, crud.go, helpers, data_source).
- **go-api-sdk-virustotal**: `virustotal/new.go`, `virustotal/interfaces/client.go`, `virustotal/client/request.go`, `virustotal/services/ioc_reputation_and_enrichment/urls/` (constants.go, crud.go, crud_test.go).
