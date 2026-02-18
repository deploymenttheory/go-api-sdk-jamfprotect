package removablestoragecontrolset

// GraphQL fragments and queries for USB Control Sets

const usbControlSetFields = `
fragment USBControlSetFields on USBControlSet {
	id
	name
	description
	defaultMountAction
	defaultMessageAction
	rules {
		mountAction
		messageAction
		type
		... on VendorRule {
			vendors
			applyTo
		}
		... on SerialRule {
			serials
			applyTo
		}
		... on ProductRule {
			products {
				vendor
				product
			}
			applyTo
		}
	}
	plans {
		id
		name
	}
	created
	updated
}
`

const createUSBControlSetMutation = `
mutation createUSBControlSet(
	$name: String!,
	$description: String,
	$defaultMountAction: USBCONTROL_MOUNT_ACTION_TYPE_ENUM!,
	$defaultMessageAction: String,
	$rules: [USBControlRuleInput!]!
) {
	createUSBControlSet(input: {
		name: $name,
		description: $description,
		defaultMountAction: $defaultMountAction,
		defaultMessageAction: $defaultMessageAction,
		rules: $rules
	}) {
		...USBControlSetFields
	}
}
` + usbControlSetFields

const getUSBControlSetQuery = `
query getUSBControlSet($id: ID!) {
	getUSBControlSet(id: $id) {
		...USBControlSetFields
	}
}
` + usbControlSetFields

const updateUSBControlSetMutation = `
mutation updateUSBControlSet(
	$id: ID!,
	$name: String!,
	$description: String,
	$defaultMountAction: USBCONTROL_MOUNT_ACTION_TYPE_ENUM!,
	$defaultMessageAction: String,
	$rules: [USBControlRuleInput!]!
) {
	updateUSBControlSet(id: $id, input: {
		name: $name,
		description: $description,
		defaultMountAction: $defaultMountAction,
		defaultMessageAction: $defaultMessageAction,
		rules: $rules
	}) {
		...USBControlSetFields
	}
}
` + usbControlSetFields

const deleteUSBControlSetMutation = `
mutation deleteUSBControlSet($id: ID!) {
	deleteUSBControlSet(id: $id) {
		id
	}
}
`

const listUSBControlSetsQuery = `
query listUSBControlSets($nextToken: String, $direction: OrderDirection!, $field: USBControlOrderField!) {
	listUSBControlSets(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...USBControlSetFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + usbControlSetFields

const listUSBControlSetNamesQuery = `
query listUsbControlNames {
	listUsbControlNames: listUSBControlSets {
		items {
			name
		}
	}
}
`
