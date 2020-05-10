import React from 'react';
import { FormControlLabel } from '@material-ui/core'
import MuiSwitch from '@material-ui/core/Switch'
import classNames from 'classnames';

function Switch(props) {
  const {label, isChecked, onChange, className} = props
  return (
    <FormControlLabel
      classes={ {
        root: classNames("makeStyles-formControl-14", className)
      } }
      control={ (
        <MuiSwitch
          checked={ isChecked }
          onChange={ (event) => onChange?.(event.target.checked) }
          classes={ {
            root: "makeStyles-root-10 makeStyles-root-16",
            switchBase: "makeStyles-switchBase-11 makeStyles-switchBase-17",
            thumb: "makeStyles-thumb-12 makeStyles-thumb-18",
            track: "makeStyles-track-13"
          } }
        />
      ) }
      label={ <span className="makeStyles-label-15">{ label }</span> }
    />
  )
}

export default Switch

const defaultProps = {
  size: 12,
  label: '',
  isChecked: false
}

Switch.defaultProps = defaultProps