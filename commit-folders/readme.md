# Commit Folders - [Download](https://github.com/nikitavoloboev/small-workflows/blob/master/commit-folders/Commit%20folders.alfredworkflow?raw=true)
> Search folders and commit changes inside the folders with commit message you like.

When you commit a folder, it would `git add` and `git commit` all files in the folder with the message you provide. And after it would also `git push` the changes.

This workflow assumes that the folder you select is already initialised with git as it is a file filter that searches through all the folders. It will however warn you if the folder you chose is not under git.

You can of course modify the scope of the search or even make it that it only searches through git initialised folders but for my use case, I use this workflow when I make a change to some README file in some folder and want to quickly commit it without going to my terminal.

For quickly editing README's of folders I use this [workflow](../folder-search) that searches through all folders and with a modifier key press will open the README in the markdown editor you set up.

There is also alternative action that will commit only the README file inside the folder with a commit message you write.

![](https://i.imgur.com/chzfkue.png)