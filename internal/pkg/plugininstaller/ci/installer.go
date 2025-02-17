package ci

import (
	"errors"

	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/pkg/util/git"
)

func PushCIFiles(options plugininstaller.RawOptions) error {
	opts, err := NewOptions(options)
	if err != nil {
		return err
	}
	// 1. get git content by config
	gitMap, err := opts.buildGitMap()
	if err != nil {
		return err
	}
	//3. init git client
	gitClient, err := opts.ProjectRepo.NewClient()
	if err != nil {
		return err
	}
	//4. push ci files to git repo
	_, err = gitClient.PushLocalFileToRepo(&git.CommitInfo{
		CommitMsg:    defaultCommitMsg,
		CommitBranch: defaultBranch,
		GitFileMap:   gitMap,
	}, true)
	return err
}

func DeleteCIFiles(options plugininstaller.RawOptions) error {
	opts, err := NewOptions(options)
	if err != nil {
		return err
	}
	// 1. get git content by config
	gitMap, err := opts.buildGitMap()
	if err != nil {
		return err
	}
	if len(gitMap) == 0 {
		return errors.New("can't get valid ci files, please check your config")
	}
	//2. init git client
	gitClient, err := opts.ProjectRepo.NewClient()
	if err != nil {
		return err
	}
	//3. delete ci files from git repo
	commitInfo := &git.CommitInfo{
		GitFileMap: gitMap,
		CommitMsg:  deleteCommitMsg,
	}
	return gitClient.DeleteFiles(commitInfo)
}
