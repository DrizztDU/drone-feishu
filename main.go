package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// Version set at compile-time
var (
	Version string
)

func action(ctx *cli.Context) error {
	p := Plugin{
		Config: Config{
			Debug:        ctx.Bool("debug"),
			Webhook:      ctx.StringSlice("webhook"),
			Message:      ctx.String("message"),
			TemplateFile: ctx.String("template.file"),
		},
		Repo: Repo{
			FullName:  ctx.String("repo.fullname"),
			Link:      ctx.String("repo.link"),
			Namespace: ctx.String("repo.namespace"),
			Name:      ctx.String("repo.name"),
		},
		Commit: Commit{
			Sha:          ctx.String("commit.sha"),
			Ref:          ctx.String("commit.ref"),
			Branch:       ctx.String("commit.branch"),
			Link:         ctx.String("commit.link"),
			AuthorName:   ctx.String("commit.author.name"),
			AuthorEmail:  ctx.String("commit.author.email"),
			AuthorAvatar: ctx.String("commit.author.avatar"),
			Message:      preprocessCommitMessage(ctx.String("commit.message")),
		},
		Build: Build{
			Tag:      ctx.String("build.tag"),
			Number:   ctx.Int("build.number"),
			Event:    ctx.String("build.event"),
			Status:   ctx.String("build.status"),
			Link:     ctx.String("build.link"),
			Started:  ctx.Int64("build.started"),
			Finished: ctx.Int64("build.finished"),
			PR:       ctx.String("pull.request"),
			DeployTo: ctx.String("deploy.to"),
		},
	}

	return p.Exec()
}

func main() {
	app := cli.App{
		Name:    "feishu plugin",
		Usage:   "feishu plugin",
		Action:  action,
		Version: Version,
	}

	app.Flags = []cli.Flag{
		// config
		&cli.BoolFlag{
			Name:    "debug",
			Usage:   "enable debug message",
			Value:   false,
			EnvVars: []string{"PLUGIN_DEBUG", "DEBUG"},
		},
		&cli.StringSliceFlag{
			Name:    "webhook",
			Usage:   "webhook url",
			EnvVars: []string{"PLUGIN_WEBHOOK", "FEISHU_WEBHOOK"},
		},
		&cli.StringFlag{
			Name:    "message",
			Usage:   "message",
			EnvVars: []string{"PLUGIN_MESSAGE", "FEISHU_MESSAGE"},
		},
		&cli.StringFlag{
			Name:    "template.file",
			Usage:   "template file path",
			EnvVars: []string{"PLUGIN_TEMPLATE_FILE", "FEISHU_TEMPLATE_FILE"},
		},
		// repo
		&cli.StringFlag{
			Name:    "repo.fullname",
			Usage:   "repository owner and repository name",
			EnvVars: []string{"DRONE_REPO"},
		},
		&cli.StringFlag{
			Name:    "repo.link",
			Usage:   "repository link",
			EnvVars: []string{"DRONE_REPO_LINK"},
		},
		&cli.StringFlag{
			Name:    "repo.namespace",
			Usage:   "repository namespace",
			EnvVars: []string{"DRONE_REPO_OWNER,DRONE_REPO_NAMESPACE"},
		},
		&cli.StringFlag{
			Name:    "repo.name",
			Usage:   "repository name",
			EnvVars: []string{"DRONE_REPO_NAME"},
		},
		// commit
		&cli.StringFlag{
			Name:    "commit.sha",
			Usage:   "git commit sha",
			EnvVars: []string{"DRONE_COMMIT_SHA"},
		},
		&cli.StringFlag{
			Name:    "commit.ref",
			Usage:   "git commit ref",
			EnvVars: []string{"DRONE_COMMIT_REF"},
		},
		&cli.StringFlag{
			Name:    "commit.branch",
			Value:   "master",
			Usage:   "git commit branch",
			EnvVars: []string{"DRONE_COMMIT_BRANCH"},
		},
		&cli.StringFlag{
			Name:    "commit.link",
			Usage:   "git commit link",
			EnvVars: []string{"DRONE_COMMIT_LINK"},
		},
		&cli.StringFlag{
			Name:    "commit.author.name",
			Usage:   "git author name",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR"},
		},
		&cli.StringFlag{
			Name:    "commit.author.email",
			Usage:   "git author email",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_EMAIL"},
		},
		&cli.StringFlag{
			Name:    "commit.author.avatar",
			Usage:   "git author avatar",
			EnvVars: []string{"DRONE_COMMIT_AUTHOR_AVATAR"},
		},
		&cli.StringFlag{
			Name:    "commit.message",
			Usage:   "commit message",
			EnvVars: []string{"DRONE_COMMIT_MESSAGE"},
		},
		// build
		&cli.StringFlag{
			Name:    "build.tag",
			Usage:   "build tag",
			EnvVars: []string{"DRONE_TAG"},
		},
		&cli.StringFlag{
			Name:    "build.event",
			Value:   "push",
			Usage:   "build event",
			EnvVars: []string{"DRONE_BUILD_EVENT"},
		},
		&cli.IntFlag{
			Name:    "build.number",
			Usage:   "build number",
			EnvVars: []string{"DRONE_BUILD_NUMBER"},
		},
		&cli.StringFlag{
			Name:    "build.status",
			Usage:   "build status",
			Value:   "success",
			EnvVars: []string{"DRONE_BUILD_STATUS"},
		},
		&cli.StringFlag{
			Name:    "build.link",
			Usage:   "build link",
			EnvVars: []string{"DRONE_BUILD_LINK"},
		},
		&cli.Int64Flag{
			Name:    "build.started",
			Usage:   "build started",
			EnvVars: []string{"DRONE_BUILD_STARTED"},
		},
		&cli.Int64Flag{
			Name:    "build.finished",
			Usage:   "build finished",
			EnvVars: []string{"DRONE_BUILD_FINISHED"},
		},
		&cli.StringFlag{
			Name:    "pull.request",
			Usage:   "pull request",
			EnvVars: []string{"DRONE_PULL_REQUEST"},
		},
		&cli.StringFlag{
			Name:    "deploy.to",
			Usage:   "Provides the target deployment environment for the running build. This value is only available to promotion and rollback pipelines.",
			EnvVars: []string{"DRONE_DEPLOY_TO"},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
