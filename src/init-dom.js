import Editor from './EditorWrapper';
import { languages, keybindings } from './lists';

const langSelect = document.getElementById('language-select');
const keybindingSelect = document.getElementById('keybinding-select');
const userCount = document.getElementById('user-count');

function populateLanguages () {
  languages.forEach(lang => {
    let opt = document.createElement('option');
    opt.setAttribute('value', lang.src);
    opt.innerHTML = lang.name;
    langSelect.appendChild(opt);
  });
  if (mode) {
    Editor.setMode(mode);
  }
}

function populateKeyBindings () {
  keybindings.forEach(kbtype => {
    let opt = document.createElement('option');
    opt.setAttribute('value', kbtype.src);
    opt.innerHTML = kbtype.name;
    keybindingSelect.appendChild(opt);
  })
}

export function changeUserCount (change) {
  let currCount = parseInt(userCount.innerHTML);
  currCount += change;
  setUserCount(currCount)
}

export function setUserCount (newCount) {
  userCount.innerHTML = newCount;
}

langSelect.onchange = function () {
  const selectedMode = langSelect.options[langSelect.selectedIndex].value;
  Editor.changeHighlighting(selectedMode);
};

keybindingSelect.onchange = function () {
  let selectedKeybinding = keybindingSelect.options[keybindingSelect.selectedIndex].value;
  if (selectedKeybinding === 'ace') {
    selectedKeybinding = undefined;
  }
  Editor.setKeyboardHandler(selectedKeybinding);
};

populateLanguages();
populateKeyBindings();
