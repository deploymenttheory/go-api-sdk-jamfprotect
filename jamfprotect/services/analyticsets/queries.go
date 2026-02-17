package analyticsets

// GraphQL fragments and queries for Analytic Sets

const analyticSetFields = `
fragment AnalyticSetFields on AnalyticSet {
	uuid
	name
	description
	analytics @skip(if: $excludeAnalytics) {
		uuid
		name
		jamf
	}
	plans @include(if: $RBAC_Plan) {
		id
		name
	}
	created
	updated
	managed
	types
}
`

const createAnalyticSetMutation = `
mutation createAnalyticSet(
	$name: String!,
	$description: String,
	$types: [ANALYTIC_SET_TYPE!],
	$analytics: [ID!]!,
	$RBAC_Plan: Boolean!,
	$excludeAnalytics: Boolean!
) {
	createAnalyticSet(input: {
		name: $name,
		description: $description,
		analytics: $analytics,
		types: $types
	}) {
		...AnalyticSetFields
	}
}
` + analyticSetFields

const getAnalyticSetQuery = `
query getAnalyticSet(
	$uuid: ID!,
	$RBAC_Plan: Boolean!,
	$excludeAnalytics: Boolean!
) {
	getAnalyticSet(uuid: $uuid) {
		...AnalyticSetFields
	}
}
` + analyticSetFields

const updateAnalyticSetMutation = `
mutation updateAnalyticSet(
	$uuid: ID!,
	$name: String!,
	$description: String,
	$types: [ANALYTIC_SET_TYPE!],
	$analytics: [ID!]!,
	$RBAC_Plan: Boolean!,
	$excludeAnalytics: Boolean!
) {
	updateAnalyticSet(uuid: $uuid, input: {
		name: $name,
		description: $description,
		analytics: $analytics,
		types: $types
	}) {
		...AnalyticSetFields
	}
}
` + analyticSetFields

const deleteAnalyticSetMutation = `
mutation deleteAnalyticSet($uuid: ID!) {
	deleteAnalyticSet(uuid: $uuid) {
		uuid
	}
}
`

const listAnalyticSetsQuery = `
query listAnalyticSets($nextToken: String, $direction: OrderDirection = DESC, $field: AnalyticSetOrderField = created, $RBAC_Plan: Boolean!, $excludeAnalytics: Boolean = false) {
	listAnalyticSets(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			uuid
			name
			description
			analytics @skip(if: $excludeAnalytics) {
				uuid
				name
				jamf
			}
			plans @include(if: $RBAC_Plan) {
				id
				name
			}
			created
			updated
			managed
			types
		}
		pageInfo {
			next
			total
		}
	}
}
`
