import React from 'react';
import { capitalize } from "../../utils/text";


export const parseField = (dataItem, fieldSpec, context) => {
  let value;
  if (fieldSpec.valueFunc) {
    value = fieldSpec.valueFunc(dataItem, fieldSpec, context);
  } else if (fieldSpec.attr) {
    value = dataItem[fieldSpec.attr];
  }
  if (fieldSpec.capitalize) {
    value = capitalize(value);
  }
  const name = fieldSpec.name || fieldSpec.attr;
  return {
    name: name,            // code name (for css class, and identifying the field in callbacks)
    value: value,          // display value
    wrap: fieldSpec.wrap,  // should the field wrap
  }
};

export const renderItemsInField = (getListItems, valueForItem) => (rowData, fieldSpec) => {
  const listItems = getListItems(rowData);
  return (
    <div className='field-items-list'>
      { listItems ?
        (
          listItems.map((item)=> {
            let value = item;
            if (valueForItem) {
              value = valueForItem(item);
            }
            return (
              <span key={ value } className='field-inner-item'>{ value }</span>
            );
          })
        ) : null
      }
    </div>
  );
};
