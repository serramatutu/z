---
layout: default
title: Contribute
permalink: /contribute/
---

# Contribute

We're glad you want to contribute to z! Before you get started, it's important you read this first.

## License agreement

By contributing to z in the form of code, documentation or ideas/suggestions, you agree that your contributions will be published under the [MIT License](https://github.com/serramatutu/z/blob/main/LICENSE). For this reason, thay may be made public and distributed in accordance with the license's terms.

## Contribution types

Your contribution will be categorized as a:
- **Major contribution:**
  - Introduces breaking changes, with public API behavior modifications in a non-backwards compatible manner;
  - Introduces a new feature such as a new command;
  - Introduces significant changes which require a huge modifications to the codebase.
- **Minor contribution:**
  - Fixes a bug;
  - Changes a help message for clarification;
  - Changes documentation;
  - Other changes.

If you intend on working on a **major contribution**, make sure you [file an issue](https://github.com/serramatutu/z/issues/new) to discuss your ideas first. Before you start to work on it, wait until it's approved. You may also indicate you're working on an [existing issue](https://github.com/serramatutu/z/issues).

**Minor contributions** don't require issues. However, if you're working on something somebody has reported through an issue, make sure to link it in your pull request.

## Working on a contribution

It's fairly easy to submit a pull request.

### 1. Fork our repository on GitHub and create a new branch

Create a fork of z by following [this tutorial](https://guides.github.com/activities/forking/).

Once you've clone the repo, install our Git hooks by running:

```
make githooks
```

Create a new branch. Give it a descriptive name that indicates clearly what you're planning to do.

```
git checkout -b issue-number
```

We ask you to only work on one issue at a time. This makes it easier for us to sort through contributions.

### 2. Work on your changes

That's pretty straight forward. Just open your favorite text editor and start coding! If you're working on a bug of feature, add comprehensive unit tests for it. Untested pull requests _will not be merged_. Make sure all tests are passing by running:
```
make test
```

If you're introducing a new feature or an API change, document your changes in the help files and in the documentation.

Once you're done, commit your changes:
```
git add .
git commit
```

Give your commit a descriptive but short title. It should have the following format:
```
tag: short description (fixes #issue-number)

Add a longer description if you think it's needed.
```

The first line must have this specific format because that's used by our build pipelines to generate changelogs. `tag` must be one of:
- `docs`: for documentation and help file changes;
- `bug`: for bug fixes;
- `feature`: for new features;
- `breaking`: for backwards-incompatible changes;
- `pipeline`: for changes to our build/test pipelines;
- `refactor`: for internal changes that will not affect the public API;
- `test`: for adding tests.

### 3. Ensure you're in sync with the global project
Before you file a pull request, make sure you've merged the most recent changes to the `main` branch of the global project. Do this by running:
```
git fetch upstream
git rebase upstream/main
```

### 4. Double-check everything
Once you've rebased, check if:
- All tests pass;
- Your code is properly formatted with gofmt;
- Your commit message follows the desired format;
- You've added tests for everything you've changed;
- You've changed the documentation if needed.

### 5. Push your changes and open a pull request
Push your changes to your fork by running
```
git push origin <name-of-your-branch>
```

If the push doesn't work because of the rebase, run a force push
```
git push --force origin <name-of-your-branch>
```

Finally, create a pull request by following [GitHub's documentation](https://help.github.com/articles/creating-a-pull-request).

### 6. Followup
Once you've created your pull request, monitor the Github Actions pipelines for it. We won't merge it while its build/test pipelines are failing.

We may ask you to change certain parts of your contribution, such as:
1. Wrong commit messages. Fix them by running `git commit --amend` and force pushing;
2. Code related stuff such as formatting, tests or even logic. Fix them by adding new commits and pushing them;
3. Synchronization issues with the `main` branch. Fix them by rebasing again. This may require a force push.

### 7. Merge
Once everything is set, we'll merge your pull request and close it for you :)
