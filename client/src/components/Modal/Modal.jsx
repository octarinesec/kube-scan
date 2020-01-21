import React from 'react';
import * as ReactDom from 'react-dom';

import classNames from 'classnames';
import { KEY_NAMES } from '../../constants/keyboard';

import './Modal.scss';

export const nullFunc = (...args) => {};
const activeReplicas = [];

const getModalRoot = () => {
  return document.getElementById('modal-root');
};

const updateRootClass = () => {
  const root = getModalRoot();
  if (!root) {
    return;
  }
  const hasActiveModals = !!activeReplicas.length;
  if (hasActiveModals) {
    if (root.classList.contains('active')) {
      return;
    }
    root.classList.add('active');
  } else {
    root.classList.remove('active');
  }
};

const registerModal = (modalReplica) => {
  activeReplicas.push(modalReplica);
  updateRootClass();
};

const unregisterModal = (modalReplica) => {
  const index = activeReplicas.indexOf(modalReplica);
  if (index > -1) {
    activeReplicas.splice(index, 1);
  }
  updateRootClass();
};


export default class Modal extends React.Component {

  static defaultProps =  {
    onBackdropClick: nullFunc,
    onEscape: nullFunc,
    modalClassNames: [],
  };

  constructor(props) {
    super(props);
    this.modalRoot = getModalRoot();
    this.el = document.createElement('div');
    this.el.className = classNames(props.modalClassNames) + ' modal-dialog';
    if (typeof props.onEscape === 'function') {
      this._addKeyUpListener();
    }
    registerModal(this);
  }

  componentDidMount() {
    if (!this.modalRoot) {
      return;
    }
    this.modalRoot.appendChild(this.el);
  }

  componentWillUnmount() {
    this.modalRoot.removeChild(this.el);
    unregisterModal(this);
  }

  render() {
    const modal = (
      <div className="modal-wrapper">
        <div className="modal-backdrop" onClick={this.onBackdropClick}></div>
        <div className="modal-content">{this.props.children}</div>
      </div>
    );
    return ReactDom.createPortal(
      modal,
      this.el,
    )
  }

  _onKeyUp = (e) => {
    if (e.key === KEY_NAMES.ESC && typeof this.props.onEscape === 'function') {
      this.props.onEscape();
    }
  };

  onBackdropClick = () => {
    this.props.onBackdropClick();
  };

  _addKeyUpListener() {
    if (window && window.document) {
      window.document.addEventListener('keyup', this._onKeyUp);
    }
  }

  _removeKeyUpListener() {
    if (window && window.document) {
      window.document.removeEventListener('keyup', this._onKeyUp);
    }
  }
}
