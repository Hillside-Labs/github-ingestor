package internal

import (
	"fmt"
	"github-ingestor-go/pkg/memphis"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/github"
)

var ghEventList []github.Event

func initEventList() {
	ghEventList = []github.Event{
		github.CheckRunEvent, github.CheckSuiteEvent,
		github.CommitCommentEvent, github.CreateEvent, github.DeleteEvent,
		github.DependabotAlertEvent,
		github.DeployKeyEvent, github.DeploymentEvent, github.DeploymentStatusEvent,
		github.ForkEvent, github.GollumEvent,
		github.InstallationEvent, github.InstallationRepositoriesEvent, github.IntegrationInstallationEvent, github.IntegrationInstallationRepositoriesEvent,
		github.IssueCommentEvent, github.IssuesEvent, github.LabelEvent,
		github.MemberEvent, github.MembershipEvent,
		github.MilestoneEvent, github.MetaEvent,
		github.OrganizationEvent, github.OrgBlockEvent,
		github.PageBuildEvent, github.PingEvent,
		github.ProjectCardEvent, github.ProjectColumnEvent, github.ProjectEvent,
		github.PublicEvent, github.PullRequestEvent, github.PullRequestReviewEvent, github.PullRequestReviewCommentEvent, github.PushEvent,
		github.ReleaseEvent,
		github.RepositoryEvent, github.RepositoryVulnerabilityAlertEvent,
		github.SecurityAdvisoryEvent, github.StatusEvent,
		github.TeamEvent, github.TeamAddEvent,
		github.WatchEvent, github.WorkflowDispatchEvent, github.WorkflowJobEvent, github.WorkflowRunEvent,
		github.GitHubAppAuthorizationEvent,
	}
}

type EventHandler struct {
	hook     *github.Webhook
	log      *log.Logger
	producer *memphis.GithubProducer
}

func NewEventHandler(hook *github.Webhook, l *log.Logger, pc *memphis.ProducerConfig) *EventHandler {

	initEventList()

	return &EventHandler{hook: hook, log: l, producer: memphis.NewProducer(pc, l)}
}

func (e *EventHandler) HandleEvents(c *gin.Context) {
	payload, err := e.hook.Parse(c.Request, ghEventList...)
	if err != nil {
		if err == github.ErrEventNotFound {
			// ok event wasn't one of the ones asked to be parsed
		} else {
			e.log.Println(err.Error())
		}
	}

	switch payload := payload.(type) {

	case github.CheckRunPayload:
		fmt.Printf("%+v\n", payload)

	case github.CheckSuitePayload:
		fmt.Printf("%+v\n", payload)

	case github.CommitCommentPayload:
		fmt.Printf("%+v\n", payload)

	case github.CreatePayload:
		fmt.Printf("%+v\n", payload)

	case github.DeletePayload:
		fmt.Printf("%+v\n", payload)

	case github.DependabotAlertPayload:
		fmt.Printf("%+v\n", payload)

	case github.DeployKeyPayload:
		fmt.Printf("%+v\n", payload)

	case github.DeploymentPayload:
		fmt.Printf("%+v\n", payload)

	case github.DeploymentStatusPayload:
		fmt.Printf("%+v\n", payload)

	case github.ForkPayload:
		fmt.Printf("%+v\n", payload)

	case github.GollumPayload:
		fmt.Printf("%+v\n", payload)

	case github.InstallationPayload:
		fmt.Printf("%+v\n", payload)

	case github.InstallationRepositoriesPayload:
		fmt.Printf("%+v\n", payload)

	case *github.InstallationPayload:
		fmt.Printf("%+v\n", payload)

	case *github.InstallationRepositoriesPayload:
		fmt.Printf("%+v\n", payload)

	case github.IssueCommentPayload:
		e.producer.PushEvent(payload)
		fmt.Printf("%+v\n", payload)

	case github.IssuesPayload:
		fmt.Printf("%+v\n", payload)
		e.producer.PushEvent(payload)

	case github.LabelPayload:
		fmt.Printf("%+v\n", payload)

	case github.MemberPayload:
		fmt.Printf("%+v\n", payload)

	case github.MembershipPayload:
		fmt.Printf("%+v\n", payload)

	case github.MilestonePayload:
		fmt.Printf("%+v\n", payload)

	case github.MetaPayload:
		fmt.Printf("%+v\n", payload)

	case github.OrganizationPayload:
		fmt.Printf("%+v\n", payload)

	case github.OrgBlockPayload:
		fmt.Printf("%+v\n", payload)

	case github.PageBuildPayload:
		fmt.Printf("%+v\n", payload)

	case github.PingPayload:
		fmt.Printf("%+v\n", payload)

	case github.ProjectCardPayload:
		fmt.Printf("%+v\n", payload)

	case github.ProjectColumnPayload:
		fmt.Printf("%+v\n", payload)

	case github.ProjectPayload:
		fmt.Printf("%+v\n", payload)

	case github.PublicPayload:
		fmt.Printf("%+v\n", payload)

	case github.PullRequestPayload:
		fmt.Printf("%+v\n", payload)

	case github.PullRequestReviewPayload:
		fmt.Printf("%+v\n", payload)

	case github.PullRequestReviewCommentPayload:
		fmt.Printf("%+v\n", payload)

	case github.PushPayload:
		fmt.Printf("%+v\n", payload)

	case github.ReleasePayload:
		fmt.Printf("%+v\n", payload)

	case github.RepositoryPayload:
		fmt.Printf("%+v\n", payload)

	case github.RepositoryVulnerabilityAlertPayload:
		fmt.Printf("%+v\n", payload)

	case github.SecurityAdvisoryPayload:
		fmt.Printf("%+v\n", payload)

	case github.StatusPayload:
		fmt.Printf("%+v\n", payload)

	case github.TeamPayload:
		fmt.Printf("%+v\n", payload)

	case github.TeamAddPayload:
		fmt.Printf("%+v\n", payload)

	case github.WatchPayload:
		fmt.Printf("%+v\n", payload)

	case github.WorkflowDispatchPayload:
		fmt.Printf("%+v\n", payload)

	case github.WorkflowJobPayload:
		fmt.Printf("%+v\n", payload)

	case github.WorkflowRunPayload:
		fmt.Printf("%+v\n", payload)

	case github.GitHubAppAuthorizationPayload:
		fmt.Printf("%+v\n", payload)

	default:
		fmt.Println("Unsupported payload type")

	}
}
