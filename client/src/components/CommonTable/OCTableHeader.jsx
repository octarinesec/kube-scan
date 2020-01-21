import React from 'react';
import classNames from 'classnames';


function OCTableHeader (props) {

  // static defaultProps = {
  // props.classNames =  props.classNames || [];
  // };

    const { children, classNames: extraClassNames } = props;
    const cNames = classNames(['oc-table-header', 'oc-table-row'].concat(extraClassNames || []));

    return (
      <div className={cNames}>
        {children}
      </div>
    );
}

export default  OCTableHeader;
