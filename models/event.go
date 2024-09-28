package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	ID        string      `json:"id"`
	Type      EventType   `json:"type"`
	Actor     User        `json:"actor"`
	Repo      EventRepo   `json:"repo"`
	Payload   interface{} `json:"payload"`
	CreatedAt time.Time   `json:"created_at"`
}

func PrintEvents(events []Event) error {
	for _, event := range events {
		err := event.print()
		if err != nil {
			return err
		}
	}
	return nil
}

func (e Event) print() error {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s - ", e.CreatedAt.Format("January 2, 2006 at 3:04 PM UTC")))
	switch e.Type {
	case EventTypeWatch:
		if err := e.formatWatchEvent(&buffer); err != nil {
			return err
		}
	case EventTypeCommitComment:
		if err := e.formatCommitCommentEvent(&buffer); err != nil {
			return err
		}
	case EventTypeCreate:
		if err := e.formatCreateEvent(&buffer); err != nil {
			return err
		}
	case EventTypeDelete:
		if err := e.formatDeleteEvent(&buffer); err != nil {
			return err
		}
	case EventTypeFork:
		if err := e.formatForkEvent(&buffer); err != nil {
			return err
		}
	case EventTypeGollum:
		if err := e.formatGollumEvent(&buffer); err != nil {
			return err
		}
	case EventTypeIssueComment:
		if err := e.formatIssueCommentEvent(&buffer); err != nil {
			return err
		}
	case EventTypeIssues:
		if err := e.formatIssuesEvent(&buffer); err != nil {
			return err
		}
	case EventTypeMember:
		if err := e.formatMemberEvent(&buffer); err != nil {
			return err
		}
	case EventTypePublic:
		if err := e.formatPublicEvent(&buffer); err != nil {
			return err
		}
	case EventTypePullRequest:
		if err := e.formatPullRequestEvent(&buffer); err != nil {
			return err
		}
	case EventTypePullRequestReview:
		if err := e.formatPullRequestReviewEvent(&buffer); err != nil {
			return err
		}
	case EventTypePullRequestReviewComment:
		if err := e.formatPullRequestReviewCommentEvent(&buffer); err != nil {
			return err
		}
	case EventTypePullRequestReviewThread:
		if err := e.formatPullRequestReviewThreadEvent(&buffer); err != nil {
			return err
		}
	case EventTypePush:
		if err := e.formatPushEvent(&buffer); err != nil {
			return err
		}
	case EventTypeRelease:
		if err := e.formatReleaseEvent(&buffer); err != nil {
			return err
		}
	}
	return nil
}

