// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gqlgen

import (
	"bytes"
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	"github.com/aklinker1/miasma/internal"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	App() AppResolver
	Health() HealthResolver
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	App struct {
		Command        func(childComplexity int) int
		CreatedAt      func(childComplexity int) int
		Env            func(childComplexity int) int
		Group          func(childComplexity int) int
		Hidden         func(childComplexity int) int
		ID             func(childComplexity int) int
		Image          func(childComplexity int) int
		ImageDigest    func(childComplexity int) int
		Instances      func(childComplexity int) int
		Name           func(childComplexity int) int
		Networks       func(childComplexity int) int
		Placement      func(childComplexity int) int
		PublishedPorts func(childComplexity int) int
		Route          func(childComplexity int) int
		SimpleRoute    func(childComplexity int) int
		Status         func(childComplexity int) int
		System         func(childComplexity int) int
		TargetPorts    func(childComplexity int) int
		UpdatedAt      func(childComplexity int) int
		Volumes        func(childComplexity int) int
	}

	AppInstances struct {
		Running func(childComplexity int) int
		Total   func(childComplexity int) int
	}

	BoundVolume struct {
		Source func(childComplexity int) int
		Target func(childComplexity int) int
	}

	ClusterInfo struct {
		CreatedAt   func(childComplexity int) int
		ID          func(childComplexity int) int
		JoinCommand func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
	}

	Health struct {
		Cluster       func(childComplexity int) int
		DockerVersion func(childComplexity int) int
		Version       func(childComplexity int) int
	}

	Mutation struct {
		CreateApp      func(childComplexity int, input internal.AppInput) int
		DeleteApp      func(childComplexity int, id string) int
		DisablePlugin  func(childComplexity int, name internal.PluginName) int
		EditApp        func(childComplexity int, id string, changes map[string]interface{}) int
		EnablePlugin   func(childComplexity int, name internal.PluginName) int
		RemoveAppRoute func(childComplexity int, appID string) int
		RestartApp     func(childComplexity int, id string) int
		SetAppEnv      func(childComplexity int, appID string, newEnv map[string]interface{}) int
		SetAppRoute    func(childComplexity int, appID string, route *internal.RouteInput) int
		StartApp       func(childComplexity int, id string) int
		StopApp        func(childComplexity int, id string) int
		UpgradeApp     func(childComplexity int, id string) int
	}

	Plugin struct {
		Enabled func(childComplexity int) int
		Name    func(childComplexity int) int
	}

	Query struct {
		GetApp      func(childComplexity int, id string) int
		GetPlugin   func(childComplexity int, pluginName string) int
		Health      func(childComplexity int) int
		ListApps    func(childComplexity int, page *int32, size *int32, showHidden *bool) int
		ListPlugins func(childComplexity int) int
	}

	Route struct {
		AppID       func(childComplexity int) int
		CreatedAt   func(childComplexity int) int
		Host        func(childComplexity int) int
		Path        func(childComplexity int) int
		TraefikRule func(childComplexity int) int
		UpdatedAt   func(childComplexity int) int
	}
}

