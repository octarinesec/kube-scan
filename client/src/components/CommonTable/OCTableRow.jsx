import React from 'react';
import classNames from 'classnames';

 function OCTableRow(props) {

   // static defaultProps = {
   // props.classNames = props.classNames || [];
   // props.wrapFields = props.wrapFields || false;
   // };
    const { children, classNames: extraClassNames, wrapFields, forwardRef, ...other } = props;
    const cNames = classNames(['oc-table-row', { 'wrap-fields': wrapFields }].concat(extraClassNames || []));

    return (
      <div className={cNames} {...other} ref={forwardRef}>
        <div className={'oc-table-row-inner'}>
          {children}
        </div>
      </div>
    );
}

export default OCTableRow;
