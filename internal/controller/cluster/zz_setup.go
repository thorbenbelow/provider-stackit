// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/v2/pkg/controller"

	applicationloadbalancer "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/alb/applicationloadbalancer"
	certificate "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/alb/certificate"
	wafconfiguration "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/alb/wafconfiguration"
	wafcustomrulegroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/alb/wafcustomrulegroup"
	wafmanagedruleset "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/alb/wafmanagedruleset"
	foldercustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/foldercustomrole"
	folderroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/folderroleassignment"
	organizationcustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/organizationcustomrole"
	organizationroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/organizationroleassignment"
	projectcustomrole "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/projectcustomrole"
	projectroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/projectroleassignment"
	serviceaccountroleassignment "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/authorization/serviceaccountroleassignment"
	customdomain "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/cdn/customdomain"
	distribution "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/cdn/distribution"
	recordset "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/dns/recordset"
	zone "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/dns/zone"
	instance "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/dremio/instance"
	user "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/dremio/user"
	instanceedgecloud "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/edgecloud/instance"
	kubeconfig "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/edgecloud/kubeconfig"
	token "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/edgecloud/token"
	instancegit "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/git/instance"
	affinitygroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/affinitygroup"
	image "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/image"
	keypair "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/keypair"
	network "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/network"
	networkarea "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/networkarea"
	networkarearegion "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/networkarearegion"
	networkarearoute "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/networkarearoute"
	networkinterface "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/networkinterface"
	publicip "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/publicip"
	publicipassociate "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/publicipassociate"
	routingtable "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/routingtable"
	routingtableroute "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/routingtableroute"
	securitygroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/securitygroup"
	securitygrouprule "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/securitygrouprule"
	server "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/server"
	serverbackupenable "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/serverbackupenable"
	serverbackupschedule "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/serverbackupschedule"
	servernetworkinterfaceattach "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/servernetworkinterfaceattach"
	serverserviceaccountattach "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/serverserviceaccountattach"
	serverupdateenable "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/serverupdateenable"
	serverupdateschedule "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/serverupdateschedule"
	servervolumeattach "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/servervolumeattach"
	volume "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/volume"
	volumeautomation "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/iaas/volumeautomation"
	runner "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/intake/runner"
	key "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/kms/key"
	keyring "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/kms/keyring"
	wrappingkey "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/kms/wrappingkey"
	loadbalancer "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/loadbalancer/loadbalancer"
	observabilitycredential "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/loadbalancer/observabilitycredential"
	credential "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/logme/credential"
	instancelogme "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/logme/instance"
	accesstoken "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/logs/accesstoken"
	instancelogs "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/logs/instance"
	credentialmariadb "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/mariadb/credential"
	instancemariadb "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/mariadb/instance"
	instancemodelexperiments "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/modelexperiments/instance"
	tokenmodelexperiments "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/modelexperiments/token"
	tokenmodelserving "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/modelserving/token"
	instancemongodbflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/mongodbflex/instance"
	usermongodbflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/mongodbflex/user"
	bucket "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/objectstorage/bucket"
	compliancelock "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/objectstorage/compliancelock"
	credentialobjectstorage "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/objectstorage/credential"
	credentialsgroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/objectstorage/credentialsgroup"
	defaultretention "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/objectstorage/defaultretention"
	alertgroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/observability/alertgroup"
	credentialobservability "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/observability/credential"
	instanceobservability "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/observability/instance"
	logalertgroup "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/observability/logalertgroup"
	scrapeconfig "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/observability/scrapeconfig"
	credentialopensearch "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/opensearch/credential"
	instanceopensearch "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/opensearch/instance"
	database "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/postgresflex/database"
	instancepostgresflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/postgresflex/instance"
	userpostgresflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/postgresflex/user"
	providerconfig "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/providerconfig"
	credentialrabbitmq "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/rabbitmq/credential"
	instancerabbitmq "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/rabbitmq/instance"
	credentialredis "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/redis/credential"
	instanceredis "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/redis/instance"
	folder "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/resourcemanager/folder"
	project "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/resourcemanager/project"
	organization "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/scf/organization"
	organizationmanager "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/scf/organizationmanager"
	instancesecretsmanager "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/secretsmanager/instance"
	instancerolebindingv1 "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/secretsmanager/instancerolebindingv1"
	secretgrouprolebindingv1 "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/secretsmanager/secretgrouprolebindingv1"
	usersecretsmanager "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/secretsmanager/user"
	federatedidentityprovider "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/federatedidentityprovider"
	keyserviceaccount "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/key"
	serviceaccount "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/serviceaccount/serviceaccount"
	exportpolicy "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sfs/exportpolicy"
	projectlock "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sfs/projectlock"
	resourcepool "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sfs/resourcepool"
	share "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sfs/share"
	cluster "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/ske/cluster"
	kubeconfigske "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/ske/kubeconfig"
	databasesqlserverflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sqlserverflex/database"
	instancesqlserverflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sqlserverflex/instance"
	usersqlserverflex "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/sqlserverflex/user"
	accesstokentelemetryrouter "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/telemetryrouter/accesstoken"
	destination "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/telemetryrouter/destination"
	instancetelemetryrouter "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/telemetryrouter/instance"
	telemetrylink "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/telemetryrouter/telemetrylink"
	credentialvalkey "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/valkey/credential"
	instancevalkey "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/valkey/instance"
	networkrange "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpc/networkrange"
	region "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpc/region"
	routingtablevpc "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpc/routingtable"
	routingtablestaticroute "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpc/routingtablestaticroute"
	vpc "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpc/vpc"
	connection "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpn/connection"
	gateway "github.com/thorbenbelow/provider-stackit/internal/controller/cluster/vpn/gateway"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationloadbalancer.Setup,
		certificate.Setup,
		wafconfiguration.Setup,
		wafcustomrulegroup.Setup,
		wafmanagedruleset.Setup,
		foldercustomrole.Setup,
		folderroleassignment.Setup,
		organizationcustomrole.Setup,
		organizationroleassignment.Setup,
		projectcustomrole.Setup,
		projectroleassignment.Setup,
		serviceaccountroleassignment.Setup,
		customdomain.Setup,
		distribution.Setup,
		recordset.Setup,
		zone.Setup,
		instance.Setup,
		user.Setup,
		instanceedgecloud.Setup,
		kubeconfig.Setup,
		token.Setup,
		instancegit.Setup,
		affinitygroup.Setup,
		image.Setup,
		keypair.Setup,
		network.Setup,
		networkarea.Setup,
		networkarearegion.Setup,
		networkarearoute.Setup,
		networkinterface.Setup,
		publicip.Setup,
		publicipassociate.Setup,
		routingtable.Setup,
		routingtableroute.Setup,
		securitygroup.Setup,
		securitygrouprule.Setup,
		server.Setup,
		serverbackupenable.Setup,
		serverbackupschedule.Setup,
		servernetworkinterfaceattach.Setup,
		serverserviceaccountattach.Setup,
		serverupdateenable.Setup,
		serverupdateschedule.Setup,
		servervolumeattach.Setup,
		volume.Setup,
		volumeautomation.Setup,
		runner.Setup,
		key.Setup,
		keyring.Setup,
		wrappingkey.Setup,
		loadbalancer.Setup,
		observabilitycredential.Setup,
		credential.Setup,
		instancelogme.Setup,
		accesstoken.Setup,
		instancelogs.Setup,
		credentialmariadb.Setup,
		instancemariadb.Setup,
		instancemodelexperiments.Setup,
		tokenmodelexperiments.Setup,
		tokenmodelserving.Setup,
		instancemongodbflex.Setup,
		usermongodbflex.Setup,
		bucket.Setup,
		compliancelock.Setup,
		credentialobjectstorage.Setup,
		credentialsgroup.Setup,
		defaultretention.Setup,
		alertgroup.Setup,
		credentialobservability.Setup,
		instanceobservability.Setup,
		logalertgroup.Setup,
		scrapeconfig.Setup,
		credentialopensearch.Setup,
		instanceopensearch.Setup,
		database.Setup,
		instancepostgresflex.Setup,
		userpostgresflex.Setup,
		providerconfig.Setup,
		credentialrabbitmq.Setup,
		instancerabbitmq.Setup,
		credentialredis.Setup,
		instanceredis.Setup,
		folder.Setup,
		project.Setup,
		organization.Setup,
		organizationmanager.Setup,
		instancesecretsmanager.Setup,
		instancerolebindingv1.Setup,
		secretgrouprolebindingv1.Setup,
		usersecretsmanager.Setup,
		federatedidentityprovider.Setup,
		keyserviceaccount.Setup,
		serviceaccount.Setup,
		exportpolicy.Setup,
		projectlock.Setup,
		resourcepool.Setup,
		share.Setup,
		cluster.Setup,
		kubeconfigske.Setup,
		databasesqlserverflex.Setup,
		instancesqlserverflex.Setup,
		usersqlserverflex.Setup,
		accesstokentelemetryrouter.Setup,
		destination.Setup,
		instancetelemetryrouter.Setup,
		telemetrylink.Setup,
		credentialvalkey.Setup,
		instancevalkey.Setup,
		networkrange.Setup,
		region.Setup,
		routingtablevpc.Setup,
		routingtablestaticroute.Setup,
		vpc.Setup,
		connection.Setup,
		gateway.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

