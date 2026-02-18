# Service Inventory and Delta: SDK vs Terraform Provider

This document inventories services in **go-api-sdk-jamfprotect** and **terraform-provider-jamfprotect-1** (internal/jamfprotect) and summarizes the delta for provider migration.

---

## 1. SDK Services (go-api-sdk-jamfprotect)

The SDK exposes **9 services** on the root `Client`. Each service has full CRUD + List where the API supports it.

| SDK Client Field       | Package (dir)           | Create | Get | Update | Delete | List |
|------------------------|-------------------------|--------|-----|--------|--------|------|
| **ActionConfigs**      | action_configuration    | ✓      | ✓   | ✓      | ✓      | ✓    |
| **Plans**              | plan (pkg: plans)       | ✓      | ✓   | ✓      | ✓      | ✓    |
| **Analytics**          | analytics               | ✓      | ✓   | ✓      | ✓      | ✓    |
| **AnalyticSets**       | analyticsets            | ✓      | ✓   | ✓      | ✓      | ✓    |
| **ExceptionSets**      | exceptionsets           | ✓      | ✓   | ✓      | ✓      | ✓    |
| **PreventLists**       | preventlists            | ✓      | ✓   | ✓      | ✓      | ✓    |
| **TelemetryV2**        | telemetryv2             | ✓      | ✓   | ✓      | ✓      | ✓    |
| **USBControlSets**     | usbcontrolsets          | ✓      | ✓   | ✓      | ✓      | ✓    |
| **UnifiedLoggingFilters** | unifiedloggingfilters | ✓   | ✓   | ✓      | ✓      | ✓    |

**SDK method signature (typical):**  
`(ctx, id/req) → (result, *interfaces.Response, error)` or for Delete: `(ctx, id) → (*interfaces.Response, error)`.

---

## 2. Provider Entities (terraform-provider-jamfprotect-1 internal/jamfprotect)

The provider has **10 entity files** (excluding `service.go` and `service_test.go`). Each defines a set of methods on the single `Service` type that wraps the provider’s GraphQL client.

| Provider File                      | Create | Get | Update | Delete | List |
|------------------------------------|--------|-----|--------|--------|------|
| action_configuration.go            | ✓      | ✓   | ✓      | ✓      | ✓    |
| analytic.go                        | ✓      | ✓   | ✓      | ✓      | ✓    |
| analytic_set.go                    | ✓      | ✓   | ✓      | ✓      | ✓    |
| exception_set.go                   | ✓      | ✓   | ✓      | ✓      | ✓    |
| custom_prevent_list.go             | ✓      | ✓   | ✓      | ✓      | ✓    |
| plan.go                            | ✓      | ✓   | ✓      | ✓      | ✓    |
| telemetry_v2.go                    | ✓      | ✓   | ✓      | ✓      | ✓    |
| unified_logging_filter.go          | ✓      | ✓   | ✓      | ✓      | ✓    |
| removable_storage_control_set.go   | ✓      | ✓   | ✓      | ✓      | ✓    |

**Provider method signature (typical):**  
`(ctx, id/input) → (result, error)` or `(*result, error)` or `error` (Delete). No `*Response` return.

---

## 3. Mapping: Provider Entity → SDK Service

| Provider entity (file)              | SDK service (Client field) | Notes |
|-------------------------------------|----------------------------|--------|
| action_configuration                | **ActionConfigs**          | 1:1    |
| analytic                            | **Analytics**              | 1:1    |
| analytic_set                        | **AnalyticSets**           | 1:1    |
| exception_set                       | **ExceptionSets**          | 1:1    |
| custom_prevent_list                 | **PreventLists**           | Same API (PreventList); “custom” is provider naming |
| plan                                | **Plans**                  | 1:1    |
| telemetry_v2                        | **TelemetryV2**            | 1:1    |
| unified_logging_filter              | **UnifiedLoggingFilters**  | 1:1    |
| removable_storage_control_set      | **USBControlSets**         | Same API (USBControlSet); provider uses “removable storage” name |

**Conclusion:** Every provider entity has a corresponding SDK service with equivalent operations. There is **no missing service** in the SDK.

---

## 4. Delta Summary

### 4.1 No missing coverage

- All 9 provider-facing domains are implemented in the SDK with Create, Get, Update, Delete, and List (where applicable).
- The provider can, in principle, stop using `internal/jamfprotect` and `internal/client` and use the SDK for all current resources and data sources.

### 4.2 Signature and return-type differences

