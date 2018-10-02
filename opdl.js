App = Application.currentApplication();
App.includeStandardAdditions = true;

do {
	answer = App.displayDialog('Please enter chapter', {
  		withTitle: 'Chapter',
	    defaultAnswer: ''
	});
	chapter = answer["textReturned"];
	isNumber = /\d+/.test(chapter);
} while(!isNumber)


var outFolder = App.chooseFolder({withPrompt: 'Select output folder - Folder for chapter will be created automatically'});

thePath = App.pathTo(this).toString();
mypath = thePath.split('/')
mypath.pop()
scriptPath = mypath.join('/')

App.doShellScript(scriptPath + '/onepiece-dl -chap=' + chapter + ' -output=' + outFolder.toString())

App.displayNotification('Download finished!',
  { withTitle: 'Success', subtitle: 'Chapter ' + chapter })