type executableSchema struct {
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e}
	_ = ec
	switch typeName + "." + field {

	case "App.command":
		if e.complexity.App.Command == nil {
			break
		}

		return e.complexity.App.Command(childComplexity), true

	case "App.createdAt":
		if e.complexity.App.CreatedAt == nil {
			break
		}

		return e.complexity.App.CreatedAt(childComplexity), true

	case "App.env":
		if e.complexity.App.Env == nil {
			break
		}

		return e.complexity.App.Env(childComplexity), true

	case "App.group":
		if e.complexity.App.Group == nil {
			break
		}

		return e.complexity.App.Group(childComplexity), true

	case "App.hidden":
		if e.complexity.App.Hidden == nil {
			break
		}

		return e.complexity.App.Hidden(childComplexity), true

	case "App.id":
		if e.complexity.App.ID == nil {
			break
		}

		return e.complexity.App.ID(childComplexity), true

	case "App.image":
		if e.complexity.App.Image == nil {
			break
		}

		return e.complexity.App.Image(childComplexity), true

	case "App.imageDigest":
		if e.complexity.App.ImageDigest == nil {
			break
		}

		return e.complexity.App.ImageDigest(childComplexity), true

	case "App.instances":
		if e.complexity.App.Instances == nil {
			break
		}

		return e.complexity.App.Instances(childComplexity), true

	case "App.name":
		if e.complexity.App.Name == nil {
			break
		}

		return e.complexity.App.Name(childComplexity), true

	case "App.networks":
		if e.complexity.App.Networks == nil {
			break
		}

		return e.complexity.App.Networks(childComplexity), true

	case "App.placement":
		if e.complexity.App.Placement == nil {
			break
		}

		return e.complexity.App.Placement(childComplexity), true

	case "App.publishedPorts":
		if e.complexity.App.PublishedPorts == nil {
			break
		}

		return e.complexity.App.PublishedPorts(childComplexity), true

	case "App.route":
		if e.complexity.App.Route == nil {
			break
		}

		return e.complexity.App.Route(childComplexity), true

	case "App.simpleRoute":
		if e.complexity.App.SimpleRoute == nil {
			break
		}

		return e.complexity.App.SimpleRoute(childComplexity), true

	case "App.status":
		if e.complexity.App.Status == nil {
			break
		}

		return e.complexity.App.Status(childComplexity), true

	case "App.system":
		if e.complexity.App.System == nil {
			break
		}

		return e.complexity.App.System(childComplexity), true

	case "App.targetPorts":
		if e.complexity.App.TargetPorts == nil {
			break
		}

		return e.complexity.App.TargetPorts(childComplexity), true

	case "App.updatedAt":
		if e.complexity.App.UpdatedAt == nil {
			break
		}

		return e.complexity.App.UpdatedAt(childComplexity), true

	case "App.volumes":
		if e.complexity.App.Volumes == nil {
			break
		}

		return e.complexity.App.Volumes(childComplexity), true

	case "AppInstances.running":
		if e.complexity.AppInstances.Running == nil {
			break
		}

		return e.complexity.AppInstances.Running(childComplexity), true

	case "AppInstances.total":
		if e.complexity.AppInstances.Total == nil {
			break
		}

		return e.complexity.AppInstances.Total(childComplexity), true

	case "BoundVolume.source":
		if e.complexity.BoundVolume.Source == nil {
			break
		}

		return e.complexity.BoundVolume.Source(childComplexity), true

	case "BoundVolume.target":
		if e.complexity.BoundVolume.Target == nil {
			break
		}

		return e.complexity.BoundVolume.Target(childComplexity), true

	case "ClusterInfo.createdAt":
		if e.complexity.ClusterInfo.CreatedAt == nil {
			break
		}

		return e.complexity.ClusterInfo.CreatedAt(childComplexity), true

	case "ClusterInfo.id":
		if e.complexity.ClusterInfo.ID == nil {
			break
		}

		return e.complexity.ClusterInfo.ID(childComplexity), true

	case "ClusterInfo.joinCommand":
		if e.complexity.ClusterInfo.JoinCommand == nil {
			break
		}

		return e.complexity.ClusterInfo.JoinCommand(childComplexity), true

	case "ClusterInfo.updatedAt":
		if e.complexity.ClusterInfo.UpdatedAt == nil {
			break
		}

		return e.complexity.ClusterInfo.UpdatedAt(childComplexity), true

	case "Health.cluster":
		if e.complexity.Health.Cluster == nil {
			break
		}

		return e.complexity.Health.Cluster(childComplexity), true

	case "Health.dockerVersion":
		if e.complexity.Health.DockerVersion == nil {
			break
		}

		return e.complexity.Health.DockerVersion(childComplexity), true

	case "Health.version":
		if e.complexity.Health.Version == nil {
			break
		}

		return e.complexity.Health.Version(childComplexity), true

	case "Mutation.createApp":
		if e.complexity.Mutation.CreateApp == nil {
			break
		}

		args, err := ec.field_Mutation_createApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateApp(childComplexity, args["input"].(internal.AppInput)), true

	case "Mutation.deleteApp":
		if e.complexity.Mutation.DeleteApp == nil {
			break
		}

		args, err := ec.field_Mutation_deleteApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DeleteApp(childComplexity, args["id"].(string)), true

	case "Mutation.disablePlugin":
		if e.complexity.Mutation.DisablePlugin == nil {
			break
		}

		args, err := ec.field_Mutation_disablePlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.DisablePlugin(childComplexity, args["name"].(internal.PluginName)), true

	case "Mutation.editApp":
		if e.complexity.Mutation.EditApp == nil {
			break
		}

		args, err := ec.field_Mutation_editApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EditApp(childComplexity, args["id"].(string), args["changes"].(map[string]interface{})), true

	case "Mutation.enablePlugin":
		if e.complexity.Mutation.EnablePlugin == nil {
			break
		}

		args, err := ec.field_Mutation_enablePlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.EnablePlugin(childComplexity, args["name"].(internal.PluginName)), true

	case "Mutation.removeAppRoute":
		if e.complexity.Mutation.RemoveAppRoute == nil {
			break
		}

		args, err := ec.field_Mutation_removeAppRoute_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveAppRoute(childComplexity, args["appId"].(string)), true

	case "Mutation.restartApp":
		if e.complexity.Mutation.RestartApp == nil {
			break
		}

		args, err := ec.field_Mutation_restartApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RestartApp(childComplexity, args["id"].(string)), true

	case "Mutation.setAppEnv":
		if e.complexity.Mutation.SetAppEnv == nil {
			break
		}

		args, err := ec.field_Mutation_setAppEnv_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppEnv(childComplexity, args["appId"].(string), args["newEnv"].(map[string]interface{})), true

	case "Mutation.setAppRoute":
		if e.complexity.Mutation.SetAppRoute == nil {
			break
		}

		args, err := ec.field_Mutation_setAppRoute_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.SetAppRoute(childComplexity, args["appId"].(string), args["route"].(*internal.RouteInput)), true

	case "Mutation.startApp":
		if e.complexity.Mutation.StartApp == nil {
			break
		}

		args, err := ec.field_Mutation_startApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StartApp(childComplexity, args["id"].(string)), true

	case "Mutation.stopApp":
		if e.complexity.Mutation.StopApp == nil {
			break
		}

		args, err := ec.field_Mutation_stopApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.StopApp(childComplexity, args["id"].(string)), true

	case "Mutation.upgradeApp":
		if e.complexity.Mutation.UpgradeApp == nil {
			break
		}

		args, err := ec.field_Mutation_upgradeApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.UpgradeApp(childComplexity, args["id"].(string)), true

	case "Plugin.enabled":
		if e.complexity.Plugin.Enabled == nil {
			break
		}

		return e.complexity.Plugin.Enabled(childComplexity), true

	case "Plugin.name":
		if e.complexity.Plugin.Name == nil {
			break
		}

		return e.complexity.Plugin.Name(childComplexity), true

	case "Query.getApp":
		if e.complexity.Query.GetApp == nil {
			break
		}

		args, err := ec.field_Query_getApp_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetApp(childComplexity, args["id"].(string)), true

	case "Query.getPlugin":
		if e.complexity.Query.GetPlugin == nil {
			break
		}

		args, err := ec.field_Query_getPlugin_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.GetPlugin(childComplexity, args["pluginName"].(string)), true

	case "Query.health":
		if e.complexity.Query.Health == nil {
			break
		}

		return e.complexity.Query.Health(childComplexity), true

	case "Query.listApps":
		if e.complexity.Query.ListApps == nil {
			break
		}

		args, err := ec.field_Query_listApps_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.ListApps(childComplexity, args["page"].(*int32), args["size"].(*int32), args["showHidden"].(*bool)), true

	case "Query.listPlugins":
		if e.complexity.Query.ListPlugins == nil {
			break
		}

		return e.complexity.Query.ListPlugins(childComplexity), true

	case "Route.appId":
		if e.complexity.Route.AppID == nil {
			break
		}

		return e.complexity.Route.AppID(childComplexity), true

	case "Route.createdAt":
		if e.complexity.Route.CreatedAt == nil {
			break
		}

		return e.complexity.Route.CreatedAt(childComplexity), true

	case "Route.host":
		if e.complexity.Route.Host == nil {
			break
		}

		return e.complexity.Route.Host(childComplexity), true

	case "Route.path":
		if e.complexity.Route.Path == nil {
			break
		}

		return e.complexity.Route.Path(childComplexity), true

	case "Route.traefikRule":
		if e.complexity.Route.TraefikRule == nil {
			break
		}

		return e.complexity.Route.TraefikRule(childComplexity), true

	case "Route.updatedAt":
		if e.complexity.Route.UpdatedAt == nil {
			break
		}

		return e.complexity.Route.UpdatedAt(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputAppInput,
		ec.unmarshalInputBoundVolumeInput,
		ec.unmarshalInputRouteInput,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Query(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(parsedSchema), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(parsedSchema, parsedSchema.Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "api/models.graphqls", Input: `"""
The info about the docker swarm if the host running miasma is apart of one.
"""
type ClusterInfo {
  id: String!
  joinCommand: String!
  createdAt: Time!
  updatedAt: Time!
}

type Health {
  "Miasma server's current version."
  version: String!
  "The version of docker running on the host, or null if docker is not running."
  dockerVersion: String!
  "The cluster versioning and information, or ` + "`" + `null` + "`" + ` if not apart of a cluster."
  cluster: ClusterInfo
}

type BoundVolume {
  "The path inside the container that the data is served from."
  target: String!
  "The volume name or directory on the host that the data is stored in."
  source: String!
}

type App {
  id: ID!
  createdAt: Time!
  updatedAt: Time!
  name: String!
  "Whether or not the application is managed by the system. You cannot edit or delete system apps."
  system: Boolean!
  group: String
  "The image and tag the application runs."
  image: String!
  """
  The currently running image digest (hash). Used internally when running
  applications instead of the tag because the when a new image is pushed, the
  tag stays the same but the digest changes.
  """
  imageDigest: String!
  "Whether or not the app is returned during regular requests."
  hidden: Boolean!
  "If the app has a route and the traefik plugin is enabled, this is it's config."
  route: Route
  "If the app has a route and the traefik plugin is enabled, this is a simple representation of it."
  simpleRoute: String
  "The environment variables configured for this app."
  env: Map
  "Whether or not the application is running, stopped, or starting up."
  status: String!
  "The number of instances running vs what should be running."
  instances: AppInstances!
  """
  The ports that the app is listening to inside the container. If no target
  ports are specified, then the container should respect the ` + "`" + `PORT` + "`" + ` env var.
  """
  targetPorts: [Int!]
  """
  The ports that you access the app through in the swarm. This field can, and
  should be left empty. Miasma automatically manages assigning published ports
  between 3001-4999. If you need to specify a port, make sure it's outside that
  range or the port has not been taken. Plugins have set ports starting with
  4000, so avoid 4000-4020 if you want to add a plugin at a later date.

  If these ports are ever cleared, the app will continue using the same ports it
  was published to before, so that the ports don't change unnecessarily. If you
  removed it to clear a port for another app/plugin, make sure to restart the
  app and a new, random port will be allocated for the app, freeing the old
  port.
  """
  publishedPorts: [Int!]
  """
  The placement constraints specifying which nodes the app will be ran on. Any
  valid value for the [` + "`" + `--constraint` + "`" + ` flag](https://docs.docker.com/engine/swarm/services/#placement-constraints)
  is valid item in this list.
  """
  placement: [String!]
  "Volume bindings for the app."
  volumes: [BoundVolume!]
  """
  A list of other apps that the service communicates with using their service
  name and docker's internal DNS. Services don't have to be two way; only the
  service that accesses the other needs the other network added.
  """
  networks: [String!]
  command: [String!]
}

input BoundVolumeInput {
  target: String!
  source: String!
}

input AppInput {
  name: String!
  image: String!
  group: String
  hidden: Boolean
  targetPorts: [Int!]
  publishedPorts: [Int!]
  placement: [String!]
  volumes: [BoundVolumeInput!]
  networks: [String!]
  command: [String!]
}

input AppChanges {
  name: String
  image: String
  group: String
  hidden: Boolean
  targetPorts: [Int!]
  publishedPorts: [Int!]
  placement: [String!]
  volumes: [BoundVolumeInput!]
  networks: [String!]
  command: [String!]
}

type Plugin {
  name: PluginName!
  "Whether or not the plugin has been enabled."
  enabled: Boolean!
}

type Route {
  appId: ID!
  createdAt: Time!
  updatedAt: Time!
  host: String
  path: String
  traefikRule: String
}

input RouteInput {
  host: String
  path: String
  traefikRule: String
}

type AppInstances {
  running: Int!
  total: Int!
}

enum PluginName {
  TRAEFIK
}
`, BuiltIn: false},
	{Name: "api/mutations.graphqls", Input: `type Mutation {
  "Create and start a new app"
  createApp(input: AppInput!): App!
  "Edit app metadata unrelated to how the container(s) that are run"
  editApp(id: ID!, changes: AppChanges!): App!
  "Stop and delete an app"
  deleteApp(id: ID!): App!
  "Start a stopped app"
  startApp(id: ID!): App!
  "Stop a running app"
  stopApp(id: ID!): App!
  "Stop and restart an app"
  restartApp(id: ID!): App!
  "Pull the latest version of the app's image and then restart"
  upgradeApp(id: ID!): App!

  "Install one of Miasma's plugins"
  enablePlugin(name: PluginName!): Plugin!
  "Disable one of Miasma's plugins"
  disablePlugin(name: PluginName!): Plugin!

  "Only available when the 'router' plugin is enabled"
  setAppEnv(appId: ID!, newEnv: Map): Map

  "Only available when the 'router' plugin is enabled"
  setAppRoute(appId: ID!, route: RouteInput): Route
  "Only available when the 'router' plugin is enabled"
  removeAppRoute(appId: ID!): Route
}
`, BuiltIn: false},
	{Name: "api/queries.graphqls", Input: `type Query {
  health: Health

  listApps(page: Int = 1, size: Int = 10, showHidden: Boolean): [App!]!
  getApp(id: ID!): App!

  listPlugins: [Plugin!]!
  getPlugin(pluginName: String!): Plugin!
}
`, BuiltIn: false},
	{Name: "api/scalars.graphqls", Input: `scalar Map
scalar Time
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
