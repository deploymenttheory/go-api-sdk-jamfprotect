package plans

// GraphQL fragments and queries for Plans

const planFields = `
fragment PlanFields on Plan {
	id
	hash
	name
	description
	created
	updated
	logLevel
	autoUpdate
	commsConfig {
		fqdn
		protocol
	}
	infoSync {
		attrs
		insightsSyncInterval
	}
	signaturesFeedConfig {
		mode
	}
	actionConfigs {
		id
		name
	}
	exceptionSets {
		uuid
		name
		managed
	}
	usbControlSet {
		id
		name
	}
	telemetry {
		id
		name
	}
	telemetryV2 {
		id
		name
	}
	analyticSets {
		type
		analyticSet {
			uuid
			name
			managed
			analytics {
				uuid
				categories
			}
		}
	}
}
`

const createPlanMutation = `
mutation createPlan(
	$name: String!,
	$description: String!,
	$logLevel: LOG_LEVEL_ENUM,
	$actionConfigs: ID!,
	$exceptionSets: [ID!],
	$telemetry: ID,
	$telemetryV2: ID,
	$analyticSets: [PlanAnalyticSetInput!],
	$usbControlSet: ID,
	$commsConfig: CommsConfigInput!,
	$infoSync: InfoSyncInput!,
	$autoUpdate: Boolean!,
	$signaturesFeedConfig: SignaturesFeedConfigInput!
) {
	createPlan(input: {
		name: $name,
		description: $description,
		logLevel: $logLevel,
		actionConfigs: $actionConfigs,
		exceptionSets: $exceptionSets,
		telemetry: $telemetry,
		telemetryV2: $telemetryV2,
		analyticSets: $analyticSets,
		usbControlSet: $usbControlSet,
		commsConfig: $commsConfig,
		infoSync: $infoSync,
		autoUpdate: $autoUpdate,
		signaturesFeedConfig: $signaturesFeedConfig
	}) {
		...PlanFields
	}
}
` + planFields

const getPlanQuery = `
query getPlan($id: ID!) {
	getPlan(id: $id) {
		...PlanFields
	}
}
` + planFields

const updatePlanMutation = `
mutation updatePlan(
	$id: ID!,
	$name: String!,
	$description: String!,
	$logLevel: LOG_LEVEL_ENUM,
	$actionConfigs: ID!,
	$exceptionSets: [ID!],
	$telemetry: ID,
	$telemetryV2: ID,
	$analyticSets: [PlanAnalyticSetInput!],
	$usbControlSet: ID,
	$commsConfig: CommsConfigInput!,
	$infoSync: InfoSyncInput!,
	$autoUpdate: Boolean!,
	$signaturesFeedConfig: SignaturesFeedConfigInput!
) {
	updatePlan(id: $id, input: {
		name: $name,
		description: $description,
		logLevel: $logLevel,
		actionConfigs: $actionConfigs,
		exceptionSets: $exceptionSets,
		telemetry: $telemetry,
		telemetryV2: $telemetryV2,
		analyticSets: $analyticSets,
		usbControlSet: $usbControlSet,
		commsConfig: $commsConfig,
		infoSync: $infoSync,
		autoUpdate: $autoUpdate,
		signaturesFeedConfig: $signaturesFeedConfig
	}) {
		...PlanFields
	}
}
` + planFields

const deletePlanMutation = `
mutation deletePlan($id: ID!) {
	deletePlan(id: $id) {
		id
	}
}
`

const listPlansQuery = `
query listPlans($nextToken: String, $direction: OrderDirection!, $field: PlanOrderField!) {
	listPlans(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...PlanFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + planFields
