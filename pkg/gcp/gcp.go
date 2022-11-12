package gcp

var Permissions = []string{
	"resourcemanager.folders.create",
	"serviceusage.services.list",
	"billing.resourceAssociations.create",
	"serviceusage.operations.cancel",
	"serviceusage.operations.delete",
	"serviceusage.operations.get",
	"serviceusage.operations.list",
	"serviceusage.quotas.get",
	"serviceusage.quotas.update",
	"serviceusage.services.disable",
	"serviceusage.services.enable",
	"serviceusage.services.get",
	"serviceusage.services.list",
	"serviceusage.services.use",
	"monitoring.timeSeries.list",
	"resourcemanager.organizations.get",
	"resourcemanager.projects.create",
	"resourcemanager.projects.update",
	"compute.organizations.enableXpnHost",
	"compute.organizations.enableXpnResource",
	"compute.organizations.disableXpnResource",
	"compute.organizations.disableXpnHost",
}