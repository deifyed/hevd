# Decisions

## Motivation

- **Anyone can quickly and easily make throw away old decisions when they are no longer relevant, i.e. premise has changed**
- Anyone at any time can understand the decisions made in the project

## How

- Document any meaningful decisions made in the project using this [template](./0000-template.md).
- Make extensive use of links. Decisions should be quick to read and easily understood, so links to other documents are
    encouraged when including information.

### Adding a new decision

1. `cp 0000-template.md 0001-my-decision.md`. The index should be the next available number and the name should reflect the
    decision.
2. Fill out the template, commit, push your proposal to a branch, open a PR
3. The PR will be considered the RFC, and will be considered accepted when and if it is merged

### Updating a decision

1. Make changes to a relevant existing decision
2. Commit, push your proposal to a branch, open a PR
3. The PR will be considered the RFC, and will be considered accepted when and if it is merged

### Removing a decision

1. Add an `-OBSOLETE` suffix to the decision file name
2. Clear the content of the decision file
3. Commit, push your proposal to a branch, open a PR
4. The PR will be considered the RFC, and will be considered obsolete when and if it is merged
