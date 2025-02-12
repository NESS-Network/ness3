'use strict'

const { app, Menu, BrowserWindow, shell, session } = require('electron');

const path = require('path');

const childProcess = require('child_process');

const axios = require('axios');

// This adds refresh and devtools console keybindings
// Page can refresh with cmd+r, ctrl+r, F5
// Devtools can be toggled with cmd+alt+i, ctrl+shift+i, F12
require('electron-debug')({enabled: true, showDevTools: false});
require('electron-context-menu')({});


global.eval = function() { throw new Error('bad!!'); }

let currentURL;
let splashLoaded = false

// Detect if the code is running with the "dev" arg. The "dev" arg is added when running npm
// start. If this is true, a local node will not be started, but one is expected to be running,
// the contents served in http://localhost:4200 will be displayed and it will be allowed to
// reload the URLs using the Electron window, so that it is easier to test the changes made to
// the UI using npm start.
let dev = process.argv.find(arg => arg === 'dev') ? true : false;

// Force everything localhost, in case of a leak
app.commandLine.appendSwitch('host-rules', 'MAP * 127.0.0.1, EXCLUDE api.coinpaprika.com, EXCLUDE teller.mdl.cx, EXCLUDE api.github.com');
app.commandLine.appendSwitch('ssl-version-fallback-min', 'tls1.2');
app.commandLine.appendSwitch('--no-proxy-server');
app.setAsDefaultProtocolClient('mdl');



// Keep a global reference of the window object, if you don't, the window will
// be closed automatically when the JavaScript object is garbage collected.
let win;

var mdl = null;

function startMDL() {
  if (!dev) {
    console.log('Starting mdl from electron');

    if (mdl) {
      console.log('MDL already running');
      app.emit('mdl-ready');
      return
    }

    var reset = () => {
      mdl = null;
    }

    // Resolve mdl binary location
    var appPath = app.getPath('exe');
    var exe = (() => {
      switch (process.platform) {
        case 'darwin':
          return path.join(appPath, '../../Resources/app/mdl');
        case 'win32':
          // Use only the relative path on windows due to short path length
          // limits
          return './resources/app/mdl.exe';
        case 'linux':
          return path.join(path.dirname(appPath), './resources/app/mdl');
        default:
          return './resources/app/mdl';
      }
    })()

    var args = [
      '-launch-browser=false',
      '-gui-dir=' + path.dirname(exe),
      '-color-log=false', // must be disabled for web interface detection
      '-logtofile=true',
      '-download-peerlist=true',
      '-enable-all-api-sets=true',
      '-enable-api-sets=INSECURE_WALLET_SEED',
      '-disable-csrf=false',
      '-reset-corrupt-db=true',
      '-enable-gui=true',
      '-web-interface-port=0' // random port assignment
      // will break
      // broken (automatically generated certs do not work):
      // '-web-interface-https=true',
    ]
    mdl = childProcess.spawn(exe, args);

    createWindow();

    mdl.on('error', (e) => {
      showError();
      app.quit();
    });

    mdl.stdout.on('data', (data) => {
      console.log(data.toString());
      if (currentURL) {
        return
      }

      const marker = 'Starting web interface on ';

      data.toString().split('\n').forEach(line => {
        if (line.indexOf(marker) !== -1) {
          currentURL = 'http://' + line.split(marker)[1].trim();
		  var id = setInterval(function() {
			// wait till the splash page loading is finished
			if (splashLoaded) {
			  app.emit('mdl-ready', { url: currentURL });
			  clearInterval(id);
			}
		  }, 500);
        }
      });
    });

    mdl.stderr.on('data', (data) => {
      console.log(data.toString());
    });

    mdl.on('close', (code) => {
      // log.info('MDL closed');
      console.log('MDL closed');
      showError();
      reset();
    });

    mdl.on('exit', (code) => {
      // log.info('MDL exited');
      console.log('MDL exited');
      showError();
      reset();
    });

  } else {
    // If in dev mode, simply open the dev server URL.
    currentURL = 'http://localhost:4200/';
    app.emit('mdl-ready', { url: currentURL });

    axios
      .get('http://localhost:4200/api/v1/wallets/folderName')
      .then(response => {
        walletsFolder = response.data.address;
      })
      .catch(() => {});
  }
}

function showError() {
  if (win) {
    win.loadURL('file://' + process.resourcesPath + '/app/dist/assets/error-alert/index.html');
    console.log('Showing the error message');
  }
}

