import { languages, EventTypes } from './lists';
import socket from './client-socket';
import SelectionManager from './SelectionManager';
const langSelect = document.getElementById('language-select');

class EditorWrapper {
  constructor (mode) {
    this.editor = ace.edit('editor');
    this.editor.setTheme('ace/theme/monokai');
    this.editor.session.setUseWorker(false);
    this.editor.session.setMode(mode);
    this.editor.$blockScrolling = Infinity;
    this.editor.getSession().setUseWrapMode(true);
    this.editor.setShowPrintMargin(false);
    let splitUrl = document.location.href.split('/');
    this.roomId = splitUrl[splitUrl.length - 1];
    this.isChanging = false;

    const self = this;

    this.editor.session.on(EventTypes.EditorChange, function (e) {
      self.userEdit(e)
    }.bind(this));

    this.selectionManager = new SelectionManager(this.editor)
  }

  setMode (newMode) {
    this.editor.session.setMode(newMode);
    let index = languages.findIndex(elem => elem.src === newMode);
    langSelect.selectedIndex = index;
  }

  changeHighlighting (newMode) {
    this.editor.session.setMode(newMode);
    if (!this.isChanging) {
      socket.emit(EventTypes.LanguageChange, JSON.stringify({
        roomId: this.roomId,
        mode: newMode
      }))
    }
  }

  serverChangeSelection (msg) {
    this.selectionManager.selectionChanged(msg)
  }

  serverChangeCursor (msg) {
    this.selectionManager.cursorChanged(msg)
  }

  removeOtherUser (clientId) {
    this.selectionManager.removeOtherUser(clientId)
  }

  setKeyboardHandler (kh) {
    this.editor.setKeyboardHandler(kh)
  }

  getRoomId () {
    return this.roomId
  }

  userEdit (e) {
    if (!this.isChanging) {
      socket.emit(EventTypes.UserEdit, JSON.stringify({
        change: e,
        text: this.editor.getValue(),
        roomId: this.roomId
      }))
    }
  }

  serverEdit (data) {
    this.isChanging = true;
    let change = data.change;
    switch (change.action) {
      case EventTypes.TextInsertion:
        this.editor.session.insert(change.start, change.lines.join('\n'))
        break;
      case EventTypes.TextRemoval:
        this.editor.session.remove(change)
    }
    this.isChanging = false
  }

  setSyntax (data) {
    this.isChanging = true;
    this.setMode(data.mode);
    this.isChanging = false;
  }
}

const Editor = new EditorWrapper(mode);
export default Editor
