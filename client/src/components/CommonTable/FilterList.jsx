import React, {useEffect} from 'react';
import { KEY_NAMES } from '../../constants/keyboard';
import cn from 'classnames';

import './FilterList.scss';

function FilterList (props) {

  // static defaultProps = {
  if (!props.items) {
    props.items = [];
  }
  //props.show =  props.show || false;
  // };

  let onItemClick = (e, item) => {
    e.preventDefault();
    e.stopPropagation();

    const { onChange, onClose } = props;
    onClose();
    onChange(item);
  };

  let onInputKeyUp = (e) => {
    e.preventDefault();
    e.stopPropagation();
    const { onChange, onClose } = props;
    switch (e.key) {
      case KEY_NAMES.ENTER:
        onChange(textInput.value);
        onClose();
        break;
      case KEY_NAMES.ESC:
        onChange('');
        onClose();
        break;
      default:
        break;
    }
  };

  let onInputChange = (e) => {
    props.onChange(e.target.value || '');
  };

  let onInputClick = (e) => {
    e.preventDefault();
    e.stopPropagation();
  };

  useEffect(() => {
    const { show } = props;

    if (textInput && show) {
      textInput.value = '';
      setTimeout(() => { if (textInput) textInput.focus(); } , 100);
    }
  });

    const { type, items, show, filterValue } = props;

    let content;

    switch (type) {
      case 'text':
        content = (
          <span>
          <div className='search icon'></div>
          <input
            type='text'
            placeholder='Type to filter'
            autoComplete="off"
            autoCorrect="off"
            autoCapitalize="off"
            spellCheck="false"
            onKeyUp={onInputKeyUp}
            onClick={onInputClick}
            onChange={onInputChange}
            value={filterValue || ''}
            ref={ input => textInput = input }
          />
        </span>
        );
        break;
      case 'enum':
        content = items.map(
          (item, i) => (
            <div
              key={`filterItem_${i}`}
              className='filterItem'
              onClick={(e) => {onItemClick(e, item); }}
            >{item}</div>
          )
        );
        break;
      default:
        content = null;
    }

    const classNames = cn([
      'FilterList',
      `type-${type}`,
      {
        show: show,
      }
    ]);

    return (
      <span className={classNames}>
        { content }
      </span>
    );

}

export default FilterList;