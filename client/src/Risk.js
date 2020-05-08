import React, {useEffect, useContext} from 'react';
import {Link} from 'react-router';
import App, {DataContext} from './App';
import './Risk.scss';
import OCTableCell from "./components/CommonTable/OCTableCell";
import K8sRisksTable from "./components/Risk/K8sRisksTable";
import { parseK8sRisksWorkloads } from "@octarine/ui-common";
import RefreshIcon from '@material-ui/icons/Refresh';
import CircularProgress from '@material-ui/core/CircularProgress';

function Risk() {
    return (
        <App>
            <DataContext.Consumer >
                {value => (
                    <div className="Risk oc-main-page">
                        <div className="Home-header">
                            <h2>K8S Risk Assessment</h2>
                            <button disabled={value.refreshing} onClick={value.refreshState} className='refresh-state-btn'>
                              {value.refreshing ? (<CircularProgress size='16px' className='refreshIcon' />) : (<RefreshIcon className='refreshIcon' />)}
                              <span>{value.refreshing ? "Refreshing..." : "Refresh"}</span>
                            </button>
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