func (e Event) formatWatchEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[WatchEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s watching for repo %s", p.Action, e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatCommitCommentEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[CommitCommentEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s a comment on commit '%s'", p.Action, p.Comment.Body))
	buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatCreateEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[CreateEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString("created ")
	switch p.RefType {
	case "branch", "tag":
		buffer.WriteString(fmt.Sprintf("%s %s", p.RefType, p.Ref))
		buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	case "repository":
		buffer.WriteString(fmt.Sprint(p.RefType))
		buffer.WriteString(fmt.Sprintf(" %s", e.Repo.Name))
	}
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatDeleteEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[DeleteEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString("deleted ")
	buffer.WriteString(fmt.Sprint(p.RefType))
	buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatForkEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[ForkEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("forked %s to %s", p.Forkee.Name, e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatGollumEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[GollumEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("created/updated following pages in repo %s:\n", e.Repo.Name))
	for _, v := range p.Pages {
		buffer.WriteString(fmt.Sprintf("%s page %s\n", v.Action, v.Name))
	}
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatIssueCommentEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[IssueCommentEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s a comment on issue #%d", p.Action, p.Issue.Number))
	buffer.WriteString(fmt.Sprintf(" '%s' in repo %s:\n", p.Issue.Title, e.Repo.Name))
	buffer.WriteString(fmt.Sprintf("'%s'", p.Comment.Body))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatIssuesEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[IssuesEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s issue #%d '%s'", p.Action, p.Issue.Number, p.Issue.Title))
	buffer.WriteString(fmt.Sprintf(" '%s' in repo %s", p.Issue.Title, e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatMemberEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[MemberEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s %s in repo %s", p.Action, p.Member.Login, e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPublicEvent(buffer *bytes.Buffer) error {
	buffer.WriteString(fmt.Sprintf("made repo %s public", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPullRequestEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[PullRequestEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s pull request #%d '%s'", p.Action, p.PullRequest.Number, p.PullRequest.Title))
	buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPullRequestReviewEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[PullRequestReviewEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s pull request #%d '%s'", p.Action, p.PullRequest.Number, p.PullRequest.Title))
	buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPullRequestReviewCommentEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[PullRequestReviewCommentEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s a comment on pull request #%d '%s'", p.Action, p.PullRequest.Number, p.PullRequest.Title))
	buffer.WriteString(fmt.Sprintf(" in repo %s:\n", e.Repo.Name))
	buffer.WriteString(fmt.Sprintf("'%s'", p.Comment.Body))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPullRequestReviewThreadEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[PullRequestReviewThreadEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s pull request review thread #%d '%s'", p.Action, p.PullRequest.Number, p.PullRequest.Title))
	buffer.WriteString(fmt.Sprintf(" in repo %s", e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatPushEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[PushEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("pushed %d ", p.Size))
	if p.Size == 1 {
		buffer.WriteString(fmt.Sprintf("commit to repo %s", e.Repo.Name))
	} else {
		buffer.WriteString(fmt.Sprintf("commits to repo %s", e.Repo.Name))

	}
	fmt.Println(buffer.String())
	return nil
}

func (e Event) formatReleaseEvent(buffer *bytes.Buffer) error {
	p, err := payloadConvertion[ReleaseEventPayload](e.Payload)
	if err != nil {
		return err
	}
	buffer.WriteString(fmt.Sprintf("%s release %s in repo %s", p.Action, p.Release.Name, e.Repo.Name))
	fmt.Println(buffer.String())
	return nil
}

func payloadConvertion[T any](payload interface{}) (*T, error) {
	var p T
	payloadRaw, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(payloadRaw, &p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

type EventType string

const (
	EventTypeWatch                    EventType = "WatchEvent"
	EventTypeCommitComment            EventType = "CommitCommentEvent"
	EventTypeCreate                   EventType = "CreateEvent"
	EventTypeDelete                   EventType = "DeleteEvent"
	EventTypeFork                     EventType = "ForkEvent"
	EventTypeGollum                   EventType = "GollumEvent"
	EventTypeIssueComment             EventType = "IssueCommentEvent"
	EventTypeIssues                   EventType = "IssuesEvent"
	EventTypeMember                   EventType = "MemberEvent"
	EventTypePublic                   EventType = "PublicEvent"
	EventTypePullRequest              EventType = "PullRequestEvent"
	EventTypePullRequestReview        EventType = "PullRequestReviewEvent"
	EventTypePullRequestReviewComment EventType = "PullRequestReviewCommentEvent"
	EventTypePullRequestReviewThread  EventType = "PullRequestReviewThreadEvent"
	EventTypePush                     EventType = "PushEvent"
	EventTypeRelease                  EventType = "ReleaseEvent"
)

type Comment struct {
	Body string `json:"body"`
}

type Issue struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type User struct {
	Login string `json:"login"`
}

type Release struct {
	Name string `json:"name"`
}

type EventRepo struct {
	Name string `json:"name"`
}

type CommitCommentEventPayload struct {
	Action  string  `json:"action"`
	Comment Comment `json:"comment"`
}

type CreateEventPayload struct {
	Ref     string `json:"ref"`
	RefType string `json:"ref_type"`
}

type DeleteEventPayload struct {
	RefType string `json:"ref_type"`
}

type ForkEventPayload struct {
	Forkee ForkEventPayloadRepo `json:"forkee"`
}

type ForkEventPayloadRepo struct {
	Name string `json:"full_name"`
}

type GollumEventPayload struct {
	Pages []GollumEventPayloadPage `json:"pages"`
}

type GollumEventPayloadPage struct {
	Name   string `json:"page_name"`
	Action string `json:"action"`
}

type IssueCommentEventPayload struct {
	Action  string  `json:"action"`
	Issue   Issue   `json:"issue"`
	Comment Comment `json:"comment"`
}

type IssuesEventPayload struct {
	Action   string `json:"action"`
	Issue    Issue  `json:"issue"`
	Assignee User   `json:"assignee"`
}

type MemberEventPayload struct {
	Action string `json:"action"`
	Member User   `json:"member"`
}

type PublicEventPayload struct{}

type PullRequestEventPayload struct {
	Action      string      `json:"action"`
	Number      int         `json:"number"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Title  string `json:"title"`
	Number int    `json:"number"`
}

type PullRequestReviewEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequestReviewCommentEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
	Comment     Comment     `json:"comment"`
}

type PullRequestReviewThreadEventPayload struct {
	Action      string      `json:"action"`
	PullRequest PullRequest `json:"pull_request"`
}

type PushEventPayload struct {
	Size int `json:"size"`
}

type ReleaseEventPayload struct {
	Action  string  `json:"action"`
	Release Release `json:"release"`
}

type WatchEventPayload struct {
	Action string `json:"action"`
}
