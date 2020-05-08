import React from 'react';
import { makeStyles, FormControlLabel } from '@material-ui/core'
import MuiSwitch from '@material-ui/core/Switch'
import classNames from 'classnames';

const useStyles = makeStyles({
  root: {
    display: 'flex',
    width: ({size}) => 2 * (size + 2),
    height: ({size}) => size + 2,
    padding: 0,
    marginRight: 5,
    overflow: 'visible'
  },
  switchBase: {
    padding: 2,
    color: 'white',
    '&.Mui-checked': {
      transform: ({size}) => `translateX(${ size }px)`,
      color: 'white',
      '& + $track': {
        opacity: 1,
        backgroundColor: '#795aa4',
        borderColor: '#795aa4'
      }
    }
  },
  thumb: {
    width: ({size}) => size,
    height: ({size}) => size,
    boxShadow: 'none'
  },
  track: {
    border: '1px solid #ccc',
    borderRadius: 50,
    opacity: 1,
    backgroundColor: '#ccc'
  },
  formControl: {
    display: 'flex',
    margin: 0
  },
  label: {
    fontSize: 14,
    color: '#4a4a4a'
  }
})

function Switch(props) {
  const classes = useStyles(props)
  const {label, isChecked, onChange, className} = props

  return (
    <FormControlLabel
      classes={ {
        root: classNames(classes.formControl, className)
      } }
      control={ (
        <MuiSwitch
          checked={ isChecked }
          onChange={ (event) => onChange?.(event.target.checked) }
          classes={ {
            root: classes.root,
            switchBase: classes.switchBase,
            thumb: classes.thumb,
            track: classes.track
          } }
        />
      ) }
      label={ <span className={ classes.label }>{ label }</span> }
    />
  )
}

export default Switch

const defaultProps = {
  size: 18,
  label: '',
  isChecked: false
}

Switch.defaultProps = defaultProps