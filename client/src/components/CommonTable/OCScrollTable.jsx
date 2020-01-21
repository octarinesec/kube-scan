import React, { useState, useEffect } from 'react';
import cn from 'classnames';

import './OCScrollTable.scss';


function OCScrollTable (props) {

  // static defaultProps = {
  //props.classNames =  props.classNames || [];
  // };


  let [state, setState] = useState({headerPadding: 0});


  function renderHeader() {
    return props.renderHeader();
  }
  function renderBody() {
    return props.renderBody();
  }
  function updateHeaderPadding() {
    // adds padding to the header to it'll match any padding added to the body from the scrollbar
    setTimeout(() => {
      if (!_tableBody) {
        return;
      }
      const scrollBounding = _yScroll && _yScroll.getBoundingClientRect && _yScroll.getBoundingClientRect();
      const bodyBounding = _tableBody && _tableBody.getBoundingClientRect && _tableBody.getBoundingClientRect();
      if (!(scrollBounding && bodyBounding)) {
        return;
      }
      const headerPadding = scrollBounding.width - bodyBounding.width;
      if (headerPadding !== state.headerPadding) {
        setState({
          headerPadding: headerPadding,
        });
      }
    }, 100);
  }

  useEffect(() => {
    updateHeaderPadding();
  });
  // componentDidMount() {
  //   this.updateHeaderPadding();
  // }
  // componentDidUpdate() {
  //   this.updateHeaderPadding();
  // }


    const header = renderHeader();
    const body = renderBody();
    let cNames = cn(['OCScrollTable'].concat(props.classNames));
    return (
      <div className={cNames}>
        <div className='xScroll'>
          <div className='xContent'>
            <div className='header-wrapper' style={{'paddingRight': state.headerPadding || 0}}>
              {header}
            </div>
            <div className='yScroll' ref={(r)=>{_yScroll = r}}>
              <div className='table-body' ref={(r)=>{_tableBody = r}}>
                {body}
              </div>
            </div>
          </div>
        </div>
      </div>
    );
}

export default  OCScrollTable;
