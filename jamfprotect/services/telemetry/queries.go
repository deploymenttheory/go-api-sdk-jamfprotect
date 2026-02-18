package telemetry

// GraphQL fragments and queries for Telemetry V2

const telemetryV2Fields = `
fragment TelemetryV2Fields on TelemetryV2 {
	id
	name
	description
	created
	updated
	logFiles
	logFileCollection
	performanceMetrics
	plans @include(if: $RBAC_Plan) {
		id
		name
	}
	events
	fileHashing
}
`

const createTelemetryV2Mutation = `
mutation createTelemetryV2(
	$name: String!,
	$description: String,
	$logFiles: [String!]!,
	$logFileCollection: Boolean!,
	$performanceMetrics: Boolean!,
	$events: [ES_EVENTS_ENUM]!,
	$fileHashing: Boolean!,
	$RBAC_Plan: Boolean!
) {
	createTelemetryV2(
		input: {name: $name, description: $description, logFiles: $logFiles, logFileCollection: $logFileCollection, performanceMetrics: $performanceMetrics, events: $events, fileHashing: $fileHashing}
	) {
		...TelemetryV2Fields
	}
}
` + telemetryV2Fields

const getTelemetryV2Query = `
query getTelemetryV2($id: ID!, $RBAC_Plan: Boolean!) {
	getTelemetryV2(id: $id) {
		...TelemetryV2Fields
	}
}
` + telemetryV2Fields

const updateTelemetryV2Mutation = `
mutation updateTelemetryV2(
	$id: ID!,
	$name: String!,
	$description: String,
	$logFiles: [String!]!,
	$logFileCollection: Boolean!,
	$performanceMetrics: Boolean!,
	$events: [ES_EVENTS_ENUM]!,
	$fileHashing: Boolean!,
	$RBAC_Plan: Boolean!
) {
	updateTelemetryV2(
		id: $id
		input: {name: $name, description: $description, logFiles: $logFiles, logFileCollection: $logFileCollection, performanceMetrics: $performanceMetrics, events: $events, fileHashing: $fileHashing}
	) {
		...TelemetryV2Fields
	}
}
` + telemetryV2Fields

const deleteTelemetryV2Mutation = `
mutation deleteTelemetryV2($id: ID!) {
	deleteTelemetryV2(id: $id) {
		id
	}
}
`

const listTelemetriesV2Query = `
query listTelemetriesV2($nextToken: String, $direction: OrderDirection!, $field: TelemetryOrderField!, $RBAC_Plan: Boolean!) {
	listTelemetriesV2(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...TelemetryV2Fields
		}
		pageInfo {
			next
			total
		}
	}
}
` + telemetryV2Fields

const telemetryV1Fields = `
fragment TelemetryFields on Telemetry {
	id
	name
	description
	verbose
	level
	created
	updated
	plans @include(if: $RBAC_Plan) {
		id
		name
	}
	performanceMetrics
	logFiles
	logFileCollection
}
`

const listTelemetriesCombinedQuery = `
query listTelemetriesCombined(
	$field: TelemetryOrderField!
	$direction: OrderDirection!
	$RBAC_Plan: Boolean!
) {
	listTelemetries(
		input: { order: { direction: $direction, field: $field }, pageSize: 100 }
	) {
		items {
			...TelemetryFields
		}
		pageInfo {
			next
			total
		}
	}
	listTelemetriesV2(
		input: { order: { direction: $direction, field: $field }, pageSize: 100 }
	) {
		items {
			id
			name
			description
			plans @include(if: $RBAC_Plan) {
				id
				name
			}
			created
			updated
		}
		pageInfo {
			next
			total
		}
	}
}
` + telemetryV1Fields
