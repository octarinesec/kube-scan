import React from 'react';
import App, {DataContext} from './App';
import './Risk.scss';
import K8sRisksTable from "./components/Risk/K8sRisksTable";
import RefreshIcon from '@material-ui/icons/Refresh';
import CircularProgress from '@material-ui/core/CircularProgress';
import Switch from "./components/Switch/Switch";

function Risk() {
    return (
        <App>
            <DataContext.Consumer >
                {value => (
                    <div className="Risk oc-main-page">
                        <div className="Home-header">
                            <h2>K8S Risk Assessment</h2>
                            <div className='right-controls'>
                              <Switch
                                className="risks-include-system-ns"
                                size={14}
                                label="Include System namespaces"
                                isChecked={value.selectedShowSystemNamespaces}
                                onChange={value.setSelectedShowSystemNamespaces}
                              />
                              <button disabled={value.refreshing} onClick={value.refreshState} className='refresh-state-btn'>
                                <div>{value.refreshing ? <CircularProgress disableShrink={true} size='16px' className='refreshIcon' /> : <RefreshIcon className='refreshIcon' />}</div>
                                <span>{value.refreshing ? "Refreshing..." : "Refresh"}</span>
                              </button>
                            </div>
                        </div>
                        <K8sRisksTable
                            risks={value.risks}
                            openPopup={value.openPopup}
                            sortFunc={value.sortFunc}
                            currentSort={value.sortField}
                            ascending={value.ascending}
                        />
                    </div>
                )}

            </DataContext.Consumer>
        </App>
    );
}

export default Risk;
