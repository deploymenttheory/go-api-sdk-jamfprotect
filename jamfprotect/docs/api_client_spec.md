# Jamf Protect API Client Specification

This document outlines the critical aspects of the Jamf Protect GraphQL API that influence SDK design and implementation.

## Authentication

### OAuth2 Client Credentials Flow

**Method**: OAuth2 Client Credentials Grant  
**Token Endpoint**: `{baseURL}/token`  
**Request Format**: JSON  
**Token Type**: Bearer (Access Token)

#### Token Request

```json
{
  "client_id": "your-client-id",
  "password": "your-client-secret"
}
```

#### Token Response

```json
{
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_in": 3600,
  "token_type": "Bearer"
}
```

**Token Characteristics**:
- Access tokens expire after the duration specified in `expires_in` (typically 3600 seconds)
- No refresh token is provided; clients must request a new access token when expired
- Token refresh should occur 60 seconds before expiry to avoid race conditions
- Authorization header format: `Authorization: {access_token}` (no "Bearer" prefix required)

**SDK Implementation**:
- Implement thread-safe token caching with automatic refresh
- Use `singleflight` to prevent concurrent token requests
- Pre-emptively refresh tokens 60 seconds before expiry
- Redact sensitive credentials in logs

## API Endpoints

### Base URL

**Production**: `https://apis.jamfprotect.cloud`

### GraphQL Endpoints

The API provides two GraphQL endpoints:

1. **Main API Endpoint** (`/app`)
   - Full API access
   - Complete schema with all types and operations
   - **Recommended for SDK usage**

2. **Limited Schema Endpoint** (`/graphql`)
   - Reduced schema for specific use cases
   - Limited type availability

**SDK Implementation**:
- Default to `/app` endpoint for all operations
- Support custom endpoint selection via method parameters
- URL construction: `{baseURL}{endpoint}`

## GraphQL Protocol

### Request Format

All GraphQL requests use HTTP POST with JSON body:

```json
{
  "query": "query getPlan($id: ID!) { getPlan(id: $id) { id name } }",
  "variables": {
    "id": "plan-123"
  }
}
```

**Headers**:
- `Content-Type: application/json`
- `Authorization: {access_token}`
- `User-Agent: go-api-sdk-jamfprotect/{version}`

### Response Format

GraphQL responses follow the standard format:

```json
{
  "data": {
    "getPlan": {
      "id": "plan-123",
      "name": "Security Plan"
    }
  },
  "errors": [
    {
      "message": "Error description",
      "locations": [{"line": 1, "column": 15}],
      "path": ["getPlan"],
      "extensions": {
        "code": "NOT_FOUND"
      }
    }
  ]
}
```

**Error Handling**:
- Errors array can contain multiple errors
- Check for "not found" indicators in error messages
- Parse error paths for field-level error context
- Extract extensions for error codes and additional context

### Query Structure

The API uses GraphQL fragments for code reusability:

```graphql
fragment PlanFields on Plan {
  id
  name
  description
  created
  updated
  logLevel
  autoUpdate
  # ... more fields
}

query getPlan($id: ID!) {
  getPlan(id: $id) {
    ...PlanFields
  }
}
```

## Pagination

### Cursor-Based Pagination

The API uses cursor-based pagination for list operations:

```graphql
query listPlans($nextToken: String, $direction: OrderDirection!, $field: PlanOrderField!) {
  listPlans(input: {
    next: $nextToken,
    order: {direction: $direction, field: $field},
    pageSize: 100
  }) {
    items {
      ...PlanFields
    }
    pageInfo {
      next
      total
    }
  }
}
```

**Pagination Parameters**:
- `next`: Cursor token for the next page (optional for first page)
- `pageSize`: Number of items per page (max 100)
- `order.direction`: Sort direction (`ASC` or `DESC`)
- `order.field`: Field to sort by (e.g., `CREATED`, `UPDATED`)

**SDK Implementation**:
- Automatically iterate through all pages
- Handle `pageInfo.next` cursor tokens
- Return consolidated results from all pages
- Expose pagination metadata via `pageInfo.total`

## Data Types

### Common Patterns

**References**: Related entities are represented as references with ID and name:

```json
{
  "actionConfigs": {
    "id": "config-123",
    "name": "Default Actions"
  }
}
```

