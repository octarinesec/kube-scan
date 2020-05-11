import React from 'react';
import { FormControlLabel } from '@material-ui/core'
import MuiSwitch from '@material-ui/core/Switch'

function Switch(props) {
  const {label, isChecked, onChange, className} = props
  return (
    <FormControlLabel
      classes={ {
        root: className
      } }
      control={ (
        <MuiSwitch
          checked={ isChecked }
          onChange={ (event) => onChange?.(event.target.checked) }
          classes={ {
            root: "switch-root",
            input: "switch-input",
            'checked': 'switch-checked',
            switchBase: "switch-base",
            thumb: "switch-thumb",
            track: "switch-track"
          } }
        />
      ) }
      label={ <span className="switch-label">{ label }</span> }
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