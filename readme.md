# Small Workflows [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> Small [Alfred](https://www.alfredforum.com/) workflows I use that don't warrant a GitHub repository of their own

## Contents
- [Search Files](./search-files) - A lot of [file filters](https://www.alfredapp.com/help/workflows/inputs/file-filter/) for various apps I use.
- [Objects library](./objects-library) - Useful premade objects for workflows.
- [Keyboard turn on/off](./keyboard-on-off) - Turn your keboard on and off.
- [Focus](./focus) - Start [Focus](https://heyfocus.com) blocking for some time that you specify.
- [Dictionary search](./dictionary-search) - Search through dictionary.
- [Useful utilities](./utilities) - Only has one utilty, to search selected text in Alfred.
- [Search for Content](./search-for-content) - Two actions to search through the insides of PDF and MindNode documents.
- [Months numbers](./months-numbers) - Search for a month and copy the month number to your clipboard.
- [WiFi Tools](/wifi) - Check wifi connection / Restart wifi / Toggle it on/off.
- [Imgur album Download](./imgur-download) - Download Imgur Album currently open in your browser to location you specify.
- [Karabiner XML Reload](./karabiner-reload) - Reload [Old Karabiner](https://github.com/tekezo/Karabiner) configuration.
- [Clean Folders](./clean-folders) - Trash items from Desktop and clean certain folders like removing .alfredworkflows from ~/Downloads.
- [Folder Search](./folder-search) - Search folders from Alfred and open them in Finder/iTerm/Editor.
- [Static Searches](./static-searches) - Make searches on popular websites. I use [Web Searches](https://github.com/nikitavoloboev/alfred-web-searches) and [Searchio](https://github.com/deanishe/alfred-searchio) for all other searches.
- [Useful File Actions](./file-actions) - Various file actions I made to operate on files and folders.
- [Search Selected Text on the Web](./search-selection) - Search selected text on various websites with hotkeys.
- [Go to Reddit Subreddit](./reddit-subs) - Go to a subreddit that you specify. For searching subreddits, use [this](https://github.com/deanishe/alfred-reddit).
- [Go code -> Go Playground](./go-play) - Create [Go Playground](https://play.golang.org) from selected Go code for sharing. Can also share to [Go Play Space](https://goplay.space/).

### Alfred Commit Folders - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/commit%20folders.alfredworkflow?raw=true)
- Simple workflow that lets you search folders and then write a commit message. It would then `git add` and `git commit` all files in the folder with the message you provide. And finally it would `git push` the changes.
- This workflow assumes that the folder you select is already initialised with git as it is a file filter that searches through all the folders. It will however warn you if the folder you chose is not under git.
- You can of course modify the scope of the search or even make it that it only searches through git initialised folders but for my use case, I use this workflow when I make a change to some README file in some folder and want to quickly commit it without going to my terminal.
- For quickly editing README's of folders I use this [workflow](https://github.com/nikitavoloboev/alfred-folder-search) that searches through all folders and with a modifier key press will open the README in the markdown editor you set up. 
- There is also alternative action that will commit only the README file inside the folder with a predefined commit message. Again all the commit messages can be customised to what you prefer.

### Dash Profile Switcher - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/dash%20profile%20switch.alfredworkflow?raw=true)
- [Dash](http://kapeli.com/dash) is a wonderful documentation tool for macOS. It allows you to create profiles to scope your global searches.
	- It also provides you to set custom URLs to switch between these profiles.
- So I made a workflow to switch between different profiles. Here is how it looks :
- You can customise the Dash profile keyword triggers in profile settings here:

<img src="https://i.imgur.com/yGvrOwE.png" width="400" alt="img">

- And then you can call these triggers using this URL scheme : `dash://go:`
	- Where you can change `go` to the trigger you set. This workflow simply lets you search through the profiles you have and will call this URL scheme for you. Adapt it to how you like it.

### Dash Search (with custom hotkeys) - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/Dash.alfredworkflow?raw=true)
- An extenstion of [official Dash workflow](https://github.com/Kapeli/Dash-Alfred-Workflow) that I extended with many custom hotkeys

### Contacts search - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/Contacts%20filter.alfredworkflow?raw=true)
- Made by [Vero](https://www.alfredforum.com/profile/1-vero/) but she has not uploaded it anywhere and I think it is super useful
- It is a file filter for contacts where you can not only search for names of contacts but also for notes attached to these contacts

### GitHub Jump - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/GitHub%20jump.alfredworkflow?raw=true)
- Not my workflow so all credit goes to [Lachlan Donald](https://github.com/lox) who made the [original workflow](https://github.com/lox/alfred-github-jump)
- This is just my augmentation of it to make an amazing workflow even better
	- you can use modifiers to do different things such as go directly to issues of the workflow, pull requests of it or even clone the repo to a specified directory

## Contribute
If you use or used any of the workflows above and think there is something awesome that can be added to them. I would love to [hear it](https://github.com/nikitavoloboev/small-workflows/issues/new).

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look [into other repositories](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) I shared. 

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)