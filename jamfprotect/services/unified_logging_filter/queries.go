package unifiedloggingfilter

// GraphQL fragments and queries for Unified Logging Filters

const unifiedLoggingFilterFields = `
fragment UnifiedLoggingFilterFields on UnifiedLoggingFilter {
	uuid
	name
	description
	created
	updated
	filter
	tags
	enabled
}
`

const createUnifiedLoggingFilterMutation = `
mutation createUnifiedLoggingFilter(
	$name: String!,
	$description: String,
	$tags: [String]!,
	$filter: String!,
	$enabled: Boolean
) {
	createUnifiedLoggingFilter(
		input: {name: $name, description: $description, tags: $tags, filter: $filter, enabled: $enabled}
	) {
		...UnifiedLoggingFilterFields
	}
}
` + unifiedLoggingFilterFields

const getUnifiedLoggingFilterQuery = `
query getUnifiedLoggingFilter($uuid: ID!) {
	getUnifiedLoggingFilter(uuid: $uuid) {
		...UnifiedLoggingFilterFields
	}
}
` + unifiedLoggingFilterFields

const updateUnifiedLoggingFilterMutation = `
mutation updateUnifiedLoggingFilter(
	$uuid: ID!,
	$name: String!,
	$description: String,
	$filter: String!,
	$tags: [String]!,
	$enabled: Boolean
) {
	updateUnifiedLoggingFilter(
		uuid: $uuid
		input: {name: $name, description: $description, filter: $filter, tags: $tags, enabled: $enabled}
	) {
		...UnifiedLoggingFilterFields
	}
}
` + unifiedLoggingFilterFields

const deleteUnifiedLoggingFilterMutation = `
mutation deleteUnifiedLoggingFilter($uuid: ID!) {
	deleteUnifiedLoggingFilter(uuid: $uuid) {
		uuid
	}
}
`

const listUnifiedLoggingFiltersQuery = `
query listUnifiedLoggingFilters($nextToken: String, $direction: OrderDirection!, $field: UnifiedLoggingFiltersOrderField!, $filter: UnifiedLoggingFiltersFilterInput!) {
	listUnifiedLoggingFilters(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100, filter: $filter}
	) {
		items {
			...UnifiedLoggingFilterFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + unifiedLoggingFilterFields
