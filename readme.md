# Small Workflows [![Workflows](https://img.shields.io/badge/More%20Workflows-🎩-purple.svg)](https://github.com/learn-anything/alfred-workflows) [![Thanks](https://img.shields.io/badge/Say%20Thanks-💗-ff69b4.svg)](https://www.patreon.com/nikitavoloboev)
> Small [Alfred](https://www.alfredforum.com/) workflows I use that don't warrant a GitHub repository of their own

## Contents
- [Search Files](/search-files) - A lot of [file filters](https://www.alfredapp.com/help/workflows/inputs/file-filter/) for various apps I use.
- [Objects library](/objects-library) - Useful premade objects for workflows.
- [Keyboard turn on / off](/keyboard-on-off) - Turn your keboard on and off.
- [Focus](/focus) - Start [Focus](https://heyfocus.com) blocking for some time that you specify.

### Focus - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/focus.alfredworkflow?raw=true)
- Allows you to start [Focus](https://heyfocus.com) blocking for some time that you specify

### Dictionary search - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/dictionary.alfredworkflow?raw=true)
- Search through dictionary from Alfred
	- I stopped using it in favour of opening Dictionary itself and then making the search
	- I also only use Dictionary to read and explore wiki

### Useful utilities - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/useful%20utilities.alfredworkflow?raw=true)
- Only has one utility so far and that is to open Alfred prompt with whatever I have selected passed in
	- I may add more in the future

### Search for content - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/search%20for%20content.alfredworkflow?raw=true)
- Two actions to search through the inside PDF and MindNode documents

### Copy months numbers - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/show%20month%20numbers.alfredworkflow?raw=true)
- A simple workflow that will let you search for a month from the 12 months in a year and copy the month number to your clipboard
- Because I always found it too hard to remember these numbers

### Check wifi connection / Restart wifi / Toggle it on/off - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/wifi.alfredworkflow?raw=true)
- It can:
	1. **Check wifi** : it pings Google.com and if there is no response, you are offline and you get notification on the top
	2. **Restart wifi** : turns the wifi on and off
	3. **Toggle wifi on/off** : if on, will turn it off, if off will turn it on

### Imgur album -> Desktop - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/imgur%20album%20downloader.alfredworkflow?raw=true)
- Will take an active url from your browser that is an imgur album and download it to your desktop

### Reload Karabiner XML - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/karabiner.alfredworkflow?raw=true)
- I use it all the time to quickly iterate on Karabiner. Only works on old [Karabiner](https://github.com/tekezo/Karabiner)

### Clean Folders - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/clean%20folders.alfredworkflow?raw=true)
- Little list filter I use to filter my file system
- Currently use it to clean all items from Desktop and removing all alfred workflows from ~/Downloads

### Folder Search - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/folder%20search.alfredworkflow?raw=true)
- Search folders from Alfred and open them in Finder / iTerm / Editor

### Static Searches - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/static%20searches.alfredworkflow?raw=true)
- In case there is no Alfred workflow dedicated to some website search that provides autosuggestions like [Searchio](https://github.com/deanishe/alfred-searchio), I use my [Web Searches](https://github.com/nikitavoloboev/alfred-web-searches) workflow, wwhich contains a large array of websites I can make searches on moderated by the community
- However certain searches I do so often that I create a separate objects for them in Alfred and attach a hotkey on them. Things like searching on GitHub or Reddit. I also have the power to modify the search with modifer keys so ⌃ + return on GitHub search will scope my query to bring me only the most recent stuff.

### Useful File Actions - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/useful%20file%20actions.alfredworkflow?raw=true)
- Here are a bunch of small file actions I use to operate on files
	1. Opening folder/file in Visual Studio Code (change for editor you like)
	2. Copying files to Desktop

### Search Selected Text on the Web - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/search%20web%20with%20selection.alfredworkflow?raw=true)
- Simple workflow I use to search highlighted text on various websites
- You can customise the hotkeys to ones you like. To make things easier for myself, I call the hotkeys from Karabiner with my custom modifier keys which you can take a look [here](https://github.com/nikitavoloboev/dotfiles) for.
- I also heavily use [Web Searches workflow](https://github.com/nikitavoloboev/alfred-web-searches) to search selected text on a wider array of websites

### Go to Reddit Subreddit from Alfred - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/go%20to%20subreddit.alfredworkflow?raw=true)
- Simple workflow that will let you go to a subreddit you specify
- If you want to search for available subreddits with autosuggestions, you can try [this workflow](https://github.com/deanishe/alfred-reddit)

### Go Play - code -> playground - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/workflows/go%20play.alfredworkflow?raw=true)
- This is a simple workflow that will take the code you select and let you create either [Official Go playground](https://play.golang.org/) or [Go play Space playground](https://goplay.space) that you can then share on IRC, Slack or with other gophers around you

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

## Contributing
If you use or used any of the workflows above and think there is something awesome that can be added to them. I would love to [hear it](https://github.com/nikitavoloboev/small-workflows/issues/new).

## Thank you 💜
You can support what I do on [Patreon](https://www.patreon.com/nikitavoloboev) or look [into other repositories](https://my.mindnode.com/ZKGETDkUaQUsL3q8q9z788CxG84oEHgDiT79GuzX#-143.5,-902.6,0) I shared. 

## License
MIT © [Nikita Voloboev](https://www.nikitavoloboev.xyz)