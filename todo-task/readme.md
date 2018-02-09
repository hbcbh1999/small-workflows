# New TODO task - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/todo-task/New%20todo%20task.alfredworkflow?raw=true)
> Write a task and save it to small file. Can then have Bitbar or Hammerspoon show contents of the task.

## Why use this?
I use this workflow to create (or update) a file named `todo` that I place in `~/app/hammerspoon/` directory.

I then execute this function from Hammersoon:
```Lua
function showTodoTask ()
hs.alert.show( ( hs.execute("~/app/hammerspoon/todo") ) )
end
hs.urlevent.bind("showTodoTask", showTodoTask)
```

Which shows the contents of this `todo` file in the middle of the screen to remind myself of what I need to do. It looks like this:

![](https://i.imgur.com/QIqdszD.png)

You can also point [BitBar](https://github.com/matryer/bitbar) or [TextBar](http://richsomerfield.com/apps/textbar/) to read the content of the file and that will show the `todo` task in your menu bar.

## Setup
You should change this folder before running the workflow and you can do it by double clicking `Run Script` object here:

![](https://i.imgur.com/gjnm7OG.png)

And then changing the folder here:

![](https://i.imgur.com/agl6XXD.png)

Default folder is `~/app/hammerspoon` which most certainly won't work for you as you won't have this folder. So change it to where you want the `todo` file to live!