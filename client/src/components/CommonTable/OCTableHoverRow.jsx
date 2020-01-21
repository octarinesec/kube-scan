import React, {useRef} from 'react';
import OCTableRow from './OCTableRow';

function OCTableHoverRow (props) {

    let isHovered = false;
    let offset = 0;

    let _refs = {
        row: useRef(null),
    };

    let {children, ...other} = props;
    for (let propName of ['onMouseEnter', 'onMouseMove', 'onMouseLeave', 'onHoverChange', 'onOffsetChange']) {
        // don't pass these to OCTableRow through 'other'
        delete other[propName];
    }

    let onMouseMove = (e) => {
        if (props.onMouseMove) {
            props.onMouseMove(e);
        }
        onHoverIn();
    };

    let onMouseEnter = (e) => {
        if (props.onMouseEnter) {
            props.onMouseEnter(e);
        }
        onHoverIn();
    };

    let onMouseLeave = (e) => {
        if (props.onMouseLeave) {
            props.onMouseLeave(e);
        }
        onHoverOut();
    };

    function calcActionButtonsOffset(rowEl) {
        if (!(rowEl && rowEl.getBoundingClientRect && window && window.document)) {
            return 0;
        }
        const elBoundingBox = rowEl.getBoundingClientRect();
        const documentWidth = window.document.body.clientWidth;
        return  elBoundingBox.right - documentWidth + 20;
    }

    function onHoverIn() {
        if (props.onOffsetChange) {
            const newOffset = calcActionButtonsOffset(_refs.row.current);
            if (offset !== newOffset) {
                offset = newOffset;
                props.onOffsetChange(newOffset);
            }
        }
        if (!isHovered) {
            isHovered = true;
            if (props.onHoverChange) {
                props.onHoverChange(true);
            }
        }
    }

    function onHoverOut() {
        if (isHovered) {
            isHovered = false;
            if (props.onHoverChange) {
                props.onHoverChange(false);
            }
        }
    }

    return <OCTableRow
        {...other}
        onMouseEnter={onMouseEnter}
        onMouseMove={onMouseMove}
        onMouseLeave={onMouseLeave}
        forwardRef={_refs.row}
    >
        {children}
    </OCTableRow>
}


export default OCTableHoverRow;