| Aspect | Provider | SDK |
|--------|----------|-----|
| Return | `(T, error)` / `(*T, error)` / `error` | `(T, *interfaces.Response, error)` / `(*T, *interfaces.Response, error)` / `(*interfaces.Response, error)` |
| Delete | `error` | `(*interfaces.Response, error)` |

**Migration:** Resource/CRUD code that today does `plan, err := r.service.GetPlan(ctx, id)` can become `plan, _, err := sdkClient.Plans.GetPlan(ctx, id)` and ignore the second return value. Delete calls can ignore the `*Response` from the SDK.

### 4.3 Naming and types

- **Prevent list:** Provider uses “CustomPreventList” / “CustomPreventListInput”; SDK uses “PreventList” / “CreatePreventListRequest” etc., aligned with the API. Migration is a rename in the provider’s mapping layer.
- **USB / removable storage:** Provider uses “RemovableStorageControlSet”; SDK uses “USBControlSet” (API name). Migration is a rename when mapping provider schema ↔ SDK types.
- **List types:** Provider uses list-item types where it only needs minimal fields (e.g. `ActionConfigListItem`, `ExceptionSetListItem`). SDK also uses list-item types where the API returns a list shape (e.g. `ListActionConfigs` → `[]ActionConfigListItem`). If the provider currently assumes a different list shape, migration may need a thin adapter or field mapping.

### 4.4 Error handling

- Provider uses `internal/common/helpers.IsNotFoundError(err)` for remove-from-state behaviour.
- SDK uses sentinel errors (e.g. `client.ErrNotFound`, `client.ErrInvalidInput`). Migration: have the provider use `errors.Is(err, client.ErrNotFound)` (or equivalent) when calling SDK methods, and keep the same CRUD behaviour.

### 4.5 Configuration and client construction

- **Provider today:** Builds `internal/client.Client` (URL, client ID, secret), passes it as provider data; each resource/data source does `jamfprotect.NewService(client)` and uses `*jamfprotect.Service`.
- **With SDK:** Build `jamfprotect.NewClient(clientID, clientSecret, jamfprotect.WithBaseURL(url))` (or `NewClientFromEnv`), pass the SDK `*jamfprotect.Client` as provider data. Resources/data sources use `sdkClient.Plans`, `sdkClient.Analytics`, etc., and no longer call `jamfprotect.NewService(...)`.

---

## 5. Per-entity migration checklist (high level)

For each resource/data source:

1. **Configure:** Accept `*jamfprotect.Client` from provider data instead of `*client.Client`; remove `jamfprotect.NewService(client)`.
2. **CRUD:** Replace `r.service.CreatePlan(...)` with `sdkClient.Plans.CreatePlan(...)` (and similarly for Get/Update/Delete/List). Adapt to `(result, *Response, error)` by ignoring or using the second return value.
3. **Types:** Map Terraform state ↔ SDK request/response types (e.g. Plan, CreatePlanRequest, PreventList, USBControlSet). Rename or adapt where the provider currently uses different names (e.g. CustomPreventList → PreventList, RemovableStorageControlSet → USBControlSet).
4. **Errors:** Replace not-found checks with `errors.Is(err, client.ErrNotFound)` (or the SDK’s equivalent) so that state removal and diagnostics stay correct.
5. **Cleanup:** Once all resources and data sources use the SDK, remove or stub out `internal/jamfprotect` and the provider’s `internal/client` GraphQL implementation.

---

## 6. Summary table: SDK vs provider operations

| Domain        | SDK service        | Provider entity              | Operations in both | Delta (for migration)        |
|---------------|--------------------|------------------------------|--------------------|------------------------------|
| Action config | ActionConfigs      | action_configuration         | C R U D L          | Signature; type names        |
| Analytic      | Analytics          | analytic                     | C R U D L          | Signature                    |
| Analytic set  | AnalyticSets       | analytic_set                 | C R U D L          | Signature                    |
| Exception set | ExceptionSets      | exception_set                | C R U D L          | Signature; list type naming  |
| Prevent list  | PreventLists       | custom_prevent_list          | C R U D L          | Signature; Custom → Prevent  |
| Plan          | Plans              | plan                         | C R U D L          | Signature                    |
| Telemetry v2  | TelemetryV2        | telemetry_v2                 | C R U D L          | Signature                    |
| Unified log   | UnifiedLoggingFilters | unified_logging_filter     | C R U D L          | Signature                    |
| USB control   | USBControlSets     | removable_storage_control_set| C R U D L          | Signature; Removable → USB   |

**Bottom line:** The SDK already has full coverage of the provider’s current API surface. The delta is signature and naming adaptation plus switching the provider to use the SDK client and error handling; no new SDK services or methods are required for the existing provider resources and data sources.
