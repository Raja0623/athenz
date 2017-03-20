//
// This file generated by rdl 1.4.11
//

package zts

import (
	rdl "github.com/ardielle/ardielle-go/rdl"
)

var schema *rdl.Schema

func init() {
	sb := rdl.NewSchemaBuilder("ZTS")
	sb.Version(1)
	sb.Namespace("com.yahoo.athenz.zts")
	sb.Comment("Copyright 2016 Yahoo Inc. Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms. The Authorization Management Service (ZTS) API")

	tSimpleName := rdl.NewStringTypeBuilder("SimpleName")
	tSimpleName.Comment("Copyright 2016 Yahoo Inc. Licensed under the terms of the Apache version 2.0 license. See LICENSE file for terms. Common name types used by several API definitions A simple identifier, an element of compound name.")
	tSimpleName.Pattern("[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tSimpleName.Build())

	tCompoundName := rdl.NewStringTypeBuilder("CompoundName")
	tCompoundName.Comment("A compound name. Most names in this API are compound names.")
	tCompoundName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tCompoundName.Build())

	tDomainName := rdl.NewStringTypeBuilder("DomainName")
	tDomainName.Comment("A domain name is the general qualifier prefix, as its uniqueness is managed.")
	tDomainName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tDomainName.Build())

	tEntityName := rdl.NewStringTypeBuilder("EntityName")
	tEntityName.Comment("An entity name is a short form of a resource name, including only the domain and entity.")
	tEntityName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tEntityName.Build())

	tServiceName := rdl.NewStringTypeBuilder("ServiceName")
	tServiceName.Comment("A service name will generally be a unique subdomain.")
	tServiceName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tServiceName.Build())

	tActionName := rdl.NewStringTypeBuilder("ActionName")
	tActionName.Comment("An action (operation) name.")
	tActionName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tActionName.Build())

	tResourceName := rdl.NewStringTypeBuilder("ResourceName")
	tResourceName.Comment("A resource name Note that the EntityName part is optional, that is, a domain name followed by a colon is valid resource name.")
	tResourceName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*(:([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*)?")
	sb.AddType(tResourceName.Build())

	tYBase64 := rdl.NewStringTypeBuilder("YBase64")
	tYBase64.Comment("The Y-specific URL-safe Base64 variant.")
	tYBase64.Pattern("[a-zA-Z0-9\\._-]+")
	sb.AddType(tYBase64.Build())

	tYEncoded := rdl.NewStringTypeBuilder("YEncoded")
	tYEncoded.Comment("YEncoded includes ybase64 chars, as well as = and %. This can represent a user cookie and URL-encoded values.")
	tYEncoded.Pattern("[a-zA-Z0-9\\._%=-]*")
	sb.AddType(tYEncoded.Build())

	tAuthorityName := rdl.NewStringTypeBuilder("AuthorityName")
	tAuthorityName.Comment("Used as the prefix in a signed assertion. This uniquely identifies a signing authority.")
	tAuthorityName.Pattern("([a-zA-Z0-9_][a-zA-Z0-9_-]*\\.)*[a-zA-Z0-9_][a-zA-Z0-9_-]*")
	sb.AddType(tAuthorityName.Build())

	tSignedToken := rdl.NewStringTypeBuilder("SignedToken")
	tSignedToken.Comment("A signed assertion if identity. i.e. the user cookie value. This token will only make sense to the authority that generated it, so it is beneficial to have something in the value that is cheaply recognized to quickly reject if it belongs to another authority. In addition to the YEncoded set our token includes ; to separate components and , to separate roles")
	tSignedToken.Pattern("[a-zA-Z0-9\\._%=;,-]*")
	sb.AddType(tSignedToken.Build())

	tResourceAccess := rdl.NewStructTypeBuilder("Struct", "ResourceAccess")
	tResourceAccess.Comment("ResourceAccess can be checked and returned as this resource. (same as ZMS.Access)")
	tResourceAccess.Field("granted", "Bool", false, nil, "true (allowed) or false (denied)")
	sb.AddType(tResourceAccess.Build())

	tPublicKeyEntry := rdl.NewStructTypeBuilder("Struct", "PublicKeyEntry")
	tPublicKeyEntry.Comment("The representation of the public key in a service identity object.")
	tPublicKeyEntry.Field("key", "String", false, nil, "the public key for the service")
	tPublicKeyEntry.Field("id", "String", false, nil, "the key identifier (version or zone name)")
	sb.AddType(tPublicKeyEntry.Build())

	tServiceIdentity := rdl.NewStructTypeBuilder("Struct", "ServiceIdentity")
	tServiceIdentity.Comment("The representation of the service identity object.")
	tServiceIdentity.Field("name", "ServiceName", false, nil, "the full name of the service, i.e. \"sports.storage\"")
	tServiceIdentity.ArrayField("publicKeys", "PublicKeyEntry", true, "array of public keys for key rotation")
	tServiceIdentity.Field("providerEndpoint", "String", true, nil, "if present, then this service can provision tenants via this endpoint.")
	tServiceIdentity.Field("modified", "Timestamp", true, nil, "the timestamp when this entry was last modified")
	tServiceIdentity.Field("executable", "String", true, nil, "the path of the executable that runs the service")
	tServiceIdentity.ArrayField("hosts", "String", true, "list of host names that this service can run on")
	tServiceIdentity.Field("user", "String", true, nil, "local (unix) user name this service can run as")
	tServiceIdentity.Field("group", "String", true, nil, "local (unix) group name this service can run as")
	sb.AddType(tServiceIdentity.Build())

	tServiceIdentityList := rdl.NewStructTypeBuilder("Struct", "ServiceIdentityList")
	tServiceIdentityList.Comment("The representation for an enumeration of services in the namespace.")
	tServiceIdentityList.ArrayField("names", "EntityName", false, "list of service names")
	sb.AddType(tServiceIdentityList.Build())

	tHostServices := rdl.NewStructTypeBuilder("Struct", "HostServices")
	tHostServices.Comment("The representation for an enumeration of services authorized to run on a specific host.")
	tHostServices.Field("host", "String", false, nil, "name of the host")
	tHostServices.ArrayField("names", "EntityName", false, "list of service names authorized to run on this host")
	sb.AddType(tHostServices.Build())

	tAssertionEffect := rdl.NewEnumTypeBuilder("Enum", "AssertionEffect")
	tAssertionEffect.Comment("Every assertion can have the effect of ALLOW or DENY.")
	tAssertionEffect.Element("ALLOW", "")
	tAssertionEffect.Element("DENY", "")
	sb.AddType(tAssertionEffect.Build())

	tAssertion := rdl.NewStructTypeBuilder("Struct", "Assertion")
	tAssertion.Comment("A representation for the encapsulation of an action to be performed on a resource by a principal.")
	tAssertion.Field("role", "String", false, nil, "the subject of the assertion, a role")
	tAssertion.Field("resource", "String", false, nil, "the object of the assertion. Must be in the local namespace. Can contain wildcards")
	tAssertion.Field("action", "String", false, nil, "the predicate of the assertion. Can contain wildcards")
	tAssertion.Field("effect", "AssertionEffect", true, ALLOW, "the effect of the assertion in the policy language")
	tAssertion.Field("id", "Int64", true, nil, "assertion id - auto generated by server")
	sb.AddType(tAssertion.Build())

	tPolicy := rdl.NewStructTypeBuilder("Struct", "Policy")
	tPolicy.Comment("The representation for a Policy with set of assertions.")
	tPolicy.Field("name", "ResourceName", false, nil, "name of the policy")
	tPolicy.Field("modified", "Timestamp", true, nil, "last modification timestamp of this policy")
	tPolicy.ArrayField("assertions", "Assertion", false, "list of defined assertions for this policy")
	sb.AddType(tPolicy.Build())

	tPolicyData := rdl.NewStructTypeBuilder("Struct", "PolicyData")
	tPolicyData.Field("domain", "DomainName", false, nil, "name of the domain")
	tPolicyData.ArrayField("policies", "Policy", false, "list of policies defined in this server")
	sb.AddType(tPolicyData.Build())

	tSignedPolicyData := rdl.NewStructTypeBuilder("Struct", "SignedPolicyData")
	tSignedPolicyData.Comment("A representation of policies object defined in a given server.")
	tSignedPolicyData.Field("policyData", "PolicyData", false, nil, "list of policies defined in a domain")
	tSignedPolicyData.Field("zmsSignature", "String", false, nil, "zms signature generated based on the domain policies object")
	tSignedPolicyData.Field("zmsKeyId", "String", false, nil, "the identifier of the zms key used to generate the signature")
	tSignedPolicyData.Field("modified", "Timestamp", false, nil, "when the domain itself was last modified")
	tSignedPolicyData.Field("expires", "Timestamp", false, nil, "timestamp specifying the expiration time for using this set of policies")
	sb.AddType(tSignedPolicyData.Build())

	tDomainSignedPolicyData := rdl.NewStructTypeBuilder("Struct", "DomainSignedPolicyData")
	tDomainSignedPolicyData.Comment("A signed bulk transfer of policies. The data is signed with server's private key.")
	tDomainSignedPolicyData.Field("signedPolicyData", "SignedPolicyData", false, nil, "policy data signed by ZMS")
	tDomainSignedPolicyData.Field("signature", "String", false, nil, "signature generated based on the domain policies object")
	tDomainSignedPolicyData.Field("keyId", "String", false, nil, "the identifier of the key used to generate the signature")
	sb.AddType(tDomainSignedPolicyData.Build())

	tRoleToken := rdl.NewStructTypeBuilder("Struct", "RoleToken")
	tRoleToken.Comment("A representation of a signed RoleToken")
	tRoleToken.Field("token", "String", false, nil, "")
	tRoleToken.Field("expiryTime", "Int64", false, nil, "")
	sb.AddType(tRoleToken.Build())

	tRoleCertificateRequest := rdl.NewStructTypeBuilder("Struct", "RoleCertificateRequest")
	tRoleCertificateRequest.Comment("RoleCertificateRequest - a certificate signing request")
	tRoleCertificateRequest.Field("csr", "String", false, nil, "")
	tRoleCertificateRequest.Field("expiryTime", "Int64", false, nil, "")
	sb.AddType(tRoleCertificateRequest.Build())

	tAccess := rdl.NewStructTypeBuilder("Struct", "Access")
	tAccess.Comment("Access can be checked and returned as this resource.")
	tAccess.Field("granted", "Bool", false, nil, "true (allowed) or false (denied)")
	sb.AddType(tAccess.Build())

	tRoleAccess := rdl.NewStructTypeBuilder("Struct", "RoleAccess")
	tRoleAccess.ArrayField("roles", "EntityName", false, "")
	sb.AddType(tRoleAccess.Build())

	tTenantDomains := rdl.NewStructTypeBuilder("Struct", "TenantDomains")
	tTenantDomains.ArrayField("tenantDomainNames", "DomainName", false, "")
	sb.AddType(tTenantDomains.Build())

	tIdentity := rdl.NewStructTypeBuilder("Struct", "Identity")
	tIdentity.Comment("Identity - a signed assertion of service or human identity, the response could be either a client certificate or just a regular NToken (depending if the request contained a csr or not).")
	tIdentity.Field("name", "CompoundName", false, nil, "name of the identity, fully qualified, i.e. my.domain.service1, or aws.1232321321312.myusername")
	tIdentity.Field("certificate", "String", true, nil, "a certificate usable for both client and server in TLS connections")
	tIdentity.Field("caCertBundle", "String", true, nil, "the CA certificate chain to use with all IMS-generated certs")
	tIdentity.Field("sshCertificate", "String", true, nil, "the SSH certificate, signed by the CA (user or host)")
	tIdentity.Field("sshCertificateSigner", "String", true, nil, "the SSH CA's public key for the sshCertificate (user or host)")
	tIdentity.Field("serviceToken", "SignedToken", true, nil, "service token instead of TLS certificate")
	tIdentity.MapField("attributes", "String", "String", true, "other config-like attributes determined at boot time")
	sb.AddType(tIdentity.Build())

	tInstanceInformation := rdl.NewStructTypeBuilder("Struct", "InstanceInformation")
	tInstanceInformation.Comment("Instance object that includes requested service details plus host document that is signed by provider as part of the host bootstrap process")
	tInstanceInformation.Field("document", "String", false, nil, "signed document containing attributes like IP address, instance-id, account#, etc.")
	tInstanceInformation.Field("signature", "String", false, nil, "the signature for the document")
	tInstanceInformation.Field("keyId", "String", false, nil, "the keyid used to sign the document")
	tInstanceInformation.Field("domain", "CompoundName", false, nil, "the domain of the instance")
	tInstanceInformation.Field("service", "SimpleName", false, nil, "the service this instance is supposed to run")
	tInstanceInformation.Field("csr", "String", false, nil, "return a certificate in the response")
	tInstanceInformation.Field("ssh", "String", true, nil, "if present, return an SSH host certificate")
	sb.AddType(tInstanceInformation.Build())

	tInstanceRefreshRequest := rdl.NewStructTypeBuilder("Struct", "InstanceRefreshRequest")
	tInstanceRefreshRequest.Comment("InstanceRefreshRequest - a certificate refresh request")
	tInstanceRefreshRequest.Field("csr", "String", false, nil, "Cert CSR if requesting TLS certificate")
	tInstanceRefreshRequest.Field("expiryTime", "Int32", true, nil, "in seconds how long token should be valid for")
	sb.AddType(tInstanceRefreshRequest.Build())

	tAWSInstanceInformation := rdl.NewStructTypeBuilder("Struct", "AWSInstanceInformation")
	tAWSInstanceInformation.Comment("AWSInstanceInformation - the information a booting EC2 instance must provide to ZTS to authenticate.")
	tAWSInstanceInformation.Field("document", "String", false, nil, "signed document containing attributes like IP address, instance-id, account#, etc.")
	tAWSInstanceInformation.Field("signature", "String", false, nil, "the signature for the document")
	tAWSInstanceInformation.Field("domain", "CompoundName", false, nil, "the domain of the instance")
	tAWSInstanceInformation.Field("service", "SimpleName", false, nil, "the service this instance is supposed to run")
	tAWSInstanceInformation.Field("csr", "String", false, nil, "return a certificate in the response")
	tAWSInstanceInformation.Field("ssh", "String", true, nil, "if present, return an SSH host certificate. Format is JSON.")
	tAWSInstanceInformation.Field("name", "CompoundName", false, nil, "the full service identity name (same as the EC2 instance profile name)")
	tAWSInstanceInformation.Field("account", "SimpleName", false, nil, "the account id (as a string) for the instance. parsed from the instance profile ARN")
	tAWSInstanceInformation.Field("cloud", "SimpleName", true, nil, "the name of the cloud (namespace) within the account, parsed from the name")
	tAWSInstanceInformation.Field("subnet", "SimpleName", true, nil, "not used")
	tAWSInstanceInformation.Field("access", "String", false, nil, "the AWS Access Key Id for the role")
	tAWSInstanceInformation.Field("secret", "String", false, nil, "the AWS Secret Access Key for the role")
	tAWSInstanceInformation.Field("token", "String", false, nil, "the AWS STS Token for the role")
	tAWSInstanceInformation.Field("expires", "Timestamp", false, nil, "the expiration time of the access keys")
	tAWSInstanceInformation.Field("modified", "Timestamp", false, nil, "the modified time of the access keys")
	tAWSInstanceInformation.Field("flavor", "String", false, nil, "the 'flavor' of the access keys, i.e. \"AWS-HMAC\"")
	sb.AddType(tAWSInstanceInformation.Build())

	tAWSCertificateRequest := rdl.NewStructTypeBuilder("Struct", "AWSCertificateRequest")
	tAWSCertificateRequest.Comment("AWSCertificateRequest - a certificate signing request")
	tAWSCertificateRequest.Field("csr", "String", true, nil, "request an X.509 certificate")
	tAWSCertificateRequest.Field("ssh", "String", true, nil, "request an SSH certificate")
	sb.AddType(tAWSCertificateRequest.Build())

	tAWSTemporaryCredentials := rdl.NewStructTypeBuilder("Struct", "AWSTemporaryCredentials")
	tAWSTemporaryCredentials.Field("accessKeyId", "String", false, nil, "")
	tAWSTemporaryCredentials.Field("secretAccessKey", "String", false, nil, "")
	tAWSTemporaryCredentials.Field("sessionToken", "String", false, nil, "")
	tAWSTemporaryCredentials.Field("expiration", "Timestamp", false, nil, "")
	sb.AddType(tAWSTemporaryCredentials.Build())

	tOSTKInstanceInformation := rdl.NewStructTypeBuilder("Struct", "OSTKInstanceInformation")
	tOSTKInstanceInformation.Comment("Instance object that includes requested service details plus host document that is signed by Openstack as part of the host bootstrap process")
	tOSTKInstanceInformation.Field("document", "String", false, nil, "signed document containing attributes like IP address, instance-id, account#, etc.")
	tOSTKInstanceInformation.Field("signature", "String", false, nil, "the signature for the document")
	tOSTKInstanceInformation.Field("keyId", "String", false, nil, "the keyid used to sign the document")
	tOSTKInstanceInformation.Field("domain", "CompoundName", false, nil, "the domain of the instance")
	tOSTKInstanceInformation.Field("service", "SimpleName", false, nil, "the service this instance is supposed to run")
	tOSTKInstanceInformation.Field("csr", "String", false, nil, "return a certificate in the response")
	sb.AddType(tOSTKInstanceInformation.Build())

	tOSTKInstanceRefreshRequest := rdl.NewStructTypeBuilder("Struct", "OSTKInstanceRefreshRequest")
	tOSTKInstanceRefreshRequest.Comment("OSTKCertificateRequest - a certificate signing request")
	tOSTKInstanceRefreshRequest.Field("csr", "String", true, nil, "request an X.509 certificate")
	sb.AddType(tOSTKInstanceRefreshRequest.Build())

	tDomainMetricType := rdl.NewEnumTypeBuilder("Enum", "DomainMetricType")
	tDomainMetricType.Comment("zpe metric attributes")
	tDomainMetricType.Element("ACCESS_ALLOWED", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DENY", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DENY_NO_MATCH", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_ALLOW", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_ERROR", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_TOKEN_INVALID", "")
	tDomainMetricType.Element("ACCESS_Allowed_TOKEN_EXPIRED", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DOMAIN_NOT_FOUND", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DOMAIN_MISMATCH", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DOMAIN_EXPIRED", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_DOMAIN_EMPTY", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_TOKEN_CACHE_FAILURE", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_TOKEN_CACHE_NOT_FOUND", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_TOKEN_CACHE_SUCCESS", "")
	tDomainMetricType.Element("ACCESS_ALLOWED_TOKEN_VALIDATE", "")
	tDomainMetricType.Element("LOAD_FILE_FAIL", "")
	tDomainMetricType.Element("LOAD_FILE_GOOD", "")
	tDomainMetricType.Element("LOAD_DOMAIN_GOOD", "")
	sb.AddType(tDomainMetricType.Build())

	tDomainMetric := rdl.NewStructTypeBuilder("Struct", "DomainMetric")
	tDomainMetric.Field("metricType", "DomainMetricType", false, nil, "")
	tDomainMetric.Field("metricVal", "Int32", false, nil, "")
	sb.AddType(tDomainMetric.Build())

	tDomainMetrics := rdl.NewStructTypeBuilder("Struct", "DomainMetrics")
	tDomainMetrics.Field("domainName", "DomainName", false, nil, "name of the domain the metrics pertain to")
	tDomainMetrics.ArrayField("metricList", "DomainMetric", false, "list of the domains metrics")
	sb.AddType(tDomainMetrics.Build())

	rGetResourceAccess := rdl.NewResourceBuilder("ResourceAccess", "GET", "/access/{action}/{resource}")
	rGetResourceAccess.Comment("Check access for the specified operation on the specified resource for the currently authenticated user. This is the slow centralized access for control-plane purposes. Use distributed mechanisms for decentralized (data-plane) access by fetching signed policies and role tokens for users. With this endpoint the resource is part of the uri and restricted to its strict definition of resource name. If needed, you can use the GetAccessExt api that allows resource name to be less restrictive.")
	rGetResourceAccess.Input("action", "ActionName", true, "", "", false, nil, "action as specified in the policy assertion, i.e. update or read")
	rGetResourceAccess.Input("resource", "ResourceName", true, "", "", false, nil, "the resource to check access against, i.e. \"media.news:articles\"")
	rGetResourceAccess.Input("domain", "DomainName", false, "domain", "", true, nil, "usually null. If present, it specifies an alternate domain for cross-domain trust relation")
	rGetResourceAccess.Input("checkPrincipal", "EntityName", false, "principal", "", true, nil, "usually null. If present, carry out the access check for this principal")
	rGetResourceAccess.Auth("", "", true, "")
	rGetResourceAccess.Exception("BAD_REQUEST", "ResourceError", "")
	rGetResourceAccess.Exception("FORBIDDEN", "ResourceError", "")
	rGetResourceAccess.Exception("NOT_FOUND", "ResourceError", "")
	rGetResourceAccess.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetResourceAccess.Build())

	rGetResourceAccessExt := rdl.NewResourceBuilder("ResourceAccess", "GET", "/access/{action}")
	rGetResourceAccessExt.Comment("Check access for the specified operation on the specified resource for the currently authenticated user. This is the slow centralized access for control-plane purposes.")
	rGetResourceAccessExt.Name("GetResourceAccessExt")
	rGetResourceAccessExt.Input("action", "ActionName", true, "", "", false, nil, "action as specified in the policy assertion, i.e. update or read")
	rGetResourceAccessExt.Input("resource", "String", false, "resource", "", false, nil, "the resource to check access against, i.e. \"media.news:articles\"")
	rGetResourceAccessExt.Input("domain", "DomainName", false, "domain", "", true, nil, "usually null. If present, it specifies an alternate domain for cross-domain trust relation")
	rGetResourceAccessExt.Input("checkPrincipal", "EntityName", false, "principal", "", true, nil, "usually null. If present, carry out the access check for this principal")
	rGetResourceAccessExt.Auth("", "", true, "")
	rGetResourceAccessExt.Exception("BAD_REQUEST", "ResourceError", "")
	rGetResourceAccessExt.Exception("FORBIDDEN", "ResourceError", "")
	rGetResourceAccessExt.Exception("NOT_FOUND", "ResourceError", "")
	rGetResourceAccessExt.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetResourceAccessExt.Build())

	rGetServiceIdentity := rdl.NewResourceBuilder("ServiceIdentity", "GET", "/domain/{domainName}/service/{serviceName}")
	rGetServiceIdentity.Comment("Get info for the specified ServiceIdentity.")
	rGetServiceIdentity.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetServiceIdentity.Input("serviceName", "ServiceName", true, "", "", false, nil, "name of the service to be retrieved")
	rGetServiceIdentity.Auth("", "", true, "")
	rGetServiceIdentity.Exception("BAD_REQUEST", "ResourceError", "")
	rGetServiceIdentity.Exception("NOT_FOUND", "ResourceError", "")
	rGetServiceIdentity.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetServiceIdentity.Build())

	rGetServiceIdentityList := rdl.NewResourceBuilder("ServiceIdentityList", "GET", "/domain/{domainName}/service")
	rGetServiceIdentityList.Comment("Enumerate services provisioned in this domain.")
	rGetServiceIdentityList.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetServiceIdentityList.Auth("", "", true, "")
	rGetServiceIdentityList.Exception("BAD_REQUEST", "ResourceError", "")
	rGetServiceIdentityList.Exception("NOT_FOUND", "ResourceError", "")
	rGetServiceIdentityList.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetServiceIdentityList.Build())

	rGetPublicKeyEntry := rdl.NewResourceBuilder("PublicKeyEntry", "GET", "/domain/{domainName}/service/{serviceName}/publickey/{keyId}")
	rGetPublicKeyEntry.Comment("Retrieve the specified public key from the service.")
	rGetPublicKeyEntry.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetPublicKeyEntry.Input("serviceName", "SimpleName", true, "", "", false, nil, "name of the service")
	rGetPublicKeyEntry.Input("keyId", "String", true, "", "", false, nil, "the identifier of the public key to be retrieved")
	rGetPublicKeyEntry.Exception("BAD_REQUEST", "ResourceError", "")
	rGetPublicKeyEntry.Exception("NOT_FOUND", "ResourceError", "")
	sb.AddResource(rGetPublicKeyEntry.Build())

	rGetHostServices := rdl.NewResourceBuilder("HostServices", "GET", "/host/{host}/services")
	rGetHostServices.Comment("Enumerate services provisioned on a specific host")
	rGetHostServices.Input("host", "String", true, "", "", false, nil, "name of the host")
	rGetHostServices.Exception("BAD_REQUEST", "ResourceError", "")
	sb.AddResource(rGetHostServices.Build())

	rGetDomainSignedPolicyData := rdl.NewResourceBuilder("DomainSignedPolicyData", "GET", "/domain/{domainName}/signed_policy_data")
	rGetDomainSignedPolicyData.Comment("Get a signed policy enumeration from the service, to transfer to a local store. An ETag is generated for the PolicyList that changes when any item in the list changes. If the If-None-Match header is provided, and it matches the ETag that would be returned, then a NOT_MODIFIED response is returned instead of the list.")
	rGetDomainSignedPolicyData.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetDomainSignedPolicyData.Input("matchingTag", "String", false, "", "If-None-Match", false, nil, "Retrieved from the previous request, this timestamp specifies to the server to return any policies modified since this time")
	rGetDomainSignedPolicyData.Output("tag", "String", "ETag", false, "The current latest modification timestamp is returned in this header")
	rGetDomainSignedPolicyData.Exception("BAD_REQUEST", "ResourceError", "")
	rGetDomainSignedPolicyData.Exception("NOT_FOUND", "ResourceError", "")
	sb.AddResource(rGetDomainSignedPolicyData.Build())

	rGetRoleToken := rdl.NewResourceBuilder("RoleToken", "GET", "/domain/{domainName}/token")
	rGetRoleToken.Comment("Return a security token for the specific role in the namespace that the principal can assume. If the role is omitted, then all roles in the namespace that the authenticated user can assume are returned. the caller can specify how long the RoleToken should be valid for by specifying the minExpiryTime and maxExpiryTime parameters. The minExpiryTime specifies that the returned RoleToken must be at least valid (min/lower bound) for specified number of seconds, while maxExpiryTime specifies that the RoleToken must be at most valid (max/upper bound) for specified number of seconds. If both values are the same, the server must return a RoleToken for that many seconds. If no values are specified, the server's default RoleToken Timeout value is used.")
	rGetRoleToken.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetRoleToken.Input("role", "EntityName", false, "role", "", true, nil, "only interested for a token for this role")
	rGetRoleToken.Input("minExpiryTime", "Int32", false, "minExpiryTime", "", true, nil, "in seconds min expiry time")
	rGetRoleToken.Input("maxExpiryTime", "Int32", false, "maxExpiryTime", "", true, nil, "in seconds max expiry time")
	rGetRoleToken.Input("proxyForPrincipal", "EntityName", false, "proxyForPrincipal", "", true, nil, "optional this request is proxy for this principal")
	rGetRoleToken.Auth("", "", true, "")
	rGetRoleToken.Exception("BAD_REQUEST", "ResourceError", "")
	rGetRoleToken.Exception("FORBIDDEN", "ResourceError", "")
	rGetRoleToken.Exception("NOT_FOUND", "ResourceError", "")
	rGetRoleToken.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetRoleToken.Build())

	rPostRoleCertificateRequest := rdl.NewResourceBuilder("RoleToken", "POST", "/domain/{domainName}/role/{roleName}/token")
	rPostRoleCertificateRequest.Comment("Return a TLS certificate for the specific role in the namespace that the principal can assume. Role certificates are valid for 30 days by default")
	rPostRoleCertificateRequest.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rPostRoleCertificateRequest.Input("roleName", "EntityName", true, "", "", false, nil, "name of role")
	rPostRoleCertificateRequest.Input("req", "RoleCertificateRequest", false, "", "", false, nil, "csr request")
	rPostRoleCertificateRequest.Auth("", "", true, "")
	rPostRoleCertificateRequest.Exception("BAD_REQUEST", "ResourceError", "")
	rPostRoleCertificateRequest.Exception("FORBIDDEN", "ResourceError", "")
	rPostRoleCertificateRequest.Exception("NOT_FOUND", "ResourceError", "")
	rPostRoleCertificateRequest.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostRoleCertificateRequest.Build())

	rGetAccess := rdl.NewResourceBuilder("Access", "GET", "/access/domain/{domainName}/role/{roleName}/principal/{principal}")
	rGetAccess.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetAccess.Input("roleName", "EntityName", true, "", "", false, nil, "name of the role to check access for")
	rGetAccess.Input("principal", "EntityName", true, "", "", false, nil, "carry out the access check for this principal")
	rGetAccess.Auth("", "", true, "")
	rGetAccess.Exception("BAD_REQUEST", "ResourceError", "")
	rGetAccess.Exception("FORBIDDEN", "ResourceError", "")
	rGetAccess.Exception("NOT_FOUND", "ResourceError", "")
	rGetAccess.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetAccess.Build())

	rGetRoleAccess := rdl.NewResourceBuilder("RoleAccess", "GET", "/access/domain/{domainName}/principal/{principal}")
	rGetRoleAccess.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain")
	rGetRoleAccess.Input("principal", "EntityName", true, "", "", false, nil, "carry out the role access lookup for this principal")
	rGetRoleAccess.Auth("", "", true, "")
	rGetRoleAccess.Exception("BAD_REQUEST", "ResourceError", "")
	rGetRoleAccess.Exception("NOT_FOUND", "ResourceError", "")
	rGetRoleAccess.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetRoleAccess.Build())

	rGetTenantDomains := rdl.NewResourceBuilder("TenantDomains", "GET", "/providerdomain/{providerDomainName}/user/{userName}")
	rGetTenantDomains.Comment("Get list of tenant domains user has access to for specified provider domain and service")
	rGetTenantDomains.Input("providerDomainName", "DomainName", true, "", "", false, nil, "name of the provider domain")
	rGetTenantDomains.Input("userName", "EntityName", true, "", "", false, nil, "name of the user to retrieve tenant domain access for")
	rGetTenantDomains.Input("roleName", "EntityName", false, "roleName", "", true, nil, "role name to filter on when looking for the tenants in provider")
	rGetTenantDomains.Input("serviceName", "ServiceName", false, "serviceName", "", true, nil, "service name to filter on when looking for the tenants in provider")
	rGetTenantDomains.Auth("", "", true, "")
	rGetTenantDomains.Exception("BAD_REQUEST", "ResourceError", "")
	rGetTenantDomains.Exception("NOT_FOUND", "ResourceError", "")
	rGetTenantDomains.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetTenantDomains.Build())

	rPostInstanceInformation := rdl.NewResourceBuilder("Identity", "POST", "/instance")
	rPostInstanceInformation.Comment("Get a cert for service being bootstrapped by supported service")
	rPostInstanceInformation.Input("info", "InstanceInformation", false, "", "", false, nil, "")
	rPostInstanceInformation.Exception("BAD_REQUEST", "ResourceError", "")
	rPostInstanceInformation.Exception("FORBIDDEN", "ResourceError", "")
	rPostInstanceInformation.Exception("INTERNAL_SERVER_ERROR", "ResourceError", "")
	rPostInstanceInformation.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostInstanceInformation.Build())

	rPostInstanceRefreshRequest := rdl.NewResourceBuilder("Identity", "POST", "/instance/{domain}/{service}/refresh")
	rPostInstanceRefreshRequest.Comment("Refresh self identity if the original identity was issued by ZTS")
	rPostInstanceRefreshRequest.Input("domain", "CompoundName", true, "", "", false, nil, "name of the domain requesting the refresh")
	rPostInstanceRefreshRequest.Input("service", "SimpleName", true, "", "", false, nil, "name of the service requesting the refresh")
	rPostInstanceRefreshRequest.Input("req", "InstanceRefreshRequest", false, "", "", false, nil, "the refresh request")
	rPostInstanceRefreshRequest.Auth("", "", true, "")
	rPostInstanceRefreshRequest.Exception("BAD_REQUEST", "ResourceError", "")
	rPostInstanceRefreshRequest.Exception("FORBIDDEN", "ResourceError", "")
	rPostInstanceRefreshRequest.Exception("INTERNAL_SERVER_ERROR", "ResourceError", "")
	rPostInstanceRefreshRequest.Exception("NOT_FOUND", "ResourceError", "")
	rPostInstanceRefreshRequest.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostInstanceRefreshRequest.Build())

	rPostAWSInstanceInformation := rdl.NewResourceBuilder("Identity", "POST", "/aws/instance")
	rPostAWSInstanceInformation.Comment("Register an instance in AWS ZTS. Whether this succeeds or not depends on the contents of the request (the request itself is not authenticated or authorized in the normal way). If successful, the Identity is returned as a x.509 client certificate (to be used in TLS operations)")
	rPostAWSInstanceInformation.Input("info", "AWSInstanceInformation", false, "", "", false, nil, "")
	rPostAWSInstanceInformation.Exception("BAD_REQUEST", "ResourceError", "")
	rPostAWSInstanceInformation.Exception("FORBIDDEN", "ResourceError", "")
	rPostAWSInstanceInformation.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostAWSInstanceInformation.Build())

	rPostAWSCertificateRequest := rdl.NewResourceBuilder("Identity", "POST", "/aws/instance/{domain}/{service}/refresh")
	rPostAWSCertificateRequest.Comment("Rotate certs. Make this request with previous cert, the result are new certs for the same identity.")
	rPostAWSCertificateRequest.Input("domain", "CompoundName", true, "", "", false, nil, "name of the domain requesting the refresh")
	rPostAWSCertificateRequest.Input("service", "SimpleName", true, "", "", false, nil, "name of the service requesting the refresh")
	rPostAWSCertificateRequest.Input("req", "AWSCertificateRequest", false, "", "", false, nil, "the refresh request")
	rPostAWSCertificateRequest.Auth("", "", true, "")
	rPostAWSCertificateRequest.Exception("BAD_REQUEST", "ResourceError", "")
	rPostAWSCertificateRequest.Exception("FORBIDDEN", "ResourceError", "")
	rPostAWSCertificateRequest.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostAWSCertificateRequest.Build())

	rGetAWSTemporaryCredentials := rdl.NewResourceBuilder("AWSTemporaryCredentials", "GET", "/domain/{domainName}/role/{role}/creds")
	rGetAWSTemporaryCredentials.Comment("perform an AWS AssumeRole of the target role and return the credentials. ZTS must have been granted the ability to assume the role in IAM, and granted the ability to ASSUME_AWS_ROLE in Athenz for this to succeed.")
	rGetAWSTemporaryCredentials.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain containing the role, which implies the target account")
	rGetAWSTemporaryCredentials.Input("role", "CompoundName", true, "", "", false, nil, "the target AWS role name in the domain account, in Athenz terms, i.e. \"the.role\"")
	rGetAWSTemporaryCredentials.Auth("", "", true, "")
	rGetAWSTemporaryCredentials.Exception("BAD_REQUEST", "ResourceError", "")
	rGetAWSTemporaryCredentials.Exception("FORBIDDEN", "ResourceError", "")
	rGetAWSTemporaryCredentials.Exception("NOT_FOUND", "ResourceError", "")
	rGetAWSTemporaryCredentials.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rGetAWSTemporaryCredentials.Build())

	rPostOSTKInstanceInformation := rdl.NewResourceBuilder("Identity", "POST", "/ostk/instance")
	rPostOSTKInstanceInformation.Comment("Get a cert for service being bootstrapped by Openstack")
	rPostOSTKInstanceInformation.Input("info", "OSTKInstanceInformation", false, "", "", false, nil, "")
	rPostOSTKInstanceInformation.Exception("BAD_REQUEST", "ResourceError", "")
	rPostOSTKInstanceInformation.Exception("FORBIDDEN", "ResourceError", "")
	rPostOSTKInstanceInformation.Exception("INTERNAL_SERVER_ERROR", "ResourceError", "")
	rPostOSTKInstanceInformation.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostOSTKInstanceInformation.Build())

	rPostOSTKInstanceRefreshRequest := rdl.NewResourceBuilder("Identity", "POST", "/ostk/instance/{domain}/{service}/refresh")
	rPostOSTKInstanceRefreshRequest.Comment("Refresh self identity if the original identity was issued by ZTS The token must include the original requestor's name and the server will verify that the service still has authorization to grant inception to the current service requesting to refresh its identity")
	rPostOSTKInstanceRefreshRequest.Input("domain", "CompoundName", true, "", "", false, nil, "name of the tenant domain")
	rPostOSTKInstanceRefreshRequest.Input("service", "SimpleName", true, "", "", false, nil, "name of the tenant service")
	rPostOSTKInstanceRefreshRequest.Input("req", "OSTKInstanceRefreshRequest", false, "", "", false, nil, "the refresh request")
	rPostOSTKInstanceRefreshRequest.Auth("", "", true, "")
	rPostOSTKInstanceRefreshRequest.Exception("BAD_REQUEST", "ResourceError", "")
	rPostOSTKInstanceRefreshRequest.Exception("FORBIDDEN", "ResourceError", "")
	rPostOSTKInstanceRefreshRequest.Exception("INTERNAL_SERVER_ERROR", "ResourceError", "")
	rPostOSTKInstanceRefreshRequest.Exception("NOT_FOUND", "ResourceError", "")
	rPostOSTKInstanceRefreshRequest.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostOSTKInstanceRefreshRequest.Build())

	rPostDomainMetrics := rdl.NewResourceBuilder("DomainMetrics", "POST", "/metrics/{domainName}")
	rPostDomainMetrics.Comment("called to post multiple zpe related metric attributes")
	rPostDomainMetrics.Input("domainName", "DomainName", true, "", "", false, nil, "name of the domain the metrics pertain to")
	rPostDomainMetrics.Input("req", "DomainMetrics", false, "", "", false, nil, "")
	rPostDomainMetrics.Exception("BAD_REQUEST", "ResourceError", "")
	rPostDomainMetrics.Exception("FORBIDDEN", "ResourceError", "")
	rPostDomainMetrics.Exception("NOT_FOUND", "ResourceError", "")
	rPostDomainMetrics.Exception("UNAUTHORIZED", "ResourceError", "")
	sb.AddResource(rPostDomainMetrics.Build())

	schema = sb.Build()
}

func ZTSSchema() *rdl.Schema {
	return schema
}
