// Diagram 1: SMESec — Logical Architecture (Clean Architecture Layers)
//
// Run:
//   cd diagrams
//   go run cmd/logical/main.go
//   dot -Tpng out/logical-architecture.dot -o out/logical-architecture.png
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
		diagram.Label("SMESec — Logical Architecture (Clean Architecture)"),
		diagram.Direction("TB"),
		diagram.Filename("out/logical-architecture"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// ─── Interface Layer (outermost) ────────────────────────────────────────
	webApp := apps.Network.Internet(diagram.NodeLabel("Web App\n(React/Next.js)"))
	mobileApp := apps.Network.Internet(diagram.NodeLabel("Mobile App\n(Flutter)"))
	browserExt := apps.Network.Internet(diagram.NodeLabel("Browser Extension\n(Chrome MV3)"))

	interfaceGroup := diagram.NewGroup("interface").
		Label("Interface Layer — Controllers & Presenters").
		Add(webApp, mobileApp, browserExt)

	// ─── API Gateway ────────────────────────────────────────────────────────
	apiGateway := apps.Network.Kong(diagram.NodeLabel("API Gateway\n(REST + gRPC + WS)\nJWT Auth · Rate Limit"))

	apiGroup := diagram.NewGroup("api").
		Label("API Layer").
		Add(apiGateway)

	// ─── Application Layer — Use Cases ──────────────────────────────────────
	assetSvc := aws.Compute.Fargate(diagram.NodeLabel("Asset Inventory\nService"))
	accessSvc := aws.Compute.Fargate(diagram.NodeLabel("Access Governance\nService"))
	playbookSvc := aws.Compute.Fargate(diagram.NodeLabel("Incident Playbook\nService"))
	complianceSvc := aws.Compute.Fargate(diagram.NodeLabel("Compliance\nService"))
	syncSvc := aws.Compute.Fargate(diagram.NodeLabel("Integration Sync\nService"))
	threatSvc := aws.Compute.Fargate(diagram.NodeLabel("Threat Detection\nService (Track 2)"))

	appGroup := diagram.NewGroup("application").
		Label("Application Layer — Use Cases").
		Add(assetSvc, accessSvc, playbookSvc, complianceSvc, syncSvc, threatSvc)

	// ─── Domain Layer (innermost) ────────────────────────────────────────────
	domainEntities := apps.Database.Postgresql(diagram.NodeLabel("Domain Entities\nAsset · TenantUser · ThreatEvent\nPlaybook · ComplianceControl\nAccessPolicy · TenantConfig"))
	domainServices := apps.Container.Docker(diagram.NodeLabel("Domain Services\nRiskScorer · AccessGovernor\nComplianceAuditor · PlaybookExecutor"))

	domainGroup := diagram.NewGroup("domain").
		Label("Domain Layer — Entities & Domain Services (zero external dependencies)").
		Add(domainEntities, domainServices)

	// ─── Infrastructure Layer ────────────────────────────────────────────────
	repositories := aws.Database.Rds(diagram.NodeLabel("Repositories\nPostgresAssetRepo · PostgresUserRepo\nPostgresPlaybookRepo\n[tenant_id RLS enforced]"))
	integAdapters := apps.Network.Internet(diagram.NodeLabel("Integration Adapters\nGoogleWorkspaceAdapter\nM365Adapter · SlackAdapter\nAWSIAMAdapter"))
	eventPublishers := aws.Integration.Eventbridge(diagram.NodeLabel("Event Publishers\nEventBridgePublisher\nSNSNotificationPublisher"))
	externalClients := apps.Network.Internet(diagram.NodeLabel("External Clients\nVantaClient · HiveModerationClient\nSageMakerClient · KeycloakClient"))

	infraGroup := diagram.NewGroup("infra").
		Label("Infrastructure Layer — Adapters & Repositories").
		Add(repositories, integAdapters, eventPublishers, externalClients)

	// ─── Connections (Dependency Rule: outer → inner) ────────────────────────
	// Interface → API
	interfaceGroup.ConnectAllTo(apiGateway.ID(), diagram.Forward())

	// API → Application services
	appGroup.ConnectAllFrom(apiGateway.ID(), diagram.Forward())

	// Application → Domain (use cases call domain services & load entities)
	appGroup.ConnectAllTo(domainEntities.ID(), diagram.Forward())
	appGroup.ConnectAllTo(domainServices.ID(), diagram.Forward())

	// Infrastructure → Domain (implements domain port interfaces — dependency inversion)
	infraGroup.ConnectAllTo(domainEntities.ID(), diagram.Forward())

	// ─── Add all groups to diagram ───────────────────────────────────────────
	d.Group(interfaceGroup).
		Group(apiGroup).
		Group(appGroup).
		Group(domainGroup).
		Group(infraGroup)

	if err := d.Render(); err != nil {
		log.Fatal(err)
	}

	log.Println("Generated: go-diagrams/out/logical-architecture.dot")
	log.Println("Render PNG: dot -Tpng go-diagrams/out/logical-architecture.dot -o go-diagrams/out/logical-architecture.png")
}
