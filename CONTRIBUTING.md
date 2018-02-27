# Contributing to Small Workflows
There are many ways you can contribute. You can:
- Make suggestions and file bugs in [Issues](../../issues/).
- Fix issues and add features to any of the workflows with [Pull Requests](../../pulls/).

## Improving the workflows
If you wish that a workflow did something that it currently doesn't do or perhaps it does something wrong and you wish it was fixed, you can open an [issue](../../issues/new) saying what it is you want fixed or added. Alternatively, you can make the change yourself in the workflow, then export the workflow with something like [transfer.sh](https://transfer.sh) and I can take a look at it and add the changes.

If you are interested in making your own workflows, I wrote about how I make and update these workflows [here](https://wiki.nikitavoloboev.xyz/macOS/apps/alfred/making-workflows.html).

## Developing Go workflows
If you want to add features and things to workflow written in Go.

It is best to use [Alfred CLI tool](https://godoc.org/github.com/jason0x43/go-alfred/alfred) which you can install by running:

`go install github.com/jason0x43/go-alfred/alfred`

You can then clone this repository. `cd` to the folder of the workflow you want to extend and run: `alfred link` inside it. This will make a symbolic link of the `workflow` directory.

You can then make changes to the code and after run `alfred build` to build the go binary to `workflow` directory. Which you can then call from inside Alfred [Script Filters](https://www.alfredapp.com/help/workflows/inputs/script-filter/).

I also wrote about my own process in making Alfred workflows [here](https://wiki.nikitavoloboev.xyz/macOS/apps/alfred/making-workflows.html).

## Submitting a Pull Request
Please go through [existing issues](../../issues/) and [pull requests](../../pulls/) to check if somebody else is already working on the issue.

#### Thank you for taking the time to contribute! 💜