import React, { useState } from 'react';
import classNames from 'classnames';


function _OCTableHeaderCell (props) {

  let onClick = (e) => {
    const { isSortable } = props;

    if (!isSortable) {
      return;
    }
    props.sortFunc(props.sortField);
  };

  function get_title() {
    return props.title || null;
  }

  function renderSortIcon() {
    const { currentSort, ascending, sortField } = props;

    if (!currentSort || currentSort !== sortField) {
      return null;
    }

    const sortIconClassName = classNames([
      'sort-icon',
      {
        small: false,
        ascending: ascending,
        descending: !ascending
      },
    ]);
    return (<span className={sortIconClassName}></span>);
  }



    const { children, classNames: extraClassNames, fieldName, as, style } = props;
    const cNames = classNames(['oc-table-cell', 'oc-table-header-cell', fieldName].concat(extraClassNames || []));
    const title = get_title();
    const sortIconElement = renderSortIcon();
    const Element = as || 'div';

    return (
      <Element
        className={ cNames }
        onClick={ onClick }
        title={ title }
        style={ style }
      >
        { children }
        { sortIconElement }
      </Element>
    );
};



export default _OCTableHeaderCell;
