import React from 'react';
import cn from 'classnames';


import './Toolbar.scss';
import ContactButton from "../ContactButton/ContactButton";

function _Toolbar ({contactLink}) {
  const cNames = cn([
    'Toolbar',
  ]);

  return (
    <div className={cNames}>
      <div className='upper-toolbar octaudit-toolbar'>
            <span className='account'>
              Kube-Scan
            </span>
        <div className='section right'>
          <ContactButton contactLink={contactLink}/>
        </div>
      </div>
    </div>
  );
}

export default _Toolbar;