// SetupGated creates all controllers with the supplied logger and adds them to
// the supplied manager gated.
func SetupGated(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		applicationloadbalancer.SetupGated,
		certificate.SetupGated,
		wafconfiguration.SetupGated,
		wafcustomrulegroup.SetupGated,
		wafmanagedruleset.SetupGated,
		foldercustomrole.SetupGated,
		folderroleassignment.SetupGated,
		organizationcustomrole.SetupGated,
		organizationroleassignment.SetupGated,
		projectcustomrole.SetupGated,
		projectroleassignment.SetupGated,
		serviceaccountroleassignment.SetupGated,
		customdomain.SetupGated,
		distribution.SetupGated,
		recordset.SetupGated,
		zone.SetupGated,
		instance.SetupGated,
		user.SetupGated,
		instanceedgecloud.SetupGated,
		kubeconfig.SetupGated,
		token.SetupGated,
		instancegit.SetupGated,
		affinitygroup.SetupGated,
		image.SetupGated,
		keypair.SetupGated,
		network.SetupGated,
		networkarea.SetupGated,
		networkarearegion.SetupGated,
		networkarearoute.SetupGated,
		networkinterface.SetupGated,
		publicip.SetupGated,
		publicipassociate.SetupGated,
		routingtable.SetupGated,
		routingtableroute.SetupGated,
		securitygroup.SetupGated,
		securitygrouprule.SetupGated,
		server.SetupGated,
		serverbackupenable.SetupGated,
		serverbackupschedule.SetupGated,
		servernetworkinterfaceattach.SetupGated,
		serverserviceaccountattach.SetupGated,
		serverupdateenable.SetupGated,
		serverupdateschedule.SetupGated,
		servervolumeattach.SetupGated,
		volume.SetupGated,
		volumeautomation.SetupGated,
		runner.SetupGated,
		key.SetupGated,
		keyring.SetupGated,
		wrappingkey.SetupGated,
		loadbalancer.SetupGated,
		observabilitycredential.SetupGated,
		credential.SetupGated,
		instancelogme.SetupGated,
		accesstoken.SetupGated,
		instancelogs.SetupGated,
		credentialmariadb.SetupGated,
		instancemariadb.SetupGated,
		instancemodelexperiments.SetupGated,
		tokenmodelexperiments.SetupGated,
		tokenmodelserving.SetupGated,
		instancemongodbflex.SetupGated,
		usermongodbflex.SetupGated,
		bucket.SetupGated,
		compliancelock.SetupGated,
		credentialobjectstorage.SetupGated,
		credentialsgroup.SetupGated,
		defaultretention.SetupGated,
		alertgroup.SetupGated,
		credentialobservability.SetupGated,
		instanceobservability.SetupGated,
		logalertgroup.SetupGated,
		scrapeconfig.SetupGated,
		credentialopensearch.SetupGated,
		instanceopensearch.SetupGated,
		database.SetupGated,
		instancepostgresflex.SetupGated,
		userpostgresflex.SetupGated,
		providerconfig.SetupGated,
		credentialrabbitmq.SetupGated,
		instancerabbitmq.SetupGated,
		credentialredis.SetupGated,
		instanceredis.SetupGated,
		folder.SetupGated,
		project.SetupGated,
		organization.SetupGated,
		organizationmanager.SetupGated,
		instancesecretsmanager.SetupGated,
		instancerolebindingv1.SetupGated,
		secretgrouprolebindingv1.SetupGated,
		usersecretsmanager.SetupGated,
		federatedidentityprovider.SetupGated,
		keyserviceaccount.SetupGated,
		serviceaccount.SetupGated,
		exportpolicy.SetupGated,
		projectlock.SetupGated,
		resourcepool.SetupGated,
		share.SetupGated,
		cluster.SetupGated,
		kubeconfigske.SetupGated,
		databasesqlserverflex.SetupGated,
		instancesqlserverflex.SetupGated,
		usersqlserverflex.SetupGated,
		accesstokentelemetryrouter.SetupGated,
		destination.SetupGated,
		instancetelemetryrouter.SetupGated,
		telemetrylink.SetupGated,
		credentialvalkey.SetupGated,
		instancevalkey.SetupGated,
		networkrange.SetupGated,
		region.SetupGated,
		routingtablevpc.SetupGated,
		routingtablestaticroute.SetupGated,
		vpc.SetupGated,
		connection.SetupGated,
		gateway.SetupGated,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