function createWindow(url) {
  // To fix appImage doesn't show icon in dock issue.
  var appPath = app.getPath('exe');
  var iconPath = (() => {
    switch (process.platform) {
      case 'linux':
        return path.join(path.dirname(appPath), './resources/icon512x512.png');
    }
  })()

  // Create the browser window.
  win = new BrowserWindow({
    width: 1200,
    height: 900,
    backgroundColor: '#000000',
    title: 'MDL Talent Hub',
    icon: iconPath,
    nodeIntegration: false,
    webPreferences: {
      webgl: false,
      webaudio: false,
      contextIsolation: false,
      webviewTag: false,
      nodeIntegration: false,
      nodeIntegrationInWorker: false,
      allowRunningInsecureContent: false,
      webSecurity: true,
      plugins: false,
      enableRemoteModule: false,
      preload: __dirname + '/electron-api.js',
    },
  });

  win.webContents.on('did-finish-load', function() {
	if (!splashLoaded) {
	  splashLoaded = true;
	}
  });

  // patch out eval
  win.eval = global.eval;
  win.webContents.executeJavaScript('window.eval = 0;');

  const ses = win.webContents.session

  ses.clearCache(function () {
    console.log('Cleared the caching of the MDL wallet.');
  });

  if (url) {
    win.loadURL(url);
  } else {
    win.loadURL('file://' + __dirname + '/splash/index.html');
  }

  // Open the DevTools.
  // win.webContents.openDevTools();

  // Emitted when the window is closed.
  win.on('closed', () => {
    // Dereference the window object, usually you would store windows
    // in an array if your app supports multi windows, this is the time
    // when you should delete the corresponding element.
    win = null;
  });

  // If in dev mode, allow to open URLs.
  if (!dev) {
    win.webContents.on('will-navigate', function(e, url) {
      e.preventDefault();
      require('electron').shell.openExternal(url);
    });
  }

  // Open links with target='_blank' in the default browser.
  win.webContents.on('new-window', function(e, url) {
    e.preventDefault();
    require('electron').shell.openExternal(url);
  });

  // create application's main menu
  var template = [{
    label: "MDL Talent Hub",
    submenu: [
      { label: "Quit", accelerator: "Command+Q", click: function() { app.quit(); } }
    ]
  }, {
    label: 'Edit',
    submenu: [
      { label: 'Undo', accelerator: 'CmdOrCtrl+Z', role: 'undo' },
      { label: 'Redo', accelerator: 'Shift+CmdOrCtrl+Z', role: 'redo' },
      { type: 'separator' },
      { label: 'Cut', accelerator: 'CmdOrCtrl+X', role: 'cut' },
      { label: 'Copy', accelerator: 'CmdOrCtrl+C', role: 'copy' },
      { label: 'Paste', accelerator: 'CmdOrCtrl+V', role: 'paste' },
      { label: 'Select All', accelerator: 'CmdOrCtrl+A', role: 'selectall' }
    ]
  }, {
    label: 'Show',
    submenu: [
      {
        label: 'Wallets folder',
        click: () => {
          if (walletsFolder) {
            shell.showItemInFolder(walletsFolder)
          } else {
            shell.showItemInFolder(path.join(app.getPath("home"), '.mdl', 'wallets'));
          }
        },
      },
      {
        label: 'Logs folder',
        click: () => {
          if (walletsFolder) {
            shell.showItemInFolder(walletsFolder.replace('wallets', 'logs'))
          } else {
            shell.showItemInFolder(path.join(app.getPath("home"), '.mdl', 'logs'));
          }
        },
      },
      {
        label: 'DevTools',
        accelerator: process.platform === 'darwin' ? 'Alt+Command+I' : 'Ctrl+Shift+I',
        click: (item, focusedWindow) => {
          if (focusedWindow) {
            focusedWindow.toggleDevTools();
          }
        }
      },
    ]
  }];

  Menu.setApplicationMenu(Menu.buildFromTemplate(template));

  session
    .fromPartition('')
    .setPermissionRequestHandler((webContents, permission, callback) => {
      return callback(false);
    });
}

const singleInstanceLockObtained = app.requestSingleInstanceLock()

if (!singleInstanceLockObtained) {
    app.quit()
    return;
} else {
    app.on('second-instance', (event, commandLine, workingDirectory) => {
        // Someone tried to run a second instance, we should focus our window.
        if (win) {
            if (win.isMinimized()) {
                win.restore();
            }
            win.focus();
        } else {
            createWindow(currentURL);
        }
    });
}

let walletsFolder = null;

// This method will be called when Electron has finished
// initialization and is ready to create browser windows.
// Some APIs can only be used after this event occurs.
app.on('ready', startMDL);

app.on('mdl-ready', (e) => {
  if (win) {
    win.loadURL(e.url);
  } else {
    createWindow(e.url);
  }

  axios
    .get(e.url + '/api/v1/wallets/folderName')
    .then(response => {
      walletsFolder = response.data.address;
      //hwCode.setWalletsFolderPath(walletsFolder);
    })
    .catch(() => {});
});

// Quit when all windows are closed.
app.on('window-all-closed', () => {
  // On OS X it is common for applications and their menu bar
  // to stay active until the user quits explicitly with Cmd + Q
  if (process.platform !== 'darwin') {
    app.quit();
  }
});

app.on('activate', () => {
  // On OS X it's common to re-create a window in the app when the
  // dock icon is clicked and there are no other windows open.
  if (win === null) {
    createWindow(currentURL);
  }
});

app.on('will-quit', () => {
  if (mdl) {
    mdl.kill('SIGINT');
  }
});

app.on('web-contents-created', (event, contents) => {
  contents.on('will-attach-webview', (event, webPreferences, params) => {
    // Strip away preload scripts if unused or verify their location is legitimate
    delete webPreferences.preload
    delete webPreferences.preloadURL

    // Disable Node.js integration
    webPreferences.nodeIntegration = false

    // Verify URL being loaded
    if (!params.src.startsWith(url)) {
      event.preventDefault();
    }
  });
});
