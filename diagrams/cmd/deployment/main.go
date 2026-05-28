// Diagram 2: SMESec — AWS Deployment Architecture
//
// Run:
//   cd diagrams
//   go run cmd/deployment/main.go
//   dot -Tpng out/deployment-architecture.dot -o out/deployment-architecture.png
package main

import (
	"log"
	"os"

	"github.com/blushft/go-diagrams/diagram"
	"github.com/blushft/go-diagrams/nodes/apps"
	"github.com/blushft/go-diagrams/nodes/aws"
)

func main() {
	if err := os.MkdirAll("go-diagrams/out", 0o755); err != nil {
		log.Fatal(err)
	}

	d, err := diagram.New(
		diagram.Label("SMESec — AWS Deployment Architecture"),
		diagram.Direction("TB"),
		diagram.Filename("out/deployment-architecture"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ─── Client Tier ─────────────────────────────────────────────────────────
	webClient := apps.Network.Internet(diagram.NodeLabel("Web App\n(React/Next.js)"))
	mobileClient := apps.Network.Internet(diagram.NodeLabel("Mobile App\n(Flutter)"))
	extensionClient := apps.Network.Internet(diagram.NodeLabel("Browser Extension\n(Chrome MV3)"))

	clientGroup := diagram.NewGroup("clients").
		Label("Clients").
		Add(webClient, mobileClient, extensionClient)

	// ─── Edge / Global ───────────────────────────────────────────────────────
	route53 := aws.Network.Route53(diagram.NodeLabel("Route 53\n(DNS + Health Check)"))
	cloudfront := aws.Network.Cloudfront(diagram.NodeLabel("CloudFront\n(CDN + SSL Termination)"))
	waf := aws.Security.Waf(diagram.NodeLabel("WAF\n(OWASP Top 10\nRate Limiting)"))
	alb := aws.Network.ElasticLoadBalancing(diagram.NodeLabel("ALB\n(Layer 7\nMulti-AZ)"))
	certMgr := aws.Security.CertificateManager(diagram.NodeLabel("ACM\n(TLS Certificates)"))

	edgeGroup := diagram.NewGroup("edge").
		Label("Edge Zone — AWS Global + us-east-1").
		Add(route53, cloudfront, waf, alb, certMgr)

	// ─── Auth Tier ───────────────────────────────────────────────────────────
	keycloak := aws.Compute.Fargate(diagram.NodeLabel("Keycloak\n(OIDC + SAML 2.0\nMFA TOTP mandatory)\nMulti-AZ"))

	authGroup := diagram.NewGroup("auth").
		Label("Auth Tier — Private Subnet").
		Add(keycloak)

	// ─── Application Tier — Track 1 ─────────────────────────────────────────
	apiGw := apps.Network.Kong(diagram.NodeLabel("API Gateway\n(JWT Validation\nRouting · Rate Limit)"))
	assetSvc := aws.Compute.Fargate(diagram.NodeLabel("Asset Inventory\nService"))
	accessSvc := aws.Compute.Fargate(diagram.NodeLabel("Access Governance\nService"))
	playbookSvc := aws.Compute.Fargate(diagram.NodeLabel("Incident Playbook\nService"))
	complianceSvc := aws.Compute.Fargate(diagram.NodeLabel("Compliance\nService"))
	syncSvc := aws.Compute.Fargate(diagram.NodeLabel("Integration Sync\nService"))

	track1Group := diagram.NewGroup("track1").
		Label("Application Tier — Track 1 Services (ECS Fargate, Private Subnet)").
		Add(apiGw, assetSvc, accessSvc, playbookSvc, complianceSvc, syncSvc)

	// ─── Application Tier — Track 2 ─────────────────────────────────────────
	threatSvc := aws.Compute.Fargate(diagram.NodeLabel("Threat Detection\nService"))
	dlpSvc := aws.Compute.Fargate(diagram.NodeLabel("LLM DLP\nService"))
	deepfakeSvc := aws.Compute.Fargate(diagram.NodeLabel("Deepfake Defense\nService"))

	track2Group := diagram.NewGroup("track2").
		Label("Application Tier — Track 2 Services (ECS Fargate, Private Subnet)").
		Add(threatSvc, dlpSvc, deepfakeSvc)

	// ─── Data Tier ───────────────────────────────────────────────────────────
	rds := aws.Database.Rds(diagram.NodeLabel("RDS PostgreSQL\n(Multi-AZ\ntenant_id RLS\ndata_residency column)"))
	redis := aws.Database.Elasticache(diagram.NodeLabel("ElastiCache Redis\n(Sessions · Cache\nRate Limit tokens)"))
	s3 := aws.Storage.SimpleStorageServiceS3(diagram.NodeLabel("S3\n(Object Lock WORM\n7-year retention\nAudit logs)"))
	ecr := aws.Compute.ElasticContainerService(diagram.NodeLabel("ECR\n(Container Registry\nImage scanning)"))

	dataGroup := diagram.NewGroup("data").
		Label("Data Tier — Private Subnet").
		Add(rds, redis, s3, ecr)

	// ─── Event / Orchestration ───────────────────────────────────────────────
	eventbridge := aws.Integration.Eventbridge(diagram.NodeLabel("EventBridge\n(Event Bus\nDomain Events routing)"))
	stepFunctions := aws.Integration.StepFunctions(diagram.NodeLabel("Step Functions\n(Playbook Engine\nWorkflow Orchestration)"))
	sns := aws.Integration.SimpleNotificationServiceSns(diagram.NodeLabel("SNS\n(Email · Slack Webhook\nPagerDuty Alerts)"))
	sqs := aws.Integration.SimpleQueueServiceSqs(diagram.NodeLabel("SQS\n(Integration Sync Queue\nRetry + DLQ)"))

	eventGroup := diagram.NewGroup("events").
		Label("Event / Orchestration — AWS Managed Services").
		Add(eventbridge, stepFunctions, sns, sqs)

	// ─── AI / ML ─────────────────────────────────────────────────────────────
	sagemaker := aws.Ml.Sagemaker(diagram.NodeLabel("SageMaker\n(Shadow AI Risk Model\nPrompt Injection BERT\nInference Endpoints)"))
	sagemakerModel := aws.Ml.SagemakerModel(diagram.NodeLabel("Model Registry\n(shadow-ai-scorer-v1\nbert-injection-v1)"))

	mlGroup := diagram.NewGroup("ml").
		Label("AI / ML — SageMaker").
		Add(sagemaker, sagemakerModel)

	// ─── Security Services ───────────────────────────────────────────────────
	secretsMgr := aws.Security.SecretsManager(diagram.NodeLabel("Secrets Manager\n(OAuth tokens\nDB passwords\nAPI keys — auto-rotate)"))
	kms := aws.Security.KeyManagementService(diagram.NodeLabel("KMS\n(CMK for RDS\nS3 SSE-KMS\nSecrets encryption)"))
	guardduty := aws.Security.Guardduty(diagram.NodeLabel("GuardDuty\n(Threat Intelligence\nAnomaly Detection)"))
	iam := aws.Security.IdentityAndAccessManagementIam(diagram.NodeLabel("IAM\n(Service Roles\nLeast Privilege\nSCPs)"))

	secGroup := diagram.NewGroup("security").
		Label("Security Services — AWS Managed").
		Add(secretsMgr, kms, guardduty, iam)

	// ─── Observability ───────────────────────────────────────────────────────
	cloudwatch := aws.Management.Cloudwatch(diagram.NodeLabel("CloudWatch\n(Metrics · Alarms\nLogs · Dashboards)"))
	cloudtrail := aws.Management.Cloudtrail(diagram.NodeLabel("CloudTrail\n(Immutable API audit\nAll regions enabled)"))

	obsGroup := diagram.NewGroup("observability").
		Label("Observability — AWS Managed").
		Add(cloudwatch, cloudtrail)

	// ─── External Services ───────────────────────────────────────────────────
	vanta := apps.Network.Internet(diagram.NodeLabel("Vanta\n(Compliance Automation\nSOC 2 + ISO 27001)"))
	hive := apps.Network.Internet(diagram.NodeLabel("Hive Moderation\n(Deepfake Detection API\nPay-per-use)"))

	externalGroup := diagram.NewGroup("external").
		Label("External Services (SaaS)").
		Add(vanta, hive)

	// ─── Connections ─────────────────────────────────────────────────────────
	// Client → Edge
	clientGroup.ConnectAllTo(cloudfront.ID(), diagram.Forward())
	d.Connect(route53, cloudfront, diagram.Forward())
	d.Connect(cloudfront, waf, diagram.Forward())
	d.Connect(waf, alb, diagram.Forward())

	// Edge → Auth + App
	d.Connect(alb, keycloak, diagram.Forward())
	d.Connect(alb, apiGw, diagram.Forward())

	// Auth → Track 1 (JWT validation)
	d.Connect(keycloak, apiGw, diagram.Forward())

	// API Gateway → Track 1 services
	track1Group.ConnectAllFrom(apiGw.ID(), diagram.Forward())

	// Track 1 services → Data
	track1Group.ConnectAllTo(rds.ID(), diagram.Forward())
	track1Group.ConnectAllTo(redis.ID(), diagram.Forward())

	// Track 1 → Event bus
	track1Group.ConnectAllTo(eventbridge.ID(), diagram.Forward())

	// Event bus → Step Functions + SNS + Track 2
	d.Connect(eventbridge, stepFunctions, diagram.Forward())
	d.Connect(eventbridge, sns, diagram.Forward())
	d.Connect(eventbridge, sqs, diagram.Forward())
	d.Connect(eventbridge, threatSvc, diagram.Forward())

	// Step Functions → Track 1 (playbook orchestration)
	d.Connect(stepFunctions, playbookSvc, diagram.Forward())

	// Track 2 services → Data + ML
	track2Group.ConnectAllTo(rds.ID(), diagram.Forward())
	track2Group.ConnectAllTo(sagemaker.ID(), diagram.Forward())
	d.Connect(sagemaker, sagemakerModel, diagram.Forward())

	// Track 2 → deepfake external
	d.Connect(deepfakeSvc, hive, diagram.Forward())

	// Compliance → Vanta
	d.Connect(complianceSvc, vanta, diagram.Forward())

	// Secrets Manager feeds all services (conceptual — shown via security group)
	track1Group.ConnectAllTo(secretsMgr.ID(), diagram.Forward())
	track2Group.ConnectAllTo(secretsMgr.ID(), diagram.Forward())

	// Audit log → S3
	d.Connect(cloudtrail, s3, diagram.Forward())
	track1Group.ConnectAllTo(cloudwatch.ID(), diagram.Forward())
	track2Group.ConnectAllTo(cloudwatch.ID(), diagram.Forward())

	// ─── Add groups to diagram ───────────────────────────────────────────────
	d.Group(clientGroup).
		Group(edgeGroup).
		Group(authGroup).
		Group(track1Group).
		Group(track2Group).
		Group(dataGroup).
		Group(eventGroup).
		Group(mlGroup).
		Group(secGroup).
		Group(obsGroup).
		Group(externalGroup)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}

	log.Println("Generated: go-diagrams/out/deployment-architecture.dot")
	log.Println("Render PNG: dot -Tpng go-diagrams/out/deployment-architecture.dot -o go-diagrams/out/deployment-architecture.png")
}
