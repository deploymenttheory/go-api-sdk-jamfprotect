package custompreventlist

// GraphQL fragments and queries for Prevent Lists

const preventListFields = `
fragment PreventListFields on PreventList {
	id
	name
	description
	type
	tags
	list
	count
	created
}
`

const createPreventListMutation = `
mutation createPreventList(
	$name: String!,
	$tags: [String]!,
	$type: PREVENT_LIST_TYPE!,
	$list: [String]!,
	$description: String
) {
	createPreventList(input: {
		name: $name,
		tags: $tags,
		type: $type,
		list: $list,
		description: $description
	}) {
		...PreventListFields
	}
}
` + preventListFields

const getPreventListQuery = `
query getPreventList($id: ID!) {
	getPreventList(id: $id) {
		...PreventListFields
	}
}
` + preventListFields

const updatePreventListMutation = `
mutation updatePreventList(
	$id: ID!,
	$name: String!,
	$tags: [String]!,
	$type: PREVENT_LIST_TYPE!,
	$list: [String]!,
	$description: String
) {
	updatePreventList(id: $id, input: {
		name: $name,
		tags: $tags,
		type: $type,
		list: $list,
		description: $description
	}) {
		...PreventListFields
	}
}
` + preventListFields

const deletePreventListMutation = `
mutation deletePreventList($id: ID!) {
	deletePreventList(id: $id) {
		id
	}
}
`

const listPreventListsQuery = `
query listPreventLists($nextToken: String, $direction: OrderDirection!, $field: PreventListOrderField!) {
	listPreventLists(
		input: {next: $nextToken, order: {direction: $direction, field: $field}, pageSize: 100}
	) {
		items {
			...PreventListFields
		}
		pageInfo {
			next
			total
		}
	}
}
` + preventListFields
