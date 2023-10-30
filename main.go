package main

import (
    "log"
	"net/http"
	"fmt"
	"github.com/go-playground/webhooks/v6/github"
)

type GithubEventHandler struct {

}

func (g *GithubEventHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}

ghEventList := []github.Event{ 
	github.CheckRunEvent, github.CheckSuiteEvent,
	github.CommitCommentEvent, github.CreateEvent, github.DeleteEvent,
	github.DependabotAlertEvent,
	github.DeployKeyEvent, github.DeploymentEvent, github.DeploymentStatusEvent,
	github.ForkEvent, github.GollumEvent,
	github.InstallationEvent,InstallationRepositoriesEvent, github.IntegrationInstallationEvent, github.IntegrationInstallationRepositoriesEvent,
	github.IssueCommentEvent, github.IssuesEvent, github.LabelEvent,
	github.MemberEvent, github.MembershipEvent,
	github.MilestoneEvent, github.MetaEvent,
	github.OrganizationEvent, github.OrgBlockEvent,
	github.PageBuildEvent, github.PingEvent,
	github.ProjectCardEvent, github.ProjectColumnEvent, github.ProjectEvent,
	github.PublicEvent, github.PullRequestEvent, github.PullRequestReviewEvent, github.PullRequestReviewCommentEvent, github.PushEvent
	github.ReleaseEvent,
	github.RepositoryEvent, github.RepositoryVulnerabilityAlertEvent,
	github.SecurityAdvisoryEvent, github.StatusEvent,
	github.TeamEvent, github.TeamAddEvent,
	github.	WatchEvent, github.WorkflowDispatchEvent,WorkflowJobEvent, github.WorkflowRunEvent,
	github.GitHubAppAuthorizationEvent
}


func main() {
	hook, err := github.New(github.Options.Secret(""))
	if err!=nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		payload, err := hook.Parse(r, ghEventList...)
		if err != nil {
			if err == github.ErrEventNotFound {
				// ok event wasn't one of the ones asked to be parsed
			}
		}
		switch payload.(type) {

		case github.ReleasePayload:
			release := payload.(github.ReleasePayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", release)

		case github.PullRequestPayload:
			pullRequest := payload.(github.PullRequestPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", pullRequest)

		case github.PushPayload:
			pushEvent := payload.(github.PushPayload)
			// Do whatever you want from here...
			fmt.Printf("%+v\n", pushEvent)
		default:
			fmt.Print("ns hjxyel")
		}
	})
	http.ListenAndServe(":3000", nil)

}
