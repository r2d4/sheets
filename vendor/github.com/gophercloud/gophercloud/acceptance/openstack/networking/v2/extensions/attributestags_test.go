// +build acceptance networking tags

package extensions

import (
	"sort"
	"testing"

	"github.com/gophercloud/gophercloud/acceptance/clients"
	networking "github.com/gophercloud/gophercloud/acceptance/openstack/networking/v2"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/attributestags"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/networks"
	th "github.com/gophercloud/gophercloud/testhelper"
)

func TestTags(t *testing.T) {
	client, err := clients.NewNetworkV2Client()
	th.AssertNoErr(t, err)

	// Create Network
	network, err := networking.CreateNetwork(t, client)
	th.AssertNoErr(t, err)
	defer networking.DeleteNetwork(t, client, network.ID)

	tagReplaceAllOpts := attributestags.ReplaceAllOpts{
		// docs say list of tags, but it's a set e.g no duplicates
		Tags: []string{"a", "b", "c"},
	}
	tags, err := attributestags.ReplaceAll(client, "networks", network.ID, tagReplaceAllOpts).Extract()
	th.AssertNoErr(t, err)
	sort.Strings(tags) // Ensure ordering, older OpenStack versions aren't sorted...
	th.AssertDeepEquals(t, []string{"a", "b", "c"}, tags)

	// Verify the tags are also set in the object Get response
	gnetwork, err := networks.Get(client, network.ID).Extract()
	th.AssertNoErr(t, err)
	tags = gnetwork.Tags
	sort.Strings(tags)
	th.AssertDeepEquals(t, []string{"a", "b", "c"}, tags)

	// Add a tag
	err = attributestags.Add(client, "networks", network.ID, "d").ExtractErr()
	th.AssertNoErr(t, err)

	// Delete a tag
	err = attributestags.Delete(client, "networks", network.ID, "a").ExtractErr()
	th.AssertNoErr(t, err)

	// Verify expected tags are set in the List response
	tags, err = attributestags.List(client, "networks", network.ID).Extract()
	th.AssertNoErr(t, err)
	sort.Strings(tags)
	th.AssertDeepEquals(t, []string{"b", "c", "d"}, tags)

	// Delete all tags
	err = attributestags.DeleteAll(client, "networks", network.ID).ExtractErr()
	th.AssertNoErr(t, err)
	tags, err = attributestags.List(client, "networks", network.ID).Extract()
	th.AssertNoErr(t, err)
	th.AssertEquals(t, 0, len(tags))
}
