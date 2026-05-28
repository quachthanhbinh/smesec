// Diagram 3: SMESec — Integration Touchpoints
//
// Run:
//   cd diagrams
//   go run cmd/integrations/main.go
//   dot -Tpng out/integration-touchpoints.dot -o out/integration-touchpoints.png
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
		diagram.Label("SMESec — Integration Touchpoints"),
		diagram.Direction("LR"),
		diagram.Filename("out/integration-touchpoints"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ─── SMESec Core Platform ────────────────────────────────────────────────
	apiGateway := apps.Network.Kong(diagram.NodeLabel("API Gateway\n(JWT · Rate Limit)"))
	integSyncSvc := aws.Compute.Fargate(diagram.NodeLabel("Integration Sync\nService"))
	assetSvc := aws.Compute.Fargate(diagram.NodeLabel("Asset Inventory\nService"))
	accessSvc := aws.Compute.Fargate(diagram.NodeLabel("Access Governance\nService"))
	threatSvc := aws.Compute.Fargate(diagram.NodeLabel("Threat Detection\nService (Track 2)"))
	complianceSvc := aws.Compute.Fargate(diagram.NodeLabel("Compliance\nService"))
	eventBus := aws.Integration.Eventbridge(diagram.NodeLabel("EventBridge\n(Event Bus)"))
	playbookEngine := aws.Integration.StepFunctions(diagram.NodeLabel("Playbook Engine\n(Step Functions)"))
	db := aws.Database.Rds(diagram.NodeLabel("PostgreSQL\n(Multi-tenant RLS)"))
	notifications := aws.Integration.SimpleNotificationServiceSns(diagram.NodeLabel("SNS\n(Alerts)"))

	coreGroup := diagram.NewGroup("core").
		Label("SMESec Core Platform").
		Add(apiGateway, integSyncSvc, assetSvc, accessSvc, threatSvc, complianceSvc, eventBus, playbookEngine, db, notifications)

	// ─── Identity & Productivity Providers ──────────────────────────────────
	googleWS := apps.Network.Internet(diagram.NodeLabel("Google Workspace\n(Admin SDK)\nOAuth 2.0 Service Account\nWorkspace Events API\n→ Users · Groups · OAuth Apps\n→ Audit Logs · Devices\nWrite: Suspend user · Revoke OAuth\nSync: 15-min delta"))
	microsoft365 := apps.Network.Internet(diagram.NodeLabel("Microsoft 365\n(Graph API + Azure AD)\nOAuth 2.0 App Registration\nDelta Link + Webhooks\n→ Users · Groups · OAuth Apps\n→ SignIn Logs · Devices\nWrite: Block signin · Revoke sessions\nSync: 15-min delta"))
	slack := apps.Network.Internet(diagram.NodeLabel("Slack\n(Admin API + Events API)\nOAuth 2.0\n→ Users · Channels · Apps · Audit Logs\nWrite: Deactivate user (Business+)\nSync: 15-min poll + webhooks"))
	awsIAM := aws.Security.IdentityAndAccessManagementIam(diagram.NodeLabel("AWS IAM\n(Assumed Role)\n→ Users · Roles · Policies\n→ Access Keys · CloudTrail\nWrite: Disable key (dry-run first)\nSync: 30-min full pull"))

	idpGroup := diagram.NewGroup("idp").
		Label("Identity & Productivity Providers").
		Add(googleWS, microsoft365, slack, awsIAM)

	// ─── Security & Compliance Services ─────────────────────────────────────
	keycloak := aws.Compute.Fargate(diagram.NodeLabel("Keycloak\n(Self-hosted ECS)\nOIDC / SAML 2.0\nGoogle + M365 federation\nMFA TOTP (mandatory)\nJWT RS256"))
	vanta := apps.Network.Internet(diagram.NodeLabel("Vanta\n(Compliance Automation)\nOAuth → AWS + GitHub\nSOC 2 Type 1 (Month 6)\nISO 27001 (Month 12)\nEvidence collection"))
	hive := apps.Network.Internet(diagram.NodeLabel("Hive Moderation\n(Deepfake Detection)\nREST API (pay-per-use)\n<$0.01 per check\nVoice + Video analysis"))
	sagemaker := aws.Ml.Sagemaker(diagram.NodeLabel("SageMaker\n(ML Inference)\nshadow-ai-scorer-v1\nbert-injection-v1\nAsync inference queue"))
	secretsMgr := aws.Security.SecretsManager(diagram.NodeLabel("Secrets Manager\n(Token Vault)\nOAuth refresh tokens\nAuto-rotation enabled\nNo plaintext env vars"))

	secGroup := diagram.NewGroup("security").
		Label("Security & Compliance Services").
		Add(keycloak, vanta, hive, sagemaker, secretsMgr)

	// ─── Client Applications ─────────────────────────────────────────────────
	webApp := apps.Network.Internet(diagram.NodeLabel("Web App\n(React/Next.js)\nAdmin dashboard\nCompliance reports\nInventory view"))
	mobileApp := apps.Network.Internet(diagram.NodeLabel("Mobile App\n(Flutter)\niOS + Android\nIncident alerts\nPlaybook approval"))
	browserExt := apps.Network.Internet(diagram.NodeLabel("Browser Extension\n(Chrome MV3 + Edge)\nLLM DLP (local inference)\nShadow AI detection\nPII interception blocking"))

	clientGroup := diagram.NewGroup("clients").
		Label("Client Applications").
		Add(webApp, mobileApp, browserExt)

	// ─── Connections ─────────────────────────────────────────────────────────
	// Clients → Core
	clientGroup.ConnectAllTo(apiGateway.ID(), diagram.Forward())

	// API Gateway → Core services
	d.Connect(apiGateway, integSyncSvc, diagram.Forward())
	d.Connect(apiGateway, assetSvc, diagram.Forward())
	d.Connect(apiGateway, accessSvc, diagram.Forward())
	d.Connect(apiGateway, complianceSvc, diagram.Forward())

	// Integration Sync ↔ Identity Providers
	d.Connect(integSyncSvc, googleWS, diagram.Forward())
	d.Connect(integSyncSvc, microsoft365, diagram.Forward())
	d.Connect(integSyncSvc, slack, diagram.Forward())
	d.Connect(integSyncSvc, awsIAM, diagram.Forward())

	// Identity Providers → Integration Sync (webhook callbacks)
	d.Connect(googleWS, integSyncSvc, diagram.Forward())
	d.Connect(microsoft365, integSyncSvc, diagram.Forward())
	d.Connect(slack, integSyncSvc, diagram.Forward())

	// Core services → Event Bus
	d.Connect(assetSvc, eventBus, diagram.Forward())
	d.Connect(accessSvc, eventBus, diagram.Forward())
	d.Connect(integSyncSvc, eventBus, diagram.Forward())

	// Event Bus → Track 2 + Playbook Engine + Notifications
	d.Connect(eventBus, threatSvc, diagram.Forward())
	d.Connect(eventBus, playbookEngine, diagram.Forward())
	d.Connect(eventBus, notifications, diagram.Forward())

	// Core ↔ Database
	coreGroup.ConnectAllTo(db.ID(), diagram.Forward())

	// Threat Detection → Hive (deepfake) + SageMaker (ML inference)
	d.Connect(threatSvc, hive, diagram.Forward())
	d.Connect(threatSvc, sagemaker, diagram.Forward())

	// Browser Extension → Threat Detection (LLM DLP events)
	d.Connect(browserExt, threatSvc, diagram.Forward())

	// Compliance → Vanta
	d.Connect(complianceSvc, vanta, diagram.Forward())

	// Auth: Keycloak validates every request via API Gateway
	d.Connect(keycloak, apiGateway, diagram.Forward())

	// Access Governance → Identity Provider write operations (offboarding)
	d.Connect(accessSvc, googleWS, diagram.Forward())
	d.Connect(accessSvc, microsoft365, diagram.Forward())
	d.Connect(accessSvc, slack, diagram.Forward())
	d.Connect(accessSvc, awsIAM, diagram.Forward())

	// Notifications → Clients (mobile push + email)
	d.Connect(notifications, mobileApp, diagram.Forward())

	// Secrets Manager → Integration Sync (OAuth token retrieval)
	d.Connect(secretsMgr, integSyncSvc, diagram.Forward())

	// ─── Add groups to diagram ───────────────────────────────────────────────
	d.Group(clientGroup).
		Group(coreGroup).
		Group(idpGroup).
		Group(secGroup)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}

	log.Println("Generated: go-diagrams/out/integration-touchpoints.dot")
	log.Println("Render PNG: dot -Tpng go-diagrams/out/integration-touchpoints.dot -o go-diagrams/out/integration-touchpoints.png")
}