**Lists**: Collections use arrays with optional metadata:

```json
{
  "exceptionSets": [
    {
      "uuid": "uuid-123",
      "name": "Standard Exceptions",
      "managed": true
    }
  ]
}
```

**Nested Objects**: Complex configuration uses nested structures:

```json
{
  "commsConfig": {
    "fqdn": "protect.example.com",
    "protocol": "HTTPS"
  }
}
```

### Null Handling

GraphQL distinguishes between:
- **Omitted field**: Field not included in request
- **Null value**: Field explicitly set to `null`
- **Empty value**: Field set to empty string/array

**SDK Implementation**:
- Use pointers for optional fields (`*string`, `*int64`)
- Provide separate `Null` flags for explicit null setting
- Example: `TelemetryV2: *string` and `TelemetryV2Null: bool`

## Mutations

### Input Types

Mutations accept input types with validation:

```graphql
mutation createPlan($name: String!, $description: String!, ...) {
  createPlan(input: {
    name: $name
    description: $description
    # ...
  }) {
    ...PlanFields
  }
}
```

**Required Fields**: Marked with `!` in schema  
**Optional Fields**: Can be omitted from input

**SDK Implementation**:
- Validate required fields before API calls
- Return descriptive validation errors
- Map Go structs to GraphQL input types
- Handle optional field presence correctly

## Rate Limiting

**Current Status**: No explicit rate limits documented  
**Best Practices**:
- Implement exponential backoff for retries
- Monitor for 429 status codes
- Consider implementing client-side rate limiting

## Error Codes

### HTTP Status Codes

- `200 OK`: Successful request (check GraphQL errors in response)
- `401 Unauthorized`: Authentication failure
- `429 Too Many Requests`: Rate limit exceeded
- `500 Internal Server Error`: Server-side error

### GraphQL Error Types

Common error patterns in `extensions.code`:
- `NOT_FOUND`: Resource does not exist
- `VALIDATION_ERROR`: Input validation failure
- `UNAUTHORIZED`: Insufficient permissions
- `INTERNAL_ERROR`: Server-side processing error

## SDK Design Implications

### Transport Layer

- **Single HTTP Client**: Reuse connections for performance
- **Timeout Configuration**: Default 60 seconds, configurable
- **Automatic Retry**: Not implemented (avoid in GraphQL mutations)
- **Logging**: Optional structured logging with credential redaction

### Service Layer

- **Type Safety**: Strong typing for all API resources
- **Validation**: Client-side validation before API calls
- **Error Wrapping**: Wrap errors with context
- **Context Support**: All operations accept `context.Context`

### Client Options

Support for customization via functional options:
- Base URL override
- Custom HTTP client
- Logger configuration
- Timeout settings
- User agent customization
- Debug mode

## Security Considerations

1. **Credential Management**:
   - Never log client secrets
   - Redact access tokens in logs
   - Store credentials securely (use environment variables)

2. **Token Security**:
   - Store tokens in memory only
   - Clear tokens on client disposal
   - Don't persist tokens to disk

3. **TLS/HTTPS**:
   - Always use HTTPS in production
   - Verify server certificates
   - Support custom CA certificates if needed

## Testing Strategy

1. **Unit Tests**:
   - Mock GraphQL responses
   - Test error handling
   - Validate input transformation

2. **Integration Tests**:
   - Use test credentials
   - Clean up test resources
   - Test pagination logic

3. **Example Programs**:
   - Demonstrate each operation
   - Show best practices
   - Provide working code samples

## Future Considerations

1. **Batch Operations**: Support for batching multiple mutations
2. **Subscriptions**: WebSocket support for real-time updates (if API adds support)
3. **Schema Introspection**: Dynamic schema discovery
4. **Code Generation**: Generate types from GraphQL schema
5. **Additional Services**: Analytics, Action Configs, Exception Sets, etc.

## References

- [Jamf Protect API Documentation](https://learn.jamf.com/bundle/jamf-protect-documentation/page/API_Documentation.html)
- [GraphQL Specification](https://spec.graphql.org/)
- [OAuth 2.0 Client Credentials](https://oauth.net/2/grant-types/client-credentials/)
- [RFC 7519 - JSON Web Tokens](https://tools.ietf.org/html/rfc7519)
