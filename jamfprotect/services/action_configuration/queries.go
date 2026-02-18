package actionconfiguration

// GraphQL fragments and queries for Action Configurations

const actionConfigFields = `
fragment ActionConfigsFields on ActionConfigs {
	id
	name
	description
	hash
	created
	updated
	alertConfig {
		data {
			binary { attrs related }
			clickEvent { attrs related }
			downloadEvent { attrs related }
			file { attrs related }
			fsEvent { attrs related }
			group { attrs related }
			procEvent { attrs related }
			process { attrs related }
			screenshotEvent { attrs related }
			usbEvent { attrs related }
			user { attrs related }
			gkEvent { attrs related }
			keylogRegisterEvent { attrs related }
			mrtEvent { attrs related }
		}
	}
	clients {
		id
		type
		supportedReports
		batchConfig {
			delimiter
			sizeIndex
			windowInSeconds
			sizeInBytes
		}
		params {
			... on JamfCloudClientParams { destinationFilter }
			... on HttpClientParams { headers { header value } method url }
			... on KafkaClientParams { host port topic clientCN serverCN }
			... on SyslogClientParams { host port scheme }
			... on LogFileClientParams { path permissions maxSizeMB ownership backups }
		}
	}
}
`

const createActionConfigMutation = `
mutation createActionConfigs(
	$name: String!,
	$description: String!,
	$alertConfig: ActionConfigsAlertConfigInput!,
	$clients: [ReportClientInput!]
) {
	createActionConfigs(input: {
		name: $name,
		description: $description,
		alertConfig: $alertConfig,
		clients: $clients
	}) {
		...ActionConfigsFields
	}
}
` + actionConfigFields

const getActionConfigQuery = `
query getActionConfigs($id: ID!) {
	getActionConfigs(id: $id) {
		...ActionConfigsFields
	}
}
` + actionConfigFields

const updateActionConfigMutation = `
mutation updateActionConfigs(
	$id: ID!,
	$name: String!,
	$description: String!,
	$alertConfig: ActionConfigsAlertConfigInput!,
	$clients: [ReportClientInput!]
) {
	updateActionConfigs(id: $id, input: {
		name: $name,
		description: $description,
		alertConfig: $alertConfig,
		clients: $clients
	}) {
		...ActionConfigsFields
	}
}
` + actionConfigFields

const deleteActionConfigMutation = `
mutation deleteActionConfigs($id: ID!) {
	deleteActionConfigs(id: $id) {
		id
	}
}
`

const listActionConfigsQuery = `
query listActionConfigs($nextToken: String, $direction: OrderDirection!, $field: ActionConfigsOrderField!) {
	listActionConfigs(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			id
			name
			description
			created
			updated
		}
		pageInfo {
			next
			total
		}
	}
}
`

const listActionConfigNamesQuery = `
query listActionConfigNames {
	listActionConfigNames: listActionConfigs {
		items {
			name
		}
	}
}
`
