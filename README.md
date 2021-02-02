# not-create-react-app

## Description:
A simple bash script to automate creating a new React based project, creating and pushing to a github repository and then opening the VS Code text editor inside your new project. This project depends on the user having the [Github CLI Tools](https://cli.github.com/) installed so that creating a repository can be done from the command line.

<br />

## Getting Started:

Clone the repository:

`git clone https://github.com/damiensedgwick/not-create-react-app.git`

Change directory into the project folder and move create script to users bin folder:

`cd not-create-react-app && mv create /Users/user/bin`

Change script permissions so that you can execute it:

`chmod +x /Users/user/bin/create`

Make sure your bin folder is exported to path:

`echo "export PATH="$PATH:/Users/user/bin/"" >> .bashrc`

You should now be able to run `create project-name` from your command line and go through the steps of choosing a framework for your project, pushing to github and opening vscode with just one command.

<br />

## Issues
If you experience any issues getting this script working please make sure you have all the the tools mentioned for it to work properly as well as making sure your path is correct and bin exported correctly. If you're still stuck, reach out and I will do my best to help.

<br />

## Contributions
If you have a suggestion for the script then please get in touch and let me know and likewise, if you would like to contribute, please do not hesitate to do so.

<br />

### NB:
This project was insired by another project automation tool which you can find here:

[Kalle Hallden: ProjectInitializationAutomation
](https://github.com/KalleHallden/ProjectInitializationAutomation)
