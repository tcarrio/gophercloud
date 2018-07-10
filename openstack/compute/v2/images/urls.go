package images

import "github.com/gophercloud/gophercloud"

func listDetailURL(client *gophercloud.ServiceClient) string {
	return client.ServiceURL("images", "detail")
}

func getURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}

func getFromFiltersURL(client *gophercloud.ServiceClient, filters map[string]string) string {
	return client.ServiceURL("images" + client.Query(filters))
}

func deleteURL(client *gophercloud.ServiceClient, id string) string {
	return client.ServiceURL("images", id)
}
