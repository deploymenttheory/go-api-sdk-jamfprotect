package exceptionsets

// GraphQL fragments and queries for Exception Sets

const exceptionSetFields = `
fragment ExceptionSetFields on ExceptionSet {
	uuid
	name
	description
	exceptions @skip(if: $minimal) {
		type
		value
		appSigningInfo {
			appId
			teamId
		}
		ignoreActivity
		analyticTypes
		analytic @include(if: $RBAC_Analytic) {
			name
			uuid
		}
	}
	esExceptions @skip(if: $minimal) {
		type
		value
		appSigningInfo {
			appId
			teamId
		}
		ignoreActivity
		ignoreListType
		ignoreListSubType
		eventType
	}
	created
	updated
	managed
}
`

const createExceptionSetMutation = `
mutation createExceptionSet(
	$name: String!,
	$description: String,
	$exceptions: [ExceptionInput!]!,
	$esExceptions: [EsExceptionInput!]!,
	$minimal: Boolean!,
	$RBAC_Analytic: Boolean!
) {
	createExceptionSet(input: {
		name: $name,
		description: $description,
		exceptions: $exceptions,
		esExceptions: $esExceptions
	}) {
		...ExceptionSetFields
	}
}
` + exceptionSetFields

const getExceptionSetQuery = `
query getExceptionSet(
	$uuid: ID!,
	$minimal: Boolean!,
	$RBAC_Analytic: Boolean!
) {
	getExceptionSet(uuid: $uuid) {
		...ExceptionSetFields
	}
}
` + exceptionSetFields

const updateExceptionSetMutation = `
mutation updateExceptionSet(
	$uuid: ID!,
	$name: String!,
	$description: String,
	$exceptions: [ExceptionInput!]!,
	$esExceptions: [EsExceptionInput!]!,
	$minimal: Boolean!,
	$RBAC_Analytic: Boolean!
) {
	updateExceptionSet(uuid: $uuid, input: {
		name: $name,
		description: $description,
		exceptions: $exceptions,
		esExceptions: $esExceptions
	}) {
		...ExceptionSetFields
	}
}
` + exceptionSetFields

const deleteExceptionSetMutation = `
mutation deleteExceptionSet($uuid: ID!) {
	deleteExceptionSet(uuid: $uuid) {
		uuid
	}
}
`

const listExceptionSetsQuery = `
query listExceptionSets($nextToken: String, $direction: OrderDirection = DESC, $field: ExceptionSetOrderField = created) {
	listExceptionSets(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			uuid
			name
			managed
		}
		pageInfo {
			next
			total
		}
	}
}
`
