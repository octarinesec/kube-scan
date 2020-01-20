import React from 'react';
import cn from 'classnames';


function OCTableCell(props){

  //static defaultProps = {
  //   props.classNames = props.classNames ||[];
  //   props.fieldName =  props.fieldName || null;
  //   props.title = props.title || null;
  //};

  const {classNames: extraClassNames, fieldName, title } = props;
  const cNames = cn(['oc-table-cell', fieldName].concat(extraClassNames || []));

  return (
    <div className={cNames} title={title}>
      {props.children}
    </div>
  );
}

export default OCTableCell;
