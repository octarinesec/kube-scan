import React, { useRef, useEffect } from 'react';
import * as PropTypes from 'prop-types';
import cn from 'classnames';
import { AutoSizer, List, CellMeasurerCache, CellMeasurer } from 'react-virtualized';
import throttle from 'lodash/throttle';
import ReactResizeDetector from 'react-resize-detector';
import { get as lookupGet } from 'lodash/object';

import './OCVirtualTable.scss';

function OCVirtualTable(props) {

  // static defaultProps
  //   props.minHeight = props.minHeight || 40;
  //   props.overscanRowCount = props.overscanRowCount || 10;
  //   props.limitWidthToHeader = props.limitWidthToHeader || true;
  //   props.headerEstimatedHeight = props.headerEstimatedHeight || 60;

    let _onResize = () => {
        if (_refs.list.current) {
            _cache.clearAll();
            _refs.list.current.recomputeRowHeights();
        }
    };

    let _cache = new CellMeasurerCache({
      fixedWidth: true,
      minHeight: props.minHeight || 40,
    });

    let onResize = throttle(_onResize, 100, { leading: true, trailing: true});

    let _refs = {
      list: useRef(null),
      headerWrapper: useRef(null),
    };

  useEffect(() => {
      setTimeout(() => {
          onResize();
      }, 100);
  });

  function get_topOffset() {
    const listRef = _refs.list.current;
    const gridRef = lookupGet(listRef, ['Grid', '_scrollingContainer']);
    if (gridRef) {
      return gridRef.offsetTop;
    }
    return props.headerEstimatedHeight || 40;
  }

  function rowDidChange(rowIndex) {
    rowsDidChange([rowIndex]);
  }

  function rowsDidChange(rowIndices) {
    for (let rowIndex of (rowIndices || [])) {
      _cache.clear(rowIndex);
    }
    if (_refs.list.current) {
      _refs.list.current.recomputeRowHeights();
    }
  }

  function get_headerWidth() {
    if (_refs.headerWrapper.current) {
      return _refs.headerWrapper.current.getBoundingClientRect().width;
    }
    return null;
  }

    const { items, classNames, overscanRowCount, limitWidthToHeader, renderEmpty } = props;
    const cNames = cn(['OCVirtualTable'].concat(classNames || []));
    const topOffset = get_topOffset();

  let _renderRow = ({index, style, key, parent: parentList}) => {
    const items = props.items;
    const item = (items && items.length && (index <= items.length) && items[index]) || null;

    return (
        <CellMeasurer
            parent={parentList}
            cache={_cache}
            columnIndex={0}
            key={key}
            rowIndex={index}
        >
          <div className='row-wrapper' style={style} key={key}>
            { props.renderRow(index, item) }
          </div>
        </CellMeasurer>
    );
  };

    return (
      <div className={cNames}>
        <ReactResizeDetector handleHeight={true} handleWidth={true} onResize={onResize}/>
        <div className='x-scroll-wrapper'>
          <div className='x-scroll-content'>
            <div className='header-wrapper' ref={_refs.headerWrapper}>
              { props.renderHeader && props.renderHeader() }
            </div>
            { !(items && items.length) && renderEmpty ? (
              renderEmpty()
            ) : (
              <AutoSizer defaultHeight={300} defaultWidth={500} style={{ width: 'auto', height: 'auto'}}>
              {({width, height}) => {
                let listWidth = width;
                let headerWidth = get_headerWidth();
                if (limitWidthToHeader && headerWidth) {
                  listWidth = Math.min(headerWidth, listWidth);
                }
                return (
                  <List
                    ref={_refs.list}
                    rowRenderer={_renderRow}
                    estimatedRowSize={60}
                    height={height - topOffset}
                    overscanRowCount={overscanRowCount || 10}
                    rowHeight={_cache.rowHeight}
                    rowCount={(items || []).length}
                    width={listWidth}
                    style={{minWidth: listWidth, width: 'auto'}}
                    className={'table-body'}
                  />
                )
              }}
              </AutoSizer>
              )}
          </div>
        </div>
      </div>
    );
}


export default OCVirtualTable;
