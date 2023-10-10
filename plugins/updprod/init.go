package updprod

import (
	"net/http"

	"github.com/infinitybotlist/eureka/crypto"
	"github.com/infinitybotlist/sysmanage-web/core/plugins"
	"github.com/infinitybotlist/sysmanage-web/plugins/actions"
	"github.com/infinitybotlist/sysmanage-web/types"
)

const ID = "updprod"

var gitRepo string
var githubUsername string
var vercelDeployHook string
var githubPat string

func InitPlugin(c *types.PluginConfig) error {
	opts, err := plugins.GetConfig(c.Name)

	if err != nil {
		return err
	}

	gitRepo, err = opts.GetString("git_repo")

	if err != nil {
		return err
	}

	githubUsername, err = opts.GetString("github_username")

	if err != nil {
		return err
	}

	githubPat, err = opts.GetString("github_pat")

	if err != nil {
		return err
	}

	vercelDeployHook, err = opts.GetString("vercel_deploy_hook")

	if err != nil {
		return err
	}

	actions.RegisterActions(&actions.Action{
		Name:        "Update Production",
		Description: "Bump the Infinity-Next repository with the latest master branch.",
		Handler: func(*actions.ActionContext) (*actions.ActionResponse, error) {
			taskId := crypto.RandString(128)

			go updateProdTask(taskId)

			return &actions.ActionResponse{
				StatusCode: http.StatusNoContent,
				TaskID:     taskId,
			}, nil
		},
	})

	return nil
}
